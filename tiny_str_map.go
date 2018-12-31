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

// StrMap stores StrTuples
//
// It behaves like a HashMap!
//
//  strMap := new(StrMap)
//  strMap.Set("foo", "bar")
//  val, err := byteMap.Get("foo")
//
//  if err != nil {
//    log.Fatal(err)
//  }
//
//  fmt.Print(val)
//
//  strMap.Delete(42)
type StrMap struct {
	// Data is public but use carefully :)
	Data []StrTuple
}

// Get fetches Tuples by Key and returns their Val
func (t StrMap) Get(key string) (string, error) {
	for _, StrTuple := range t.Data {
		if StrTuple.Key == key {
			return StrTuple.Val, nil
		}
	}

	var errVal string
	errMsg := fmt.Sprintf("No such key (%s) - string default returned", key)
	err := errors.New(errMsg)

	return errVal, err
}

// Set will update or add Tuples.
// If StrTuple.Key already exists, only the StrTuple.Val is updated.
// Otherwise a new StrTuple is inserted into the Data slice.
func (t *StrMap) Set(key string, val string) {
	for i, StrTuple := range t.Data {
		if StrTuple.Key == key {
			StrTuple.Val = val

			t.Data[i] = StrTuple

			return
		}
	}

	StrTuple := StrTuple{Key: key, Val: val}

	t.Data = append(t.Data, StrTuple)
}

// Delete removes Tuples.
// Returns true if deleted.
// Returns false if the key was not found.
func (t *StrMap) Delete(key string) bool {
	for i, StrTuple := range t.Data {
		if StrTuple.Key == key {
			t.Data[i] = t.Data[len(t.Data)-1]
			t.Data[len(t.Data)-1] = StrTuple
			t.Data = t.Data[:len(t.Data)-1]

			return true
		}
	}

	return false
}
