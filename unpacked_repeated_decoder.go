package codec

type unpackedRepeatedDecoder func([][]byte) (interface{}, error)

// 标记为packed的unpackedRepeatedDecoder，仅适用于wire_type为VARINT的数字类型

// UnpackedRepeatedInt32Decoder 解码repeated int32
var UnpackedRepeatedInt32Decoder unpackedRepeatedDecoder = func(b [][]byte) (interface{}, error) {
	result := make([]int32, 0, len(b))
	for i := range b {
		val, err := PackedRepeatedInt32Decoder(b[i])
		if err != nil {
			return nil, err
		}
		result = append(result, val.([]int32)...)
	}
	return result, nil
}

// UnpackedRepeatedInt64Decoder 解码repeated int64
var UnpackedRepeatedInt64Decoder unpackedRepeatedDecoder = func(b [][]byte) (interface{}, error) {
	result := make([]int64, 0, len(b))
	for i := range b {
		val, err := PackedRepeatedInt64Decoder(b[i])
		if err != nil {
			return nil, err
		}
		result = append(result, val.([]int64)...)
	}
	return result, nil
}

// UnpackedRepeatedUint32Decoder 解码repeated uint32
var UnpackedRepeatedUint32Decoder unpackedRepeatedDecoder = func(b [][]byte) (interface{}, error) {
	result := make([]uint32, 0, len(b))
	for i := range b {
		val, err := PackedRepeatedUint32Decoder(b[i])
		if err != nil {
			return nil, err
		}
		result = append(result, val.([]uint32)...)
	}
	return result, nil
}

// UnpackedRepeatedUint64Decoder 解码repeated uint64
var UnpackedRepeatedUint64Decoder unpackedRepeatedDecoder = func(b [][]byte) (interface{}, error) {
	result := make([]uint64, 0, len(b))
	for i := range b {
		val, err := PackedRepeatedUint64Decoder(b[i])
		if err != nil {
			return nil, err
		}
		result = append(result, val.([]uint64)...)
	}
	return result, nil
}

// UnpackedRepeatedSint32Decoder 解码repeated sint32
var UnpackedRepeatedSint32Decoder unpackedRepeatedDecoder = func(b [][]byte) (interface{}, error) {
	result := make([]int32, 0, len(b))
	for i := range b {
		val, err := PackedRepeatedSint32Decoder(b[i])
		if err != nil {
			return nil, err
		}
		result = append(result, val.([]int32)...)
	}
	return result, nil
}

// UnpackedRepeatedSint64Decoder 解码repeated sint64
var UnpackedRepeatedSint64Decoder unpackedRepeatedDecoder = func(b [][]byte) (interface{}, error) {
	result := make([]int64, 0, len(b))
	for i := range b {
		val, err := PackedRepeatedSint64Decoder(b[i])
		if err != nil {
			return nil, err
		}
		result = append(result, val.([]int64)...)
	}
	return result, nil
}

// UnpackedRepeatedBoolDecoder 解码repeated bool
var UnpackedRepeatedBoolDecoder unpackedRepeatedDecoder = func(b [][]byte) (interface{}, error) {
	result := make([]bool, 0, len(b))
	for i := range b {
		val, err := PackedRepeatedBoolDecoder(b[i])
		if err != nil {
			return nil, err
		}
		result = append(result, val.([]bool)...)
	}
	return result, nil
}

// UnpackedRepeatedEnumDecoder 解码repeated enum（本质是[]int32）
var UnpackedRepeatedEnumDecoder unpackedRepeatedDecoder = func(b [][]byte) (interface{}, error) {
	return UnpackedRepeatedInt32Decoder(b)
}

// UnpackedRepeatedFixed64Decoder 解码repeated fixed64
var UnpackedRepeatedFixed64Decoder unpackedRepeatedDecoder = func(b [][]byte) (interface{}, error) {
	result := make([]uint64, 0, len(b))
	for i := range b {
		val, err := PackedRepeatedFixed64Decoder(b[i])
		if err != nil {
			return nil, err
		}
		result = append(result, val.([]uint64)...)
	}
	return result, nil
}

// UnpackedRepeatedSfixed64Decoder 解码repeated sfixed64
var UnpackedRepeatedSfixed64Decoder unpackedRepeatedDecoder = func(b [][]byte) (interface{}, error) {
	result := make([]int64, 0, len(b))
	for i := range b {
		val, err := PackedRepeatedSfixed64Decoder(b[i])
		if err != nil {
			return nil, err
		}
		result = append(result, val.([]int64)...)
	}
	return result, nil
}

// UnpackedRepeatedDoubleDecoder 解码repeated double
var UnpackedRepeatedDoubleDecoder unpackedRepeatedDecoder = func(b [][]byte) (interface{}, error) {
	result := make([]float64, 0, len(b))
	for i := range b {
		val, err := PackedRepeatedDoubleDecoder(b[i])
		if err != nil {
			return nil, err
		}
		result = append(result, val.([]float64)...)
	}
	return result, nil
}

// UnpackedRepeatedStringDecoder 解码repeated string
var UnpackedRepeatedStringDecoder unpackedRepeatedDecoder = func(b [][]byte) (interface{}, error) {
	result := make([]string, 0, len(b))
	for i := range b {
		if len(b[i]) == 0 {
			result = append(result, "")
			continue
		}
		result = append(result, string(b[i]))
	}
	return result, nil
}

// UnpackedRepeatedBytesDecoder 解码repeated bytes
var UnpackedRepeatedBytesDecoder unpackedRepeatedDecoder = func(b [][]byte) (interface{}, error) {
	result := make([][]byte, 0, len(b))
	for i := range b {
		if len(b[i]) == 0 {
			result = append(result, []byte{})
			continue
		}
		result = append(result, b[i])
	}
	return result, nil
}

// UnpackedRepeatedMessageDecoder 解码repeated message
var UnpackedRepeatedMessageDecoder unpackedRepeatedDecoder = func(b [][]byte) (interface{}, error) {
	result := make([]ProtoMessage, 0, len(b))
	for i := range b {
		if len(b[i]) == 0 {
			result = append(result, make(ProtoMessage))
			continue
		}
		msg, err := DecodeBinaryData(b[i])
		if err != nil {
			return nil, err
		}
		result = append(result, msg)
	}
	return result, nil
}

// UnpackedRepeatedFixed32Decoder 解码repeated fixed32
var UnpackedRepeatedFixed32Decoder unpackedRepeatedDecoder = func(b [][]byte) (interface{}, error) {
	result := make([]uint32, 0, len(b))
	for i := range b {
		val, err := PackedRepeatedFixed32Decoder(b[i])
		if err != nil {
			return nil, err
		}
		result = append(result, val.([]uint32)...)
	}
	return result, nil
}

// UnpackedRepeatedSfixed32Decoder 解码repeated sfixed32
var UnpackedRepeatedSfixed32Decoder unpackedRepeatedDecoder = func(b [][]byte) (interface{}, error) {
	result := make([]int32, 0, len(b))
	for i := range b {
		val, err := PackedRepeatedSfixed32Decoder(b[i])
		if err != nil {
			return nil, err
		}
		result = append(result, val.([]int32)...)
	}
	return result, nil
}

// UnpackedRepeatedFloatDecoder 解码repeated float
var UnpackedRepeatedFloatDecoder unpackedRepeatedDecoder = func(b [][]byte) (interface{}, error) {
	result := make([]float32, 0, len(b))
	for i := range b {
		val, err := PackedRepeatedFloatDecoder(b[i])
		if err != nil {
			return nil, err
		}
		result = append(result, val.([]float32)...)
	}
	return result, nil
}
