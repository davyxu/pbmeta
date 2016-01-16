package pbmeta

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
)

var (
	errFieldNotFound = errors.New("field not found")
	errFieldTagError = errors.New("field tag error")
)

func getFieldNumber(src interface{}, name string) (int32, error) {

	tdef := reflect.TypeOf(src).Elem()

	structFieldInfo, found := tdef.FieldByName(name)
	if !found {
		return 0, errFieldNotFound
	}

	paramList := strings.Split(structFieldInfo.Tag.Get("protobuf"), ",")

	if len(paramList) <= 1 {
		return 0, errFieldTagError
	}

	fieldNum, err := strconv.Atoi(paramList[1])
	if err != nil {
		return 0, err
	}

	return int32(fieldNum), nil

}
