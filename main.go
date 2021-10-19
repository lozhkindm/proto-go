package main

import (
	"fmt"
	"github.com/lozhkindm/proto-go/src/simplepb"
	"google.golang.org/protobuf/proto"
	"io/ioutil"
	"log"
)

func main() {
	sm := doSimple()

	readAndWrite(sm)
}

func readAndWrite(sm proto.Message) {
	sm2 := &simplepb.SimpleMessage{}

	if err := writeToFile("simple.bin", sm); err != nil {
		panic(err)
	}

	if err := readFromFile("simple.bin", sm2); err != nil {
		panic(err)
	}
}

func readFromFile(filename string, pb proto.Message) error {
	bytes, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatalln("Can't read the file", err)
		return err
	}

	if err := proto.Unmarshal(bytes, pb); err != nil {
		log.Fatalln("Can't unmarshal bytes", err)
		return err
	}

	return nil
}

func writeToFile(filename string, pb proto.Message) error {
	bytes, err := proto.Marshal(pb)

	if err != nil {
		log.Fatalln("Can't serialise to bytes", err)
		return err
	}

	if err := ioutil.WriteFile(filename, bytes, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}

	fmt.Println("Data has been written")

	return nil
}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         100500,
		IsSimple:   true,
		Name:       "testing name",
		SampleList: []int32{100, 500},
	}

	sm.Name = "New name"

	return &sm
}
