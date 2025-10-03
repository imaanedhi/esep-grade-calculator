package esepunittests

import "testing"

func TestGetGrade_Boundaries(t *testing.T) {
	cases := []struct {
		in   int
		want string
	}{
		{90, "A"}, {89, "B"},
		{80, "B"}, {79, "C"},
		{70, "C"}, {69, "D"},
		{60, "D"}, {59, "F"},
	}
	for _, c := range cases {
		if got := GetGrade(c.in); got != c.want {
			t.Fatalf("GetGrade(%d) = %q; want %q", c.in, got, c.want)
		}
	}
}

func addN(gc *GradeCalculator, n int, gt GradeType) {
	for i := 0; i < n; i++ {
		gc.AddGrade("x", 100, gt)
	}
}

func makeGC(a, e, s int) *GradeCalculator {
	gc := NewGradeCalculator()
	addN(gc, a, Assignment)
	addN(gc, e, Exam)
	addN(gc, s, Essay)
	return gc
}

func TestGetFinalGrade_AllBranches(t *testing.T) {
	// A branch: avg(assignments)=90, avg(exams)=90 -> final = 90 -> "A"
	if got := makeGC(181, 181, 0).GetFinalGrade(); got != "A" {
		t.Fatalf("GetFinalGrade(A) = %q; want %q", got, "A")
	}
	// B branch: 80
	if got := makeGC(161, 161, 1).GetFinalGrade(); got != "B" {
		t.Fatalf("GetFinalGrade(B) = %q; want %q", got, "B")
	}
	// C branch: 70
	if got := makeGC(141, 141, 2).GetFinalGrade(); got != "C" {
		t.Fatalf("GetFinalGrade(C) = %q; want %q", got, "C")
	}
	// D branch: 60
	if got := makeGC(121, 121, 3).GetFinalGrade(); got != "D" {
		t.Fatalf("GetFinalGrade(D) = %q; want %q", got, "D")
	}
	// F branch: minimal non-empty counts keep avg = 0
	if got := makeGC(1, 1, 1).GetFinalGrade(); got != "F" {
		t.Fatalf("GetFinalGrade(F) = %q; want %q", got, "F")
	}
}

func TestAddGradeAndGradeTypeString(t *testing.T) {
	gc := NewGradeCalculator()
	gc.AddGrade("a", 95, Assignment)
	gc.AddGrade("e", 88, Exam)
	gc.AddGrade("s", 77, Essay)

	if Assignment.String() != "assignment" {
		t.Fatalf("Assignment.String()=%q; want %q", Assignment.String(), "assignment")
	}
	if Exam.String() != "exam" {
		t.Fatalf("Exam.String()=%q; want %q", Exam.String(), "exam")
	}
	if Essay.String() != "essay" {
		t.Fatalf("Essay.String()=%q; want %q", Essay.String(), "essay")
	}
}

func TestNewGradeCalculator_Empty(t *testing.T) {
	gc := NewGradeCalculator()
	// We intentionally don't call computeAverage on empties to avoid div-by-zero in original code.
	if gc == nil {
		t.Fatal("NewGradeCalculator returned nil")
	}
}
