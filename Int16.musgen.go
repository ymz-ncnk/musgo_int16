// Code generated by musgen. DO NOT EDIT.

package musgo_int16

import "github.com/ymz-ncnk/musgo/errs"

// MarshalMUS fills buf with the MUS encoding of v.
func (v Int16) MarshalMUS(buf []byte) int {
	i := 0
	{
		uv := uint16(v)
		if v < 0 {
			uv = ^(uv << 1)
		} else {
			uv = uv << 1
		}
		{
			for uv >= 0x80 {
				buf[i] = byte(uv) | 0x80
				uv >>= 7
				i++
			}
			buf[i] = byte(uv)
			i++
		}
	}
	return i
}

// UnmarshalMUS parses the MUS-encoded buf, and sets the result to *v.
func (v *Int16) UnmarshalMUS(buf []byte) (int, error) {
	i := 0
	var err error
	{
		var uv uint16
		{
			if i > len(buf)-1 {
				return i, errs.ErrSmallBuf
			}
			shift := 0
			done := false
			for l, b := range buf[i:] {
				if l == 2 && b > 3 {
					return i, errs.ErrOverflow
				}
				if b < 0x80 {
					uv = uv | uint16(b)<<shift
					done = true
					i += l + 1
					break
				}
				uv = uv | uint16(b&0x7F)<<shift
				shift += 7
			}
			if !done {
				return i, errs.ErrSmallBuf
			}
		}
		if uv&1 == 1 {
			uv = ^(uv >> 1)
		} else {
			uv = uv >> 1
		}
		(*v) = Int16(uv)
	}
	return i, err
}

// SizeMUS returns the size of the MUS-encoded v.
func (v Int16) SizeMUS() int {
	size := 0
	{
		uv := uint16(v<<1) ^ uint16(v>>15)
		{
			for uv >= 0x80 {
				uv >>= 7
				size++
			}
			size++
		}
	}
	return size
}
