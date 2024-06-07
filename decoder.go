package codec

import (
	"math"

	"google.golang.org/protobuf/encoding/protowire"
)

// ParseInt32 将底层数据尝试解析为int32
func (p ProtoValue) ParseInt32() (int32, error) {
	val, err := p.parseVariant()
	if err != nil {
		return 0, err
	}
	return int32(val), nil
}

// ParseInt64 将底层数据尝试解析为int64
func (p ProtoValue) ParseInt64() (int64, error) {
	val, err := p.parseVariant()
	if err != nil {
		return 0, err
	}
	return int64(val), nil
}

// ParseUint32 将底层数据尝试解析为uint32
func (p ProtoValue) ParseUint32() (uint32, error) {
	val, err := p.parseVariant()
	if err != nil {
		return 0, err
	}
	return uint32(val), nil
}

// ParseUint64 将底层数据尝试解析为uint64
func (p ProtoValue) ParseUint64() (uint64, error) {
	val, err := p.parseVariant()
	if err != nil {
		return 0, err
	}
	return uint64(val), nil
}

// ParseSint32 将底层数据尝试解析为sint32（ZigZag解码后）
func (p ProtoValue) ParseSint32() (int32, error) {
	val, err := p.parseVariant()
	if err != nil {
		return 0, err
	}
	return int32(protowire.DecodeZigZag(val)), nil
}

// ParseSint64 将底层数据尝试解析为sint64（ZigZag解码后）
func (p ProtoValue) ParseSint64() (int64, error) {
	val, err := p.parseVariant()
	if err != nil {
		return 0, err
	}
	return protowire.DecodeZigZag(val), nil
}

// ParseBool 将底层数据尝试解析为bool
func (p ProtoValue) ParseBool() (bool, error) {
	val, err := p.parseVariant()
	if err != nil {
		return false, err
	}
	return protowire.DecodeBool(val), nil
}

// ParseEnum 将底层数据尝试解析为enum（enum底层是int32类型）
func (p ProtoValue) ParseEnum() (int32, error) {
	val, err := p.parseVariant()
	if err != nil {
		return 0, err
	}
	return int32(val), nil
}

// ParseFixed64 将底层数据尝试解析为fixed64
func (p ProtoValue) ParseFixed64() (uint64, error) {
	val, err := p.parseI64()
	if err != nil {
		return 0, err
	}
	return val, nil
}

// ParseSfixed64 将底层数据尝试解析为sfixed64
func (p ProtoValue) ParseSfixed64() (int64, error) {
	val, err := p.parseI64()
	if err != nil {
		return 0, err
	}
	return int64(val), nil
}

// ParseDouble 将底层数据尝试解析为double（对应golang float64类型）
func (p ProtoValue) ParseDouble() (float64, error) {
	val, err := p.parseI64()
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(val), nil
}

// ParseString 将底层数据尝试解析为string
func (p ProtoValue) ParseString() (string, error) {
	val, err := p.parseLen()
	if err != nil {
		return "", err
	}
	return string(val), nil
}

// ParseBytes 将底层数据尝试解析为[]byte
func (p ProtoValue) ParseBytes() ([]byte, error) {
	val, err := p.parseLen()
	if err != nil {
		return nil, err
	}
	return val, nil
}

// ParseEmbeddedMsg 将底层数据尝试解析为嵌套proto message
func (p ProtoValue) ParseEmbeddedMsg() (ProtoMessage, error) {
	val, err := p.parseLen()
	if err != nil {
		return nil, err
	}
	return DecodeBinaryData(val)
}

// ParseMap 将底层数据尝试解析为嵌套proto map类型
func (p ProtoValue) ParseMap(keyDec keyDecoder, valDec valueDecoder) ([]ProtoMapElem, error) {
	msg, err := p.parseUnpackedRepeated()
	if err != nil {
		return nil, err
	}
	m := make([]ProtoMapElem, 0, len(msg))
	for i := 0; i < len(msg); i++ {
		var key ProtoMapKey
		var value ProtoMapValue
		var err error
		msg[i], key, err = keyDec(msg[i])
		if err != nil {
			return nil, err
		}
		msg[i], value, err = valDec(msg[i])
		if err != nil {
			return nil, err
		}
		m = append(m, ProtoMapElem{
			Key:   ProtoMapKey(key),
			Value: ProtoMapValue(value),
		})
	}
	return m, nil
}

