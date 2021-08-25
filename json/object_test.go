package json_test

import (
	"testing"

	"github.com/fdschonborn/go-sandbox/json"
)

func TestObject_Get(t *testing.T) {
	o := json.Object[bool]{"result": true}
	t.Log(o)

	var result bool
	if err := o.Get("result", &result); err != nil {
		t.Fatal(err)
	}

	if !result {
		t.Fatal("Expected result to be true")
	}
	t.Log(result)
}
