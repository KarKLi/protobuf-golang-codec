package codec

import (
	"math"
	"math/rand"
	"testing"

	"github.com/KarKLi/protobuf-golang-codec/internal/proto3_test"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
)

var testBin []byte

var testMsg = &proto3_test.Msg{
	I_1:  -rand.Int31(),
	I_2:  rand.Int63(),
	U_3:  rand.Uint32(),
	U_4:  rand.Uint64(),
	S_5:  rand.Int31(),
	S_6:  rand.Int63(),
	B_7:  rand.Intn(2) == 1,
	E_8:  proto3_test.TestEnum(rand.Intn(len(proto3_test.TestEnum_value))),
	F_9:  rand.Uint64(),
	S_10: -rand.Int63(),
	D_11: rand.Float64(),
	S_12: "this is s_12",
	B_13: []byte{1, 2, 3, 4, 5},
	M_14: &proto3_test.Embeeded{
		I_1: -rand.Int31(),
		F_2: rand.Uint64(),
		S_3: "你好",
		F_4: rand.Uint32(),
	},
	F_15: rand.Uint32(),
	S_16: rand.Int31(),
	F_17: -rand.Float32(),
}

var testPackedRepeatedBin []byte

var testPackedRepeatedMsg = &proto3_test.RepeatedMsgWithPacked{
	I_1:  []int32{0, math.MaxInt32, rand.Int31(), -rand.Int31()},
	I_2:  []int64{0, math.MaxInt64, rand.Int63(), -rand.Int63()},
	U_3:  []uint32{0, math.MaxInt32, rand.Uint32()},
	U_4:  []uint64{0, math.MaxInt64, rand.Uint64()},
	S_5:  []int32{0, math.MaxInt32, rand.Int31(), -rand.Int31()},
	S_6:  []int64{0, math.MaxInt64, rand.Int63(), -rand.Int63()},
	B_7:  []bool{true, false, true, true, false},
	E_8:  []proto3_test.TestEnum{proto3_test.TestEnum_ZERO, proto3_test.TestEnum_ONE, proto3_test.TestEnum_TWO},
	F_9:  []uint64{0, math.MaxInt64, rand.Uint64()},
	S_10: []int64{0, math.MaxInt64, rand.Int63(), -rand.Int63()},
	D_11: []float64{0, math.MaxFloat64, rand.Float64(), -rand.Float64()},
	F_12: []uint32{0, math.MaxInt32, rand.Uint32()},
	S_13: []int32{0, math.MaxInt32, rand.Int31(), -rand.Int31()},
	F_14: []float32{0, math.MaxFloat32, rand.Float32(), -rand.Float32()},
}

var testPackedRepeatedDecoderMap = map[protowire.Number]packedRepeatedDecoder{
	1:  PackedRepeatedInt32Decoder,
	2:  PackedRepeatedInt64Decoder,
	3:  PackedRepeatedUint32Decoder,
	4:  PackedRepeatedUint64Decoder,
	5:  PackedRepeatedSint32Decoder,
	6:  PackedRepeatedSint64Decoder,
	7:  PackedRepeatedBoolDecoder,
	8:  PackedRepeatedEnumDecoder,
	9:  PackedRepeatedFixed64Decoder,
	10: PackedRepeatedSfixed64Decoder,
	11: PackedRepeatedDoubleDecoder,
	12: PackedRepeatedFixed32Decoder,
	13: PackedRepeatedSfixed32Decoder,
	14: PackedRepeatedFloatDecoder,
}

var testUnpackedRepeatedBin []byte

