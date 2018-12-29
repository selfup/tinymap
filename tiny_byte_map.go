package tinymap

import (
	"errors"
	"fmt"
)

// ByteTuple is a basic struct
//
//  ByteTuple{Key: 42, Val: []byte("9000")}
type ByteTuple struct {
	Key int
	Val []byte
}

// ByteMap stores ByteTuples
//
// It behaves like a HashMap!
//
//  byteMap := new(ByteMap)
//  byteMap.Set(42, []byte("9000"))
//  val, err := byteMap.Get(42)
//
//  if err != nil {
//    log.Fatal(err)
//  }
//
//  fmt.Print(val)
//
//  byteMap.Delete(42)
type ByteMap struct {
	data []ByteTuple
}

// Get fetches data.
func (t ByteMap) Get(key int) ([]byte, error) {
	for _, ByteTuple := range t.data {
		if ByteTuple.Key == key {
			return ByteTuple.Val, nil
		}
	}

	var errByte []byte
	errMsg := fmt.Sprintf("No such key (%d) - byte default returned", key)
	err := errors.New(errMsg)

	return errByte, err
}

// Set will update or add data.
// If ByteTuple.Key already exists, only the ByteTuple.Val is updated.
// Otherwise a new ByteTuple is inserted into the data slice.
func (t *ByteMap) Set(key int, val []byte) {
	for i, ByteTuple := range t.data {
		if ByteTuple.Key == key {
			ByteTuple.Val = val

			// assign element by index the updated ByteTuple
			t.data[i] = ByteTuple
		}
	}

	ByteTuple := ByteTuple{Key: key, Val: val}

	t.data = append(t.data, ByteTuple)
}

// Delete removes data.
// Returns true if deleted.
// Returns false if the key was not found.
func (t *ByteMap) Delete(key int) bool {
	for i, ByteTuple := range t.data {
		if ByteTuple.Key == key {
			t.data[i] = t.data[len(t.data)-1]
			t.data[len(t.data)-1] = ByteTuple
			t.data = t.data[:len(t.data)-1]

			return true
		}
	}

	return false
}
