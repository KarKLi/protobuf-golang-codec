# Protobuf wire format data decode library

This library can decode protobuf wire format data without any .proto file.
## Usage
just:
```go
import (
  "github.com/KarKLi/protobuf-golang-codec"
  "fmt"
)

func main() {
  var wireData []byte
  // fill data into wireData variable...
  msg, err := Decode(wireData)
  if err != nil {
    // err handle
  }
  for tag, v := range msg {
    // and you can use DecodeXXX to decode real data
    // for example, tag == 1 is a uint64 data
    if tag == 1 {
      data, err := v.DecodeUint64()
      if err != nil {
        // error handle
      }
      fmt.Printf("tag=%d, value=%d\n", tag, data)
    }
  }
  return 0
}
```
And there is a specific method called `DecodeMap` which returns a `ProtoMapElem` array, caller can call `FillMapFromProtoMapElem` function to transfer it into `reflect.Value`. After getting this value, caller can just call `reflect.Value`'s method `Interface()` for getting the interface{} (underlying type is the real map type), for example:
```go
import (
  "github.com/KarKLi/protobuf-golang-codec"
  "fmt"
)

func main() {
  var wireData []byte
  // fill data into wireData variable...
  msg, err := Decode(wireData)
  if err != nil {
    // err handle
  }
  for tag, v := range msg {
    // and you can use DecodeXXX to decode real data
    // for example, tag == 1's type is map<int32,string>.
    if tag == 1 {
      mapElems, err := v.DecodeMap()
      if err != nil {
        // error handle
      }
      mapRefVal, err := FillMapFromProtoMapElem(mapElems)
      if err != nil {
        // error handle
      }
      m := mapRefVal.Interface().(map[int32]string) // omit type assertion
      fmt.Printf("tag=%d, value=%+v\n", tag, m)
    }
  }
  return 0
}
```
For more usage and example, just check the `codec_test.go` for decode, parse and assert.

## Benchmark
```
goos: linux
goarch: amd64
pkg: github.com/KarKLi/protobuf-golang-codec
cpu: AMD EPYC 7K62 48-Core Processor
BenchmarkNonRepeatedBaseline
BenchmarkNonRepeatedBaseline-8          	 1499984	       799.0 ns/op	     304 B/op	       5 allocs/op
BenchmarkDecodeNonRepeatedData
BenchmarkDecodeNonRepeatedData-8        	 1274136	       957.5 ns/op	    1696 B/op	      17 allocs/op
BenchmarkPackedRepeatedBaseline
BenchmarkPackedRepeatedBaseline-8       	  752754	      1487 ns/op	     677 B/op	      15 allocs/op
BenchmarkDecodePackedRepeatedData
BenchmarkDecodePackedRepeatedData-8     	  256388	      4605 ns/op	    2144 B/op	      76 allocs/op
BenchmarkUnpackedRepeatedBaseline
BenchmarkUnpackedRepeatedBaseline-8     	  763533	      1499 ns/op	     677 B/op	      15 allocs/op
BenchmarkDecodeUnpackedRepeatedData
BenchmarkDecodeUnpackedRepeatedData-8   	   16624	     72084 ns/op	   71776 B/op	     933 allocs/op
PASS
ok  	github.com/KarKLi/protobuf-golang-codec	9.644s
```
Benchmark code can be referenced at `codec_benchmark_test.go`.