var testUnpackedRepeatedMsg = &proto3_test.RepeatedMsgWithUnpacked{
	I_1:  []int32{0, math.MaxInt32, rand.Int31(), -rand.Int31()},
	I_2:  []int64{0, math.MaxInt64, rand.Int63(), -rand.Int63()},
	U_3:  []uint32{0, math.MaxInt32, rand.Uint32()},
	U_4:  []uint64{0, math.MaxInt64, rand.Uint64()},
	S_5:  []int32{0, math.MaxInt32, rand.Int31(), -rand.Int31()},
	S_6:  []int64{0, math.MaxInt64, rand.Int63(), -rand.Int63()},
	B_7:  []bool{true, false, true, true, false},
	E_8:  []proto3_test.TestEnum{proto3_test.TestEnum_ZERO, proto3_test.TestEnum_ONE, proto3_test.TestEnum_TWO},
	F_9:  []uint64{0, math.MaxInt64, rand.Uint64()},
	S_10: []int64{0, math.MaxInt64, rand.Int63(), -rand.Int63()},
	D_11: []float64{0, math.MaxFloat64, rand.Float64(), -rand.Float64()},
	F_12: []uint32{0, math.MaxInt32, rand.Uint32()},
	S_13: []int32{0, math.MaxInt32, rand.Int31(), -rand.Int31()},
	F_14: []float32{0, math.MaxFloat32, rand.Float32(), -rand.Float32()},
	S_15: []string{"", "aaa", "bbbbb", "cc$%^", "123", "1-2+3", "你好"},
	B_16: [][]byte{nil, []byte(""), []byte("aaa"), []byte("bbbbb"), []byte("cc$%^"), []byte("123"), []byte("1-2+3"), []byte("你好")},
	M_17: []*proto3_test.Embeeded{
		{},
		{
			I_1: rand.Int31(),
			F_2: rand.Uint64(),
			S_3: "aa",
			F_4: rand.Uint32(),
		},
		{
			I_1: rand.Int31(),
			F_2: rand.Uint64(),
			S_3: "",
			F_4: rand.Uint32(),
		},
		{
			I_1: -rand.Int31(),
			F_2: rand.Uint64(),
			S_3: "aa",
			F_4: rand.Uint32(),
		},
		{
			I_1: -rand.Int31(),
			F_2: rand.Uint64(),
			S_3: "",
			F_4: rand.Uint32(),
		},
	},
	M_18: map[int32]string{0: "", 1: "", 2: "aa", 3: "你好"},
	M_19: map[string]int32{"": 0, "a": 1, "aa": 2, "你好": 3},
	M_20: map[string]*proto3_test.Embeeded{
		"":  {I_1: rand.Int31()},
		"a": {},
		"aa": {
			I_1: rand.Int31(),
			F_2: rand.Uint64(),
			S_3: "aa",
			F_4: rand.Uint32(),
		},
		"你好": {
			I_1: rand.Int31(),
			F_2: rand.Uint64(),
			S_3: "",
			F_4: rand.Uint32(),
		},
	},
}

var testUnpackedRepeatedDecoderMap = map[protowire.Number]unpackedRepeatedDecoder{
	1:  UnpackedRepeatedInt32Decoder,
	2:  UnpackedRepeatedInt64Decoder,
	3:  UnpackedRepeatedUint32Decoder,
	4:  UnpackedRepeatedUint64Decoder,
	5:  UnpackedRepeatedSint32Decoder,
	6:  UnpackedRepeatedSint64Decoder,
	7:  UnpackedRepeatedBoolDecoder,
	8:  UnpackedRepeatedEnumDecoder,
	9:  UnpackedRepeatedFixed64Decoder,
	10: UnpackedRepeatedSfixed64Decoder,
	11: UnpackedRepeatedDoubleDecoder,
	12: UnpackedRepeatedFixed32Decoder,
	13: UnpackedRepeatedSfixed32Decoder,
	14: UnpackedRepeatedFloatDecoder,
	15: UnpackedRepeatedStringDecoder,
	16: UnpackedRepeatedBytesDecoder,
	17: UnpackedRepeatedMessageDecoder,
}

var testMapDecoderMap = map[protowire.Number]struct {
	k keyDecoder
	v valueDecoder
}{
	18: {k: Int32KeyDecoder, v: StringValueDecoder},
	19: {k: StringKeyDecoder, v: Int32ValueDecoder},
	20: {k: StringKeyDecoder, v: MessageValueDecoder},
}

