package json

import (
	"encoding/json"
	"io"
)

func Parse(data []byte) (value interface{}, err error) {
	err = json.Unmarshal(data, &value)
	return value, err
}

func ParseReader(r io.Reader) (value interface{}, err error) {
	err = json.NewDecoder(r).Decode(&value)
	return value, err
}

func ParseObject(data []byte) (value map[string]interface{}, err error) {
	err = json.Unmarshal(data, &value)
	return value, err
}

func ParseObjectReader(r io.Reader) (value map[string]interface{}, err error) {
	err = json.NewDecoder(r).Decode(&value)
	return value, err
}

func ParseArray(data []byte) (value []map[string]interface{}, err error) {
	err = json.Unmarshal(data, &value)
	return value, err
}

func ParseArrayReader(r io.Reader) (value []map[string]interface{}, err error) {
	err = json.NewDecoder(r).Decode(&value)
	return value, err
}

func StringMarshal(v interface{}) (s string, err error) {
	b, err := json.Marshal(v)
	return string(b), err
}

func StringUnmarshal(s string) (i interface{}, err error) {
	err = json.Unmarshal([]byte(s), &i)
	return i, err
}
