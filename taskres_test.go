package gosync_test

import (
	"github.com/askolesov/gosync"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestGoRes(t *testing.T) {
	ts := gosync.GoRes(func() int {
		time.Sleep(100 * time.Millisecond)
		return 1
	})

	time.Sleep(10 * time.Millisecond)
	require.False(t, ts.IsDone())

	_, err := ts.WaitTimeout(50 * time.Millisecond)
	require.Error(t, err)

	res := ts.Wait()
	require.Equal(t, 1, res)
	require.True(t, ts.IsDone())

	// test that Wait() can be called multiple times
	res, err = ts.WaitTimeout(50 * time.Millisecond)
	require.NoError(t, err)
	require.Equal(t, 1, res)
}

func TestWaitAllRes(t *testing.T) {
	ts1 := gosync.GoRes(func() int {
		time.Sleep(50 * time.Millisecond)
		return 1
	})
	ts2 := gosync.GoRes(func() int {
		time.Sleep(100 * time.Millisecond)
		return 2
	})
	ts3 := gosync.GoRes(func() int {
		time.Sleep(150 * time.Millisecond)
		return 3
	})

	start := time.Now()

	res := gosync.WaitAllRes(ts1, ts2, ts3)

	require.InDelta(t, 150*time.Millisecond, time.Since(start), float64(20*time.Millisecond))
	require.Equal(t, []int{1, 2, 3}, res)
}

func TestWaitAllResTimeout(t *testing.T) {
	ts1 := gosync.GoRes(func() int {
		time.Sleep(50 * time.Millisecond)
		return 1
	})
	ts2 := gosync.GoRes(func() int {
		time.Sleep(100 * time.Millisecond)
		return 2
	})
	ts3 := gosync.GoRes(func() int {
		time.Sleep(150 * time.Millisecond)
		return 3
	})

	start := time.Now()

	_, err := gosync.WaitAllResTimeout(75*time.Millisecond, ts1, ts2, ts3) // ts3 should not finish
	require.Error(t, err)

	res, err := gosync.WaitAllResTimeout(500*time.Millisecond, ts1, ts2, ts3) // ts3 should finish
	require.NoError(t, err)

	require.InDelta(t, 150*time.Millisecond, time.Since(start), float64(20*time.Millisecond))
	require.Equal(t, []int{1, 2, 3}, res)
}

func TestWaitAnyRes(t *testing.T) {
	ts1 := gosync.GoRes(func() int {
		time.Sleep(50 * time.Millisecond)
		return 1
	})
	ts2 := gosync.GoRes(func() int {
		time.Sleep(100 * time.Millisecond)
		return 2
	})
	ts3 := gosync.GoRes(func() int {
		time.Sleep(150 * time.Millisecond)
		return 3
	})

	start := time.Now()

	gosync.WaitAnyRes(ts1, ts2, ts3)

	require.InDelta(t, 50*time.Millisecond, time.Since(start), float64(20*time.Millisecond))
}

func TestWaitAnyResTimeout(t *testing.T) {
	ts1 := gosync.GoRes(func() int {
		time.Sleep(50 * time.Millisecond)
		return 1
	})
	ts2 := gosync.GoRes(func() int {
		time.Sleep(100 * time.Millisecond)
		return 2
	})
	ts3 := gosync.GoRes(func() int {
		time.Sleep(150 * time.Millisecond)
		return 3
	})

	start := time.Now()

	_, err := gosync.WaitAnyResTimeout(10*time.Millisecond, ts1, ts2, ts3) // ts1 should not finish
	require.Error(t, err)

	res, err := gosync.WaitAnyResTimeout(500*time.Millisecond, ts1, ts2, ts3) // ts1 should finish
	require.NoError(t, err)

	require.InDelta(t, 50*time.Millisecond, time.Since(start), float64(20*time.Millisecond))
	require.Equal(t, 1, res)
}
