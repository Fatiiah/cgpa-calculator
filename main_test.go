package main

import (
	"math"
	"testing"
)

// this function tests the gradeToPoints function with various grades and checks if the returned points are correct
func TestGradeToPoints(t *testing.T) {
	tests := []struct {
		grade    string
		expected float64
		valid    bool
	}{
		{"A", 5.0, true},
		{"B", 4.0, true},
		{"C", 3.0, true},
		{"D", 2.0, true},
		{"F", 0.0, true},
		{"C+", 2.3, false},
		{"Z", 0.0, false},
		{"a", 5.0, true}, 
	}

	for _, tt := range tests {
		points, ok := gradeToPoints(tt.grade)
		if ok != tt.valid {
			t.Errorf("gradeToPoints(%q): expected valid=%v, got valid=%v", tt.grade, tt.valid, ok)
		}
		if ok && points != tt.expected {
			t.Errorf("gradeToPoints(%q): expected %.1f, got %.1f", tt.grade, tt.expected, points)
		}
	}
}

// this function tests the calculateCGPA function with a set of courses and checks if the calculated CGPA is correct
func TestCalculateCGPA(t *testing.T) {
	courses := []Course{
		{Name: "Maths", Grade: "A", Units: 3, Points: 4.0},
		{Name: "English", Grade: "B", Units: 2, Points: 3.0},
		{Name: "Physics", Grade: "C", Units: 3, Points: 2.0},
	}
	// (4*3 + 3*2 + 2*3) / (3+2+3) = (12+6+6)/8 = 24/8 = 3.0
	expected := 3.0
	got := calculateCGPA(courses)
	if math.Abs(got-expected) > 0.001 {
		t.Errorf("calculateCGPA: expected %.4f, got %.4f", expected, got)
	}
}

// this function tests the calculateCGPA function with an empty course list and expects a CGPA of 0.0
func TestCalculateCGPAEmpty(t *testing.T) {
	got := calculateCGPA([]Course{})
	if got != 0.0 {
		t.Errorf("calculateCGPA(empty): expected 0.0, got %.4f", got)
	}
}

// this function tests the cgpaToClass function with various CGPA values and checks if the returned class is correct
func TestCgpaToClass(t *testing.T) {
	tests := []struct {
		cgpa     float64
		expected string
	}{
		{5.0, "First Class"},
		{4.6, "First Class"},
		{4.49, "Second Class Upper"},
		{4.0, "Second Class Upper"},
		{3.49, "Second Class Lower"},
		{3.0, "Second Class Lower"},
		{2.49, "Third Class"},
		{1.0, "Pass"},
		{0.5, "Pass"},
	}

	for _, tt := range tests {
		got := cgpaToClass(tt.cgpa)
		if got != tt.expected {
			t.Errorf("cgpaToClass(%.1f): expected %q, got %q", tt.cgpa, tt.expected, got)
		}
	}
}

// this function tests the parsing of course arguments
func TestParseCourseArg(t *testing.T) {
	c, err := parseCourseArg("Mathematics:A:3")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if c.Name != "Mathematics" || c.Grade != "A" || c.Units != 3 || c.Points != 5.0 {
		t.Errorf("parseCourseArg: unexpected result %+v", c)
	}
}

// this function tests for missing units, bad grade, bad units, negative units and zero units
func TestParseCourseArgInvalid(t *testing.T) {
	cases := []string{
		"Maths:A",         
		"Maths:Z:3",        
		"Maths:A:abc",      
		"Maths:A:-1",       
		"Maths:A:0",        
	}
	for _, c := range cases {
		_, err := parseCourseArg(c)
		if err == nil {
			t.Errorf("parseCourseArg(%q): expected error, got nil", c)
		}
	}
}

// this function tests if the truncate function correctly truncates long strings and leaves short strings unchanged
func TestTruncate(t *testing.T) {
	if truncate("Hello", 10) != "Hello" {
		t.Error("truncate short string failed")
	}
	result := truncate("This is a very long course name", 10)
	if len(result) != 10 || !endsWith(result, "...") {
		t.Errorf("truncate long string failed: got %q", result)
	}
}

func endsWith(s, suffix string) bool {
	return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}