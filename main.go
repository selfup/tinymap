package main

import "fmt"

// StrTuple is a basic struct that holds a key and a value
type StrTuple struct {
	Key string
	Val string
}

// TinyStrMap is the str/str interface for TinyStrMap
type TinyStrMap struct {
	Slice []StrTuple
}

// Get fetches the StrTuple.Val from the TinyStrMap Slice
func (t TinyStrMap) Get(key string) (string, bool) {
	for _, StrTuple := range t.Slice {
		if StrTuple.Key == key {
			return StrTuple.Val, true
		}
	}

	return "", false
}

// Set appends a new StrTuple to the Slice Slice
func (t *TinyStrMap) Set(key string, val string) bool {
	for i, StrTuple := range t.Slice {
		if StrTuple.Key == key {
			StrTuple.Val = val

			// assign element by index the updated StrTuple
			t.Slice[i] = StrTuple

			return true
		}
	}

	StrTuple := StrTuple{Key: key, Val: val}

	t.Slice = append(t.Slice, StrTuple)

	return true
}

// Delete removes the StrTuple from t.Slice
func (t *TinyStrMap) Delete(key string) bool {
	for i, StrTuple := range t.Slice {
		if StrTuple.Key == key {
			// deleting from a Slice is always fun
			t.Slice[i] = t.Slice[len(t.Slice)-1]
			t.Slice[len(t.Slice)-1] = StrTuple
			t.Slice = t.Slice[:len(t.Slice)-1]

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
