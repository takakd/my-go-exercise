package sandbox

type StructName struct {
	Value int
}

func NewStructName1() *StructName {
	s := &StructName{}
	s.Value = 123
	return s
}

// this is same as NewStructName1
func NewStructName2() *StructName {
	s := new(StructName)
	s.Value = 123
	return s
}

// Basically use NewStructName1, but in special cases use this.
// @ref: https://medium.com/a-journey-with-go/go-should-i-use-a-pointer-instead-of-a-copy-of-my-struct-44b43b104963
func NewStructName3() StructName {
	s := StructName{}
	s.Value = 123
	return s
}

// with parameter
func NewStructName4(value int) *StructName {
	s := &StructName{Value: value}
	// or
	s = &StructName{}
	s.Value = value
	return s
}
