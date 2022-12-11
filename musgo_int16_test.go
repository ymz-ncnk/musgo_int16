package musgo_int16

import "testing"

func TestMusgoInt16(t *testing.T) {
	var n int16 = 5
	buf := make([]byte, Int16(n).SizeMUS())
	Int16(n).MarshalMUS(buf)

	var an int16
	(*Int16)(&an).UnmarshalMUS(buf)

	if n != an {
		t.Fatal("something went wrong")
	}
}
