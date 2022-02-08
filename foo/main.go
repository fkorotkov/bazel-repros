package main

import (
	"github.com/fkorotkov/bazel-repros/proto/foo"
	"google.golang.org/protobuf/encoding/protojson"
)

func main() {
	f := foo.Foo{}
	err := protojson.Unmarshal([]byte("{\"timestamp\": {}}}"), &f)
	if err != nil {
		return
	}
	println(f.Timestamp.AsTime())
}
