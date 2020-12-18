// Copyright (c) 2020, devgo.club
// All rights reserved.

package golog

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

// log level
const (
	_DebugLevel = logLevel(iota + 1)
	_InfoLevel
	_WarnLevel
	_ErrorLevel
)

// log level
const (
	DebugLevel = "debug"
	InfoLevel  = "info"
	WarnLevel  = "warn"
	ErrorLevel = "error"
)

// variable
var (
	ErrParam      = errors.New("invalid parameter")
	currentDate   = uint32(0)
	currentTime   = uint32(0)
	currentLogger *logger
	currentPool   = &bufferPool{
		Pool: sync.Pool{
			New: func() interface{} {
				return bytes.NewBuffer(make([]byte, 0))
			},
		},
	}
	cancelCh = make(chan struct{}, 1)
	doneCh   = make(chan struct{}, 1)
)

func init() {
	update()
	go func() {
		var timer = time.NewTicker(time.Second)
		for {
			select {
			case <-timer.C:
				update()
			}
		}
	}()
}

func update() {
	var now = time.Now()
	var date = uint32(now.Year()%100*10000 + int(now.Month())*100 + now.Day())
	var time = uint32(now.Hour()*10000 + now.Minute()*100 + now.Second())
	atomic.StoreUint32(&currentDate, date)
	atomic.StoreUint32(&currentTime, time)
}

// logLevel : log level
type logLevel uint8

// abbr : abbr to one letter
func (level logLevel) abbr() byte {
	switch level {
	case _DebugLevel:
		return 'D'
	case _InfoLevel:
		return 'I'
	case _WarnLevel:
		return 'W'
	case _ErrorLevel:
		return 'E'
	default:
		return 'U'
	}
}

// convertLogLevel : convert log level from string
func convertLogLevel(level string) (logLevel, error) {
	switch level {
	case DebugLevel:
		return _DebugLevel, nil
	case InfoLevel:
		return _InfoLevel, nil
	case WarnLevel:
		return _WarnLevel, nil
	case ErrorLevel:
		return _ErrorLevel, nil
	default:
		return logLevel(0), ErrParam
	}
}

// logger : log object
type logger struct {
	minLevel  logLevel
	formatter formatter
	devices   []device
}

// bgWork : flush every second
func (log *logger) bgWork() {
	go func() {
		var timer = time.NewTicker(time.Second)
		for {
			select {
			case <-timer.C:
				for _, dev := range log.devices {
					dev.Flush()
				}
			case <-cancelCh:
				for _, dev := range log.devices {
					dev.Flush()
				}
				doneCh <- struct{}{}
				break
			}
		}
	}()
}

// output : format and output log message
func (log *logger) output(pool *bufferPool, level logLevel, format string, a ...interface{}) {
	if level < log.minLevel {
		return
	}
	var msg string
	if len(a) == 0 {
		msg = format
	} else {
		msg = fmt.Sprintf(format, a...)
	}
	var buf = log.formatter.Format(pool, level, msg)
	for _, dev := range log.devices {
		dev.Write(buf.Bytes())
	}
	pool.put(buf)
}

type bufferPool struct {
	sync.Pool
}

func (pool *bufferPool) get() *bytes.Buffer {
	return pool.Get().(*bytes.Buffer)
}

func (pool *bufferPool) put(buf *bytes.Buffer) {
	buf.Reset()
	pool.Pool.Put(buf)
}

// formatter : custom formatter
type formatter struct{}

// Format : format message
func (f *formatter) Format(pool *bufferPool, level logLevel, msg string) *bytes.Buffer {
	var buf = pool.get()
	buf.WriteString(fmt.Sprintf("%c %06v", level.abbr(), atomic.LoadUint32(&currentTime)))
	_, file, line, ok := runtime.Caller(3)
	if ok {
		var i int
		for i = len(file) - 1; i >= 0; i-- {
			if file[i] == '/' {
				if i == len(file)-1 {
					break
				} else {
					i++
					break
				}
			}
		}
		buf.WriteString(fmt.Sprintf(" [%v:%v] %v\n", file[i:], line, msg))
	}
	return buf
}

