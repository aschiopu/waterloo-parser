package printers

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/waterloo-parser/models"
)

// Take a map of students and pretty print to terminal
func PrettyPrint(students map[uint64]models.Student) {
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 0, '-', tabwriter.AlignRight|tabwriter.Debug)
	fmt.Fprintln(w, "ID\t", "Name\t", "Grade Average\t", "Outstandings\t", "Excellent\t", "Very Good\t", "Good\t", "Satisfactory\t", "Unsatisfactory\t", "Work Terms")
	for _, student := range students {

		fmt.Fprintln(w,
			student.ID, "\t",
			student.Name, "\t",
			student.GradeAverage, "\t",
			student.RatingOutstanding, "\t",
			student.RatingExcellent, "\t",
			student.RatingVeryGood, "\t",
			student.RatingGood, "\t",
			student.RatingSatisfactory, "\t",
			student.RatingUnsatisfactory, "\t",
			countWorkTerms(student), "\t",
		)
	}
	w.Flush()
}
