package codec

import (
	"bytes"
	"math"
	"math/rand"
	"reflect"
	"testing"

	"github.com/KarKLi/protobuf-golang-codec/internal/proto3_test"
	"google.golang.org/protobuf/proto"
)

func TestDecodeNonRepeatedData(t *testing.T) {
	testMsg := &proto3_test.Msg{
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
	bin, err := proto.Marshal(testMsg)
	if err != nil {
		t.Fatalf("can not marshal test proto message, err: %+v", err)
	}
	m, err := Decode(bin, NotSort)
	if err != nil {
		t.Fatalf("decode test proto message into map failed, err: %+v", err)
	}
	v1, err := m.GetData(1)
	if err != nil {
		t.Fatalf("can not get tag=1's data, err: %+v", err)
	}
	realI1, err := v1.DecodeInt32()
	if err != nil {
		t.Fatalf("can not parse tag 1, err: %+v", err)
	}
	if realI1 != testMsg.I_1 {
		t.Fatalf("parse result %d != real val %d", realI1, testMsg.I_1)
	}

	v2, err := m.GetData(2)
	if err != nil {
		t.Fatalf("can not get tag=2's data, err: %+v", err)
	}
	realI2, err := v2.DecodeInt64()
	if err != nil {
		t.Fatalf("can not parse tag 2, err: %+v", err)
	}
	if realI2 != testMsg.I_2 {
		t.Fatalf("parse result %d != real val %d", realI1, testMsg.I_1)
	}

	v3, err := m.GetData(3)
	if err != nil {
		t.Fatalf("can not get tag=3's data, err: %+v", err)
	}
	realU3, err := v3.DecodeUint32()
	if err != nil {
		t.Fatalf("can not parse tag 3, err: %+v", err)
	}
	if realU3 != testMsg.U_3 {
		t.Fatalf("parse result %d != real val %d", realU3, testMsg.U_3)
	}

	v4, err := m.GetData(4)
	if err != nil {
		t.Fatalf("can not get tag=4's data, err: %+v", err)
	}
	realU4, err := v4.DecodeUint64()
	if err != nil {
		t.Fatalf("can not parse tag 4, err: %+v", err)
	}
	if realU4 != testMsg.U_4 {
		t.Fatalf("parse result %d != real val %d", realU4, testMsg.U_4)
	}

	v5, err := m.GetData(5)
	if err != nil {
		t.Fatalf("can not get tag=5's data, err: %+v", err)
	}
	realS5, err := v5.DecodeSint32()
	if err != nil {
		t.Fatalf("can not parse tag 5, err: %+v", err)
	}
	if realS5 != testMsg.S_5 {
		t.Fatalf("parse result %d != real val %d", realS5, testMsg.S_5)
	}

	v6, err := m.GetData(6)
	if err != nil {
		t.Fatalf("can not get tag=6's data, err: %+v", err)
	}
	realS6, err := v6.DecodeSint64()
	if err != nil {
		t.Fatalf("can not parse tag 6, err: %+v", err)
	}
	if realS6 != testMsg.S_6 {
		t.Fatalf("parse result %d != real val %d", realS6, testMsg.S_6)
	}

	v7, err := m.GetData(7)
	if err != nil {
		t.Fatalf("can not get tag=7's data, err: %+v", err)
	}
	realB7, err := v7.DecodeBool()
	if err != nil {
		t.Fatalf("can not parse tag 7, err: %+v", err)
	}
	if realB7 != testMsg.B_7 {
		t.Fatalf("parse result %v != real val %v", realB7, testMsg.B_7)
	}

	v8, err := m.GetData(8)
	if err != nil {
		t.Fatalf("can not get tag=8's data, err: %+v", err)
	}
	realE8, err := v8.DecodeEnum()
	if err != nil {
		t.Fatalf("can not parse tag 8, err: %+v", err)
	}
	if proto3_test.TestEnum(realE8) != testMsg.E_8 {
		t.Fatalf("parse result %d != real val %d", realE8, testMsg.E_8)
	}

	v9, err := m.GetData(9)
	if err != nil {
		t.Fatalf("can not get tag=9's data, err: %+v", err)
	}
	realF9, err := v9.DecodeFixed64()
	if err != nil {
		t.Fatalf("can not parse tag 9, err: %+v", err)
	}
	if realF9 != testMsg.F_9 {
		t.Fatalf("parse result %d != real val %d", realF9, testMsg.F_9)
	}

	v10, err := m.GetData(10)
	if err != nil {
		t.Fatalf("can not get tag=10's data, err: %+v", err)
	}
	realF10, err := v10.DecodeSfixed64()
	if err != nil {
		t.Fatalf("can not parse tag 10, err: %+v", err)
	}
	if realF10 != testMsg.S_10 {
		t.Fatalf("parse result %d != real val %d", realF10, testMsg.S_10)
	}

	v11, err := m.GetData(11)
	if err != nil {
		t.Fatalf("can not get tag=11's data, err: %+v", err)
	}
	realD11, err := v11.DecodeDouble()
	if err != nil {
		t.Fatalf("can not parse tag 11, err: %+v", err)
	}
	if realD11 != testMsg.D_11 {
		t.Fatalf("parse result %f != real val %f", realD11, testMsg.D_11)
	}

	v12, err := m.GetData(12)
	if err != nil {
		t.Fatalf("can not get tag=12's data, err: %+v", err)
	}
	realS12, err := v12.DecodeString()
	if err != nil {
		t.Fatalf("can not parse tag 12, err: %+v", err)
	}
	if realS12 != testMsg.S_12 {
		t.Fatalf("parse result %s != real val %s", realS12, testMsg.S_12)
	}

	v13, err := m.GetData(13)
	if err != nil {
		t.Fatalf("can not get tag=13's data, err: %+v", err)
	}
	realB13, err := v13.DecodeBytes()
	if err != nil {
		t.Fatalf("can not parse tag 13, err: %+v", err)
	}
	if !bytes.Equal(realB13, testMsg.B_13) {
		t.Fatalf("parse result %v != real val %v", realB13, testMsg.B_13)
	}

	v14, err := m.GetData(14)
	if err != nil {
		t.Fatalf("can not get tag=14's data, err: %+v", err)
	}
	realM14, err := v14.DecodeEmbeddedMsg(NotSort)
	if err != nil {
		t.Fatalf("can not parse tag 14, err: %+v", err)
	}
	checkEmbeededMsgEqual(t, realM14, testMsg.M_14)
	v15, err := m.GetData(15)
	if err != nil {
		t.Fatalf("can not get tag=15's data, err: %+v", err)
	}
	realF15, err := v15.DecodeFixed32()
	if err != nil {
		t.Fatalf("can not parse tag 15, err: %+v", err)
	}
	if realF15 != testMsg.F_15 {
		t.Fatalf("parse result %d != real val %d", realF15, testMsg.F_15)
	}
	v16, err := m.GetData(16)
	if err != nil {
		t.Fatalf("can not get tag=16's data, err: %+v", err)
	}
	realS16, err := v16.DecodeSfixed32()
	if err != nil {
		t.Fatalf("can not parse tag 16, err: %+v", err)
	}
	if realS16 != testMsg.S_16 {
		t.Fatalf("parse result %d != real val %d", realS16, testMsg.S_16)
	}
	v17, err := m.GetData(17)
	if err != nil {
		t.Fatalf("can not get tag=17's data, err: %+v", err)
	}
	realF17, err := v17.DecodeFloat()
	if err != nil {
		t.Fatalf("can not parse tag 17, err: %+v", err)
	}
	if realF17 != testMsg.F_17 {
		t.Fatalf("parse result %f != real val %f", realF17, testMsg.F_17)
	}
}

func TestDecodePackedRepeatedData(t *testing.T) {
	testMsg := &proto3_test.RepeatedMsgWithPacked{
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
	bin, err := proto.Marshal(testMsg)
	if err != nil {
		t.Fatalf("can not marshal test proto message, err: %+v", err)
	}
	m, err := Decode(bin, NotSort)
	if err != nil {
		t.Fatalf("decode test proto message into map failed, err: %+v", err)
	}

	realI1, err := m.DecodePackedRepeated(1, PackedRepeatedInt32Decoder)
	if err != nil {
		t.Fatalf("can not parse tag 1, err: %+v", err)
	}
	if !reflect.DeepEqual(realI1.([]int32), testMsg.I_1) {
		t.Fatalf("parse result %v != real val %v", realI1.([]int32), testMsg.I_1)
	}
	realI2, err := m.DecodePackedRepeated(2, PackedRepeatedInt64Decoder)
	if err != nil {
		t.Fatalf("can not parse tag 2, err: %+v", err)
	}
	if !reflect.DeepEqual(realI2.([]int64), testMsg.I_2) {
		t.Fatalf("parse result %v != real val %v", realI2.([]int64), testMsg.I_2)
	}
	realU3, err := m.DecodePackedRepeated(3, PackedRepeatedUint32Decoder)
	if err != nil {
		t.Fatalf("can not parse tag 3, err: %+v", err)
	}
	if !reflect.DeepEqual(realU3.([]uint32), testMsg.U_3) {
		t.Fatalf("parse result %v != real val %v", realU3.([]uint32), testMsg.U_3)
	}
	realU4, err := m.DecodePackedRepeated(4, PackedRepeatedUint64Decoder)
	if err != nil {
		t.Fatalf("can not parse tag 4, err: %+v", err)
	}
	if !reflect.DeepEqual(realU4.([]uint64), testMsg.U_4) {
		t.Fatalf("parse result %v != real val %v", realU4.([]uint64), testMsg.U_4)
	}
	realS5, err := m.DecodePackedRepeated(5, PackedRepeatedSint32Decoder)
	if err != nil {
		t.Fatalf("can not parse tag 5, err: %+v", err)
	}
	if !reflect.DeepEqual(realS5.([]int32), testMsg.S_5) {
		t.Fatalf("parse result %v != real val %v", realS5.([]int32), testMsg.S_5)
	}
	realS6, err := m.DecodePackedRepeated(6, PackedRepeatedSint64Decoder)
	if err != nil {
		t.Fatalf("can not parse tag 6, err: %+v", err)
	}
	if !reflect.DeepEqual(realS6.([]int64), testMsg.S_6) {
		t.Fatalf("parse result %v != real val %v", realS6.([]int64), testMsg.S_6)
	}
	realB7, err := m.DecodePackedRepeated(7, PackedRepeatedBoolDecoder)
	if err != nil {
		t.Fatalf("can not parse tag 7, err: %+v", err)
	}
	if !reflect.DeepEqual(realB7.([]bool), testMsg.B_7) {
		t.Fatalf("parse result %v != real val %v", realB7.([]bool), testMsg.B_7)
	}
	realE8, err := m.DecodePackedRepeated(8, PackedRepeatedEnumDecoder)
	if err != nil {
		t.Fatalf("can not parse tag 8, err: %+v", err)
	}
	tempE8 := make([]int32, 0, len(testMsg.E_8))
	for _, e := range testMsg.E_8 {
		tempE8 = append(tempE8, int32(e))
	}
	if !reflect.DeepEqual(realE8.([]int32), tempE8) {
		t.Fatalf("parse result %v != real val %v", realE8.([]int32), tempE8)
	}
	realF9, err := m.DecodePackedRepeated(9, PackedRepeatedFixed64Decoder)
	if err != nil {
		t.Fatalf("can not parse tag 9, err: %+v", err)
	}
	if !reflect.DeepEqual(realF9.([]uint64), testMsg.F_9) {
		t.Fatalf("parse result %v != real val %v", realF9.([]uint64), testMsg.F_9)
	}
	realS10, err := m.DecodePackedRepeated(10, PackedRepeatedSfixed64Decoder)
	if err != nil {
		t.Fatalf("can not parse tag 10, err: %+v", err)
	}
	if !reflect.DeepEqual(realS10.([]int64), testMsg.S_10) {
		t.Fatalf("parse result %v != real val %v", realS10.([]int64), testMsg.S_10)
	}
	realD11, err := m.DecodePackedRepeated(11, PackedRepeatedDoubleDecoder)
	if err != nil {
		t.Fatalf("can not parse tag 11, err: %+v", err)
	}
	if !reflect.DeepEqual(realD11.([]float64), testMsg.D_11) {
		t.Fatalf("parse result %v != real val %v", realD11.([]float64), testMsg.D_11)
	}
	realF12, err := m.DecodePackedRepeated(12, PackedRepeatedFixed32Decoder)
	if err != nil {
		t.Fatalf("can not parse tag 12, err: %+v", err)
	}
	if !reflect.DeepEqual(realF12.([]uint32), testMsg.F_12) {
		t.Fatalf("parse result %v != real val %v", realF12.([]uint32), testMsg.F_12)
	}
	realS13, err := m.DecodePackedRepeated(13, PackedRepeatedSfixed32Decoder)
	if err != nil {
		t.Fatalf("can not parse tag 13, err: %+v", err)
	}
	if !reflect.DeepEqual(realS13.([]int32), testMsg.S_13) {
		t.Fatalf("parse result %v != real val %v", realS13.([]int32), testMsg.S_13)
	}
	realF14, err := m.DecodePackedRepeated(14, PackedRepeatedFloatDecoder)
	if err != nil {
		t.Fatalf("can not parse tag 14, err: %+v", err)
	}
	if !reflect.DeepEqual(realF14.([]float32), testMsg.F_14) {
		t.Fatalf("parse result %v != real val %v", realF14.([]float32), testMsg.F_14)
	}
}

func TestDecodeUnpackedRepeatedData(t *testing.T) {
	testMsg := &proto3_test.RepeatedMsgWithUnpacked{
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
	bin, err := proto.Marshal(testMsg)
	if err != nil {
		t.Fatalf("can not marshal test proto message, err: %+v", err)
	}
	m, err := Decode(bin, NotSort)
	if err != nil {
		t.Fatalf("decode test proto message into map failed, err: %+v", err)
	}

	realI1, err := m.DecodeUnpackedRepeated(1, UnpackedRepeatedInt32Decoder)
	if err != nil {
		t.Fatalf("can not parse tag 1, err: %+v", err)
	}
	if !reflect.DeepEqual(realI1.([]int32), testMsg.I_1) {
		t.Fatalf("parse result %v != real val %v", realI1.([]int32), testMsg.I_1)
	}
	realI2, err := m.DecodeUnpackedRepeated(2, UnpackedRepeatedInt64Decoder)
	if err != nil {
		t.Fatalf("can not parse tag 2, err: %+v", err)
	}
	if !reflect.DeepEqual(realI2.([]int64), testMsg.I_2) {
		t.Fatalf("parse result %v != real val %v", realI2.([]int64), testMsg.I_2)
	}
	realU3, err := m.DecodeUnpackedRepeated(3, UnpackedRepeatedUint32Decoder)
	if err != nil {
		t.Fatalf("can not parse tag 3, err: %+v", err)
	}
	if !reflect.DeepEqual(realU3.([]uint32), testMsg.U_3) {
		t.Fatalf("parse result %v != real val %v", realU3.([]uint32), testMsg.U_3)
	}
	realU4, err := m.DecodeUnpackedRepeated(4, UnpackedRepeatedUint64Decoder)
	if err != nil {
		t.Fatalf("can not parse tag 4, err: %+v", err)
	}
	if !reflect.DeepEqual(realU4.([]uint64), testMsg.U_4) {
		t.Fatalf("parse result %v != real val %v", realU4.([]uint64), testMsg.U_4)
	}
	realS5, err := m.DecodeUnpackedRepeated(5, UnpackedRepeatedSint32Decoder)
	if err != nil {
		t.Fatalf("can not parse tag 5, err: %+v", err)
	}
	if !reflect.DeepEqual(realS5.([]int32), testMsg.S_5) {
		t.Fatalf("parse result %v != real val %v", realS5.([]int32), testMsg.S_5)
	}
	realS6, err := m.DecodeUnpackedRepeated(6, UnpackedRepeatedSint64Decoder)
	if err != nil {
		t.Fatalf("can not parse tag 6, err: %+v", err)
	}
	if !reflect.DeepEqual(realS6.([]int64), testMsg.S_6) {
		t.Fatalf("parse result %v != real val %v", realS6.([]int64), testMsg.S_6)
	}
	realB7, err := m.DecodeUnpackedRepeated(7, UnpackedRepeatedBoolDecoder)
	if err != nil {
		t.Fatalf("can not parse tag 7, err: %+v", err)
	}
	if !reflect.DeepEqual(realB7.([]bool), testMsg.B_7) {
		t.Fatalf("parse result %v != real val %v", realB7.([]bool), testMsg.B_7)
	}
	realE8, err := m.DecodeUnpackedRepeated(8, UnpackedRepeatedEnumDecoder)
	if err != nil {
		t.Fatalf("can not parse tag 8, err: %+v", err)
	}
	tempE8 := make([]int32, 0, len(testMsg.E_8))
	for _, e := range testMsg.E_8 {
		tempE8 = append(tempE8, int32(e))
	}
	if !reflect.DeepEqual(realE8.([]int32), tempE8) {
		t.Fatalf("parse result %v != real val %v", realE8.([]int32), tempE8)
	}
	realF9, err := m.DecodeUnpackedRepeated(9, UnpackedRepeatedFixed64Decoder)
	if err != nil {
		t.Fatalf("can not parse tag 9, err: %+v", err)
	}
	if !reflect.DeepEqual(realF9.([]uint64), testMsg.F_9) {
		t.Fatalf("parse result %v != real val %v", realF9.([]uint64), testMsg.F_9)
	}
	realS10, err := m.DecodeUnpackedRepeated(10, UnpackedRepeatedSfixed64Decoder)
	if err != nil {
		t.Fatalf("can not parse tag 10, err: %+v", err)
	}
	if !reflect.DeepEqual(realS10.([]int64), testMsg.S_10) {
		t.Fatalf("parse result %v != real val %v", realS10.([]int64), testMsg.S_10)
	}
	realD11, err := m.DecodeUnpackedRepeated(11, UnpackedRepeatedDoubleDecoder)
	if err != nil {
		t.Fatalf("can not parse tag 11, err: %+v", err)
	}
	if !reflect.DeepEqual(realD11.([]float64), testMsg.D_11) {
		t.Fatalf("parse result %v != real val %v", realD11.([]float64), testMsg.D_11)
	}
	realF12, err := m.DecodeUnpackedRepeated(12, UnpackedRepeatedFixed32Decoder)
	if err != nil {
		t.Fatalf("can not parse tag 12, err: %+v", err)
	}
	if !reflect.DeepEqual(realF12.([]uint32), testMsg.F_12) {
		t.Fatalf("parse result %v != real val %v", realF12.([]uint32), testMsg.F_12)
	}
	realS13, err := m.DecodeUnpackedRepeated(13, UnpackedRepeatedSfixed32Decoder)
	if err != nil {
		t.Fatalf("can not parse tag 13, err: %+v", err)
	}
	if !reflect.DeepEqual(realS13.([]int32), testMsg.S_13) {
		t.Fatalf("parse result %v != real val %v", realS13.([]int32), testMsg.S_13)
	}
	realF14, err := m.DecodeUnpackedRepeated(14, UnpackedRepeatedFloatDecoder)
	if err != nil {
		t.Fatalf("can not parse tag 14, err: %+v", err)
	}
	if !reflect.DeepEqual(realF14.([]float32), testMsg.F_14) {
		t.Fatalf("parse result %v != real val %v", realF14.([]float32), testMsg.F_14)
	}
	realS15, err := m.DecodeUnpackedRepeated(15, UnpackedRepeatedStringDecoder)
	if err != nil {
		t.Fatalf("can not parse tag 15, err: %+v", err)
	}
	if !reflect.DeepEqual(realS15.([]string), testMsg.S_15) {
		t.Fatalf("parse result %v != real val %v", realS15.([]string), testMsg.S_15)
	}
	realB16, err := m.DecodeUnpackedRepeated(16, UnpackedRepeatedBytesDecoder)
	if err != nil {
		t.Fatalf("can not parse tag 16, err: %+v", err)
	}
	realB16Arr := realB16.([][]byte)
	if len(realB16Arr) != len(testMsg.B_16) {
		t.Fatalf("parse result %v != real val %v", realB16.([][]byte), testMsg.B_16)
	}
	for i := range realB16Arr {
		if !bytes.Equal(realB16Arr[i], testMsg.B_16[i]) {
			t.Fatalf("[idx %d] parse result %v != real val %v", i, realB16.([][]byte), testMsg.B_16)
		}
	}
	realM17, err := m.DecodeUnpackedRepeated(17, UnpackedRepeatedMessageDecoder)
	if err != nil {
		t.Fatalf("can not parse tag 17, err: %+v", err)
	}
	realM17Arr := realM17.([]ProtoMessage)
	for i, msg := range testMsg.M_17 {
		checkEmbeededMsgEqual(t, realM17Arr[i], msg)
	}
	realM18, err := m.DecodeMap(18, Int32KeyDecoder, StringValueDecoder)
	if err != nil {
		t.Fatalf("can not parse tag 18, err: %+v", err)
	}
	realM18ReflectMap, err := FillMapFromProtoMapElem(realM18)
	if err != nil {
		t.Fatalf("can not convert tag 18 proto map elems into reflect map, err: %+v", err)
	}
	realM18Map := realM18ReflectMap.Interface().(map[int32]string)
	if !reflect.DeepEqual(realM18Map, testMsg.M_18) {
		t.Fatalf("parse result %v != real val %v", realM18Map, testMsg.M_18)
	}
	realM19, err := m.DecodeMap(19, StringKeyDecoder, Int32ValueDecoder)
	if err != nil {
		t.Fatalf("can not parse tag 19, err: %+v", err)
	}
	realM19ReflectMap, err := FillMapFromProtoMapElem(realM19)
	if err != nil {
		t.Fatalf("can not convert tag 19 proto map elems into reflect map, err: %+v", err)
	}
	realM19Map := realM19ReflectMap.Interface().(map[string]int32)
	if !reflect.DeepEqual(realM19Map, testMsg.M_19) {
		t.Fatalf("parse result %v != real val %v", realM19Map, testMsg.M_19)
	}
	realM20, err := m.DecodeMap(20, StringKeyDecoder, MessageValueDecoder)
	if err != nil {
		t.Fatalf("can not parse tag 20, err: %+v", err)
	}
	realM20ReflectMap, err := FillMapFromProtoMapElem(realM20)
	if err != nil {
		t.Fatalf("can not convert tag 20 proto map elems into reflect map, err: %+v", err)
	}
	realM20Map := realM20ReflectMap.Interface().(map[string]ProtoMessage)
	if len(realM20Map) != len(testMsg.M_20) {
		t.Fatalf("parse result %v != real val %v", realM20Map, testMsg.M_20)
	}
	for k, v := range realM20Map {
		checkEmbeededMsgEqual(t, v, testMsg.M_20[k])
	}
}

func checkEmbeededMsgEqual(t *testing.T, m1 ProtoMessage, m2 *proto3_test.Embeeded) {
	v1, err := m1.GetData(1)
	if err != nil {
		t.Fatalf("can not get tag=1's data, err: %+v", err)
	}
	realM14I1, err := v1.DecodeInt32()
	if err != nil {
		t.Fatalf("can not parse tag 14_tag 1, err: %+v", err)
	}
	if realM14I1 != m2.I_1 {
		t.Fatalf("parse result %d != real val %d", realM14I1, m2.I_1)
	}
	v2, err := m1.GetData(2)
	if err != nil {
		t.Fatalf("can not get tag=2's data, err: %+v", err)
	}
	realM14F2, err := v2.DecodeFixed64()
	if err != nil {
		t.Fatalf("can not parse tag 14_tag 2, err: %+v", err)
	}
	if realM14F2 != m2.F_2 {
		t.Fatalf("parse result %d != real val %d", realM14F2, m2.F_2)
	}
	v3, err := m1.GetData(3)
	if err != nil {
		t.Fatalf("can not get tag=3's data, err: %+v", err)
	}
	realM14S3, err := v3.DecodeString()
	if err != nil {
		t.Fatalf("can not parse tag 14_tag 3, err: %+v", err)
	}
	if realM14S3 != m2.S_3 {
		t.Fatalf("parse result %s != real val %s", realM14S3, m2.S_3)
	}
	v4, err := m1.GetData(4)
	if err != nil {
		t.Fatalf("can not get tag=4's data, err: %+v", err)
	}
	realM14F4, err := v4.DecodeFixed32()
	if err != nil {
		t.Fatalf("can not parse tag 14_tag 4, err: %+v", err)
	}
	if realM14F4 != m2.F_4 {
		t.Fatalf("parse result %d != real val %d", realM14F4, m2.F_4)
	}
}
