package codec

import (
	"math"

	"google.golang.org/protobuf/encoding/protowire"
)

// DecodeInt32 将底层数据尝试解析为int32
func (p ProtoValue) DecodeInt32() (int32, error) {
	val, err := p.parseVariant()
	if err != nil {
		return 0, err
	}
	return int32(val), nil
}

// DecodeInt64 将底层数据尝试解析为int64
func (p ProtoValue) DecodeInt64() (int64, error) {
	val, err := p.parseVariant()
	if err != nil {
		return 0, err
	}
	return int64(val), nil
}

// DecodeUint32 将底层数据尝试解析为uint32
func (p ProtoValue) DecodeUint32() (uint32, error) {
	val, err := p.parseVariant()
	if err != nil {
		return 0, err
	}
	return uint32(val), nil
}

// DecodeUint64 将底层数据尝试解析为uint64
func (p ProtoValue) DecodeUint64() (uint64, error) {
	val, err := p.parseVariant()
	if err != nil {
		return 0, err
	}
	return uint64(val), nil
}

// DecodeSint32 将底层数据尝试解析为sint32（ZigZag解码后）
func (p ProtoValue) DecodeSint32() (int32, error) {
	val, err := p.parseVariant()
	if err != nil {
		return 0, err
	}
	return int32(protowire.DecodeZigZag(val)), nil
}

// DecodeSint64 将底层数据尝试解析为sint64（ZigZag解码后）
func (p ProtoValue) DecodeSint64() (int64, error) {
	val, err := p.parseVariant()
	if err != nil {
		return 0, err
	}
	return protowire.DecodeZigZag(val), nil
}

// DecodeBool 将底层数据尝试解析为bool
func (p ProtoValue) DecodeBool() (bool, error) {
	val, err := p.parseVariant()
	if err != nil {
		return false, err
	}
	return protowire.DecodeBool(val), nil
}

// DecodeEnum 将底层数据尝试解析为enum（enum底层是int32类型）
func (p ProtoValue) DecodeEnum() (int32, error) {
	val, err := p.parseVariant()
	if err != nil {
		return 0, err
	}
	return int32(val), nil
}

// DecodeFixed64 将底层数据尝试解析为fixed64
func (p ProtoValue) DecodeFixed64() (uint64, error) {
	val, err := p.parseI64()
	if err != nil {
		return 0, err
	}
	return val, nil
}

// DecodeSfixed64 将底层数据尝试解析为sfixed64
func (p ProtoValue) DecodeSfixed64() (int64, error) {
	val, err := p.parseI64()
	if err != nil {
		return 0, err
	}
	return int64(val), nil
}

// DecodeDouble 将底层数据尝试解析为double（对应golang float64类型）
func (p ProtoValue) DecodeDouble() (float64, error) {
	val, err := p.parseI64()
	if err != nil {
		return 0, err
	}
	return math.Float64frombits(val), nil
}

// DecodeString 将底层数据尝试解析为string
func (p ProtoValue) DecodeString() (string, error) {
	val, err := p.parseLen()
	if err != nil {
		return "", err
	}
	return string(val), nil
}

// DecodeBytes 将底层数据尝试解析为[]byte
func (p ProtoValue) DecodeBytes() ([]byte, error) {
	val, err := p.parseLen()
	if err != nil {
		return nil, err
	}
	return val, nil
}

// DecodeEmbeddedMsg 将底层数据尝试解析为嵌套proto message
func (p ProtoValue) DecodeEmbeddedMsg(sortType MessageSortType) (ProtoMessage, error) {
	val, err := p.parseLen()
	if err != nil {
		return ProtoMessage{}, err
	}
	return Decode(val, sortType)
}

// DecodeMap 将底层数据尝试解析为嵌套proto map类型
func (p ProtoMessage) DecodeMap(tag protowire.Number, keyDec keyDecoder, valDec valueDecoder) ([]ProtoMapElem, error) {
	idxs, err := p.GetRepeatedData(tag)
	if err != nil {
		return nil, err
	}
	m := make([]ProtoMapElem, 0, len(idxs))
	for i := 0; i < len(idxs); i++ {
		var key ProtoMapKey
		var value ProtoMapValue
		var err error

		payload := p.Values[idxs[i]].val.([]byte)
		payload, key, err = keyDec(payload)
		if err != nil {
			return nil, err
		}
		_, value, err = valDec(payload)
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

// DecodePackedRepeated 将底层数据尝试解析为[packed=true]的repeated字段数据
//
// 用于非repeated string, repeated bytes和repeated message
func (p ProtoMessage) DecodePackedRepeated(tag protowire.Number, decoder packedRepeatedDecoder) (interface{}, error) {
	idxs, err := p.GetRepeatedData(tag)
	if err != nil {
		return nil, err
	}
	if len(idxs) == 0 {
		return decoder(ProtoValue{})
	}
	// 数据是variant类型数据的集合，通过传入的decoder进行解码
	return decoder(p.Values[idxs[0]])
}

// DecodeUnpackedRepeated 将底层数据尝试解析为[packed=false]的repeated字段数据
//
// 也可用来解析非repeated scalar type（即repeated数字类型）的数据
func (p ProtoMessage) DecodeUnpackedRepeated(tag protowire.Number, decoder unpackedRepeatedDecoder) (interface{}, error) {
	idxs, err := p.GetRepeatedData(tag)
	if err != nil {
		return nil, err
	}
	// 数据是variant类型数据的集合，通过传入的decoder进行解码
	return decoder(p, idxs)
}

// DecodeFixed32 将底层数据尝试解析为fixed32字段数据
func (p ProtoValue) DecodeFixed32() (uint32, error) {
	val, err := p.parseI32()
	if err != nil {
		return 0, err
	}
	return val, nil
}

// DecodeSfixed32 将底层数据尝试解析为sfixed32字段数据
func (p ProtoValue) DecodeSfixed32() (int32, error) {
	val, err := p.parseI32()
	if err != nil {
		return 0, err
	}
	return int32(val), nil
}

// DecodeFloat 将底层数据尝试解析为float字段数据
func (p ProtoValue) DecodeFloat() (float32, error) {
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
