package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"proto/example/internal/proto/entity"
	"proto/example/internal/proto/proto_msg"

	"google.golang.org/protobuf/types/known/timestamppb"
)

const FNAME = "person_file"

func main() {
	error_log := log.New(os.Stderr, "ERROR: ", 0)
	info_log := log.New(os.Stdout, "INFO: ", 0)

	entity := &entity.Person{
		Id:    1234,
		Name:  "Test",
		Email: "email@test.com",
		Phones: []*entity.Person_PhoneNumber{
			{
				Number: "77668899",
				Type:   entity.Person_MOBILE,
			},
		},
		LastUpdated: timestamppb.New(time.Now()),
	}

	info_log.Println("Writting to file")

	err := proto_msg.WriteTo(FNAME, entity)

	if err != nil {
		error_log.Printf("%v", err)
	}

	info_log.Println("Writting to file DONE")

	info_log.Println("Reading from file")

	p, err := proto_msg.ReadFrom(FNAME)

	if err != nil {
		error_log.Printf("%v", err)
		return
	}

	info_log.Println("Reading from file DONE")

	txt, _ := json.Marshal(p)

	fmt.Printf("%s\n", txt)
}
