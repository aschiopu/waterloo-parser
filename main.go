package main

import (
	"flag"

	"github.com/ledongthuc/pdf"
	"github.com/waterloo-parser/models"
	"github.com/waterloo-parser/parsers"
	"github.com/waterloo-parser/printers"
)

// TODO: add tests

func main() {

	exportPointer := flag.Bool("export", false, "export to csv")
	filePointer := flag.String("file", "samples/test1.pdf", "filepath to parse")

	flag.Parse()
	students := processPDF(*filePointer) // Read local pdf file

	if *exportPointer {
		printers.CSV(students)
	} else {
		printers.PrettyPrint(students)
	}
}

func processPDF(path string) map[uint64]models.Student {
	// TODO resave waterloo pdfs becase they are protected so errors when opening
	f, r, err := pdf.Open(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	students := map[uint64]models.Student{}

	ch := make(chan models.Student)

	for i := 1; i < r.NumPage(); i++ {
		rawText, err := r.Page(i).GetPlainText(nil)
		if err != nil {
			panic(err)
		}

		if parsers.IsWorkTermPage(rawText) {
			go updateStudentWorkTerms(rawText, ch)
			prccessStudentUpdate(<-ch, students)
		}

		if parsers.IsInitialGradePage(rawText) {
			pageCount, err := parsers.GetGradePageCount(rawText)
			if err != nil {
				panic(err)
			}
			gradeText := rawText

			if pageCount > 1 {
				secondPageText, _ := r.Page(i + 1).GetPlainText(nil)
				gradeText = gradeText + secondPageText
			}
			go updateStudentGrade(gradeText, ch)
			prccessStudentUpdate(<-ch, students)
		}
	}

	return students
}

func prccessStudentUpdate(student models.Student, students map[uint64]models.Student) {
	if _, ok := students[student.ID]; !ok {
		students[student.ID] = student
	} else {
		studentToUpdate := students[student.ID]
		studentToUpdate.GradeAverage = student.GradeAverage
		students[student.ID] = studentToUpdate
	}
}

func updateStudentGrade(rawText string, ch chan models.Student) {
	studentID, err := parsers.GetStudentID(rawText)
	if err != nil {
		panic(err)
	}
	ch <- models.Student{
		ID:           studentID,
		GradeAverage: parsers.GetStudentAverage(rawText),
	}
}

func updateStudentWorkTerms(rawText string, ch chan models.Student) {
	ratings := parsers.GetWorkTermRatings(rawText)
	studentID, err := parsers.GetStudentID(rawText)
	if err != nil {
		panic(err)
	}

	ch <- models.Student{
		ID:                   studentID,
		Name:                 parsers.GetStudentName(rawText),
		RatingOutstanding:    ratings["outstanding"],
		RatingExcellent:      ratings["excellent"],
		RatingVeryGood:       ratings["veryGood"],
		RatingGood:           ratings["good"],
		RatingSatisfactory:   ratings["satisfactory"],
		RatingUnsatisfactory: ratings["unsatisfactory"],
	}
}
