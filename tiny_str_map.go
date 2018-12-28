package tinymap

import (
	"errors"
	"fmt"
)

// StrTuple is a basic struct
//
//  StrTuple{Key: "foo", Val: "bar"}
type StrTuple struct {
	Key string
	Val string
}

// TinyStrMap stores StrTuples
//
// It behaves like a HashMap!
//
//  tinyStrMap := new(TinyStrMap)
//  tinyStrMap.Set("foo", "bar")
//  val, err := tinyByteMap.Get("foo")
//
//  if err != nil {
//    log.Fatal(err)
//  }
//
//  fmt.Print(val)
//
//  tinyStrMap.Delete(42)
type TinyStrMap struct {
	data []StrTuple
}

// Get fetches data
func (t TinyStrMap) Get(key string) (string, error) {
	for _, StrTuple := range t.data {
		if StrTuple.Key == key {
			return StrTuple.Val, nil
		}
	}

	var errVal string
	errMsg := fmt.Sprintf("No such key (%s) - string default returned", key)
	err := errors.New(errMsg)

	return errVal, err
}

// Set will update or add data.
// If StrTuple.Key already exists, only the StrTuple.Val is updated.
// Otherwise a new StrTuple is inserted into the data slice.
func (t *TinyStrMap) Set(key string, val string) bool {
	for i, StrTuple := range t.data {
		if StrTuple.Key == key {
			StrTuple.Val = val

			// assign element by index the updated StrTuple
			t.data[i] = StrTuple

			return true
		}
	}

	StrTuple := StrTuple{Key: key, Val: val}

	t.data = append(t.data, StrTuple)

	return true
}

// Delete removes data.
// Returns true if deleted.
// Returns false if the key was not found.
func (t *TinyStrMap) Delete(key string) bool {
	for i, StrTuple := range t.data {
		if StrTuple.Key == key {
			t.data[i] = t.data[len(t.data)-1]
			t.data[len(t.data)-1] = StrTuple
			t.data = t.data[:len(t.data)-1]

			return true
		}
	}

	return false
}
