package codec

import (
	"math"

	"google.golang.org/protobuf/encoding/protowire"
)

type packedRepeatedDecoder func(ProtoValue) (interface{}, error)

// 标记为packed的packedRepeatedDecoder，仅适用于wire_type为VARINT的数字类型

// PackedRepeatedInt32Decoder 解码repeated int32
var PackedRepeatedInt32Decoder packedRepeatedDecoder = func(p ProtoValue) (interface{}, error) {
	result := []int32{}

	payload := p.val.([]byte)
	for len(payload) > 0 {
		val, n := protowire.ConsumeVarint(payload)
		if n < 0 {
			return nil, protowire.ParseError(n)
		}
		result = append(result, int32(val))
		payload = payload[n:]
	}
	return result, nil
}

// PackedRepeatedInt64Decoder 解码repeated int64
var PackedRepeatedInt64Decoder packedRepeatedDecoder = func(p ProtoValue) (interface{}, error) {
	result := []int64{}

	payload := p.val.([]byte)
	for len(payload) > 0 {
		val, n := protowire.ConsumeVarint(payload)
		if n < 0 {
			return nil, protowire.ParseError(n)
		}
		result = append(result, int64(val))
		payload = payload[n:]
	}
	return result, nil
}

// PackedRepeatedUint32Decoder 解码repeated uint32
var PackedRepeatedUint32Decoder packedRepeatedDecoder = func(p ProtoValue) (interface{}, error) {
	result := []uint32{}

	payload := p.val.([]byte)
	for len(payload) > 0 {
		val, n := protowire.ConsumeVarint(payload)
		if n < 0 {
			return nil, protowire.ParseError(n)
		}
		result = append(result, uint32(val))
		payload = payload[n:]
	}
	return result, nil
}

// PackedRepeatedUint64Decoder 解码repeated uint64
var PackedRepeatedUint64Decoder packedRepeatedDecoder = func(p ProtoValue) (interface{}, error) {
	result := []uint64{}

	payload := p.val.([]byte)
	for len(payload) > 0 {
		val, n := protowire.ConsumeVarint(payload)
		if n < 0 {
			return nil, protowire.ParseError(n)
		}
		result = append(result, uint64(val))
		payload = payload[n:]
	}
	return result, nil
}

// PackedRepeatedSint32Decoder 解码repeated sint32
var PackedRepeatedSint32Decoder packedRepeatedDecoder = func(p ProtoValue) (interface{}, error) {
	result := []int32{}

	payload := p.val.([]byte)
	for len(payload) > 0 {
		val, n := protowire.ConsumeVarint(payload)
		if n < 0 {
			return nil, protowire.ParseError(n)
		}
		result = append(result, int32(protowire.DecodeZigZag(val)))
		payload = payload[n:]
	}
	return result, nil
}

// PackedRepeatedSint64Decoder 解码repeated sint64
var PackedRepeatedSint64Decoder packedRepeatedDecoder = func(p ProtoValue) (interface{}, error) {
	result := []int64{}

	payload := p.val.([]byte)
	for len(payload) > 0 {
		val, n := protowire.ConsumeVarint(payload)
		if n < 0 {
			return nil, protowire.ParseError(n)
		}
		result = append(result, protowire.DecodeZigZag(val))
		payload = payload[n:]
	}
	return result, nil
}

// PackedRepeatedBoolDecoder 解码repeated bool
var PackedRepeatedBoolDecoder packedRepeatedDecoder = func(p ProtoValue) (interface{}, error) {
	result := []bool{}

	payload := p.val.([]byte)
	for len(payload) > 0 {
		val, n := protowire.ConsumeVarint(payload)
		if n < 0 {
			return nil, protowire.ParseError(n)
		}
		result = append(result, protowire.DecodeBool(val))
		payload = payload[n:]
	}
	return result, nil
}

// PackedRepeatedEnumDecoder 解码repeated enum（本质是[]int32）
var PackedRepeatedEnumDecoder packedRepeatedDecoder = func(p ProtoValue) (interface{}, error) {
	return PackedRepeatedInt32Decoder(p)
}

// PackedRepeatedFixed64Decoder 解码repeated fixed64
var PackedRepeatedFixed64Decoder packedRepeatedDecoder = func(p ProtoValue) (interface{}, error) {
	result := []uint64{}

	payload := p.val.([]byte)
	for len(payload) > 0 {
		val, n := protowire.ConsumeFixed64(payload)
		if n < 0 {
			return nil, protowire.ParseError(n)
		}
		result = append(result, val)
		payload = payload[n:]
	}
	return result, nil
}

// PackedRepeatedSfixed64Decoder 解码repeated sfixed64
var PackedRepeatedSfixed64Decoder packedRepeatedDecoder = func(p ProtoValue) (interface{}, error) {
	result := []int64{}

	payload := p.val.([]byte)
	for len(payload) > 0 {
		val, n := protowire.ConsumeFixed64(payload)
		if n < 0 {
			return nil, protowire.ParseError(n)
		}
		result = append(result, int64(val))
		payload = payload[n:]
	}
	return result, nil
}

// PackedRepeatedDoubleDecoder 解码repeated double
var PackedRepeatedDoubleDecoder packedRepeatedDecoder = func(p ProtoValue) (interface{}, error) {
	result := []float64{}

	payload := p.val.([]byte)
	for len(payload) > 0 {
		val, n := protowire.ConsumeFixed64(payload)
		if n < 0 {
			return nil, protowire.ParseError(n)
		}
		result = append(result, math.Float64frombits(val))
		payload = payload[n:]
	}
	return result, nil
}

// PackedRepeatedFixed32Decoder 解码repeated fixed32
var PackedRepeatedFixed32Decoder packedRepeatedDecoder = func(p ProtoValue) (interface{}, error) {
	result := []uint32{}

	payload := p.val.([]byte)
	for len(payload) > 0 {
		val, n := protowire.ConsumeFixed32(payload)
		if n < 0 {
			return nil, protowire.ParseError(n)
		}
		result = append(result, val)
		payload = payload[n:]
	}
	return result, nil
}

// PackedRepeatedSfixed32Decoder 解码repeated sfixed32
var PackedRepeatedSfixed32Decoder packedRepeatedDecoder = func(p ProtoValue) (interface{}, error) {
	result := []int32{}

	payload := p.val.([]byte)
	for len(payload) > 0 {
		val, n := protowire.ConsumeFixed32(payload)
		if n < 0 {
			return nil, protowire.ParseError(n)
		}
		result = append(result, int32(val))
		payload = payload[n:]
	}
	return result, nil
}

// PackedRepeatedFloatDecoder 解码repeated float
var PackedRepeatedFloatDecoder packedRepeatedDecoder = func(p ProtoValue) (interface{}, error) {
	result := []float32{}

	payload := p.val.([]byte)
	for len(payload) > 0 {
		val, n := protowire.ConsumeFixed32(payload)
		if n < 0 {
			return nil, protowire.ParseError(n)
		}
		result = append(result, math.Float32frombits(val))
		payload = payload[n:]
	}
	return result, nil
}
