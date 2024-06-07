package codec

import (
	"errors"
	"io"
	"reflect"

	"google.golang.org/protobuf/encoding/protowire"
)

type ProtoMapKey ProtoValue
type ProtoMapValue ProtoValue
type ProtoMapElem struct {
	Key   ProtoMapKey
	Value ProtoMapValue
}

var (
	ErrEmptyProtoMapElems = errors.New("can not fill map from empty proto map elems")
)

// FillMapFromProtoMapElem 将ProtoMapElem数组转换成一个可由反射获取值的map
func FillMapFromProtoMapElem(elems []ProtoMapElem) (reflect.Value, error) {
	if len(elems) == 0 {
		return reflect.Zero(reflect.TypeOf(nil)), ErrEmptyProtoMapElems
	}
	mType := reflect.MapOf(reflect.TypeOf(elems[0].Key.val), reflect.TypeOf(elems[0].Value.val))
	m := reflect.MakeMap(mType)
	for _, ele := range elems {
		m.SetMapIndex(reflect.ValueOf(ele.Key.val), reflect.ValueOf(ele.Value.val))
	}
	return m, nil
}

// where the key_type can be any integral or string type (so, any scalar type except for floating point types and bytes)
type keyDecoder func([]byte) ([]byte, ProtoMapKey, error)

type valueDecoder func([]byte) ([]byte, ProtoMapValue, error)

var Int32KeyDecoder keyDecoder = func(b []byte) ([]byte, ProtoMapKey, error) {
	n, v, err := variantDecoder(b, keyTag)
	if err != nil {
		return nil, ProtoMapKey{}, err
	}
	return b[n:], ProtoMapKey{_type: protowire.VarintType, val: int32(v)}, nil
}

var Int32ValueDecoder valueDecoder = func(b []byte) ([]byte, ProtoMapValue, error) {
	n, v, err := variantDecoder(b, valTag)
	if err != nil {
		return nil, ProtoMapValue{}, err
	}
	return b[n:], ProtoMapValue{_type: protowire.VarintType, val: int32(v)}, nil
}

var Int64KeyDecoder keyDecoder = func(b []byte) ([]byte, ProtoMapKey, error) {
	n, v, err := variantDecoder(b, keyTag)
	if err != nil {
		return nil, ProtoMapKey{}, err
	}
	return b[n:], ProtoMapKey{_type: protowire.VarintType, val: int64(v)}, nil
}

var Int64ValueDecoder valueDecoder = func(b []byte) ([]byte, ProtoMapValue, error) {
	n, v, err := variantDecoder(b, valTag)
	if err != nil {
		return nil, ProtoMapValue{}, err
	}
	return b[n:], ProtoMapValue{_type: protowire.VarintType, val: int64(v)}, nil
}

var Uint32KeyDecoder keyDecoder = func(b []byte) ([]byte, ProtoMapKey, error) {
	n, v, err := variantDecoder(b, keyTag)
	if err != nil {
		return nil, ProtoMapKey{}, err
	}
	return b[n:], ProtoMapKey{_type: protowire.VarintType, val: uint32(v)}, nil
}

var Uint32ValueDecoder valueDecoder = func(b []byte) ([]byte, ProtoMapValue, error) {
	n, v, err := variantDecoder(b, valTag)
	if err != nil {
		return nil, ProtoMapValue{}, err
	}
	return b[n:], ProtoMapValue{_type: protowire.VarintType, val: uint32(v)}, nil
}

var Uint64KeyDecoder keyDecoder = func(b []byte) ([]byte, ProtoMapKey, error) {
	n, v, err := variantDecoder(b, keyTag)
	if err != nil {
		return nil, ProtoMapKey{}, err
	}
	return b[n:], ProtoMapKey{_type: protowire.VarintType, val: v}, nil
}

var Uint64ValueDecoder valueDecoder = func(b []byte) ([]byte, ProtoMapValue, error) {
	n, v, err := variantDecoder(b, valTag)
	if err != nil {
		return nil, ProtoMapValue{}, err
	}
	return b[n:], ProtoMapValue{_type: protowire.VarintType, val: v}, nil
}

var Sint32KeyDecoder keyDecoder = func(b []byte) ([]byte, ProtoMapKey, error) {
	n, v, err := variantDecoder(b, keyTag)
	if err != nil {
		return nil, ProtoMapKey{}, err
	}
	return b[n:], ProtoMapKey{_type: protowire.VarintType, val: int32(protowire.DecodeZigZag(v))}, nil
}

var Sint32ValueDecoder valueDecoder = func(b []byte) ([]byte, ProtoMapValue, error) {
	n, v, err := variantDecoder(b, valTag)
	if err != nil {
		return nil, ProtoMapValue{}, err
	}
	return b[n:], ProtoMapValue{_type: protowire.VarintType, val: int32(protowire.DecodeZigZag(v))}, nil
}

var Sint64KeyDecoder keyDecoder = func(b []byte) ([]byte, ProtoMapKey, error) {
	n, v, err := variantDecoder(b, keyTag)
	if err != nil {
		return nil, ProtoMapKey{}, err
	}
	return b[n:], ProtoMapKey{_type: protowire.VarintType, val: protowire.DecodeZigZag(v)}, nil
}

var Sint64ValueDecoder valueDecoder = func(b []byte) ([]byte, ProtoMapValue, error) {
	n, v, err := variantDecoder(b, valTag)
	if err != nil {
		return nil, ProtoMapValue{}, err
	}
	return b[n:], ProtoMapValue{_type: protowire.VarintType, val: protowire.DecodeZigZag(v)}, nil
}

