package codec

import (
	"math/rand"
	"testing"

	"github.com/KarKLi/protobuf-golang-codec/internal/proto3_test"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/anypb"
)

var bin []byte

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

func initTestData(b *testing.B) {
	var err error
	bin, err = proto.Marshal(testMsg)
	if err != nil {
		b.Fatalf("can not marshal test proto message, err: %+v", err)
	}
}

func BenchmarkBaseline(b *testing.B) {
	initTestData(b)
	for i := 0; i < b.N; i++ {
		anypbObj := &anypb.Any{
			TypeUrl: "Msg",
			Value:   bin,
		}
		if _, err := anypbObj.UnmarshalNew(); err != nil {
			b.Fatalf("can not unmarshal test proto message, err: %+v", err)
		}
	}
}

func BenchmarkDecodeNonRepeatedData(b *testing.B) {
	initTestData(b)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		m, err := Decode(bin, NotSort)
		if err != nil {
			b.Fatalf("decode test proto message into map failed, err: %+v", err)
		}
		_ = m
	}
}
