// Package array implements simple arrays structures:
// 1) array with duplicates
// 2) array with no duplicates

// O(x):
// Operation    No duplicates     Duplicates
// Search       N/2               N
// Insertion    N, 1              1, 1
// Deletion     N/2, N/2          N, N/2

package array

import (
	"fmt"
	ref "reflect"
)

const (
	defaultOffsetB = 3
	defaultOffsetE = 4
)

// Array is structure for simple array data structure
type Array struct {
	data       []interface{}
	size       int
	duplicates bool
}

type PrintOptions struct {
	offsetBegin int
	offsetEnd   int
}

type emptyContainerError struct {
	err string
}

func (e *emptyContainerError) Error() string {
	return fmt.Sprintf("creating error: %v", e.err)
}

// Create creates new array structure filled by values...
func Create(duplicates bool, values ...interface{}) *Array {
	array := &Array{duplicates: duplicates}
	array.Add(values...)
	return array
}

// Add adds new elements to array instance and change capacity of
// data field if needed.
func (a *Array) Add(values ...interface{}) {
	if len(values) < 1 {
		panic(fmt.Errorf("empty values container for insert: %v",
			values))
	}
	checkCap(a, a.size+len(values))
	if len(values) == 1 {
		a.AddSlice(values[0])
	} else {
		a.AddAbstractData(values...)
	}
}

// checkCap checks capability of array's data
func checkCap(array *Array, newSize int) {
	if newSize > cap(array.data) {
		array.IncreaseCap(&newSize)
	}
	return
}

// IncreaseCap increases cap of target array
func (a *Array) IncreaseCap(newSize *int) {
	newCap := int(2 * *newSize)
	newData := make([]interface{}, newCap, newCap)
	copy(newData, a.data)
	a.data = newData
}

// AddSlice adds element in array.data with types changes to max bits usage
func (a *Array) AddSlice(values interface{}) {
	valuesContainer := ref.ValueOf(values)
	for i := 0; i < valuesContainer.Len(); i++ {
		if !a.duplicates {
			value := valuesContainer.Index(i)
			if ok := a.checkDuplicates(value); !ok {
				continue
			}
			switch value.Kind() {
			case ref.Float32, ref.Float64:
				a.data[a.size] = value.Float()
				a.size += 1
			case ref.String:
				a.data[a.size] = value.String()
				a.size += 1
			case ref.Int, ref.Int8, ref.Int16, ref.Int32, ref.Int64:
				a.data[a.size] = value.Int()
				a.size += 1
			case ref.Uint, ref.Uint8, ref.Uint16, ref.Uint32, ref.Uint64:
				a.data[a.size] = value.Uint()
				a.size += 1
			}
		}
	}
}


func (a *Array) AddAbstractData(values ...interface{}) {
	for _, value := range values {
		if !a.duplicates {
			if ok := a.checkDuplicates(value); !ok {
				continue
			}
		}
		a.data[a.size] = value
		a.size += 1
	}
}

// checkDuplicates checks is value in a.data
func (a *Array) checkDuplicates(value interface{}) (ok bool) {
	if ref.ValueOf(value).Kind().String() == "slice" {
		vals := ref.ValueOf(value)
		for i := 0; i < len(a.data); i++ {
			if a.data[i] != nil && ref.TypeOf(a.data[i]).Kind().String() == "slice" && ref.ValueOf(a.data[i]).Len() == vals.Len() {
				maxLen := ref.ValueOf(a.data[i]).Len()
				for j := 0; j < maxLen; j++  {
					if vals.Index(j) == ref.ValueOf(a.data[i]).Index(j) {
						return false
						}
					}
				return true
				}
			}
		}
	for _, element := range a.data {
		if element != nil {
			if ref.TypeOf(element).Kind().String() == "slice" {
				continue
			}
			if value == element {
				return false
			}
		}
	}
	return true
}

// Print prints to array stdout
func (a *Array) Print(options PrintOptions) {
	if options == (PrintOptions{}) {
		for index, element := range a.data {
			if (index < defaultOffsetB || index > a.size-defaultOffsetE) && element != nil {
				fmt.Printf("%d: %v\n", index, element)
			}
		}
	} else {
		for index, element := range a.data {
			if (index < options.offsetBegin || index > a.size-options.offsetEnd-1) && element != nil {
				fmt.Printf("%d: %v\n", index, element)
			}
		}
	}
}
