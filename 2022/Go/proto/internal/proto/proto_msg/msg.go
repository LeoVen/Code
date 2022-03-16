package proto_msg

import (
	"io/ioutil"
	"proto/example/internal/proto/entity"

	"google.golang.org/protobuf/proto"
)

func WriteTo(fname string, person *entity.Person) error {
	out, err := proto.Marshal(person)

	if err != nil {
		return err
	}

	if err := ioutil.WriteFile(fname, out, 0644); err != nil {
		return err
	}

	return nil
}

func ReadFrom(fname string) (*entity.Person, error) {
	in, err := ioutil.ReadFile(fname)

	if err != nil {
		return nil, err
	}

	person := &entity.Person{}

	if err := proto.Unmarshal(in, person); err != nil {
		return nil, err
	}

	return person, nil
}
