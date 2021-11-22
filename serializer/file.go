package serializer

import (
	"fmt"
	"io/ioutil"

	"github.com/golang/protobuf/proto"
)

// WriteProtobufToBinaryFile writes protocol buffer message to a binary file
func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto to binary: %w", err)
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("cannot write binary to a file: %w", err)
	}
	return nil
}

// ReadProtobufFromBinaryFile ...
func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("cannot read binary file %w", err)
	}
	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("cannot unmarshal binary to proto message %w", err)
	}
	return nil
}

// WriteProtobufToJSONFile writes protocol buffer message to JSON file
func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	data, err := ProtobufToJSON(message)
	if err != nil {
		return fmt.Errorf("cannot marshal proto message to json file %w", err)
	}
	err = ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("cannot write JSON data to file %w", err)
	}
	return nil
}
