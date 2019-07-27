package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	enumpb "github.com/sagarsahu27/myCode/src/enum_example"
	simplepb "github.com/sagarsahu27/myCode/src/simple"
)

func main() {
	sm := doSimple()

	readAndWriteDemo(sm)
	jsonDemo(sm)

	enumDemo(sm)
}

func enumDemo(pb proto.Message) {
	ep := enumpb.EnumMessage{
		Id:           42,
		DayOfTheWeek: enumpb.DayOfTheWeek_SATURDAY,
	}
	fmt.Println(ep)
}

func jsonDemo(pb proto.Message) {

	str := toJSON(pb)
	fmt.Println(str)

	sm2 := &simplepb.SimpleMessage{}
	fromJSON(str, sm2)
	fmt.Println("read protobuf message is:", sm2)

}

func fromJSON(str string, pb proto.Message) {
	err := jsonpb.UnmarshalString(str, pb)
	if err != nil {
		log.Println("Couldn't read from json")
	}
}

func toJSON(pb proto.Message) string {
	marshaler := jsonpb.Marshaler{}
	str, err := marshaler.MarshalToString(pb)
	if err != nil {
		log.Fatalln("marshalling to json failed", err)
		return ""
	}
	return str
}

func readAndWriteDemo(sm proto.Message) {
	writeToFile("simple.bin", sm)

	sm2 := &simplepb.SimpleMessage{}
	readFromFile("simple.bin", sm2)

	fmt.Println(sm2)
}

func writeToFile(fname string, pb proto.Message) error {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't serialise to byte", err)
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		log.Fatalln("Can't write to file", err)
		return err
	}
	fmt.Println("Serialize data written successfully to file")
	return nil
}

func readFromFile(fname string, pb proto.Message) error {
	in, err := ioutil.ReadFile(fname)
	if err != nil {
		log.Fatalln("Cannot read from file", err)
		return err
	}

	err = proto.Unmarshal(in, pb)
	if err != nil {
		log.Fatalln("Unmarshalling failed !!!", err)
		return err
	}

	return nil

}

func doSimple() *simplepb.SimpleMessage {
	sm := simplepb.SimpleMessage{
		Id:         12345,
		IsSimple:   true,
		Name:       "Sam",
		SampleList: []int32{1, 2, 3, 4, 5},
	}
	fmt.Println(sm)
	sm.Name = "Bhai Sahab !!!"
	fmt.Println(sm)
	fmt.Println("Product ID is ", sm.GetId())

	return &sm
}
