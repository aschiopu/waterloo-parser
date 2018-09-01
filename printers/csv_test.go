package printers

import (
	"testing"

	"github.com/aschiopu/waterloo-parser/models"
)

var prepStudentsForCSVTests = []struct {
	name string
	in   models.Student // input
	out  [][]string     //output

}{
	{
		name: "tests student id export",
		in:   models.Student{ID: 123, Name: "Alexei", GradeAverage: 12},
		out:  [][]string{{"ID", "Name", "Grade Average"}, {"123", "Alexei", "12"}},
	},
}

func TestPrepStudentsForCSV(t *testing.T) {
	for _, tt := range prepStudentsForCSVTests {
		t.Run(tt.name, func(t *testing.T) {
			students := map[uint64]models.Student{12: tt.in}
			s := prepStudentsForCSV(students)

			for i := 0; i < 2; i++ {
				if s[1][i] != tt.out[1][i] {
					t.Errorf("got %q, want %q", s, tt.out)
				}
			}
		})
	}
}
