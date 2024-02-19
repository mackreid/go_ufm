package gofieldmasker

import (
	"reflect"
	"slices"
	"strings"
)

var (
	GFM_TAG = "gmm"
)

func GetFieldMaskerValues(object any, mask []string) map[string]any {
	m := make(map[string]any)
	val := reflect.ValueOf(object)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	t := val.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag := field.Tag.Get(GFM_TAG)
		split := strings.Split(tag, ">")
		if slices.Contains(mask, split[0]) {
			iVal := val.Field(i)
			m[split[len(split)-1]] = iVal.Interface()
		}
	}
	return m
}
