package chap2_2

import (
	"reflect"
)

type Result struct {
	integer         []any
	positiveInteger []any
}

func P211(input []any) Result {
	var integer []any
	var positiveInteger []any
	for _, v := range input {
		switch reflect.TypeOf(v).Kind() {
		case reflect.Int:
			integer = append(integer, v)
			if v.(int) > 0 {
				positiveInteger = append(positiveInteger, v)
			}
		}

	}
	return Result{
		integer:         integer,
		positiveInteger: positiveInteger,
	}
}
