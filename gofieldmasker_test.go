package go_ufm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type FullTestStruct struct {
	ID        string `db:"id"`
	FirstName string `db:"first_name" ufm:"firstName"`
	LastName  string `db:"last_name" ufm:"lastName"`
	Email     string `db:"email" ufm:"email"`
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

func TestIDPass(t *testing.T) {
	test := &FullTestStruct{
		ID:        "1",
		FirstName: "Foo",
		LastName:  "Bar",
		Email:     "foobar@email.com",
	}
	mask := []string{"id", "firstName"}
	update := GetFieldMaskerValues(test, mask)
	expected := map[string]any{
		"first_name": "Foo",
	}
	assert.Equal(t, expected, update)
}

func TestNoMask(t *testing.T) {
	test := &FullTestStruct{
		ID:        "1",
		FirstName: "Foo",
		LastName:  "Bar",
		Email:     "foobar@email.com",
	}
	mask := []string{}
	update := GetFieldMaskerValues(test, mask)
	expected := map[string]any{}
	assert.Equal(t, expected, update)
}
