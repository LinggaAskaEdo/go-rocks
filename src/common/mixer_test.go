package common

import (
	"testing"
)

var (
	input1          int64  = 1
	expectedResult1 string = "rghtzlbmgjroc2fq"
	input2          int64  = 2
	expectedResult2 string = "rghtzlbmgjroc29q"
	input3          int64  = 10001
	expectedResult3 string = "jwzthgobh8m5tofq"
	input4          int64  = 10040
	expectedResult4 string = "j4z3hgobh8m5tofq"
)

func TestMixerEncode(t *testing.T) {
	result := MixerEncode(input1)

	if result != expectedResult1 {
		t.Errorf("SALAH !!! harusnya %s", result)
	} else {
		t.Log("BENAR !!!")
	}
}

func TestMixerEncode2(t *testing.T) {
	result := MixerEncode(input2)

	if result != expectedResult2 {
		t.Errorf("SALAH !!! harusnya %s", result)
	} else {
		t.Log("BENAR !!!")
	}
}

func TestMixerEncode3(t *testing.T) {
	result := MixerEncode(input3)

	if result != expectedResult3 {
		t.Errorf("SALAH !!! harusnya %s", result)
	} else {
		t.Log("BENAR !!!")
	}
}

func TestMixerEncode4(t *testing.T) {
	result := MixerEncode(input4)

	if result != expectedResult4 {
		t.Errorf("SALAH !!! harusnya %s", result)
	} else {
		t.Log("BENAR !!!")
	}
}