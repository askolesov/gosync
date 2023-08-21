package gosync_test

import (
	"github.com/askolesov/gosync"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGo(t *testing.T) {
	ts := gosync.Go(func() {
		time.Sleep(100 * time.Millisecond)
	})

	time.Sleep(10 * time.Millisecond)
	require.False(t, ts.IsDone())

	err := ts.WaitTimeout(50 * time.Millisecond)
	require.Error(t, err)

	ts.Wait()
	require.True(t, ts.IsDone())

	// test that Wait() can be called multiple times
	err = ts.WaitTimeout(50 * time.Millisecond)
	require.NoError(t, err)
}

func TestWaitAll(t *testing.T) {
	ts1 := gosync.Go(func() {
		time.Sleep(50 * time.Millisecond)
	})
	ts2 := gosync.Go(func() {
		time.Sleep(100 * time.Millisecond)
	})
	ts3 := gosync.Go(func() {
		time.Sleep(150 * time.Millisecond)
	})

	start := time.Now()
	gosync.WaitAll(ts1, ts2, ts3)
	require.InDelta(t, 150*time.Millisecond, time.Since(start), float64(20*time.Millisecond))
}

func TestWaitAllTimeout(t *testing.T) {
	ts1 := gosync.Go(func() {
		time.Sleep(50 * time.Millisecond)
	})
	ts2 := gosync.Go(func() {
		time.Sleep(100 * time.Millisecond)
	})
	ts3 := gosync.Go(func() {
		time.Sleep(150 * time.Millisecond)
	})

	start := time.Now()

	err := gosync.WaitAllTimeout(75*time.Millisecond, ts1, ts2, ts3) // ts3 should not finish
	require.Error(t, err)

	err = gosync.WaitAllTimeout(500*time.Millisecond, ts1, ts2, ts3) // ts3 should finish
	require.NoError(t, err)

	require.InDelta(t, 150*time.Millisecond, time.Since(start), float64(20*time.Millisecond))
}

func TestWaitAny(t *testing.T) {
	ts1 := gosync.Go(func() {
		time.Sleep(50 * time.Millisecond)
	})
	ts2 := gosync.Go(func() {
		time.Sleep(100 * time.Millisecond)
	})
	ts3 := gosync.Go(func() {
		time.Sleep(150 * time.Millisecond)
	})

	start := time.Now()

	gosync.WaitAny(ts1, ts2, ts3)

	require.InDelta(t, 50*time.Millisecond, time.Since(start), float64(20*time.Millisecond))
}

func TestWaitAnyTimeout(t *testing.T) {
	ts1 := gosync.Go(func() {
		time.Sleep(50 * time.Millisecond)
	})
	ts2 := gosync.Go(func() {
		time.Sleep(100 * time.Millisecond)
	})
	ts3 := gosync.Go(func() {
		time.Sleep(150 * time.Millisecond)
	})

	start := time.Now()

	err := gosync.WaitAnyTimeout(10*time.Millisecond, ts1, ts2, ts3) // ts1 should not finish
	require.Error(t, err)

	err = gosync.WaitAnyTimeout(500*time.Millisecond, ts1, ts2, ts3) // ts1 should finish
	require.NoError(t, err)

	require.InDelta(t, 50*time.Millisecond, time.Since(start), float64(20*time.Millisecond))
}
