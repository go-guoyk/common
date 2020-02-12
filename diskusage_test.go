package common

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestDiskUsage(t *testing.T) {
	du := NewDiskUsage(".")
	t.Logf("DiskUsage: %.2f%%", du.Usage()*100)
	require.NotEmpty(t, du.Usage())
}
