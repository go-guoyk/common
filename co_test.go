package common

import (
	"context"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

type coTestRunnable struct {
	c      Co
	done   bool
	doneCh chan interface{}
}

func (ct *coTestRunnable) Run(ctx context.Context) error {
	ct.c.Take()
	defer ct.c.Return()
	time.Sleep(time.Second)
	ct.done = true
	ct.doneCh <- nil
	return nil
}

func TestNewCo(t *testing.T) {
	c := NewCo(-1)
	c.Take()

	c = NewCo(2)
	done := make(chan interface{}, 4)

	rs := []Runnable{
		&coTestRunnable{c: c, doneCh: done},
		&coTestRunnable{c: c, doneCh: done},
		&coTestRunnable{c: c, doneCh: done},
		&coTestRunnable{c: c, doneCh: done},
	}

	RunAsync(context.Background(), nil, nil, rs...)
	<-done
	<-done

	count := 0
	for _, r := range rs {
		if r.(*coTestRunnable).done {
			count++
		}
	}

	time.Sleep(time.Millisecond * 500)

	require.Equal(t, 2, count, "must 2 complete")

	<-done
	<-done

	count = 0
	for _, r := range rs {
		if r.(*coTestRunnable).done {
			count++
		}
	}

	require.Equal(t, 4, count, "must 4 complete")
}
