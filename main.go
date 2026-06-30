package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Course holds a single course entry
type Course struct {
	Name   string
	Grade  string
	Units  float64
	Points float64
}

// this function converts a letter grade to grade points (5.0 scale)
func gradeToPoints(grade string) (float64, bool) {
	gradeMap := map[string]float64{
		
		"A": 5.0,
		"B": 4.0,
		"C": 3.0,
		"D": 2.0,
		"F": 0.0,
	}
	points, ok := gradeMap[strings.ToUpper(grade)]
	return points, ok
}

// this function returns the degree class for a CGPA
func cgpaToClass(cgpa float64) string {
	switch {
	case cgpa >= 4.50:
		return "First Class"
	case cgpa >= 3.50:
		return "Second Class Upper"
	case cgpa >= 2.50:
		return "Second Class Lower"
	case cgpa >= 1.50:
		return "Third Class"
	default:
		return "Pass"
	}
}

// this function computes CGPA from a list of courses
func calculateCGPA(courses []Course) float64 {
	totalPoints := 0.0
	totalUnits := 0.0
	for _, c := range courses {
		totalPoints += c.Points * c.Units
		totalUnits += c.Units
	}
	if totalUnits == 0 {
		return 0.0
	}
	return totalPoints / totalUnits
}

// this function shows how to use the program
func printUsage() {
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("  Single CGPA calculation:")
	fmt.Println("    go run . --gpa \"CourseName:Grade:Units\" \"CourseName:Grade:Units\" ...")
	fmt.Println()
	fmt.Println("  Cumulative CGPA (multiple semesters):")
	fmt.Println("    go run . --cgpa \"semester1.txt\" \"semester2.txt\" ...")
	fmt.Println()
	fmt.Println("  Transcript summary:")
	fmt.Println("    go run . --transcript \"semester1.txt\" \"semester2.txt\" ...")
	fmt.Println()
	fmt.Println("Course format:  CourseName:Grade:Units")
	fmt.Println("Grades:         A, B, C, D, F")
	fmt.Println()
	fmt.Println("Example:")
	fmt.Println("  go run . --gpa \"Maths:A:3\" \"English:B:2\" \"Physics:C:3\"")
	fmt.Println()
	fmt.Println("File format (semester.txt):")
	fmt.Println("  Maths:A:3")
	fmt.Println("  English:B:2")
	fmt.Println("  Physics:C:3")
}

// this function parses "CourseName:Grade:Units" string
func parseCourseArg(arg string) (Course, error) {
	parts := strings.Split(arg, ":")
	if len(parts) != 3 {
		return Course{}, fmt.Errorf("invalid course format %q, expected Name:Grade:Units", arg)
	}

	name := strings.TrimSpace(parts[0])
	grade := strings.TrimSpace(parts[1])
	unitsStr := strings.TrimSpace(parts[2])

	units, err := strconv.ParseFloat(unitsStr, 64)
	if err != nil || units <= 0 {
		return Course{}, fmt.Errorf("invalid units %q for course %q", unitsStr, name)
	}

	points, ok := gradeToPoints(grade)
	if !ok {
		return Course{}, fmt.Errorf("unknown grade %q for course %q", grade, name)
	}

	return Course{Name: name, Grade: strings.ToUpper(grade), Units: units, Points: points}, nil
}

// this function reads courses from a text file
func parseFile(filename string) ([]Course, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("cannot read file %q: %v", filename, err)
	}

	lines := strings.Split(strings.TrimSpace(string(data)), "\n")
	var courses []Course
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		c, err := parseCourseArg(line)
		if err != nil {
			return nil, fmt.Errorf("line %d in %s: %v", i+1, filename, err)
		}
		courses = append(courses, c)
	}
	return courses, nil
}

// this function prints a formatted table for a semester
func printSemesterTable(semesterName string, courses []Course) float64 {
	fmt.Printf("\n┌─────────────────────────────────────────────────────┐\n")
	fmt.Printf("│  %-51s│\n", semesterName)
	fmt.Printf("├──────────────────────────┬────────┬───────┬─────────┤\n")
	fmt.Printf("│ %-24s │ Grade  │ Units │  Points │\n", "Course")
	fmt.Printf("├──────────────────────────┼────────┼───────┼─────────┤\n")

	gpa := calculateCGPA(courses)
	for _, c := range courses {
		qualityPoints := c.Points * c.Units
		fmt.Printf("│ %-24s │  %-5s │  %-4.0f │  %5.2f  │\n",
			truncate(c.Name, 24), c.Grade, c.Units, qualityPoints)
	}

	fmt.Printf("├──────────────────────────┴────────┴───────┴─────────┤\n")
	fmt.Printf("│  Semester GPA: %-36s│\n", fmt.Sprintf("%.2f  (%s)", gpa, cgpaToClass(gpa)))
	fmt.Printf("└─────────────────────────────────────────────────────┘\n")

	return gpa
}

