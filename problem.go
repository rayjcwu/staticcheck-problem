package problem

import proto "github.com/golang/protobuf/proto"

func Marshal(msg proto.Message) ([]byte, error) {
        return proto.Marshal(msg)
}
