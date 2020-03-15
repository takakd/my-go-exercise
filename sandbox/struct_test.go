package sandbox

import (
	"testing"
)

func TestStruct(t *testing.T) {
	s1 := NewStructName1()
	s1.Value = 123
	s2 := NewStructName2()
	s2.Value = 234
	s3 := NewStructName3()
	s3.Value = 345
	s4 := NewStructName4(456)

	// check
	if s1.Value != 123 {
		t.Errorf("result was incorrect, got: %d, want: %d.", s1.Value, 123)
	}
	if s2.Value != 234 {
		t.Errorf("result was incorrect, got: %d, want: %d.", s2.Value, 234)
	}
	if s3.Value != 345 {
		t.Errorf("result was incorrect, got: %d, want: %d.", s3.Value, 345)
	}
	if s4.Value != 456 {
		t.Errorf("result was incorrect, got: %d, want: %d.", s4.Value, 456)
	}
}
