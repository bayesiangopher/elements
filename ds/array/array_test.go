package array

import (
	"testing"
)

func TestCreate(t *testing.T) {
	arr := Create(false, 0,1,2,3,4,5,4,6,7,8,9,3,2,1)
	arr.Print(PrintOptions{offsetBegin:3,offsetEnd:3})
}

func TestAdd(t *testing.T) {
	arr := Create(false, 0,1,2,3,4,5,4,6,7,8,9,3,2,1)
	arr.Print(PrintOptions{offsetBegin:3,offsetEnd:3})

	arr.Add(99,98,3,4,5,6)
	arr.Print(PrintOptions{})
}