package DogPack

import (
	"fmt"
	"unsafe"
)

type Dog struct {
	name string
}

func New(name string) Dog {
	d := Dog{name}
	fmt.Println(&d.name)
	return d
}

func GetNameStr(d *Dog) *string {
	pDog := &d
	ptrDog := uintptr(unsafe.Pointer(pDog))
	ptrName := ptrDog + unsafe.Offsetof(d.name)
	pName := (*string)(unsafe.Pointer(ptrName))
	return pName
}

func GetNameStrA(d *Dog) *string {
	return &(d.name)
}