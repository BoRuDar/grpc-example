package main

import (
	"encoding/json"
	"testing"

	"github.com/BoRuDar/grpc-example/internal/models/api"

	"github.com/golang/protobuf/proto"
)

func TestSize(t *testing.T) {
	payload := &api.Request{
		A:  133333333333.2,
		B:  3.4,
		Op: api.OP_MUL,
	}

	jsonBytes, _ := json.Marshal(payload)
	protoBytes, _ := proto.Marshal(payload)

	t.Logf("JSON: %d bytes", len(jsonBytes))
	t.Logf("Proto: %d bytes", len(protoBytes))
}
