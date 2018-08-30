// waterloo parsing library
package main

import (
	"regexp"
	"strconv"
	"strings"
)

// all pages
func getStudentID(rawText string) uint64 {
	validStudentID := regexp.MustCompile(`[0-9]{8}`)
	studentIDString := validStudentID.FindStringSubmatch(rawText)[0]
	studentID, err := strconv.ParseUint(studentIDString, 10, 64)
	if err != nil {
		panic(err)
	}

	return studentID
}

// work term page
func isWorkTermPage(rawText string) bool {
	workTermPage, err := regexp.MatchString("University of WaterlooCo-operative Work Terms", rawText)
	if err != nil {
		panic(err)
	}

	return workTermPage
}

// Returns a map of ratings and their count a student received.
func getWorkTermRatings(rawText string) map[string]int {
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
func getStudentName(rawText string) string {
	rawText = strings.TrimPrefix(rawText, "University of WaterlooCo-operative Work Terms")

	validName := regexp.MustCompile(`^\D+`)
	return validName.FindStringSubmatch(rawText)[0]
}

// grade page
func isInitialGradePage(rawText string) bool {
	initialGradePage, err := regexp.MatchString("UNIVERSITY OF WATERLOOUNOFFICIAL GRADE REPORT", rawText)
	if err != nil {
		panic(err)
	}

	return initialGradePage
}

func getGradePageCount(rawText string) int64 {
	validLastPageCount := regexp.MustCompile(`[0-9]+$`)
	pageCountString := validLastPageCount.FindStringSubmatch(rawText)[0]
	pageCount, err := strconv.ParseInt(pageCountString, 10, 64)
	if err != nil {
		panic(err)
	}

	return pageCount
}

func getStudentAverage(rawText string) int {
	validTermAverages := regexp.MustCompile(`Term Average:Decision:([0-9]{2})`)
	termAverages := validTermAverages.FindAllStringSubmatch(rawText, -1)
	averageTotal := 0

	for _, num := range termAverages {
		// TODO: I should be able to do this w/ regex to only return the subset - the grade
		numText := strings.TrimPrefix(num[0], "Term Average:Decision:")
		numInt, _ := strconv.ParseInt(numText, 10, 64)
		averageTotal = averageTotal + int(numInt)
	}

	return averageTotal / len(termAverages)
}