package util

import (
	"reflect"
	"strconv"
	"strings"
)

func StringConversionInt8(data string) int8 {
	v, err := strconv.ParseInt(data, 10, 8)
	if err != nil {
		panic(err)
	}
	return int8(v)
}

func StringConversionInt16(data string) int16 {
	v, err := strconv.ParseInt(data, 10, 16)
	if err != nil {
		panic(err)
	}
	return int16(v)
}

func StringConversionInt(data string) int {
	v, err := strconv.ParseInt(data, 10, 32)
	if err != nil {
		panic(err)
	}
	return int(v)
}

func StringConversionInt32(data string) int32 {
	v, err := strconv.ParseInt(data, 10, 32)
	if err != nil {
		panic(err)
	}
	return int32(v)
}

func StringConversionInt64(data string) int64 {
	v, err := strconv.ParseInt(data, 10, 64)
	if err != nil {
		panic(err)
	}
	return v
}

func StringConversionUint8(data string) uint8 {
	v, err := strconv.ParseUint(data, 10, 8)
	if err != nil {
		panic(err)
	}
	return uint8(v)
}

func StringConversionUint16(data string) uint16 {
	v, err := strconv.ParseUint(data, 10, 16)
	if err != nil {
		panic(err)
	}
	return uint16(v)
}

func StringConversionUint(data string) uint {
	v, err := strconv.ParseUint(data, 10, 32)
	if err != nil {
		panic(err)
	}
	return uint(v)
}

func StringConversionUint32(data string) uint32 {
	v, err := strconv.ParseUint(data, 10, 32)
	if err != nil {
		panic(err)
	}
	return uint32(v)
}

func StringConversionUint64(data string) uint64 {
	v, err := strconv.ParseUint(data, 10, 64)
	if err != nil {
		panic(err)
	}
	return v
}

func StringConversionFloat32(data string) float32 {
	v, err := strconv.ParseFloat(data, 32)
	if err != nil {
		panic(err)
	}
	return float32(v)
}

func StringConversionFloat64(data string) float64 {
	v, err := strconv.ParseFloat(data, 64)
	if err != nil {
		panic(err)
	}
	return v
}

func StringConversionType(reflectType reflect.Type, data string) reflect.Value {
	var v interface{}
	switch reflectType.String() {
	case "uint8":
		v = StringConversionUint8(data)
	case "uint16":
		v = StringConversionUint16(data)
	case "uint":
		v = StringConversionUint(data)
	case "uint32":
		v = StringConversionUint32(data)
	case "uint64":
		v = StringConversionUint64(data)
	case "int8":
		v = StringConversionInt8(data)
	case "int16":
		v = StringConversionInt16(data)
	case "int":
		v = StringConversionInt(data)
	case "int32":
		v = StringConversionInt32(data)
	case "int64":
		v = StringConversionInt64(data)
	case "float32":
		v = StringConversionFloat32(data)
	case "float64":
		v = StringConversionFloat64(data)
	case "string":
		v = data
	}
	return reflect.ValueOf(v)
}

func UnmarshalReq(values map[string][]string,pkt interface{}) interface{} {
	k := reflect.TypeOf(pkt)
	v := reflect.ValueOf(pkt).Elem()
	if k.Kind() == reflect.Ptr {
		k = k.Elem()
	}
	for i := 0; i < k.NumField(); i++ {
		value,ok := values[strings.ToLower(k.Field(i).Name)]
		if ok && len(value) > 0{
			v.FieldByName(k.Field(i).Name).Set(StringConversionType(v.FieldByName(k.Field(i).Name).Type(),value[0]))
		}
	}
	return pkt
}