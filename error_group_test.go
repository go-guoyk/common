package common

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestErrorGroup(t *testing.T) {
	eg := NewSafeErrorGroup()
	eg.Add(nil)
	assert.Equal(t, nil, eg.Err(), "should be nil")
	err1 := errors.New("err1")
	eg.Add(err1)
	assert.Equal(t, err1, eg.Err(), "should be err1")
	err2 := errors.New("err2")
	eg.Add(err2)
	assert.NotEqual(t, err1, eg.Err(), "should be no longer err1")
	assert.Equal(t, "err1; err2", eg.Err().Error(), "should compose Error()")
}
