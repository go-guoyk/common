package common

import (
	"context"
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type testRunnable struct {
	shouldFail bool

	completed bool
	succeeded bool
	cancelled bool
}

func (t *testRunnable) Run(ctx context.Context) error {
	if t.shouldFail {
		time.Sleep(time.Second / 2)
		t.completed = true
		return errors.New("failed")
	}

	m := time.NewTimer(time.Second)
	select {
	case <-ctx.Done():
		t.cancelled = true
	case <-m.C:
		t.succeeded = true
	}
	t.completed = true

	return nil
}

func TestRun(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error)

	r1 := &testRunnable{shouldFail: true}
	r2 := &testRunnable{shouldFail: false}
	r3 := &testRunnable{shouldFail: false}

	RunAsync(ctx, cancel, done, r1, r2, r3)

	err := <-done

	assert.Error(t, err, "should failed")
	assert.True(t, r1.completed, "r1 should complete")
	assert.True(t, r2.completed, "r2 should complete")
	assert.True(t, r3.completed, "r3 should complete")

	assert.False(t, r1.succeeded, "r1 should not success")
	assert.False(t, r1.cancelled, "r1 should not cancel")
	assert.False(t, r2.succeeded, "r2 should not success")
	assert.True(t, r2.cancelled, "r2 should cancel")
	assert.False(t, r3.succeeded, "r3 should not success")
	assert.True(t, r3.cancelled, "r3 should cancel")
}

func TestRun1(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan error)
	done2 := make(chan interface{})

	r1 := &testRunnable{shouldFail: true}
	r2 := &testRunnable{shouldFail: false}

	RunAsync(ctx, cancel, done, r1)

	go func() {
		r2.Run(ctx)
		close(done2)
	}()

	err := <-done
	<-done2

	assert.Error(t, err, "should failed")
	assert.True(t, r1.completed, "r1 should complete")
	assert.True(t, r2.completed, "r2 should complete")

	assert.False(t, r1.succeeded, "r1 should not success")
	assert.False(t, r1.cancelled, "r1 should not cancel")
	assert.False(t, r2.succeeded, "r2 should not success")
	assert.True(t, r2.cancelled, "r2 should cancel")
}

func TestRunWithNoCancel(t *testing.T) {
	ctx := context.Background()

	r1 := &testRunnable{shouldFail: false}
	r2 := &testRunnable{shouldFail: false}

	err := Run(ctx, nil, nil, r1, r2)

	assert.NoError(t, err, "should not failed")
	assert.True(t, r1.completed, "r1 should complete")
	assert.True(t, r2.completed, "r2 should complete")
}
