package tcp

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/zusi/zusi-go/tcp/message"
)

func TestMarshal(t *testing.T) {
	hm := helloTestMessage{
		ProtokollVersion: 2,
		ClientTyp:        2,
		Name:             "Fahrpult",
		Version:          "2.0",
	}

	res, err := Marshal(hm)

	assert.NoError(t, err)
	assert.Equal(t, helloTestMessageBytes, res)
}

func TestEncapsulateUint(t *testing.T) {
	res, err := encapsulate(0x0002, reflect.ValueOf(uint16(5)))
	assert.NoError(t, err)

	expected := []byte{
		0x04, 0x00, 0x00, 0x00,
		0x02, 0x00,
		0x05, 0x00,
	}

	assert.EqualValues(t, expected, res)
}

func TestEncapsulateString(t *testing.T) {
	res, err := encapsulate(0x0003, reflect.ValueOf("Fahrpult"))
	assert.NoError(t, err)

	expected := []byte{
		0x0A, 0x00, 0x00, 0x00, // Länge 10 Bytes → es folgt ein Attribut
		0x03, 0x00, // ID x0003: Klartextstring
		0x46, 0x61, 0x68, 0x72, 0x70, 0x75, 0x6C, 0x74, // String „Fahrpult“ (8 Zeichen, da 2 bytes für die ID)
	}

	assert.EqualValues(t, expected, res)
}

func TestEncapsulateAttributeSlice(t *testing.T) {
	slice := []uint16{2, 5}

	res, err := encapsulate(0x0001, reflect.ValueOf(slice))
	assert.NoError(t, err)

	expected := []byte{
		0x04, 0x00, 0x00, 0x00,
		0x01, 0x00,
		0x02, 0x00,
		0x04, 0x00, 0x00, 0x00,
		0x01, 0x00,
		0x05, 0x00,
	}

	assert.EqualValues(t, expected, res)
}

func TestEncapsulateAttributesHandledRightLength(t *testing.T) {
	t.Run("byte", func(t *testing.T) { testReflectionAssertLength(t, byte(5), 7) })
	t.Run("int8", func(t *testing.T) { testReflectionAssertLength(t, int8(5), 7) })
	t.Run("uint16", func(t *testing.T) { testReflectionAssertLength(t, uint16(5), 8) })
	t.Run("int16", func(t *testing.T) { testReflectionAssertLength(t, int16(5), 8) })
	t.Run("uint32", func(t *testing.T) { testReflectionAssertLength(t, uint32(5), 10) })
	t.Run("int32", func(t *testing.T) { testReflectionAssertLength(t, int32(5), 10) })
	t.Run("uint64", func(t *testing.T) { testReflectionAssertLength(t, uint64(5), 14) })
	t.Run("int64", func(t *testing.T) { testReflectionAssertLength(t, int64(5), 14) })
	t.Run("float32", func(t *testing.T) { testReflectionAssertLength(t, float32(5.1), 10) })
	t.Run("float64", func(t *testing.T) { testReflectionAssertLength(t, float64(5.1), 14) })
	t.Run("string", func(t *testing.T) { testReflectionAssertLength(t, "hello", 11) })
}

func testReflectionAssertLength(t *testing.T, value interface{}, expectedLength int) {
	res, err := encapsulate(0x0002, reflect.ValueOf(value))

	assert.NoError(t, err)
	assert.Len(t, res, expectedLength)
}

func BenchmarkEncapsulateUint32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encapsulate(0x0002, reflect.ValueOf(uint32(i%500)))
	}
}

func BenchmarkEncapsulateInt32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encapsulate(0x0002, reflect.ValueOf(int32(i%500)))
	}
}

func BenchmarkEncapsulateInt64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encapsulate(0x0002, reflect.ValueOf(int64(i%500)))
	}
}

func BenchmarkEncapsulateFloat64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encapsulate(0x0002, reflect.ValueOf(float64(i%500)))
	}
}

func BenchmarkEncapsulateString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		encapsulate(0x0002, reflect.ValueOf(strconv.Itoa(i)))
	}
}

func TestMarshalBeispiel(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		marshalBeispiel(t, beispiel1msg, beispiel1bytes)
	})

	t.Run("2", func(t *testing.T) {
		marshalBeispiel(t, beispiel2msg, beispiel2bytes)
	})

	t.Run("3", func(t *testing.T) {
		marshalBeispiel(t, beispiel3msg, beispiel3bytes)
	})

	t.Run("4", func(t *testing.T) {
		marshalBeispiel(t, beispiel4msg, beispiel4bytes)
	})

	t.Run("5", func(t *testing.T) {
		marshalBeispiel(t, beispiel5msg, beispiel5bytes)
	})
}

func marshalBeispiel(t *testing.T, message message.Message, expected []byte) {
	res, err := MarshalMessage(message)
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestMarshalStructSlice(t *testing.T) {
	res, err := Marshal(structWithSliceExample)

	assert.NoError(t, err)
	assert.Equal(t, structWithSliceBytes, res)
}

func TestMarshalStructWithPtr(t *testing.T) {
	strct := StructWithPtr{Float32(5.3), nil}

	res, err := Marshal(strct)
	expected := []byte{
		0x06, 0x00, 0x00, 0x00,
		0x01, 0x00,
		0x9a, 0x99, 0xa9, 0x40,
	}

	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

type StructWithPtr struct {
	Float32Ptr    *float32 `zusi:"0001"`
	Float32NilPtr *float32 `zusi:"0002"`
}
