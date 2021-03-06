package common

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRunnableGroup_Run(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan interface{})

	r1 := &testRunnable{shouldFail: true}
	r2 := &testRunnable{shouldFail: false}
	r3 := &testRunnable{shouldFail: false}

	rg := NewRunnableGroup()
	rg.Add(r1)
	rg.Add(r2)
	rg.Add(r3)

	var err error

	go func() {
		err = rg.Run(ctx, cancel, done)
	}()

	<-done

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

func TestRunnableGroup_Run1(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	done := make(chan interface{})
	done2 := make(chan interface{})

	r1 := &testRunnable{shouldFail: true}
	r2 := &testRunnable{shouldFail: false}

	rg := NewRunnableGroup()
	rg.Add(r1)

	var err error

	go func() {
		err = rg.Run(ctx, cancel, done)
	}()

	go func() {
		r2.Run(ctx)
		close(done2)
	}()

	<-done
	<-done2

	assert.Error(t, err, "should failed")
	assert.True(t, r1.completed, "r1 should complete")
	assert.True(t, r2.completed, "r2 should complete")

	assert.False(t, r1.succeeded, "r1 should not success")
	assert.False(t, r1.cancelled, "r1 should not cancel")
	assert.False(t, r2.succeeded, "r2 should not success")
	assert.True(t, r2.cancelled, "r2 should cancel")
}
