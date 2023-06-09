package main

import (
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func toJson(pb proto.Message) string {
	option := protojson.MarshalOptions{
		Multiline: true,
	}
	output, err := option.Marshal(pb)
	if err != nil {
		log.Fatalln("Can't convert to JSON", err)
		return ""
	}

	return string(output)
}

func fromJson(input string, pb proto.Message) {
	option := protojson.UnmarshalOptions{
		DiscardUnknown: true,
	}

	if err := option.Unmarshal([]byte(input), pb); err != nil {
		log.Fatalln("Couldn't unmarshal to JSON", err)
		return
	}
}
