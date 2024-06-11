package codec

import (
	"errors"
	"fmt"
	"sort"

	"google.golang.org/protobuf/encoding/protowire"
)

var (
	ErrTypeMismatch        = errors.New("proto defined type not equal to expected type")
	ErrAssertTypeFailed    = errors.New("assert value type failed")
	ErrDataNotRepeatedData = errors.New("expected repeated data, got singular data")
	ErrDataNotSingularData = errors.New("expected singular data, got repeated data")
)

type ProtoValue struct {
	_type protowire.Type
	/*
		由type的类型决定，可能为：

		uint32（Fixed32Type）

		uint64（VarintType/Fixed64Type）

		[]byte（BytesType）
	*/
	val interface{}
	// tag 该字段的实际tag
	tag protowire.Number
}

type ProtoMessage struct {
	Values   []ProtoValue
	sortType MessageSortType
}

// GetRepeatedData 返回ProtoMessage中所有满足传入tag的底层数据索引
func (p *ProtoMessage) GetRepeatedData(tag protowire.Number) ([]int, error) {
	idxs := make([]int, 0, len(p.Values))
	if p.sortType != NotSort {
		// 二分法寻找
		idx := sort.Search(len(p.Values), func(i int) bool {
			return int32(p.Values[i].tag) >= int32(tag)
		})
		if idx == -1 {
			return nil, nil
		}
		var i int
		for i := idx; i < len(p.Values); i++ {
			if p.Values[i].tag != tag {
				break
			}
		}
		if len(p.Values[idx:i]) == 1 {
			return nil, ErrDataNotRepeatedData
		}
		for j := idx; j < i; j++ {
			idxs = append(idxs, j)
		}
		return idxs, nil
	}
	for i := 0; i < len(p.Values); i++ {
		if p.Values[i].tag == tag {
			idxs = append(idxs, i)
		}
	}
	return idxs, nil
}

func (p *ProtoMessage) GetData(tag protowire.Number) (ProtoValue, error) {
	if p.sortType != NotSort {
		// 二分法寻找
		idx := sort.Search(len(p.Values), func(i int) bool {
			return int32(p.Values[i].tag) >= int32(tag)
		})
		if idx == -1 {
			return ProtoValue{}, nil
		}
		if idx+1 < len(p.Values) && p.Values[idx+1].tag == tag {
			return ProtoValue{}, ErrDataNotSingularData
		}
	}
	// 退化为顺序查找
	for i := 0; i < len(p.Values); i++ {
		if p.Values[i].tag == tag {
			return p.Values[i], nil
		}
	}
	return ProtoValue{}, nil
}

type MessageSortType int

const (
	// NotSort 不排序
	NotSort MessageSortType = iota
	// Asc 升序
	Asc
	// Desc 降序
	Desc
)

// Decode 解析proto二进制流数据
func Decode(b []byte, sortType MessageSortType) (ProtoMessage, error) {
	m := ProtoMessage{
		Values:   make([]ProtoValue, 0, 16),
		sortType: sortType,
	}
	for len(b) > 0 {
		var n int
		num, typ, n := protowire.ConsumeTag(b)
		if n < 0 {
			return ProtoMessage{}, protowire.ParseError(n)
		}
		b = b[n:]
		var val interface{}
		switch typ {
		case protowire.VarintType:
			val, n = protowire.ConsumeVarint(b)
		case protowire.Fixed32Type:
			val, n = protowire.ConsumeFixed32(b)
		case protowire.Fixed64Type:
			val, n = protowire.ConsumeFixed64(b)
		case protowire.BytesType:
			val, n = protowire.ConsumeBytes(b)
		default:
			return ProtoMessage{}, fmt.Errorf("not support proto data type %d", typ)
		}
		if n < 0 {
			return ProtoMessage{}, protowire.ParseError(n)
		}
		m.Values = append(m.Values, ProtoValue{_type: typ, val: val, tag: num})
		b = b[n:]
	}
	switch sortType {
	case NotSort:
	case Asc:
		sort.Slice(m.Values, func(i, j int) bool {
			return m.Values[i].tag < m.Values[j].tag
		})
	case Desc:
		sort.Slice(m.Values, func(i, j int) bool {
			return m.Values[i].tag > m.Values[j].tag
		})
	}
	return m, nil
}
