## CGPA CALCULATOR
# CGPA Calculator

A command-line transcript and CGPA calculator written in Golang. Calculate your semester GPA, cumulative CGPA, and generate a full academic transcript, all from your terminal.

---

## Requirements

- Go 1.22 or higher

---

## Installation

**Clone or download the project:**
```console
git clone https://github.com/Fatiiah/cgpa-calculator.git
cd cgpa-calculator
```

**Initialize the module:**
```console
go mod init cgpa-calculator
```

**Build the binary:**
```console
go build -o cgpa-calculator .
```

---

## Usage

### Option 1 — Run directly with Go
```console
go run . [flag] [arguments]
```

### Option 2 — Run the binary (after building)
```console
./cgpa-calculator [flag] [arguments]
```

---

## Flags

| Flag | Description | Input |
|------|-------------|-------|
| `--gpa` | Calculate GPA from inline course arguments | Course strings |
| `--cgpa` | Calculate cumulative GPA across multiple semesters | `.txt` files |
| `--transcript` | Generate a full academic transcript with summary | `.txt` files |
| `--help` | Show usage instructions | — |

---

## Course Format

All courses follow this format:
```
CourseName:Grade:Units
```

**Example:**
```
Mathematics:A:3
English:B:2
Physics:A:3
```

### Supported Grades

| Grade | Points | 
|-------|--------|
| A     | 5.0    |
| B     | 4.0    | 
| C     | 3.0    |
| D     | 2.0    | 
| F     | 0.0    | 
|       |        | 

---

## Examples

### Quick GPA (inline)
```console
./cgpa-calculator --gpa "Maths:A:3" "English:B:2" "Physics:A:3"
```

### CGPA from semester files
```console
./cgpa-calculator --cgpa semester1.txt semester2.txt
```

### Full transcript
```console
./cgpa-calculator --transcript semester1.txt semester2.txt
```

### Show help
```console
./cgpa-calculator --help
```

---

## Semester File Format

Create a `.txt` file for each semester. Each line is one course in `Name:Grade:Units` format.
Lines starting with `#` are treated as comments and ignored.

---

## Degree Classification

| CGPA Range | Class |
|------------|-------|
| 4.50 – 5.0  | First Class |
| 3.50 – 4.49  | Second Class Upper |
| 2.40 – 3.49  | Second Class Lower |
| 1.50 – 2.39  | Third Class |
| 1.00 – 1.49  | Pass |

---

## How CGPA is Calculated

CGPA uses the **weighted average** formula:

```
CGPA = Total Quality Points / Total Credit Units

Quality Points = Grade Points × Credit Units
```

**Example:**
```
Maths    → A  (5.0) × 3 units = 10.0
English  → B  (4.0) × 2 units =  8.0
Physics  → A- (5.0) × 3 units = 10.0
                                ──────
Total Quality Points = 28.0
Total Units          = 8
CGPA                 = 28.0 /  = 3.5 → Second Class Upper
```

---

## AUTHOR
Fatiah Shehu 