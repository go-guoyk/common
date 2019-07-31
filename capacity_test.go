package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseCapacity(t *testing.T) {
	var err error
	var c Capacity

	c, err = ParseCapacity(" 10 B ")
	assert.NoError(t, err, "no error")
	assert.Equal(t, Byte*10, c, "10B")

	c, err = ParseCapacity("   10  ")
	assert.NoError(t, err, "no error")
	assert.Equal(t, Byte*10, c, "10B")

	c, err = ParseCapacity("10 GB")
	assert.NoError(t, err, "no error")
	assert.Equal(t, Gigabyte*10, c, "10 GB")

	c, err = ParseCapacity("10.20GB")
	assert.NoError(t, err, "no error")
	assert.Equal(t, Gigabyte*10+Megabyte*200, c, "10.20GB")

	c, err = ParseCapacity(" 10.20 GB")
	assert.NoError(t, err, "no error")
	assert.Equal(t, Gigabyte*10+Megabyte*200, c, "10.20GB")

	c, err = ParseCapacity(" 10.20 GB ")
	assert.NoError(t, err, "no error")
	assert.Equal(t, Gigabyte*10+Megabyte*200, c, "10.20GB")

	c, err = ParseCapacity(" 10.20 giB ")
	assert.NoError(t, err, "no error")
	assert.Equal(t, Gibibyte*1020/100, c, "10.20GiB")
}

func TestParseCapacityUnit(t *testing.T) {
	var err error
	var c Capacity

	c, err = ParseCapacityUnit("kib")
	assert.NoError(t, err, "kib not parse")
	assert.Equal(t, Kibibyte, c, "kib not match")
	c, err = ParseCapacityUnit("Mib")
	assert.NoError(t, err, "mib not parse")
	assert.Equal(t, Mebibyte, c, "mib not match")
	c, err = ParseCapacityUnit("gib")
	assert.NoError(t, err, "gib not parse")
	assert.Equal(t, Gibibyte, c, "gib not match")
	c, err = ParseCapacityUnit("Tib")
	assert.NoError(t, err, "Tib not parse")
	assert.Equal(t, Tebibyte, c, "Tib not match")
	c, err = ParseCapacityUnit("pib")
	assert.NoError(t, err, "pib not parse")
	assert.Equal(t, Pebibyte, c, "pib not match")
	c, err = ParseCapacityUnit("eib")
	assert.NoError(t, err, "eib not parse")
	assert.Equal(t, Exbibyte, c, "eib not match")

	c, err = ParseCapacityUnit("kb")
	assert.NoError(t, err, "kb not parse")
	assert.Equal(t, Kilobyte, c, "kb not match")
	c, err = ParseCapacityUnit("Mb")
	assert.NoError(t, err, "mb not parse")
	assert.Equal(t, Megabyte, c, "mb not match")
	c, err = ParseCapacityUnit("gb")
	assert.NoError(t, err, "gb not parse")
	assert.Equal(t, Gigabyte, c, "gb not match")
	c, err = ParseCapacityUnit("Tb")
	assert.NoError(t, err, "Tb not parse")
	assert.Equal(t, Terabyte, c, "Tb not match")
	c, err = ParseCapacityUnit("pb")
	assert.NoError(t, err, "pb not parse")
	assert.Equal(t, Petabyte, c, "pb not match")
	c, err = ParseCapacityUnit("eb")
	assert.NoError(t, err, "eb not parse")
	assert.Equal(t, Exabyte, c, "eb not match")

	c, err = ParseCapacityUnit("kib")
	assert.NoError(t, err, "kib not parse")
	assert.Equal(t, Kibibyte, c, "kib not match")
	c, err = ParseCapacityUnit("Mi")
	assert.NoError(t, err, "mib not parse")
	assert.Equal(t, Mebibyte, c, "mib not match")
	c, err = ParseCapacityUnit("gi")
	assert.NoError(t, err, "gib not parse")
	assert.Equal(t, Gibibyte, c, "gib not match")
	c, err = ParseCapacityUnit("Ti")
	assert.NoError(t, err, "Tib not parse")
	assert.Equal(t, Tebibyte, c, "Tib not match")
	c, err = ParseCapacityUnit("pi")
	assert.NoError(t, err, "pib not parse")
	assert.Equal(t, Pebibyte, c, "pib not match")
	c, err = ParseCapacityUnit("ei")
	assert.NoError(t, err, "eib not parse")
	assert.Equal(t, Exbibyte, c, "eib not match")

	c, err = ParseCapacityUnit("k")
	assert.NoError(t, err, "kb not parse")
	assert.Equal(t, Kilobyte, c, "kb not match")
	c, err = ParseCapacityUnit("M")
	assert.NoError(t, err, "mb not parse")
	assert.Equal(t, Megabyte, c, "mb not match")
	c, err = ParseCapacityUnit("g")
	assert.NoError(t, err, "gb not parse")
	assert.Equal(t, Gigabyte, c, "gb not match")
	c, err = ParseCapacityUnit("T")
	assert.NoError(t, err, "Tb not parse")
	assert.Equal(t, Terabyte, c, "Tb not match")
	c, err = ParseCapacityUnit("p")
	assert.NoError(t, err, "pb not parse")
	assert.Equal(t, Petabyte, c, "pb not match")
	c, err = ParseCapacityUnit("e")
	assert.NoError(t, err, "eb not parse")
	assert.Equal(t, Exabyte, c, "eb not match")
}
