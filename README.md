# Protobuf wire format data decode library

This library can decode protobuf wire format data without any .proto file, just:
```go
import (
  "github.com/KarKLi/protobuf-golang-codec"
)

func main() {
  var wireData []byte
  // fill data into wireData variable...
  msg, err := DecodeBinaryData(wireData)
  if err != nil {
    // err handle
  }
  for k,v := range msg {
    fmt.Printf("tag=%d, value=%+v\n", k, v)
  }
  return 0
}
```
