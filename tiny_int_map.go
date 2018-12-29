package tinymap

import (
	"errors"
	"fmt"
)

// IntTuple is a basic struct
//
//  IntTuple{Key: 42, Val: 9000}
type IntTuple struct {
	Key int
	Val int
}

// IntMap stores IntTuples
//
// It behaves like a HashMap!
//
//  intMap := new(IntMap)
//  intMap.Set(42, 9000)
//  val, err := byteMap.Get(42)
//
//  if err != nil {
//    log.Fatal(err)
//  }
//
//  fmt.Print(val)
//
//  intMap.Delete(42)
type IntMap struct {
	data []IntTuple
}

// Get fetches data
func (t IntMap) Get(key int) (int, error) {
	for _, IntTuple := range t.data {
		if IntTuple.Key == key {
			return IntTuple.Val, nil
		}
	}

	var errVal int
	errMsg := fmt.Sprintf("No such key (%d) - int default returned", key)
	err := errors.New(errMsg)

	return errVal, err
}

// Set will update or add data.
// If IntTuple.Key already exists, only the IntTuple.Val is updated.
// Otherwise a new IntTuple is inserted into the data slice.
func (t *IntMap) Set(key int, val int) bool {
	for i, IntTuple := range t.data {
		if IntTuple.Key == key {
			IntTuple.Val = val

			// assign element by index the updated IntTuple
			t.data[i] = IntTuple

			return true
		}
	}

	IntTuple := IntTuple{Key: key, Val: val}

	t.data = append(t.data, IntTuple)

	return true
}

// Delete removes data.
// Returns true if deleted.
// Returns false if the key was not found.
func (t *IntMap) Delete(key int) bool {
	for i, IntTuple := range t.data {
		if IntTuple.Key == key {
			t.data[i] = t.data[len(t.data)-1]
			t.data[len(t.data)-1] = IntTuple
			t.data = t.data[:len(t.data)-1]

			return true
		}
	}

	return false
}
