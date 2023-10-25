package grades

import (
	"encoding/csv"
	"fmt"
	"io"
)

type Record struct {
	Student string
	Subject string
	Grade   string
}

type Gradebook []Record

func NewGradebook(csvFile io.Reader) (Gradebook, error) {
	var gradebook Gradebook
	reader := csv.NewReader(csvFile)

	for {
		line, err := reader.Read()

		if err == io.EOF {
			break
		}

		if err != nil {
			return gradebook, err
		}

		if len(line) < 3 {
			return gradebook, fmt.Errorf("Invalid csv")
		}

		gradebook = append(gradebook, Record{
			Student: line[0],
			Subject: line[1],
			Grade:   line[2],
		})
	}

	return gradebook, nil
}

func (gb *Gradebook) FindByStudent(student string) []Record {
	var result []Record
	for _, record := range *gb {
		if record.Student == student {
			result = append(result, record)
		}
	}
	return result
}
