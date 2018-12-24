package tinymap

// IntTuple is a basic struct that holds a key and a value
type IntTuple struct {
	Key int
	Val int
}

// TinyIntMap is the str/str interface for TinyIntMap
type TinyIntMap struct {
	data []IntTuple
}

// Get fetches the IntTuple.Val from the TinyIntMap data
func (t TinyIntMap) Get(key int) (int, bool) {
	for _, IntTuple := range t.data {
		if IntTuple.Key == key {
			return IntTuple.Val, true
		}
	}

	return 0, false
}

// Set appends a new IntTuple to the data data
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
