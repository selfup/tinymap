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
//  val, err := intMap.Get(42)
//
//  if err != nil {
//    log.Fatal(err)
//  }
//
//  fmt.Print(val)
//
//  intMap.Delete(42)
type IntMap struct {
	// Data is public but use carefully :)
	Data []IntTuple
}

// Get fetches Tuples by Key and returns their Val
func (t IntMap) Get(key int) (int, error) {
	for _, IntTuple := range t.Data {
		if IntTuple.Key == key {
			return IntTuple.Val, nil
		}
	}

	var errVal int
	errMsg := fmt.Sprintf("No such key (%d) - int default returned", key)
	err := errors.New(errMsg)

	return errVal, err
}

// Set will update or add Tuples.
// If IntTuple.Key already exists, only the IntTuple.Val is updated.
// Otherwise a new IntTuple is inserted into the Data slice.
func (t *IntMap) Set(key int, val int) {
	for i, IntTuple := range t.Data {
		if IntTuple.Key == key {
			IntTuple.Val = val

			t.Data[i] = IntTuple

			return
		}
	}

	intTuple := IntTuple{Key: key, Val: val}

	t.Data = append(t.Data, intTuple)
}

// Delete removes Tuples.
// Returns true if deleted.
// Returns false if the key was not found.
func (t *IntMap) Delete(key int) bool {
	for i, IntTuple := range t.Data {
		if IntTuple.Key == key {
			t.Data[i] = t.Data[len(t.Data)-1]
			t.Data[len(t.Data)-1] = IntTuple
			t.Data = t.Data[:len(t.Data)-1]

			return true
		}
	}

	return false
}
