# MusGo_int16
Provides serialization of the Golang's `int16` type.

# How to use
Simply cast `int16` to `musgo_int16.Int16`:
```go
	var n int16 = 5
	// Marshal
	buf := make([]byte, musgo_int16.Int16(n).SizeMUS())
	musgo_int16.Int16(n).MarshalMUS(buf)
	// Unmarshal
	_, err := (*musgo_int16.Int16)(&n).UnmarshalMUS(buf)
	if err != nil {
		panic(err)
	}
```

# More info
You can find at [github.com/ymz-ncnk/musgo](https://github.com/ymz-ncnk/musgo).