// device : log device
type device interface {
	Write([]byte)
	Flush()
}

// fileDevice : device, print log message to file
type fileDevice struct {
	locker *sync.Mutex
	prefix string
	date   uint32
	file   *os.File
	writer *bufio.Writer
}

func (f *fileDevice) Flush() {
	f.locker.Lock()
	if f.writer != nil {
		f.writer.Flush()
	}
	f.locker.Unlock()
}

func (f *fileDevice) Write(msg []byte) {
	var err error
	var date = atomic.LoadUint32(&currentDate)
	f.locker.Lock()
	// flush old data and set file to nil if a new day comes
	if f.date != date {
		if f.file != nil {
			f.writer.Flush()
			if err = f.file.Close(); err != nil {
				fmt.Printf("[ERROR] golog cannot close file: %v\n", err)
			}
			f.file = nil
		}
	}
	// create a new log file if file is nil
	if f.file == nil {
		var fileName = fmt.Sprintf("%v_%v.log", f.prefix, date)
		f.file, err = os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Printf("[ERROR] golog cannot open file: %v\n", err)
			f.locker.Unlock()
			return
		}
		f.writer = bufio.NewWriter(f.file)
		f.date = date
	}

	// write message
	if _, err = f.writer.Write(msg); err != nil {
		fmt.Printf("[ERROR] golog cannot write file: %v\n", err)
	}
	f.locker.Unlock()
	return
}

func newFileDevice(prefix string) (device, error) {
	path, err := filepath.Abs(filepath.Dir(prefix))
	if err != nil {
		return nil, err
	}
	if err = os.MkdirAll(path, os.ModePerm); err != nil {
		return nil, err
	}
	return &fileDevice{
		locker: new(sync.Mutex),
		prefix: prefix,
	}, nil
}

// stdoutDevice : device, print log message to stdout
type stdoutDevice struct {
	locker *sync.Mutex
	writer *bufio.Writer
}

func (s *stdoutDevice) Flush() {
	s.locker.Lock()
	s.writer.Flush()
	s.locker.Unlock()
}

func (s *stdoutDevice) Write(msg []byte) {
	s.locker.Lock()
	s.writer.Write(msg)
	s.locker.Unlock()
}

func newStdoutDevice() device {
	return &stdoutDevice{
		locker: new(sync.Mutex),
		writer: bufio.NewWriter(os.Stdout),
	}
}

// Init : initialize golog
func Init(level string, enableStdout bool, enableFile bool, filePrefix string) error {
	var minLevel, err = convertLogLevel(level)
	if err != nil {
		return err
	}
	if logLevel(minLevel) < _DebugLevel || logLevel(minLevel) > _ErrorLevel {
		return ErrParam
	}
	if !enableStdout && !enableFile {
		return ErrParam
	}
	var devs = make([]device, 0, 2)
	if enableStdout {
		devs = append(devs, newStdoutDevice())
	}
	if enableFile {
		dev, err := newFileDevice(filePrefix)
		if err != nil {
			return err
		}
		devs = append(devs, dev)
	}
	currentLogger = &logger{
		minLevel: logLevel(minLevel),
		devices:  devs,
	}
	currentLogger.bgWork()
	return nil
}

// Close : close golog and flush buffer
func Close() {
	cancelCh <- struct{}{}
	<-doneCh
}

// Debug : output "debug" message
func Debug(format string, a ...interface{}) {
	if currentLogger == nil {
		return
	}
	currentLogger.output(currentPool, _DebugLevel, format, a...)
}

// Info : output "info" message
func Info(format string, a ...interface{}) {
	if currentLogger == nil {
		return
	}
	currentLogger.output(currentPool, _InfoLevel, format, a...)
}

// Warn : output "warn" message
func Warn(format string, a ...interface{}) {
	if currentLogger == nil {
		return
	}
	currentLogger.output(currentPool, _WarnLevel, format, a...)
}

// Error : output "error" message
func Error(format string, a ...interface{}) {
	if currentLogger == nil {
		return
	}
	currentLogger.output(currentPool, _ErrorLevel, format, a...)
}
