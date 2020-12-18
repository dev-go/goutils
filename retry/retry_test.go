// Copyright (c) 2020, devgo.club
// All rights reserved.

package retry

import (
	"context"
	"errors"
	"testing"
	"time"
)

func TestRetry(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	var err error
	go func() {
		err = Do(ctx, 10, func() error {
			t.Logf("time=%v", time.Now())
			return errors.New("fuck you")
		}, func(attampt int) time.Duration {
			return time.Duration(attampt) * time.Second
		})
	}()
	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
	t.Logf("%v", err)
	return
}
