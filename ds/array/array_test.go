package array

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCreate(t *testing.T) {
	arr := Create(false, 0, 1, 2, 3, 4, 5, 4, 6, 7, 8, 9, 3, 2, 1)
	arr.Print(PrintOptions{offsetBegin: 3, offsetEnd: 3})

	vecString := []string{"a", "b", "c"}
	arr.Add(vecString)
	vecInt := []int16{42, 432, 423}
	arr.Add(vecInt)
	vecFloat := []float32{1.1, 1.2, 1.3}
	arr.Add(vecFloat)

	arr.AddAbstractData([]float64{1,2,3}, []string{"234", "abc"})
	arr.AddAbstractData([]float64{1,2,3}, []string{"234", "abc"})
	arr.Add(vecInt)

	for _, element := range arr.data {
		fmt.Println(reflect.TypeOf(element))
	}
}

func TestAdd(t *testing.T) {
	arr := Create(false, 0, 1, 2, 3, 4, 5, 4, 6, 7, 8, 9, 3, 2, 1)
	arr.Print(PrintOptions{offsetBegin: 3, offsetEnd: 3})

	arr.Add(99, 98, 3, 4, 5, 6)
	arr.Print(PrintOptions{})
}
