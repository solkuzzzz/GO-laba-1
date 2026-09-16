package employee

import (
	"errors"
	"fmt"
	"strings"
)

type Department string
type Grade int
type Money int64
type Employee struct {
	ID         int
	Name       string
	Department Department
	Salary     float64
}

const (
	Engineering Department = "Engineering"
	Sales       Department = "Sales"
	HR          Department = "HR"
	Marketing   Department = "Marketing"
)

const (
	Junior Grade = iota
	Middle
	Senior
)

func (d Department) Valid() bool {
	switch d {
	case Engineering, Sales, HR, Marketing:
		return true
	default:
		return false
	}
}

func (m Money) String() string {
	return fmt.Sprintf("%d.%02d ₽", m/100, m%100)
}

func New(id int, name string, dept Department, salary float64) (Employee, error) {
	if id <= 0 {
		return Employee{}, fmt.Errorf("employee: id must be positive, got %d", id)
	}
	name = strings.TrimSpace(name)
	if name == "" {
		return Employee{}, errors.New("employee: name must not be empty")
	}

	if salary < 0 {
		return Employee{}, fmt.Errorf("employee: salary must not be negative, got %.2f", salary)
	}

	if !dept.Valid() {
		return Employee{}, fmt.Errorf("employee: unknown department %q", dept)
	}

	return Employee{ID: id, Name: name, Department: dept, Salary: salary}, nil
}

// ToString возвращает человекочитаемое представление одной записи.
func (e Employee) ToString() string {
	return fmt.Sprintf(
		"#%-4d %-20s | %-12s | %10.2f ₽",
		e.ID, e.Name, e.Department, e.Salary,
	)
}

// String реализует fmt.Stringer, чтобы Employee красиво печатался через fmt.
func (e Employee) String() string {
	return e.ToString()
}
