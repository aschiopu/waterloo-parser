package parsers

import (
	"errors"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// all pages
func GetStudentID(rawText string) (uint64, error) {
	validStudentID := regexp.MustCompile(`[0-9]{8}`)
	studentIDString := validStudentID.FindStringSubmatch(rawText)
	if len(studentIDString) == 0 {
		return 0, errors.New("No student id found")
	}
	studentID, err := strconv.ParseUint(studentIDString[0], 10, 64)
	if err != nil {
		panic(err)
	}

	return studentID, nil
}

// work term page
func IsWorkTermPage(rawText string) bool {
	workTermPage, err := regexp.MatchString("University of WaterlooCo-operative Work Terms", rawText)
	if err != nil {
		panic(err)
	}

	return workTermPage
}

// Returns a map of ratings and their count a student received.
func GetWorkTermRatings(rawText string) map[string]int {
	rawText = strings.TrimPrefix(rawText, "University of WaterlooCo-operative Work Terms")

	validRatingOutstanding := regexp.MustCompile(`OUTSTANDING`)
	validRatingExcellent := regexp.MustCompile(`EXCELLENT`)
	validRatingVeryGood := regexp.MustCompile(`VERY GOOD`)
	validRatingGood := regexp.MustCompile(`GOOD`)
	validRatingSatisfactory := regexp.MustCompile(`SATISFACTORY`)
	validRatingUnsatisfactory := regexp.MustCompile(`UNSATISFACTORY`)

	ratings := make(map[string]int)
	ratings["outstanding"] = len(validRatingOutstanding.FindAllStringIndex(rawText, -1))
	ratings["excellent"] = len(validRatingExcellent.FindAllStringIndex(rawText, -1))
	ratings["veryGood"] = len(validRatingVeryGood.FindAllStringIndex(rawText, -1))
	ratings["good"] = len(validRatingGood.FindAllStringIndex(rawText, -1))
	ratings["satisfactory"] = len(validRatingSatisfactory.FindAllStringIndex(rawText, -1))
	ratings["unsatisfactory"] = len(validRatingUnsatisfactory.FindAllStringIndex(rawText, -1))

	return ratings
}

// Extract the student name from the work term page
func GetStudentName(rawText string) string {
	rawText = strings.TrimPrefix(rawText, "University of WaterlooCo-operative Work Terms")

	validName := regexp.MustCompile(`^\D+`)
	return validName.FindStringSubmatch(rawText)[0]
}

// grade page
func IsInitialGradePage(rawText string) bool {
	initialGradePage, err := regexp.MatchString("UNIVERSITY OF WATERLOOUNOFFICIAL GRADE REPORT", rawText)
	if err != nil {
		panic(err)
	}

	return initialGradePage
}

func GetGradePageCount(rawText string) (int64, error) {
	validLastPageCount := regexp.MustCompile(`[0-9]+$`)
	pageCountString := validLastPageCount.FindStringSubmatch(rawText)
	if len(pageCountString) == 0 {
		return 0, errors.New("Couldn't find page number")
	}
	pageCount, err := strconv.ParseInt(pageCountString[0], 10, 64)
	if err != nil {
		panic(err)
	}

	return pageCount, nil
}

func GetStudentAverage(rawText string) int {
	validTermAverages := regexp.MustCompile(`Term Average:Decision:([0-9]{2})`)
	termAverages := validTermAverages.FindAllStringSubmatch(rawText, -1)
	numOfTerms := math.Max(float64(len(termAverages)), 1)
	averageTotal := 0

	for _, num := range termAverages {
		// TODO: I should be able to do this w/ regex to only return the subset - the grade
		numText := strings.TrimPrefix(num[0], "Term Average:Decision:")
		numInt, _ := strconv.ParseInt(numText, 10, 64)
		averageTotal = averageTotal + int(numInt)
	}

	return int(float64(averageTotal) / numOfTerms)
}
