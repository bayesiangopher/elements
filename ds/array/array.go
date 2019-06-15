// Package array implements simple arrays structures:
// 1) array with duplicates
// 2) array with no duplicates

// O(x):
// Operation		No duplicates		Duplicates
// Search			N/2					N
// Insertion		N, 1				1, 1
// Deletion			N/2, N/2			N, N/2

package array

import (
	"fmt"
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

// Create creates new array structure filled by values
func Create(duplicates bool, values ...interface{}) *Array {
	array := &Array{duplicates: duplicates}
	if len(values) > 0 {
		array.Add(values...)
	}
	return array
}

// Add adds new elements to array instance and change capacity of
// data field if needed.
func (a *Array) Add(values ...interface{}) {
	if len(values) < 1 {
		panic(fmt.Errorf("empty values container for insert: %v",
			values))
	}
	if newSize := a.size + len(values); newSize > cap(a.data) {
		newCap := int(2 * newSize)
		newData := make([]interface{}, newCap, newCap)
		copy(newData, a.data)
		a.data = newData
	}
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
	for _, element := range a.data {
		if value == element {
			return false
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
