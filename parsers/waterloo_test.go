package parsers

import "testing"

var studentIDTests = []struct {
	in  string // input
	out uint64 //output

}{
	{"this won't work", 0},
	{"1234567", 0},
	{"1234a568", 0},
	{"12345678", 12345678},
	{"a12345678", 12345678},
	{"12345678a", 12345678},
	{"123456789", 12345678},
}

func TestGetStudentID(t *testing.T) {
	for _, tt := range studentIDTests {
		t.Run(tt.in, func(t *testing.T) {
			s, err := GetStudentID(tt.in)

			if s == 0 && err == nil {
				t.Errorf("got 0 - but no error")
			}

			if s != tt.out {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}

var GetGradePageCountTests = []struct {
	in  string // input
	out int64  //output

}{
	{"page 1 of 1", 1},
	{"page 1 of 2", 2},
	{"page 2 of 2", 2},
	{"page 1 of 22", 22},
	{"page 1 of a", 0},
	{"invalid input", 0},
}

func TestGetGradePageCount(t *testing.T) {
	for _, tt := range GetGradePageCountTests {
		t.Run(tt.in, func(t *testing.T) {
			s, err := GetGradePageCount(tt.in)

			if s == 0 && err == nil {
				t.Errorf("got 0 - but no error")
			}

			if s != tt.out {
				t.Errorf("got %q, want %q", s, tt.out)
			}
		})
	}
}
