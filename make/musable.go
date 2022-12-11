//go:build ignore

package main

import (
	"reflect"

	mgi "github.com/ymz-ncnk/musgo_int16"

	"github.com/ymz-ncnk/dotmusgo"
)

func main() {
	musGo, err := dotmusgo.New()
	if err != nil {
		panic(err)
	}
	unsafe := false
	err = musGo.Generate(reflect.TypeOf((*mgi.Int16)(nil)).Elem(), unsafe)
	if err != nil {
		panic(err)
	}
}
