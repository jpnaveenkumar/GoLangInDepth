package main

import (
	"fmt"
	"unsafe"
)

type biodata struct {
	name    string // 16
	address string // 16
	isMale  bool   // 1
}

type StructA struct {
	a int8              // 1
	b string            // 16
	c int64             // 8
	d bool              // 1
	f map[string]string // 8
}

type StructB struct {
	b string            // 16
	c int64             // 8
	f map[string]string // 8
	d bool              // 1
	a int8              // 1
}

type BadStruct struct {
	a int8              // 1
	b string            // 16
	c int64             // 8
	d bool              // 1
	e biodata           // 40
	f map[string]string // 8
	g []biodata         // 24
}

type GoodStruct struct {
	e biodata           // 40
	g []biodata         // 24
	b string            // 16
	f map[string]string // 8
	c int64             // 8
	a int8              // 1
	d bool              // 1
}

func main() {
	badStruct := BadStruct{
		a: 0,
		b: "",
		c: 0,
		d: false,
		e: biodata{},
		f: nil,
		g: []biodata{
			{
				name:   "naveen",
				isMale: true,
			},
			{
				name:   "kumar",
				isMale: true,
			},
		},
	}

	goodStruct := GoodStruct{
		e: biodata{},
		g: nil,
		b: "",
		f: nil,
		c: 0,
		a: 0,
		d: false,
	}

	fmt.Println("Size of a : ", unsafe.Sizeof(badStruct.a))
	fmt.Println("Size of b : ", unsafe.Sizeof(badStruct.b))
	fmt.Println("Size of c : ", unsafe.Sizeof(badStruct.c))
	fmt.Println("Size of d : ", unsafe.Sizeof(badStruct.d))
	fmt.Println("Size of e : ", unsafe.Sizeof(badStruct.e))
	fmt.Println("Size of f : ", unsafe.Sizeof(badStruct.f))
	fmt.Println("Size of g : ", unsafe.Sizeof(badStruct.g))

	fmt.Println("Size of bad struct : ", unsafe.Sizeof(badStruct))

	fmt.Println("Size of good struct : ", unsafe.Sizeof(goodStruct))

	//structA := StructA{}
	//fmt.Println("Size of structA : ", unsafe.Sizeof(structA))
	//structB := StructB{}
	//fmt.Println("Size of structA : ", unsafe.Sizeof(structB))
}
