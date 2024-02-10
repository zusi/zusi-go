package tcp

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"reflect"

	"github.com/zusi/zusi-go/tcp/message"
)

func MarshalMessage(message message.Message) ([]byte, error) {
	return Marshal(message)
}

func Marshal(v interface{}) ([]byte, error) {
	var bytes []byte

	t := reflect.TypeOf(v)
	r := reflect.ValueOf(v)

	for i := 0; i < r.NumField(); i++ {
		f := r.Field(i)

		val, ok := t.Field(i).Tag.Lookup("zusi")
		if !ok {
			continue
		}

		bts, err := hex.DecodeString(val)
		if err != nil {
			return nil, fmt.Errorf("unable to convert zusi annotation to valid element ID: %w", err)
		}

		elId := binary.BigEndian.Uint16(bts)

		fieldBytes, err := encapsulate(elId, f)
		if err != nil {
			return nil, fmt.Errorf("error encapsulating attribute %v: %w", t.Field(i).Name, err)
		}

		bytes = append(bytes, fieldBytes...)
	}

	return bytes, nil
}

func encapsulate(elementId uint16, value reflect.Value) ([]byte, error) {
	if value.Kind() == reflect.Ptr {
		if value.IsNil() {
			return []byte{}, nil
		}

		value = value.Elem()
	}

	var bytes []byte //:= make([]byte, 6)

	switch value.Kind() {
	case reflect.Struct:
		bytes = make([]byte, 6)
		binary.LittleEndian.PutUint16(bytes[4:6], elementId)
		marshal, err := Marshal(value.Interface())

		bytes = append(bytes, marshal...)
		bytes = append(bytes, 0xFF, 0xFF, 0xFF, 0xFF)

		return bytes, err // can't use length calculation because length should be 0x0000 for structs
	case reflect.Int8:
		bytes = make([]byte, 7)
		bytes[6] = byte(value.Interface().(int8))
	case reflect.Uint8: // byte
		bytes = make([]byte, 7)
		bytes[6] = value.Interface().(byte)
	case reflect.Int16:
		bytes = make([]byte, 8)
		binary.LittleEndian.PutUint16(bytes[6:8], uint16(value.Interface().(int16)))
	case reflect.Uint16: // word
		bytes = make([]byte, 8)
		binary.LittleEndian.PutUint16(bytes[6:8], value.Interface().(uint16))
	case reflect.Int32:
		bytes = make([]byte, 10)
		binary.LittleEndian.PutUint32(bytes[6:10], uint32(value.Interface().(int32)))
	case reflect.Uint32:
		bytes = make([]byte, 10)
		binary.LittleEndian.PutUint32(bytes[6:10], value.Interface().(uint32))
	case reflect.Int64:
		bytes = make([]byte, 14)
		binary.LittleEndian.PutUint64(bytes[6:14], uint64(value.Interface().(int64)))
	case reflect.Uint64:
		bytes = make([]byte, 14)
		binary.LittleEndian.PutUint64(bytes[6:14], value.Interface().(uint64))
	case reflect.Float32: // single
		bytes = make([]byte, 10)
		asUint := math.Float32bits(value.Interface().(float32))
		binary.LittleEndian.PutUint32(bytes[6:10], asUint)
	case reflect.Float64: // double
		bytes = make([]byte, 14)
		asUint := math.Float64bits(value.Interface().(float64))
		binary.LittleEndian.PutUint64(bytes[6:14], asUint)
	case reflect.String:
		str := value.Interface().(string)
		bytes = make([]byte, len(str)+6)
		copy(bytes[6:], str)
	case reflect.Slice:
		sliceType := value.Type().Elem()
		if sliceType.Kind() == reflect.Ptr {
			sliceType = sliceType.Elem()
		}
		switch sliceType.Kind() {
		case reflect.Uint16:
			asSlice := value.Interface().([]uint16)
			var bts []byte
			for _, val := range asSlice {
				tile, err := encapsulate(elementId, reflect.ValueOf(val))
				if err != nil {
					return nil, fmt.Errorf("error encapsulating part of slice: %w", err)
				}

				bts = append(bts, tile...)
			}

			return bts, nil
		case reflect.Struct:
			var bts []byte

			for i := 0; i < value.Len(); i++ {
				node, err := encapsulate(elementId, value.Index(i))
				if err != nil {
					return nil, fmt.Errorf("error encapsulating part of slice: %w", err)
				}

				bts = append(bts, node...)
			}

			return bts, nil
		default:
			return nil, errors.New("unhandled slice type")
		}
	default:
		return nil, fmt.Errorf("could not encapsulate attribute type %v into bytes", value.Type())
	}

	binary.LittleEndian.PutUint32(bytes[0:4], uint32(len(bytes)-4))
	binary.LittleEndian.PutUint16(bytes[4:6], elementId)

	return bytes, nil
}
