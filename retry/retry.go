// Copyright (c) 2020, devgo.club
// All rights reserved.

package retry

import (
	"context"
	"fmt"
	"time"
)

// Errors : errors
type Errors []error

func (errs Errors) Error() string {
	var str = ""
	for i, e := range errs {
		str += fmt.Sprintf("[%d] %s", i, e)
		if i != len(errs)-1 {
			str += " | "
		}
	}
	return str
}

// Do : try to execute a function, specify the number of attempts, and the execution interval
func Do(ctx context.Context, attempts int, f func() error, delay func(attempt int) time.Duration) error {
	var errs = make(Errors, 0, attempts)
	var i = 0
	var t = time.NewTimer(delay(0))
	for {
		select {
		case <-ctx.Done():
			return fmt.Errorf("cancel: attempt=%v, errs=%v", i, errs)
		case <-t.C:
			var err error
			if err = f(); err == nil {
				return nil
			}
			errs = append(errs, err)
			i++
			if i > attempts {
				return errs
			}
			t.Reset(delay(i))
		}
	}
}
