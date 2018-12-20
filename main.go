package main

import "fmt"

// StrTuple is a basic struct that holds a key and a value
type StrTuple struct {
	Key string
	Val string
}

// TinyStrMap is the str/str interface for TinyStrMap
type TinyStrMap struct {
	data []StrTuple
}

// Get fetches the StrTuple.Val from the TinyStrMap data
func (t TinyStrMap) Get(key string) (string, bool) {
	for _, StrTuple := range t.data {
		if StrTuple.Key == key {
			return StrTuple.Val, true
		}
	}

	return "", false
}

// Set appends a new StrTuple to the data data
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

// Delete removes the StrTuple from t.data
func (t *TinyStrMap) Delete(key string) bool {
	for i, StrTuple := range t.data {
		if StrTuple.Key == key {
			// deleting from a data is always fun
			t.data[i] = t.data[len(t.data)-1]
			t.data[len(t.data)-1] = StrTuple
			t.data = t.data[:len(t.data)-1]

			return true
		}
	}

	return false
}

func main() {
	t := new(TinyStrMap)

	t.Set("wow", "omg")
	a, b := t.Get("wow")
	fmt.Println(a, b)

	t.Set("wow", "lol")
	aa, bb := t.Get("wow")
	fmt.Println(aa, bb)

	t.Delete("wow")
	_, bbb := t.Get("wow")
	fmt.Println(bbb)
}
