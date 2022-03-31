package grades_test

import (
	"fixtures_example/grades"
	"os"
	"testing"
)

func TestNewGrade_ErrorHandling(t *testing.T) {
	cases := []struct {
		fixture   string
		returnErr bool
		name      string
	}{
		{
			fixture:   "testdata/grades/empty.csv",
			returnErr: false,
			name:      "EmptyFile",
		},
		{
			fixture:   "testdata/grades/invalid.csv",
			returnErr: true,
			name:      "InvalidFile",
		},
		{
			fixture:   "testdata/grades/valid.csv",
			returnErr: false,
			name:      "ValidFile",
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			file, err := os.Open(tc.fixture)
			if err != nil {
				t.Fatal(err)
			}
			_, err = grades.NewGradebook(file)
			returnedErr := err != nil

			if returnedErr != tc.returnErr {
				t.Fatalf("Expected returnErr: %v, got: %v", tc.returnErr, returnedErr)
			}
		})
	}
}

func TestFindByStudent(t *testing.T) {
	cases := []struct {
		fixture string
		student string
		want    grades.Gradebook
		name    string
	}{
		{
			name:    "EmptyGradebook",
			fixture: "testdata/grades/empty.csv",
			student: "Jane",
			want:    grades.Gradebook{},
		},
		{
			name:    "NonEmptyGradebook_Found",
			fixture: "testdata/grades/valid.csv",
			student: "Jane",
			want: grades.Gradebook{
				grades.Record{
					Student: "Jane",
					Subject: "Chemistry",
					Grade:   "A",
				},
				grades.Record{
					Student: "Jane",
					Subject: "Algebra",
					Grade:   "B",
				},
			},
		},
		{
			name:    "NonEmptyGradebook_NotFound",
			fixture: "testdata/grades/valid.csv",
			student: "Missing",
			want:    grades.Gradebook{},
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			file, err := os.Open(tc.fixture)
			if err != nil {
				t.Fatal(err)
			}
			gradebook, err := grades.NewGradebook(file)
			if err != nil {
				t.Fatalf("Cannot create gradebook: %v", err)
			}

			got := gradebook.FindByStudent(tc.student)
			for idx, gotGrade := range got {
				wantedGrade := tc.want[idx]
				if gotGrade != wantedGrade {
					t.Errorf("Expected: %v, got: %v", wantedGrade, gotGrade)
				}
			}

		})
	}
}
