package models

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
