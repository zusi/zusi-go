package gen

import (
	"os"
	"reflect"
	"testing"
)

func TestWrite(t *testing.T) {
	msg := Message{
		TypeName: "Message",
		PkgPath:  "github.com/ebd/test",
		Fields: []Field{
			{Name: "Field1", Id: 0x000A, FieldType: reflect.Uint16.String(), IdX: "000A"},
			{Name: "Field2", Id: 0x0010, FieldType: reflect.Uint16.String(), IdX: "0010"},
			{Name: "Field2", Id: 0x0010, FieldType: reflect.Uint16.String(), IdX: "0010", Slice: true},
			{Name: "Field2", Id: 0x0010, FieldType: reflect.Ptr.String(), IdX: "0010", StructPkgPath: PkgPath, StructType: "ZusiMessage", Str: true},
		},
	}

	Generate(msg, os.Stdout)
}
