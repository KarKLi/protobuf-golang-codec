package codec

import (
	"math"

	"google.golang.org/protobuf/encoding/protowire"
)

type unpackedRepeatedDecoder func(ProtoMessage, []int) (interface{}, error)

// 标记为packed的unpackedRepeatedDecoder，仅适用于wire_type为VARINT的数字类型

// UnpackedRepeatedInt32Decoder 解码repeated int32
var UnpackedRepeatedInt32Decoder unpackedRepeatedDecoder = func(m ProtoMessage, idxs []int) (interface{}, error) {
	result := make([]int32, 0, len(idxs))
	for i := range idxs {
		result = append(result, int32(m.Values[idxs[i]].val.(uint64)))
	}
	return result, nil
}

// UnpackedRepeatedInt64Decoder 解码repeated int64
var UnpackedRepeatedInt64Decoder unpackedRepeatedDecoder = func(m ProtoMessage, idxs []int) (interface{}, error) {
	result := make([]int64, 0, len(idxs))
	for i := range idxs {
		result = append(result, int64(m.Values[idxs[i]].val.(uint64)))
	}
	return result, nil
}

// UnpackedRepeatedUint32Decoder 解码repeated uint32
var UnpackedRepeatedUint32Decoder unpackedRepeatedDecoder = func(m ProtoMessage, idxs []int) (interface{}, error) {
	result := make([]uint32, 0, len(idxs))
	for i := range idxs {
		result = append(result, uint32(m.Values[idxs[i]].val.(uint64)))
	}
	return result, nil
}

// UnpackedRepeatedUint64Decoder 解码repeated uint64
var UnpackedRepeatedUint64Decoder unpackedRepeatedDecoder = func(m ProtoMessage, idxs []int) (interface{}, error) {
	result := make([]uint64, 0, len(idxs))
	for i := range idxs {
		result = append(result, uint64(m.Values[idxs[i]].val.(uint64)))
	}
	return result, nil
}

// UnpackedRepeatedSint32Decoder 解码repeated sint32
var UnpackedRepeatedSint32Decoder unpackedRepeatedDecoder = func(m ProtoMessage, idxs []int) (interface{}, error) {
	result := make([]int32, 0, len(idxs))
	for i := range idxs {
		result = append(result, int32(protowire.DecodeZigZag(m.Values[idxs[i]].val.(uint64))))
	}
	return result, nil
}

// UnpackedRepeatedSint64Decoder 解码repeated sint64
var UnpackedRepeatedSint64Decoder unpackedRepeatedDecoder = func(m ProtoMessage, idxs []int) (interface{}, error) {
	result := make([]int64, 0, len(idxs))
	for i := range idxs {
		result = append(result, int64(protowire.DecodeZigZag(m.Values[idxs[i]].val.(uint64))))
	}
	return result, nil
}

// UnpackedRepeatedBoolDecoder 解码repeated bool
var UnpackedRepeatedBoolDecoder unpackedRepeatedDecoder = func(m ProtoMessage, idxs []int) (interface{}, error) {
	result := make([]bool, 0, len(idxs))
	for i := range idxs {
		result = append(result, protowire.DecodeBool(m.Values[idxs[i]].val.(uint64)))
	}
	return result, nil
}

// UnpackedRepeatedEnumDecoder 解码repeated enum（本质是[]int32）
var UnpackedRepeatedEnumDecoder unpackedRepeatedDecoder = func(m ProtoMessage, idxs []int) (interface{}, error) {
	return UnpackedRepeatedInt32Decoder(m, idxs)
}

// UnpackedRepeatedFixed64Decoder 解码repeated fixed64
var UnpackedRepeatedFixed64Decoder unpackedRepeatedDecoder = func(m ProtoMessage, idxs []int) (interface{}, error) {
	result := make([]uint64, 0, len(idxs))
	for i := range idxs {
		result = append(result, uint64(m.Values[idxs[i]].val.(uint64)))
	}
	return result, nil
}

// UnpackedRepeatedSfixed64Decoder 解码repeated sfixed64
var UnpackedRepeatedSfixed64Decoder unpackedRepeatedDecoder = func(m ProtoMessage, idxs []int) (interface{}, error) {
	result := make([]int64, 0, len(idxs))
	for i := range idxs {
		result = append(result, int64(m.Values[idxs[i]].val.(uint64)))
	}
	return result, nil
}

// UnpackedRepeatedDoubleDecoder 解码repeated double
var UnpackedRepeatedDoubleDecoder unpackedRepeatedDecoder = func(m ProtoMessage, idxs []int) (interface{}, error) {
	result := make([]float64, 0, len(idxs))
	for i := range idxs {
		result = append(result, math.Float64frombits(m.Values[idxs[i]].val.(uint64)))
	}
	return result, nil
}

// UnpackedRepeatedStringDecoder 解码repeated string
var UnpackedRepeatedStringDecoder unpackedRepeatedDecoder = func(m ProtoMessage, idxs []int) (interface{}, error) {
	result := make([]string, 0, len(idxs))
	for i := range idxs {
		payload := m.Values[idxs[i]].val.([]byte)
		if len(payload) == 0 {
			result = append(result, "")
			continue
		}
		result = append(result, string(payload))
	}
	return result, nil
}

// UnpackedRepeatedBytesDecoder 解码repeated bytes
var UnpackedRepeatedBytesDecoder unpackedRepeatedDecoder = func(m ProtoMessage, idxs []int) (interface{}, error) {
	result := make([][]byte, 0, len(idxs))
	for i := range idxs {
		payload := m.Values[idxs[i]].val.([]byte)
		if len(payload) == 0 {
			result = append(result, []byte{})
			continue
		}
		result = append(result, payload)
	}
	return result, nil
}

// UnpackedRepeatedMessageDecoder 解码repeated message
var UnpackedRepeatedMessageDecoder unpackedRepeatedDecoder = func(m ProtoMessage, idxs []int) (interface{}, error) {
	result := make([]ProtoMessage, 0, len(idxs))
	for i := range idxs {
		payload := m.Values[idxs[i]].val.([]byte)
		if len(payload) == 0 {
			result = append(result, ProtoMessage{})
			continue
		}
		msg, err := Decode(payload, Asc)
		if err != nil {
			return nil, err
		}
		result = append(result, msg)
	}
	return result, nil
}

// UnpackedRepeatedFixed32Decoder 解码repeated fixed32
var UnpackedRepeatedFixed32Decoder unpackedRepeatedDecoder = func(m ProtoMessage, idxs []int) (interface{}, error) {
	result := make([]uint32, 0, len(idxs))
	for i := range idxs {
		result = append(result, uint32(m.Values[idxs[i]].val.(uint32)))
	}
	return result, nil
}

// UnpackedRepeatedSfixed32Decoder 解码repeated sfixed32
var UnpackedRepeatedSfixed32Decoder unpackedRepeatedDecoder = func(m ProtoMessage, idxs []int) (interface{}, error) {
	result := make([]int32, 0, len(idxs))
	for i := range idxs {
		result = append(result, int32(m.Values[idxs[i]].val.(uint32)))
	}
	return result, nil
}

// UnpackedRepeatedFloatDecoder 解码repeated float
var UnpackedRepeatedFloatDecoder unpackedRepeatedDecoder = func(m ProtoMessage, idxs []int) (interface{}, error) {
	result := make([]float32, 0, len(idxs))
	for i := range idxs {
		result = append(result, math.Float32frombits(m.Values[idxs[i]].val.(uint32)))
	}
	return result, nil
}
