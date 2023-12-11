package utils

import (
	"fmt"
	"reflect"
	"strconv"
)

type ArrayUtils struct{}

func (o ArrayUtils) ToSliceString(sliceValue reflect.Value) (*[]string, error) {
	//+ Si no es un slice retornamos error
	if sliceValue.Elem().Kind() != reflect.Slice {
		return nil, fmt.Errorf("sliceValue no es un slice")
	}

	//+ Si el slice está vació retornamos un slice vació
	if sliceValue.Elem().Len() == 0 {
		return &[]string{}, nil
	}

	//+ Comprobamos que el slice contenga valores enteros
	elemKind := sliceValue.Elem().Index(0).Kind()
	if !(elemKind >= reflect.Int && elemKind <= reflect.Uint64) {
		return sliceValue.Interface().(*[]string), nil
	}

	//+ Comprobamos si es un entero unsigned o no
	var unsigned bool
	if elemKind >= reflect.Uint && elemKind <= reflect.Uint64 {
		unsigned = true
	}

	//+ Convertimos el slice a un slice de int64
	var stringSlice []string = make([]string, sliceValue.Elem().Len())
	for i := 0; i < sliceValue.Elem().Len(); i++ {
		if unsigned {
			stringSlice[i] = strconv.FormatUint(sliceValue.Elem().Index(i).Uint(), 10)
		} else {
			stringSlice[i] = strconv.FormatInt(sliceValue.Elem().Index(i).Int(), 10)
		}
	}
	return &stringSlice, nil
}
