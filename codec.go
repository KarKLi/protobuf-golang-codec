package codec

import (
	"errors"
	"fmt"

	"google.golang.org/protobuf/encoding/protowire"
)

var (
	ErrTypeMismatch     = errors.New("proto defined type not equal to expected type")
	ErrAssertTypeFailed = errors.New("assert value type failed")
)

type ProtoValue struct {
	_type protowire.Type
	/*
		由type的类型决定，可能为：

		uint32（Fixed32Type）

		uint64（VarintType/Fixed64Type）

		[]byte（BytesType）

		[][]byte（BytesType&&repeated field）
	*/
	val interface{}
	// raw val的字节表示法
	raw []byte
}

type ProtoMessage map[protowire.Number]ProtoValue

/* Decode 解析proto二进制流数据*/
func Decode(b []byte) (ProtoMessage, error) {
	m := make(map[protowire.Number]ProtoValue)
	for len(b) > 0 {
		var n int
		num, typ, n := protowire.ConsumeTag(b)
		if n < 0 {
			return nil, protowire.ParseError(n)
		}
		b = b[n:]
		var val interface{}
		var raw []byte
		switch typ {
		case protowire.VarintType:
			val, n = protowire.ConsumeVarint(b)
			raw = b[:n]
		case protowire.Fixed32Type:
			val, n = protowire.ConsumeFixed32(b)
			raw = b[:n]
		case protowire.Fixed64Type:
			val, n = protowire.ConsumeFixed64(b)
			raw = b[:n]
		case protowire.BytesType:
			val, n = protowire.ConsumeBytes(b)
			raw = val.([]byte)
		default:
			return nil, fmt.Errorf("not support proto data type %d", typ)
		}
		if n < 0 {
			return nil, protowire.ParseError(n)
		}
		/*
			如果m[num]已经存在，则数据是repeated的。因为是LEN类型数据，所以val的长度就是payload的长度（全部数据）
		*/
		if ele, exist := m[num]; exist {
			ele._type = protowire.BytesType // 强制为BytesType（因为是packed=true的数据）
			switch oldVal := ele.val.(type) {
			case []byte:
				// 这是第二个元素，且第一个元素已经是BytesType
				ele.val = [][]byte{oldVal, raw}
			case [][]byte:
				// 这是至少第三个元素
				ele.val = append(oldVal, raw)
			default:
				// 第二个元素，且第一个元素为非BytesType的情况
				ele.val = [][]byte{ele.raw, raw}
			}
			m[num] = ele
			b = b[n:]
			continue
		}
		m[num] = ProtoValue{_type: typ, val: val, raw: raw}
		b = b[n:]
	}
	return m, nil
}
