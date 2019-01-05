package tinymap

import (
	"errors"
	"fmt"
)

// StrByteTuple is a basic struct
//
//  StrByteTuple{Key: "foo", Val: "bar"}
type StrByteTuple struct {
	Key string
	Val []byte
}

// StrByteMap stores StrByteTuples
//
// It behaves like a HashMap!
//
//  strByteMap := new(StrByteMap)
//  strByteMap.Set("foo", []byte("bar"))
//  val, err := strByteMap.Get("foo")
//
//  if err != nil {
//    log.Print(err)
//  }
//
//  fmt.Print(val)
//
//  strByteMap.Delete(42)
type StrByteMap struct {
	// Data is public but use carefully :)
	Data []StrByteTuple
}

// Get fetches Tuples by Key and returns their Val
func (t StrByteMap) Get(key string) ([]byte, error) {
	for _, StrByteTuple := range t.Data {
		if StrByteTuple.Key == key {
			return StrByteTuple.Val, nil
		}
	}

	var errVal []byte
	errMsg := fmt.Sprintf("No such key (%v) - []byte default returned", key)
	err := errors.New(errMsg)

	return errVal, err
}

// Set will update or add Tuples.
// If StrByteTuple.Key already exists, only the StrByteTuple.Val is updated.
// Otherwise a new StrByteTuple is inserted into the Data slice.
func (t *StrByteMap) Set(key string, val []byte) {
	for i, StrByteTuple := range t.Data {
		if StrByteTuple.Key == key {
			StrByteTuple.Val = val

			t.Data[i] = StrByteTuple

			return
		}
	}

	StrByteTuple := StrByteTuple{Key: key, Val: val}

	t.Data = append(t.Data, StrByteTuple)
}

// Delete removes Tuples.
// Returns true if deleted.
// Returns false if the key was not found.
func (t *StrByteMap) Delete(key string) bool {
	for i, StrByteTuple := range t.Data {
		if StrByteTuple.Key == key {
			t.Data[i] = t.Data[len(t.Data)-1]
			t.Data[len(t.Data)-1] = StrByteTuple
			t.Data = t.Data[:len(t.Data)-1]

			return true
		}
	}

	return false
}
