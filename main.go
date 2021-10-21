package main

import (
	"fmt"
	"github.com/lozhkindm/proto-go/src/addresspb"
	"github.com/lozhkindm/proto-go/src/complexpb"
	"github.com/lozhkindm/proto-go/src/enumpb"
	"github.com/lozhkindm/proto-go/src/simplepb"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io/ioutil"
	"log"
)

func main() {
	sm := doSimple()

	readAndWrite(sm)
	toAndFromJSON(sm)

	enum()
	complx()
	address()
}

func address() {
	ab := addresspb.AddressBook{
		People: []*addresspb.Person{
			{
				Name:  "P1",
				Id:    1,
				Email: "p1@fia.com",
				Phones: []*addresspb.Person_PhoneNumber{
					{
						Number: "P1",
						Type:   addresspb.Person_MOBILE,
					},
				},
				LastUpdated: &timestamppb.Timestamp{
					Seconds: 100,
					Nanos:   323,
				},
			},
			{
				Name:  "P2",
				Id:    2,
				Email: "p2@fia.com",
				Phones: []*addresspb.Person_PhoneNumber{
					{
						Number: "P2",
						Type:   addresspb.Person_HOME,
					},
				},
				LastUpdated: &timestamppb.Timestamp{
					Seconds: 155,
					Nanos:   877,
				},
			},
		},
	}

	fmt.Println(ab)
}

func complx() {
	cm := complexpb.ComplexMessage{
		OneDummy: &complexpb.DummyMessage{
			Id:   32,
			Name: "inner dummy",
		},
		MultipleDummy: []*complexpb.DummyMessage{
			{
				Id:   66,
				Name: "slice dummy",
			},
			{
				Id:   70,
				Name: "slice dummy 2",
			},
		},
	}

	fmt.Println(cm)
}

func enum() {
	em := enumpb.EnumMessage{
		Id:           18,
		DayOfTheWeek: enumpb.DayOfTheWeek_FRIDAY,
	}

	em.DayOfTheWeek = enumpb.DayOfTheWeek_MONDAY

	fmt.Println(em)
}

func toAndFromJSON(sm proto.Message) {
	sm2 := &simplepb.SimpleMessage{}

	json := toJSON(sm)

	if err := fromJSON(json, sm2); err != nil {
		panic(err)
	}
}

func toJSON(pb proto.Message) string {
	json, err := protojson.Marshal(pb)

	if err != nil {
		log.Fatalln("Can't marshal to json", err)
		return ""
	}

	return string(json)
}

func fromJSON(json string, pb proto.Message) error {
	err := protojson.Unmarshal([]byte(json), pb)

	if err != nil {
		log.Fatalln("Can't unmarshal from json", err)
		return err
	}

	return nil
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
