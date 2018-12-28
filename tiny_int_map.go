package tinymap

import (
	"errors"
	"fmt"
)

// IntTuple is a basic struct that holds an string Key and a string Val
//
// It is exposed in case a user just needs a simple tuple :)
//
//  IntTuple{Key: "foo", Val: "bar"}
type IntTuple struct {
	Key int
	Val int
}

// TinyIntMap stores StrTuples
//
// It behaves like a HashMap!
//
//  tinyIntMap := new(TinyIntMap)
//  tinyIntMap.Set(42, 9000)
//  val, err := tinyByteMap.Get(42)
//
//  if err != nil {
//    log.Fatal(err)
//  }
//
//  fmt.Print(val)
//
//  tinyIntMap.Delete(42)
type TinyIntMap struct {
	data []IntTuple
}

// Get fetches the IntTuple.Val when given an int key
func (t TinyIntMap) Get(key int) (int, error) {
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

// Set will update or add data based on existence of an int key
//
// If IntTuple.Key already exists, only the IntTuple.Val is updated
//
// Otherwise a new IntTuple is inserted into the data slice
func (t *TinyIntMap) Set(key int, val int) bool {
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

// Delete removes the IntTuple from t.data
//
// Returns true if deleted
//
// Returns false if the key was not found
func (t *TinyIntMap) Delete(key int) bool {
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
