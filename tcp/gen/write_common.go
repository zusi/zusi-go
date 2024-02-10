package gen

import (
	"fmt"
	"reflect"
	"strings"

	. "github.com/dave/jennifer/jen"
)

type types struct {
	int8
	int16
	int32
	int64
	uint8
	uint16
	uint32
	uint64
	float32
	float64
	string
}

func generateCommon() string {
	f := NewFilePath(CommonPath)

	typ := types{}
	ty := reflect.TypeOf(typ)

	for i := 0; i < ty.NumField(); i++ {
		readBaseType(f, ty.Field(i).Type)
		f.Empty()
	}

	return fmt.Sprintf("%#v", f)
}

func readBaseType(f *File, t reflect.Type) {
	f.Func().Id("Read"+strings.Title(t.Name())).Params(Id("data").Qual("io", "Reader"), Id("length").Uint32()).Params(Op("*").Id(t.Name()), Error()).Block(
		Id("bts").Op(":=").Make(Index().Byte(), Id("length")),
		Empty(),
		List(Id("n"), Err()).Op(":=").Qual("io", "ReadFull").Call(Id("data"), Id("bts")),
		If(Uint32().Parens(Id("n")).Op("!=").Id("length")).Block(
			Return(Nil(), Qual(PkgErrs, "New").Call(Lit("invalid payload length received"))),
		),
		Empty(),
		Empty(),
		Return(Op("&").Id("res"), Err()),
	)
}
