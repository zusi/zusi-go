package gen

import (
	"github.com/stretchr/testify/assert"
	"github.com/zusi/zusi-go/tcp/message"
	"reflect"
	"testing"
)

type blaStruct struct {
	Field1    *string
	Field2    *int32      `zusi:"0001"`
	Field3    *int32      `zusi:"abc"`
	Sub1      *subStruct  `zusi:"0002"`
	SubSlice1 []subStruct `zusi:"0003"`
	Field4    []int32     `zusi:"0004"`
}

type testStruct struct {
	Field1 *int32     `zusi:"0001"`
	Sub1   *subStruct `zusi:"0002"`
}

type subStruct struct {
	Field1 *string `zusi:"0001"`
}

func TestForField(t *testing.T) {
	sub := blaStruct{}
	ty := reflect.TypeOf(sub)
	tv := reflect.ValueOf(sub)

	t.Run("field without tag", func(t *testing.T) {
		field, err := forField(ty, ty.Field(0), tv.Field(0))
		assert.NoError(t, err)
		assert.Nil(t, field)
	})

	t.Run("field with valid tag", func(t *testing.T) {
		field, err := forField(ty, ty.Field(1), tv.Field(1))
		assert.NoError(t, err)
		assert.NotEmpty(t, field)

		assert.Equal(t, "Field2", field.Name)
		assert.Equal(t, uint16(1), field.Id)
		assert.True(t, field.Ptr)
		assert.False(t, field.Slice)
		assert.Equal(t, "int32", field.FieldType)
	})

	t.Run("field with invalid tag should fail", func(t *testing.T) {
		field, err := forField(ty, ty.Field(2), tv.Field(2))
		assert.Error(t, err)
		assert.Nil(t, field)
	})

	t.Run("field with substruct", func(t *testing.T) {
		field, err := forField(ty, ty.Field(3), tv.Field(3))
		assert.NoError(t, err)
		assert.NotEmpty(t, field)

		assert.Equal(t, "Sub1", field.Name)
		assert.Equal(t, uint16(2), field.Id)
		assert.True(t, field.Ptr)
		assert.False(t, field.Slice)
		assert.Equal(t, "struct", field.FieldType)

		assert.Equal(t, "subStruct", field.StructType)
		assert.Equal(t, "github.com/zusi/zusi-go/tcp/gen", field.StructPkgPath)
	})

	t.Run("field with substruct slice", func(t *testing.T) {
		field, err := forField(ty, ty.Field(4), tv.Field(4))
		assert.NoError(t, err)
		assert.NotEmpty(t, field)

		assert.Equal(t, "SubSlice1", field.Name)
		assert.Equal(t, uint16(3), field.Id)
		assert.False(t, field.Ptr)
		assert.True(t, field.Slice)
		assert.Equal(t, "struct", field.FieldType)

		assert.Equal(t, "subStruct", field.StructType)
		assert.Equal(t, "github.com/zusi/zusi-go/tcp/gen", field.StructPkgPath)
	})

	t.Run("field with int32 slice", func(t *testing.T) {
		field, err := forField(ty, ty.Field(5), tv.Field(5))
		assert.NoError(t, err)
		assert.NotEmpty(t, field)

		assert.Equal(t, "Field4", field.Name)
		assert.Equal(t, uint16(4), field.Id)
		assert.False(t, field.Ptr)
		assert.True(t, field.Slice)
		assert.Equal(t, "int32", field.FieldType)

		assert.Equal(t, "", field.StructType)
		assert.Equal(t, "", field.StructPkgPath)
	})
}

func TestFindStructs(t *testing.T) {
	m := message.Message{}

	res := FindStructsToReflect(m)

	t.Log(res)
}

func TestReflectMessage(t *testing.T) {
	sub := testStruct{}

	res, err := Reflect(sub)

	assert.NoError(t, err)
	assert.Len(t, res.Fields, 2)
	assert.Equal(t, "testStruct", res.TypeName)
	assert.Equal(t, "github.com/zusi/zusi-go/tcp/gen", res.PkgPath)
}