var BoolKeyDecoder keyDecoder = func(b []byte) ([]byte, ProtoMapKey, error) {
	n, v, err := variantDecoder(b, keyTag)
	if err != nil {
		return nil, ProtoMapKey{}, err
	}
	return b[n:], ProtoMapKey{_type: protowire.VarintType, val: protowire.DecodeBool(v)}, nil
}

var BoolValueDecoder valueDecoder = func(b []byte) ([]byte, ProtoMapValue, error) {
	n, v, err := variantDecoder(b, valTag)
	if err != nil {
		return nil, ProtoMapValue{}, err
	}
	return b[n:], ProtoMapValue{_type: protowire.VarintType, val: protowire.DecodeBool(v)}, nil
}

var Fixed64KeyDecoder keyDecoder = func(b []byte) ([]byte, ProtoMapKey, error) {
	n, v, err := i64Decoder(b, keyTag)
	if err != nil {
		return nil, ProtoMapKey{}, err
	}
	return b[n:], ProtoMapKey{_type: protowire.VarintType, val: v}, nil
}

var Fixed64ValueDecoder valueDecoder = func(b []byte) ([]byte, ProtoMapValue, error) {
	n, v, err := i64Decoder(b, valTag)
	if err != nil {
		return nil, ProtoMapValue{}, err
	}
	return b[n:], ProtoMapValue{_type: protowire.VarintType, val: v}, nil
}

var Sfixed64KeyDecoder keyDecoder = func(b []byte) ([]byte, ProtoMapKey, error) {
	n, v, err := i64Decoder(b, keyTag)
	if err != nil {
		return nil, ProtoMapKey{}, err
	}
	return b[n:], ProtoMapKey{_type: protowire.VarintType, val: int64(v)}, nil
}

var Sfixed64ValueDecoder valueDecoder = func(b []byte) ([]byte, ProtoMapValue, error) {
	n, v, err := i64Decoder(b, valTag)
	if err != nil {
		return nil, ProtoMapValue{}, err
	}
	return b[n:], ProtoMapValue{_type: protowire.VarintType, val: int64(v)}, nil
}

var StringKeyDecoder keyDecoder = func(b []byte) ([]byte, ProtoMapKey, error) {
	n, v, err := lenDecoder(b, keyTag)
	if err != nil {
		return nil, ProtoMapKey{}, err
	}
	return b[n:], ProtoMapKey{_type: protowire.VarintType, val: string(v)}, nil
}

var StringValueDecoder valueDecoder = func(b []byte) ([]byte, ProtoMapValue, error) {
	n, v, err := lenDecoder(b, valTag)
	if err != nil {
		return nil, ProtoMapValue{}, err
	}
	return b[n:], ProtoMapValue{_type: protowire.VarintType, val: string(v)}, nil
}

var BytesKeyDecoder keyDecoder = func(b []byte) ([]byte, ProtoMapKey, error) {
	n, v, err := lenDecoder(b, keyTag)
	if err != nil {
		return nil, ProtoMapKey{}, err
	}
	return b[n:], ProtoMapKey{_type: protowire.VarintType, val: v}, nil
}

var BytesValueDecoder valueDecoder = func(b []byte) ([]byte, ProtoMapValue, error) {
	n, v, err := lenDecoder(b, valTag)
	if err != nil {
		return nil, ProtoMapValue{}, err
	}
	return b[n:], ProtoMapValue{_type: protowire.VarintType, val: v}, nil
}

var MessageValueDecoder valueDecoder = func(b []byte) ([]byte, ProtoMapValue, error) {
	n, v, err := lenDecoder(b, valTag)
	if err != nil {
		return nil, ProtoMapValue{}, err
	}
	msg, err := DecodeBinaryData(v)
	if err != nil {
		return nil, ProtoMapValue{}, err
	}
	return b[n:], ProtoMapValue{_type: protowire.VarintType, val: msg}, nil
}

const (
	keyTag = protowire.Number(1)
	valTag = protowire.Number(2)
)

func baseDecoder(b []byte, _typ protowire.Type, _tag protowire.Number) (int, error) {
	if len(b) == 0 {
		return 0, io.EOF
	}
	// decode tag
	v, n := protowire.ConsumeVarint(b)
	if n < 0 {
		return 0, protowire.ParseError(n)
	}
	tag, typ := protowire.DecodeTag(v)
	if typ != _typ {
		return 0, errors.New("current payload type invalid")
	}
	if tag != _tag {
		return 0, errors.New("current payload's data tag valid")
	}
	return n, nil
}

func variantDecoder(b []byte, tag protowire.Number) (int, uint64, error) {
	m, err := baseDecoder(b, protowire.VarintType, tag)
	if err != nil {
		return 0, 0, err
	}
	b = b[m:]
	v, n := protowire.ConsumeVarint(b)
	if n < 0 {
		return 0, 0, protowire.ParseError(n)
	}
	return m + n, v, nil
}

func i64Decoder(b []byte, tag protowire.Number) (int, uint64, error) {
	m, err := baseDecoder(b, protowire.Fixed64Type, tag)
	if err != nil {
		return 0, 0, err
	}
	b = b[m:]
	v, n := protowire.ConsumeFixed64(b)
	if n < 0 {
		return 0, 0, protowire.ParseError(n)
	}
	return m + n, v, nil
}

func lenDecoder(b []byte, tag protowire.Number) (int, []byte, error) {
	m, err := baseDecoder(b, protowire.BytesType, tag)
	if err != nil {
		return 0, nil, err
	}
	b = b[m:]
	// 解出长度
	l, n := protowire.ConsumeVarint(b)
	if n < 0 {
		return 0, nil, protowire.ParseError(n)
	}
	b = b[n:]
	return m + n + int(l), b[:l], nil
}
