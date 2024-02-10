package tcp

import (
	"encoding/binary"
	"io"
	"math"

	"github.com/pkg/errors"
)

func ReadHeader(bytes io.Reader) (uint32, uint16, error) {
	// get 4-byte length
	bts := make([]byte, 4)

	n, err := io.ReadFull(bytes, bts)
	if err != nil {
		return 0, 0, errors.Wrap(err, "error reading length of attribute")
	}
	if n != 4 {
		return 0, 0, errors.New("error reading length of attribute")
	}
	length := binary.LittleEndian.Uint32(bts)
	if length == 0xFFFFFFFF {
		return length, 0, nil
	}

	// get 2-byte attribute ID
	bts = make([]byte, 2)
	n, err = io.ReadFull(bytes, bts)
	if err != nil {
		return 0, 0, errors.Wrap(err, "error reading attribute Id")
	}
	if n != 2 {
		return 0, 0, errors.New("error reading attribute Id")
	}
	attributeId := binary.LittleEndian.Uint16(bts)

	return length, attributeId, nil
}

func ReadInt8(data io.Reader, length uint32) (*int8, error) {
	bts := make([]byte, length)

	n, err := io.ReadFull(data, bts)
	if uint32(n) != length {
		return nil, errors.New("invalid payload length received")
	}

	res := int8(bts[0])

	return &res, err
}

func ReadInt16(data io.Reader, length uint32) (*int16, error) {
	bts := make([]byte, length)

	n, err := io.ReadFull(data, bts)
	if uint32(n) != length {
		return nil, errors.New("invalid payload length received")
	}

	res := int16(binary.LittleEndian.Uint16(bts))

	return &res, err
}

func ReadInt32(data io.Reader, length uint32) (*int32, error) {
	bts := make([]byte, length)

	n, err := io.ReadFull(data, bts)
	if uint32(n) != length {
		return nil, errors.New("invalid payload length received")
	}

	res := int32(binary.LittleEndian.Uint32(bts))

	return &res, err
}

func ReadInt64(data io.Reader, length uint32) (*int64, error) {
	bts := make([]byte, length)

	n, err := io.ReadFull(data, bts)
	if uint32(n) != length {
		return nil, errors.New("invalid payload length received")
	}

	res := int64(binary.LittleEndian.Uint64(bts))

	return &res, err
}

func ReadUint8(data io.Reader, length uint32) (*uint8, error) {
	if length == 0 {
		return Uint8(0), nil
	}

	bts := make([]byte, length)

	n, err := io.ReadFull(data, bts)
	if uint32(n) != length {
		return nil, errors.New("invalid payload length received")
	}

	res := uint8(bts[0])

	return &res, err
}

func ReadUint16(data io.Reader, length uint32) (*uint16, error) {
	bts := make([]byte, length)

	n, err := io.ReadFull(data, bts)
	if uint32(n) != length {
		return nil, errors.New("invalid payload length received")
	}

	res := binary.LittleEndian.Uint16(bts)

	return &res, err
}

func ReadUint32(data io.Reader, length uint32) (*uint32, error) {
	bts := make([]byte, length)

	n, err := io.ReadFull(data, bts)
	if uint32(n) != length {
		return nil, errors.New("invalid payload length received")
	}

	res := binary.LittleEndian.Uint32(bts)

	return &res, err
}

func ReadUint64(data io.Reader, length uint32) (*uint64, error) {
	bts := make([]byte, length)

	n, err := io.ReadFull(data, bts)
	if uint32(n) != length {
		return nil, errors.New("invalid payload length received")
	}

	res := binary.LittleEndian.Uint64(bts)

	return &res, err
}

func ReadFloat32(data io.Reader, length uint32) (*float32, error) {
	bts := make([]byte, length)

	n, err := io.ReadFull(data, bts)
	if uint32(n) != length {
		return nil, errors.New("invalid payload length received")
	}

	res := math.Float32frombits(binary.LittleEndian.Uint32(bts))

	return &res, err
}

func ReadFloat64(data io.Reader, length uint32) (*float64, error) {
	bts := make([]byte, length)

	n, err := io.ReadFull(data, bts)
	if uint32(n) != length {
		return nil, errors.New("invalid payload length received")
	}

	res := math.Float64frombits(binary.LittleEndian.Uint64(bts))

	return &res, err
}

func ReadString(data io.Reader, length uint32) (*string, error) {
	bts := make([]byte, length)

	n, err := io.ReadFull(data, bts)
	if uint32(n) != length {
		return nil, errors.New("invalid payload length received")
	}

	res := string(bts)

	return &res, err
}
