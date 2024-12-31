package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/golang/protobuf/proto"
	"github.com/jinlongchen/golang-utilities/protobuf"
)

func main() {
	var (
		inputFile   string
		outputFile  string
		messageType string
	)

	flag.StringVar(&inputFile, "input", "", "Input proto file")
	flag.StringVar(&outputFile, "output", "", "Output proto file")
	flag.StringVar(&messageType, "type", "", "Protobuf message type")
	flag.Parse()

	if inputFile == "" || outputFile == "" || messageType == "" {
		flag.Usage()
		os.Exit(1)
	}

	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Failed to read input file: %v", err)
	}

	var message proto.Message
	switch messageType {
	case "MyProtoMessage":
		message = &MyProtoMessage{}
	// Add more cases here for different message types
	default:
		log.Fatalf("Unknown message type: %s", messageType)
	}

	err = proto.Unmarshal(data, message)
	if err != nil {
		log.Fatalf("Failed to unmarshal proto message: %v", err)
	}

	err = protobuf.MarshalToFile(message, outputFile)
	if err != nil {
		log.Fatalf("Failed to write output file: %v", err)
	}

	fmt.Printf("Successfully processed proto file: %s\n", outputFile)
}
