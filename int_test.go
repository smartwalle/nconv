package nconv_test

import (
	"github.com/smartwalle/nconv"
	"math"
	"testing"
)

func TestInt(t *testing.T) {
	var tests = []struct {
		v interface{}
		r int
	}{
		{1, 1},
		{9, 9},
		{9999, 9999},
		{uint64(9999), 9999},
		{int64(9999), 9999},
		{1.119, 1},
		{9.119, 9},
		{9.999, 9},
		{-2, -2},
		{"1", 1},
		{"-1", -1},
		{"999999", 999999},
		{"999.999", 999},
		{"999.1", 999},
		{"1.1111", 1},
		{"1.9999", 1},
	}

	for _, tt := range tests {
		if actual := nconv.Int(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 int, 期望获得 %d, 实际获得  %d", tt.v, tt.r, actual)
		}
	}
}

func TestInt64(t *testing.T) {
	var tests = []struct {
		v interface{}
		r int64
	}{
		{1, 1},
		{9, 9},
		{9999, 9999},
		{uint64(9999), 9999},
		{int64(9999), 9999},
		{1.119, 1},
		{9.119, 9},
		{9.999, 9},
		{"1", 1},
		{"999999", 999999},
		{"999.999", 999},
		{"999.1", 999},
		{"1.1111", 1},
		{"1.9999", 1},
		{"", 0},
		{"3414416614257328130", 3414416614257328130},
	}

	for _, tt := range tests {
		if actual := nconv.Int64(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 int64, 期望获得 %d, 实际获得  %d", tt.v, tt.r, actual)
		}
	}
}

func TestUint(t *testing.T) {
	var tests = []struct {
		v interface{}
		r uint
	}{
		{"1", 1},
		{"999999", 999999},
		{"999.999", 999},
		{"999.1", 999},
		{"1.1111", 1},
		{"1.9999", 1},
		{1000, 1000},
		{9991, 9991},
		{1.119, 1},
		{9.119, 9},
		{9.999, 9},
		{uint(9991), 9991},
	}

	for _, tt := range tests {
		if actual := nconv.Uint(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 uint32, 期望获得 %d, 实际获得  %d", tt.v, tt.r, actual)
		}
	}
}

func TestUint8(t *testing.T) {
	var tests = []struct {
		v interface{}
		r uint8
	}{
		{"1", 1},
		{"65", 65},
		{"9.999", 9},
		{"99.1", 99},
		{"1.1111", 1},
		{"1.9999", 1},
		{100, 100},
		{99, 99},
		{1.119, 1},
		{9.119, 9},
		{9.999, 9},
		{uint(99), 99},
		{uint8(math.MaxUint8), math.MaxUint8},
	}

	for _, tt := range tests {
		if actual := nconv.Uint8(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 uint8, 期望获得 %d, 实际获得  %d", tt.v, tt.r, actual)
		}
	}
}

func TestUint16(t *testing.T) {
	var tests = []struct {
		v interface{}
		r uint16
	}{
		{"1", 1},
		{"65535", 65535},
		{"999.999", 999},
		{"999.1", 999},
		{"1.1111", 1},
		{"1.9999", 1},
		{1000, 1000},
		{9991, 9991},
		{1.119, 1},
		{9.119, 9},
		{9.999, 9},
		{uint(9991), 9991},
		{uint16(math.MaxUint16), math.MaxUint16},
	}

	for _, tt := range tests {
		if actual := nconv.Uint16(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 uint16, 期望获得 %d, 实际获得  %d", tt.v, tt.r, actual)
		}
	}
}

func TestUint32(t *testing.T) {
	var tests = []struct {
		v interface{}
		r uint32
	}{
		{"1", 1},
		{"999999", 999999},
		{"999.999", 999},
		{"999.1", 999},
		{"1.1111", 1},
		{"1.9999", 1},
		{1000, 1000},
		{9991, 9991},
		{1.119, 1},
		{9.119, 9},
		{9.999, 9},
		{uint(9991), 9991},
		{uint32(math.MaxUint32), math.MaxUint32},
	}

	for _, tt := range tests {
		if actual := nconv.Uint32(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 uint32, 期望获得 %d, 实际获得  %d", tt.v, tt.r, actual)
		}
	}
}

func TestUint64(t *testing.T) {
	var tests = []struct {
		v interface{}
		r uint64
	}{
		{"1", 1},
		{"999999", 999999},
		{"999.999", 999},
		{"999.1", 999},
		{"1.1111", 1},
		{"1.9999", 1},
		{1000, 1000},
		{9991, 9991},
		{1.119, 1},
		{9.119, 9},
		{9.999, 9},
		{uint(9991), 9991},
		{uint64(math.MaxUint64), math.MaxUint64},
	}

	for _, tt := range tests {
		if actual := nconv.Uint64(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 Uint64, 期望获得 %d, 实际获得  %d", tt.v, tt.r, actual)
		}
	}
}
