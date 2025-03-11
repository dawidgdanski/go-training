package generate

import (
	"fmt"
	"go-training/generate/data"

	"google.golang.org/protobuf/proto"
)

//go:generate protoc -I=data --go_out=data --go_opt=module=go-training/generate/data --go_opt=Mperson.proto=go-training/generate/data person.proto
func ProtobufExample() {
	p := &data.Person{
		Name:  "Bob Bobson",
		Id:    20,
		Email: "bob@bobson.com",
	}
	fmt.Println(p)
	protoBytes, _ := proto.Marshal(p)
	fmt.Println(protoBytes)
	var p2 data.Person
	err := proto.Unmarshal(protoBytes, &p2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(&p2)
	}
}
