package gofieldmasker

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestStruct struct {
	Name string `gmm:"name>first_name"`
	Age  int    `gmm:"age"`
}

func TestMain(t *testing.T) {
	s := &TestStruct{
		Name: "Foo",
		Age:  28,
	}
	mask := []string{"name", "age"}
	updateMask := GetFieldMaskerValues(s, mask)
	expected := map[string]any{
		"first_name": "Foo",
		"age":        28,
	}
	assert.Equal(t, expected, updateMask)
}

type FullTestStruct struct {
	ID        string
	FirstName string `gmm:"firstName>first_name"`
	LastName  string `gmm:"lastName>last_name"`
	Email     string `gmm:"email"`
}

func TestFullStruct(t *testing.T) {
	test := &FullTestStruct{
		ID:        "1",
		FirstName: "Foo",
		LastName:  "Bar",
		Email:     "foobar@email.com",
	}
	mask := []string{"firstName", "lastName", "email"}
	updateM := GetFieldMaskerValues(test, mask)
	expected := map[string]any{
		"first_name": "Foo",
		"last_name":  "Bar",
		"email":      "foobar@email.com",
	}
	assert.Equal(t, expected, updateM)
}
