package gen

import (
	"encoding/binary"
	"encoding/hex"
	"errors"
	"log/slog"
	"path"
	"reflect"
)

type Message struct {
	TypeName string
	PkgPath  string
	Fields   []Field
}

type Field struct {
	Name          string
	Id            uint16
	IdX           string
	FieldType     string
	StructType    string
	RefStructType reflect.Type
	StructPkgPath string
	Ptr           bool
	Slice         bool
	Str           bool
}

func FindStructsToReflect(i interface{}) map[string]reflect.Type {
	ty := reflect.TypeOf(i)
	tv := reflect.ValueOf(i)

	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
		tv = reflect.Indirect(reflect.ValueOf(i))
	}

	var structs = map[string]reflect.Type{}

	structs[path.Join(ty.PkgPath(), ty.Name())] = reflect.PtrTo(ty)

	for i := 0; i < ty.NumField(); i++ {
		field := ty.Field(i)
		vField := tv.Field(i)

		if (field.Type.Kind() == reflect.Ptr || field.Type.Kind() == reflect.Slice) && vField.Type().Elem().Kind() == reflect.Struct {
			structType := vField.Type().Elem().Name()
			structPkgPath := vField.Type().Elem().PkgPath()

			structs[path.Join(structPkgPath, structType)] = vField.Type()

			v := reflect.New(vField.Type().Elem())
			res := FindStructsToReflect(v.Interface())

			for k, v := range res {
				structs[k] = v
			}
		}
	}

	return structs
}

func Reflect(i interface{}) (*Message, error) {
	ty := reflect.TypeOf(i)
	tv := reflect.ValueOf(i)

	if ty.Kind() == reflect.Ptr {
		ty = ty.Elem()
		tv = reflect.Indirect(reflect.ValueOf(i))
	}

	msgType := Message{
		TypeName: ty.Name(),
		PkgPath:  ty.PkgPath(),
		Fields:   []Field{},
	}

	for i := 0; i < ty.NumField(); i++ {
		field, err := forField(ty, ty.Field(i), tv.Field(i))
		if err != nil {
			return nil, err
		}

		msgType.Fields = append(msgType.Fields, *field)
	}

	slog.With("type", ty).
		Info("Sucessfully reflected")

	return &msgType, nil
}

func forField(parent reflect.Type, field reflect.StructField, value reflect.Value) (*Field, error) {
	ctxLogger := slog.With("parent", parent.Name(), "field", field.Name)

	tag, ok := field.Tag.Lookup("zusi")
	if !ok {
		ctxLogger.Warn("Field has no Zusi Tag")
		return nil, nil
	}

	ctxLogger = ctxLogger.With("tag", tag)

	bts, err := hex.DecodeString(tag)
	if err != nil {
		ctxLogger.
			Error("Couldn't decode Zusi Tag")

		return nil, errors.New("invalid Zusi Tag")
	}

	elementId := binary.BigEndian.Uint16(bts)

	fld := &Field{Name: field.Name, Id: elementId, IdX: tag}

	if field.Type.Kind() == reflect.Ptr {
		fld.Ptr = true
		fld.FieldType = field.Type.Elem().Kind().String()
	} else {
		fld.FieldType = field.Type.Kind().String()
	}

	if field.Type.Kind() == reflect.Slice {
		fld.Slice = true
		fld.FieldType = value.Type().Elem().Kind().String()
	}

	if (field.Type.Kind() == reflect.Ptr || field.Type.Kind() == reflect.Slice) && value.Type().Elem().Kind() == reflect.Struct {
		fld.Str = true
		fld.RefStructType = value.Type()
		fld.StructType = value.Type().Elem().Name()
		fld.StructPkgPath = value.Type().Elem().PkgPath()
	}

	return fld, nil
}
