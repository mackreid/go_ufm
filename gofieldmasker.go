package go_ufm

import (
	"reflect"
	"slices"
)

var (
	gfm_tag = "ufm"
	db_tag  = "db"
)

func GetFieldMaskerValues(object any, mask []string) map[string]any {
	updateMask := make(map[string]any)
	val := reflect.ValueOf(object)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}
	t := val.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		tag, ok := field.Tag.Lookup(gfm_tag)
		if !ok {
			continue
		}
		db := field.Tag.Get(db_tag)
		if slices.Contains(mask, tag) {
			updateMask[db] = val.Field(i).Interface()
		}

	}
	return updateMask
}
