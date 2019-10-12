package conv4go

import "testing"

func TestBool(t *testing.T) {
	var tests = []struct {
		v interface{}
		r bool
	}{
		{"true", true},
		{"yes", true},
		{"on", true},
		{"t", true},
		{"y", true},
		{"1", true},
		{"0", false},
		{"2", false},
		{"false", false},
		{"no", false},
	}

	for _, tt := range tests {
		if actual := Bool(tt.v); actual != tt.r {
			t.Errorf("把 %v 转换为 bool, 期望获得 %v, 实际获得  %v", tt.v, tt.r, actual)
		}
	}
}
