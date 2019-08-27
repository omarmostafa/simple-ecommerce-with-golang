package Transformers

import (
	"reflect"
	"simple-ecommerce/app/infrastructures"
	"strings"
)

func TransformItem(transformer interface{},data interface{}) {
	typeOfTransformer := reflect.ValueOf(transformer).Elem().Type()
	for i := 0; i < typeOfTransformer.NumField(); i++ {
		fieldName := typeOfTransformer.Field(i)
		funcName :="Set"+strings.Title(fieldName.Name)
		infrastructures.CallMethodOfInterfaceByReflect(transformer,funcName,data)
	}
}

