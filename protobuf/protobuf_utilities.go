package protobuf

import (
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

// MarshalToFile marshals a protobuf message to a file.
func MarshalToFile(pb proto.Message, filename string) error {
	data, err := proto.Marshal(pb)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}

// UnmarshalFromFile unmarshals a protobuf message from a file.
func UnmarshalFromFile(pb proto.Message, filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	return proto.Unmarshal(data, pb)
}

// MarshalToBytes marshals a protobuf message to a byte slice.
func MarshalToBytes(pb proto.Message) ([]byte, error) {
	return proto.Marshal(pb)
}

// UnmarshalFromBytes unmarshals a protobuf message from a byte slice.
func UnmarshalFromBytes(pb proto.Message, data []byte) error {
	return proto.Unmarshal(data, pb)
}

// Example usage:
//
// message := &MyProtoMessage{}
// err := protobuf.MarshalToFile(message, "message.bin")
// if err != nil {
//     log.Fatal(err)
// }
//
// var newMessage MyProtoMessage
// err = protobuf.UnmarshalFromFile(&newMessage, "message.bin")
// if err != nil {
//     log.Fatal(err)
// }