// ParsePackedRepeated 将底层数据尝试解析为[packed=true]的repeated字段数据
//
// 用于非repeated string, repeated bytes和repeated message
func (p ProtoValue) ParsePackedRepeated(decoder packedRepeatedDecoder) (interface{}, error) {
	val, err := p.parsePackedRepeated()
	if err != nil {
		return nil, err
	}
	// 数据是variant类型数据的集合，通过传入的decoder进行解码
	return decoder(val)
}

// ParseUnpackedRepeated 将底层数据尝试解析为[packed=false]的repeated字段数据
//
// 也可用来解析非repeated scalar type（即repeated数字类型）的数据
func (p ProtoValue) ParseUnpackedRepeated(decoder unpackedRepeatedDecoder) (interface{}, error) {
	val, err := p.parseUnpackedRepeated()
	if err != nil {
		return nil, err
	}
	// 数据是variant类型数据的集合，通过传入的decoder进行解码
	return decoder(val)
}

// ParseFixed32 将底层数据尝试解析为fixed32字段数据
func (p ProtoValue) ParseFixed32() (uint32, error) {
	val, err := p.parseI32()
	if err != nil {
		return 0, err
	}
	return val, nil
}

// ParseSfixed32 将底层数据尝试解析为sfixed32字段数据
func (p ProtoValue) ParseSfixed32() (int32, error) {
	val, err := p.parseI32()
	if err != nil {
		return 0, err
	}
	return int32(val), nil
}

// ParseFloat 将底层数据尝试解析为float字段数据
func (p ProtoValue) ParseFloat() (float32, error) {
	val, err := p.parseI32()
	if err != nil {
		return 0, err
	}
	return math.Float32frombits(val), nil
}

func (p ProtoValue) parseVariant() (uint64, error) {
	if p.val == nil {
		// 零值情况
		return 0, nil
	}
	if p._type != protowire.VarintType {
		return 0, ErrTypeMismatch
	}
	val, ok := p.val.(uint64)
	if !ok {
		return 0, ErrAssertTypeFailed
	}
	return val, nil
}

func (p ProtoValue) parseI64() (uint64, error) {
	if p.val == nil {
		// 零值情况
		return 0, nil
	}
	if p._type != protowire.Fixed64Type {
		return 0, ErrTypeMismatch
	}
	val, ok := p.val.(uint64)
	if !ok {
		return 0, ErrAssertTypeFailed
	}
	return val, nil
}

func (p ProtoValue) parseLen() ([]byte, error) {
	if p.val == nil {
		// 零值情况
		return []byte{}, nil
	}
	if p._type != protowire.BytesType {
		return nil, ErrTypeMismatch
	}
	val, ok := p.val.([]byte)
	if !ok {
		return nil, ErrAssertTypeFailed
	}
	return val, nil
}

func (p ProtoValue) parseUnpackedRepeated() ([][]byte, error) {
	if p.val == nil {
		// 零值情况
		return [][]byte{}, nil
	}
	if p._type != protowire.BytesType {
		return nil, ErrTypeMismatch
	}
	val, ok := p.val.([][]byte)
	if !ok {
		return nil, ErrAssertTypeFailed
	}
	return val, nil
}

func (p ProtoValue) parsePackedRepeated() ([]byte, error) {
	if p.val == nil {
		// 零值情况
		return []byte{}, nil
	}
	if p._type != protowire.BytesType {
		return nil, ErrTypeMismatch
	}
	val, ok := p.val.([]byte)
	if !ok {
		return nil, ErrAssertTypeFailed
	}
	return val, nil
}

func (p ProtoValue) parseI32() (uint32, error) {
	if p.val == nil {
		// 零值情况
		return 0, nil
	}
	if p._type != protowire.Fixed32Type {
		return 0, ErrTypeMismatch
	}
	val, ok := p.val.(uint32)
	if !ok {
		return 0, ErrAssertTypeFailed
	}
	return val, nil
}
