package tinymap

import (
	"errors"
	"fmt"
)

// IntStrTuple is a basic struct
//
//  IntStrTuple{Key: 42, Val: "9000"}
type IntStrTuple struct {
	Key int
	Val string
}

// IntStrMap stores IntStrTuples
//
// It behaves like a HashMap!
//
//  intStrMap := new(IntStrMap)
//  intStrMap.Set(42, "9000")
//  val, err := intStrMap.Get(42)
//
//  if err != nil {
//    log.Fatal(err)
//  }
//
//  fmt.Print(val)
//
//  intStrMap.Delete(42)
type IntStrMap struct {
	// Data is public but use carefully :)
	Data []IntStrTuple
}

// Get fetches Tuples by Key and returns their Val
func (t IntStrMap) Get(key int) (string, error) {
	for _, IntStrTuple := range t.Data {
		if IntStrTuple.Key == key {
			return IntStrTuple.Val, nil
		}
	}

	var errVal string
	errMsg := fmt.Sprintf("No such key (%d) - string default returned", key)
	err := errors.New(errMsg)

	return errVal, err
}

// Set will update or add Tuples.
// If IntStrTuple.Key already exists, only the IntStrTuple.Val is updated.
// Otherwise a new IntStrTuple is inserted into the Data slice.
func (t *IntStrMap) Set(key int, val string) {
	for i, intStrTuple := range t.Data {
		if intStrTuple.Key == key {
			intStrTuple.Val = val

			t.Data[i] = intStrTuple

			return
		}
	}

	intStrTuple := IntStrTuple{Key: key, Val: val}

	t.Data = append(t.Data, intStrTuple)
}

// Delete removes Tuples.
// Returns true if deleted.
// Returns false if the key was not found.
func (t *IntStrMap) Delete(key int) bool {
	for i, IntStrTuple := range t.Data {
		if IntStrTuple.Key == key {
			t.Data[i] = t.Data[len(t.Data)-1]
			t.Data[len(t.Data)-1] = IntStrTuple
			t.Data = t.Data[:len(t.Data)-1]

			return true
		}
	}

	return false
}
