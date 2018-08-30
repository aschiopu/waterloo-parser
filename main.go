package main

import (
	"flag"

	"github.com/ledongthuc/pdf"
)

// TODO: add tests

// Student is the representation of all information on a student.
type Student struct {
	ID                   uint64
	Name                 string
	GradeAverage         int
	RatingOutstanding    int
	RatingExcellent      int
	RatingVeryGood       int
	RatingGood           int
	RatingSatisfactory   int
	RatingUnsatisfactory int
}

func main() {

	exportPointer := flag.Bool("export", false, "a bool")
	filePointer := flag.String("file", "samples/test1.pdf", "file")

	flag.Parse()
	students := processPDF(*filePointer) // Read local pdf file

	if *exportPointer {
		printToCSV(students)
	} else {
		prettyPrint(students)
	}
}

func processPDF(path string) map[uint64]Student {
	// TODO resave pdfs becase they are protected so errors when opening
	f, r, err := pdf.Open(path)
	defer f.Close()
	if err != nil {
		panic(err)
	}
	students := map[uint64]Student{}

	ch := make(chan Student)

	for i := 1; i < r.NumPage(); i++ {
		rawText, err := r.Page(i).GetPlainText(nil)
		if err != nil {
			panic(err)
		}

		if isWorkTermPage(rawText) {
			go updateStudentWorkTerms(rawText, ch)
			prccessStudentUpdate(<-ch, students)
		}

		if isInitialGradePage(rawText) {
			pageCount := getGradePageCount(rawText)
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

func prccessStudentUpdate(student Student, students map[uint64]Student) {
	if _, ok := students[student.ID]; !ok {
		students[student.ID] = student
	} else {
		studentToUpdate := students[student.ID]
		studentToUpdate.GradeAverage = student.GradeAverage
		students[student.ID] = studentToUpdate
	}
}

func updateStudentGrade(rawText string, ch chan Student) {
	ch <- Student{
		ID:           getStudentID(rawText),
		GradeAverage: getStudentAverage(rawText),
	}
}

func updateStudentWorkTerms(rawText string, ch chan Student) {
	ratings := getWorkTermRatings(rawText)

	ch <- Student{
		ID:                   getStudentID(rawText),
		Name:                 getStudentName(rawText),
		RatingOutstanding:    ratings["outstanding"],
		RatingExcellent:      ratings["excellent"],
		RatingVeryGood:       ratings["veryGood"],
		RatingGood:           ratings["good"],
		RatingSatisfactory:   ratings["satisfactory"],
		RatingUnsatisfactory: ratings["unsatisfactory"],
	}
}
