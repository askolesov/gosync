package gosync

import (
	"context"
	"time"
)

type Task interface {
	IsDone() bool

	Wait()
	WaitCtx(ctx context.Context) error
	WaitTimeout(timeout time.Duration) error
}

var _ Task = (*task)(nil)

type task struct {
	doneCh chan struct{}
}

func newTask() *task {
	return &task{
		doneCh: make(chan struct{}),
	}
}

func (t *task) done() {
	close(t.doneCh)
}

func (t *task) IsDone() bool {
	select {
	case <-t.doneCh:
		return true
	default:
		return false
	}
}

func (t *task) Wait() {
	<-t.doneCh
}

func (t *task) WaitCtx(ctx context.Context) error {
	select {
	case <-t.doneCh:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func (t *task) WaitTimeout(timeout time.Duration) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return t.WaitCtx(ctx)
}

// Go runs fn in a goroutine and returns a Task that
// can be used to wait for the goroutine to finish.
func Go(fn func()) Task {
	t := newTask()
	go func() {
		fn()
		t.done()
	}()
	return t
}

func WaitAll(tasks ...Task) {
	for _, t := range tasks {
		t.Wait()
	}
}

func WaitAllCtx(ctx context.Context, tasks ...Task) error {
	done := make(chan struct{})
	go func() {
		WaitAll(tasks...)
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func WaitAllTimeout(timeout time.Duration, tasks ...Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return WaitAllCtx(ctx, tasks...)
}

func WaitAny(tasks ...Task) {
	done := make(chan struct{})
	for _, t := range tasks {
		go func(t Task) {
			t.Wait()
			done <- struct{}{}
		}(t)
	}
	<-done
}

func WaitAnyCtx(ctx context.Context, tasks ...Task) error {
	done := make(chan struct{})
	go func() {
		WaitAny(tasks...)
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

func WaitAnyTimeout(timeout time.Duration, tasks ...Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return WaitAnyCtx(ctx, tasks...)
}
