package gosync

import (
	"context"
	"time"
)

type TaskRes[R any] interface {
	IsDone() bool

	Wait() R
	WaitCtx(ctx context.Context) (R, error)
	WaitTimeout(timeout time.Duration) (R, error)
}

var _ TaskRes[interface{}] = (*taskRes[interface{}])(nil)

type taskRes[R any] struct {
	doneCh chan struct{}
	res    R
}

func newTaskR[R any]() *taskRes[R] {
	return &taskRes[R]{
		doneCh: make(chan struct{}),
	}
}

func (t *taskRes[R]) done(res R) {
	t.res = res
	close(t.doneCh)
}

func (t *taskRes[R]) IsDone() bool {
	select {
	case <-t.doneCh:
		return true
	default:
		return false
	}
}

func (t *taskRes[R]) Wait() R {
	<-t.doneCh
	return t.res
}

func (t *taskRes[R]) WaitCtx(ctx context.Context) (R, error) {
	var nilVal R

	select {
	case <-ctx.Done():
		return nilVal, ctx.Err()
	case <-t.doneCh:
		return t.res, nil
	}
}

func (t *taskRes[R]) WaitTimeout(timeout time.Duration) (R, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return t.WaitCtx(ctx)
}

// GoRes runs fn in a goroutine and returns a TaskRes that
// can be used to wait for the result.
func GoRes[R any](fn func() R) TaskRes[R] {
	t := newTaskR[R]()
	go func() {
		res := fn()
		t.done(res)
	}()
	return t
}

func WaitAllRes[R any](tasks ...TaskRes[R]) []R {
	res := make([]R, len(tasks))
	for i, t := range tasks {
		res[i] = t.Wait()
	}
	return res
}

func WaitAllResCtx[R any](ctx context.Context, tasks ...TaskRes[R]) ([]R, error) {
	res := make(chan []R)
	go func() {
		res <- WaitAllRes(tasks...)
	}()

	var nilVal []R

	select {
	case <-ctx.Done():
		return nilVal, ctx.Err()
	case r := <-res:
		return r, nil
	}
}

func WaitAllResTimeout[R any](timeout time.Duration, tasks ...TaskRes[R]) ([]R, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return WaitAllResCtx(ctx, tasks...)
}

func WaitAnyRes[R any](tasks ...TaskRes[R]) R {
	ch := make(chan R)
	for _, t := range tasks {
		go func(t TaskRes[R]) {
			ch <- t.Wait()
		}(t)
	}
	return <-ch
}

func WaitAnyResCtx[R any](ctx context.Context, tasks ...TaskRes[R]) (R, error) {
	ch := make(chan R)
	go func() {
		ch <- WaitAnyRes(tasks...)
	}()

	var nilVal R

	select {
	case <-ctx.Done():
		return nilVal, ctx.Err()
	case r := <-ch:
		return r, nil
	}
}

func WaitAnyResTimeout[R any](timeout time.Duration, tasks ...TaskRes[R]) (R, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return WaitAnyResCtx(ctx, tasks...)
}
