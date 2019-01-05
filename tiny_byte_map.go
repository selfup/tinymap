package tinymap

import (
	"bytes"
	"errors"
	"fmt"
)

// ByteTuple is a basic struct
//
//  ByteTuple{Key: []byte("42"), Val: []byte("9000")}
type ByteTuple struct {
	Key []byte
	Val []byte
}

// ByteMap stores ByteTuples
//
// It behaves like a HashMap!
//
//  byteMap := new(ByteMap)
//  byteMap.Set([]byte("42"), []byte("9000"))
//  val, err := byteMap.Get([]byte("42"))
//
//  if err != nil {
//    log.Print(err)
//  }
//
//  fmt.Print(val)
//
//  byteMap.Delete([]byte("42"))
type ByteMap struct {
	// Data is public but use carefully :)
	Data []ByteTuple
}

// Get fetches Tuples by Key and returns their Val.
func (t ByteMap) Get(key []byte) ([]byte, error) {
	for _, ByteTuple := range t.Data {
		if bytes.Equal(ByteTuple.Key, key) {
			return ByteTuple.Val, nil
		}
	}

	var errByte []byte
	errMsg := fmt.Sprintf("No such key (%d) - byte default returned", key)
	err := errors.New(errMsg)

	return errByte, err
}

// Set will update or add Tuples.
// If ByteTuple.Key already exists, only the ByteTuple.Val is updated.
// Otherwise a new ByteTuple is inserted into the Data slice.
func (t *ByteMap) Set(key []byte, val []byte) {
	for i, ByteTuple := range t.Data {
		if bytes.Equal(ByteTuple.Key, key) {
			ByteTuple.Val = val

			t.Data[i] = ByteTuple

			return
		}
	}

	ByteTuple := ByteTuple{Key: key, Val: val}

	t.Data = append(t.Data, ByteTuple)
}

// Delete removes Tuples.
// Returns true if deleted.
// Returns false if the key was not found.
func (t *ByteMap) Delete(key []byte) bool {
	for i, ByteTuple := range t.Data {
		if bytes.Equal(ByteTuple.Key, key) {
			t.Data[i] = t.Data[len(t.Data)-1]
			t.Data[len(t.Data)-1] = ByteTuple
			t.Data = t.Data[:len(t.Data)-1]

			return true
		}
	}

	return false
}
