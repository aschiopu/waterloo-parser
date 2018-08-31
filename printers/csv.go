package printers

// Take a map of students and export all the data to csv
import (
	"encoding/csv"
	"os"
	"strconv"

	"github.com/waterloo-parser/models"
)

func CSV(students map[uint64]models.Student) {
	data := [][]string{}

	data = append(data, []string{"ID", "Name", "Grade Average", "Outstandings", "Excellent", "Very Good", "Good", "Satisfactory", "Unsatisfactory", "Work Terms"})

	for _, student := range students {

		data = append(data, []string{
			strconv.FormatUint(student.ID, 10),
			student.Name,
			strconv.Itoa(student.GradeAverage),
			strconv.Itoa(student.RatingOutstanding),
			strconv.Itoa(student.RatingExcellent),
			strconv.Itoa(student.RatingVeryGood),
			strconv.Itoa(student.RatingGood),
			strconv.Itoa(student.RatingSatisfactory),
			strconv.Itoa(student.RatingUnsatisfactory),
			strconv.Itoa(countWorkTerms(student)),
		})
	}

	csvExport(data, "results.csv")
}

// Take in a nested array of data and write to csv file
func csvExport(data [][]string, filename string) {
	file, _ := os.Create(filename)
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, row := range data {
		err := writer.Write(row)
		if err != nil {
			panic(err)
		}
	}
}

func countWorkTerms(student models.Student) int {
	return student.RatingOutstanding + student.RatingExcellent + student.RatingVeryGood + student.RatingGood + student.RatingUnsatisfactory
}
