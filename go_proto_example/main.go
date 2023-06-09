package main

import (
	"fmt"
	"reflect"

	pb "github.com/here-bossnanda/go_proto_example/proto"
	"google.golang.org/protobuf/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          100,
		IsSimple:    true,
		Name:        "nanda",
		SampleLists: []int32{1, 2, 3, 4, 5},
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		OneDummy: &pb.Dummy{
			Id:   1,
			Name: "nanda",
		},
		MultipleDummies: []*pb.Dummy{
			{
				Id:   2,
				Name: "nanda1",
			},
			{
				Id:   3,
				Name: "nanda2",
			},
		},
	}
}

func doEnum() *pb.EnumEye {
	return &pb.EnumEye{
		EyeColor: pb.EyeColor_EYE_COLOR_GREEN,
	}
}

func doOneOfs(message interface{}) {
	switch x := message.(type) {
	case *pb.Result_Id:
		fmt.Println(message.(*pb.Result_Id).Id)
	case *pb.Result_Message:
		fmt.Println(message.(*pb.Result_Message).Message)
	default:
		fmt.Printf("message has unexpected type: %v", x)
	}
}

func doMap() *pb.MapExample {
	return &pb.MapExample{
		Ids: map[string]*pb.IDWrapper{
			"id_custom":  {Id: 1},
			"id_custom2": {Id: 2},
			"id_custom3": {Id: 3},
		},
	}
}

func doFile(pm proto.Message) {
	path := "simple.bin"
	writeToFile(path, pm)
	message := &pb.Simple{}
	readFromFile(path, message)

	fmt.Println(message)
}

func doToJSON(p proto.Message) string {
	jsonString := toJson(p)
	return jsonString
}

func doFromJSON(jsonString string, t reflect.Type) proto.Message {
	message := reflect.New(t).Interface().(proto.Message)
	fromJson(jsonString, message)
	return message
}

func main() {
	// fmt.Println(doSimple())
	// fmt.Println(doComplex())
	// fmt.Println(doEnum())
	// fmt.Println("this should be an id")
	// doOneOfs(&pb.Result_Id{Id: 1})

	// fmt.Println("this should be an message")
	// doOneOfs(&pb.Result_Message{Message: "a message"})

	// fmt.Println("this should be an error")
	// doOneOfs(true)

	// fmt.Println(doMap())

	// doFile(doSimple())

	jsonString := doToJSON(doSimple())
	message := doFromJSON(jsonString, reflect.TypeOf(pb.Simple{}))

	fmt.Println(jsonString)
	fmt.Println(message)
}