// this function shortens a string with ellipsis if too long
func truncate(s string, max int) string {
	if len(s) <= max {
		return s
	}
	return s[:max-3] + "..."
}

// this function processes the --gpa flag (inline course args)
func handleGPA(args []string) {
	if len(args) == 0 {
		fmt.Println("Error: --gpa requires at least one course argument.")
		fmt.Println("Example: go run . --gpa \"Maths:A:3\" \"English:B:2\"")
		os.Exit(1)
	}

	var courses []Course
	for _, arg := range args {
		c, err := parseCourseArg(arg)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		courses = append(courses, c)
	}

	printSemesterTable("GPA Calculation", courses)
}

// this function processes the --cgpa flag (multiple semester files)
func handleCGPA(files []string) {
	if len(files) == 0 {
		fmt.Println("Error: --cgpa requires at least one semester file.")
		fmt.Println("Example: go run . --cgpa semester1.txt semester2.txt")
		os.Exit(1)
	}

	var allCourses []Course
	for i, file := range files {
		courses, err := parseFile(file)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		printSemesterTable(fmt.Sprintf("Semester %d — %s", i+1, file), courses)
		allCourses = append(allCourses, courses...)
	}

	cgpa := calculateCGPA(allCourses)
	fmt.Printf("\n╔═════════════════════════════════════════════════════╗\n")
	fmt.Printf("║  CUMULATIVE GPA (CGPA)                              ║\n")
	fmt.Printf("╠═════════════════════════════════════════════════════╣\n")
	fmt.Printf("║  CGPA  : %-42s║\n", fmt.Sprintf("%.4f", cgpa))
	fmt.Printf("║  Class : %-42s║\n", cgpaToClass(cgpa))
	fmt.Printf("╚═════════════════════════════════════════════════════╝\n\n")
}

// this function processes the --transcript flag with full summary
func handleTranscript(files []string) {
	if len(files) == 0 {
		fmt.Println("Error: --transcript requires at least one semester file.")
		os.Exit(1)
	}

	fmt.Printf("\n╔═════════════════════════════════════════════════════╗\n")
	fmt.Printf("║           OFFICIAL ACADEMIC TRANSCRIPT              ║\n")
	fmt.Printf("╚═════════════════════════════════════════════════════╝\n")

	var allCourses []Course
	semesterGPAs := []float64{}

	for i, file := range files {
		courses, err := parseFile(file)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			os.Exit(1)
		}
		gpa := printSemesterTable(fmt.Sprintf("Semester %d — %s", i+1, file), courses)
		semesterGPAs = append(semesterGPAs, gpa)
		allCourses = append(allCourses, courses...)
	}

	cgpa := calculateCGPA(allCourses)

	// Summary table
	fmt.Printf("\n┌─────────────────────────────────────────────────────┐\n")
	fmt.Printf("│  SEMESTER GPA SUMMARY                               │\n")
	fmt.Printf("├────────────────────────────────────┬────────────────┤\n")
	fmt.Printf("│ Semester                           │ GPA            │\n")
	fmt.Printf("├────────────────────────────────────┼────────────────┤\n")
	for i, gpa := range semesterGPAs {
		label := fmt.Sprintf("Semester %d", i+1)
		fmt.Printf("│ %-34s │ %-14.4f │\n", label, gpa)
	}
	fmt.Printf("├────────────────────────────────────┼────────────────┤\n")
	fmt.Printf("│ %-34s │ %-14.4f │\n", "CUMULATIVE GPA (CGPA)", cgpa)
	fmt.Printf("│ %-34s │ %-14s │\n", "Degree Class", cgpaToClass(cgpa))
	fmt.Printf("└────────────────────────────────────┴────────────────┘\n\n")
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printUsage()
		os.Exit(0)
	}

	flag := args[0]
	rest := args[1:]

	switch flag {
	case "--gpa":
		handleGPA(rest)
	case "--cgpa":
		handleCGPA(rest)
	case "--transcript":
		handleTranscript(rest)
	case "--help", "-h":
		printUsage()
	default:
		fmt.Printf("Error: unknown flag %q\n\n", flag)
		printUsage()
		os.Exit(1)
	}
}