func initTestData(b *testing.B) {
	var err error
	testBin, err = proto.Marshal(testMsg)
	if err != nil {
		b.Fatalf("can not marshal Msg message, err: %+v", err)
	}
	testPackedRepeatedBin, err = proto.Marshal(testPackedRepeatedMsg)
	if err != nil {
		b.Fatalf("can not marshal RepeatedMsgWithPacked message, err: %+v", err)
	}
	testUnpackedRepeatedBin, err = proto.Marshal(testUnpackedRepeatedMsg)
	if err != nil {
		b.Fatalf("can not marshal RepeatedMsgWithPacked message, err: %+v", err)
	}
}

func BenchmarkNonRepeatedBaseline(b *testing.B) {
	initTestData(b)
	for i := 0; i < b.N; i++ {
		if err := proto.Unmarshal(testBin, &proto3_test.Msg{}); err != nil {
			b.Fatalf("can not unmarshal test proto message, err: %+v", err)
		}
	}
}

func BenchmarkDecodeNonRepeatedData(b *testing.B) {
	initTestData(b)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		m, err := Decode(testBin, NotSort)
		if err != nil {
			b.Fatalf("decode Msg proto message into ProtoMessage struct failed, err: %+v", err)
		}
		_ = m
	}
}

func BenchmarkPackedRepeatedBaseline(b *testing.B) {
	initTestData(b)
	for i := 0; i < b.N; i++ {
		if err := proto.Unmarshal(testPackedRepeatedBin, &proto3_test.RepeatedMsgWithPacked{}); err != nil {
			b.Fatalf("can not unmarshal test proto message, err: %+v", err)
		}
	}
}

func BenchmarkDecodePackedRepeatedData(b *testing.B) {
	initTestData(b)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		m, err := Decode(testPackedRepeatedBin, Asc)
		if err != nil {
			b.Fatalf("decode RepeatedMsgWithPacked proto message into ProtoMessage struct failed, err: %+v", err)
		}
		// 对所有repeated字段继续解码
		for i := range m.Values {
			tag := m.Values[i].tag
			_, err := m.DecodePackedRepeated(tag, testPackedRepeatedDecoderMap[tag])
			if err != nil {
				b.Fatalf("decode RepeatedMsgWithPacked %d field failed, err: %+v", tag, err)
			}
		}
	}
}

func BenchmarkUnpackedRepeatedBaseline(b *testing.B) {
	initTestData(b)
	for i := 0; i < b.N; i++ {
		if err := proto.Unmarshal(testUnpackedRepeatedBin, &proto3_test.RepeatedMsgWithUnpacked{}); err != nil {
			b.Fatalf("can not unmarshal test proto message, err: %+v", err)
		}
	}
}

func BenchmarkDecodeUnpackedRepeatedData(b *testing.B) {
	initTestData(b)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		m, err := Decode(testUnpackedRepeatedBin, Asc)
		if err != nil {
			b.Fatalf("decode RepeatedMsgWithPacked proto message into ProtoMessage struct failed, err: %+v", err)
		}
		// 对所有repeated字段继续解码
		for i := range m.Values {
			tag := m.Values[i].tag
			if tag == 18 || tag == 19 || tag == 20 {
				elems, err := m.DecodeMap(tag, testMapDecoderMap[tag].k, testMapDecoderMap[tag].v)
				if err != nil {
					b.Fatalf("decode RepeatedMsgWithPacked %d field failed, err: %+v", tag, err)
				}
				_, err = FillMapFromProtoMapElem(elems)
				if err != nil {
					b.Fatalf("decode RepeatedMsgWithPacked %d field into map failed, err: %+v", tag, err)
				}
				continue
			}
			_, err := m.DecodeUnpackedRepeated(tag, testUnpackedRepeatedDecoderMap[tag])
			if err != nil {
				b.Fatalf("decode RepeatedMsgWithPacked %d field failed, err: %+v", tag, err)
			}

		}
	}
}
