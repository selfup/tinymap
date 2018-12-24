package tinymap

// ByteTuple is a basic struct that holds a key and a value
type ByteTuple struct {
	Key int
	Val []byte
}

// TinyByteMap is the str/str interface for TinyByteMap
type TinyByteMap struct {
	data []ByteTuple
}

// Get fetches the ByteTuple.Val from the TinyByteMap data
func (t TinyByteMap) Get(key int) ([]byte, bool) {
	for _, ByteTuple := range t.data {
		if ByteTuple.Key == key {
			return ByteTuple.Val, true
		}
	}

	return []byte("0"), false
}

// Set appends a new ByteTuple to the data data
func (t *TinyByteMap) Set(key int, val []byte) bool {
	for i, ByteTuple := range t.data {
		if ByteTuple.Key == key {
			ByteTuple.Val = val

			// assign element by index the updated ByteTuple
			t.data[i] = ByteTuple

			return true
		}
	}

	ByteTuple := ByteTuple{Key: key, Val: val}

	t.data = append(t.data, ByteTuple)

	return true
}

// Delete removes the ByteTuple from t.data
func (t *TinyByteMap) Delete(key int) bool {
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
