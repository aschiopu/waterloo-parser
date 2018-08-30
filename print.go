package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"text/tabwriter"
)

// Take a map of students and pretty print to terminal
func prettyPrint(students map[uint64]Student) {
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

// Take a map of students and export all the data to csv
func printToCSV(students map[uint64]Student) {
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

func countWorkTerms(student Student) int {
	return student.RatingOutstanding + student.RatingExcellent + student.RatingVeryGood + student.RatingGood + student.RatingUnsatisfactory
}
