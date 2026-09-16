package employee

import (
	"math"
	"testing"
)

func mustNew(t *testing.T, id int, name string, d Department, s float64) Employee {
	t.Helper()
	e, err := New(id, name, d, s)
	if err != nil {
		t.Fatalf("New: %v", err)
	}
	return e
}

func sample(t *testing.T) Employees {
	t.Helper()
	return Employees{
		mustNew(t, 1, "Анна", Engineering, 150),
		mustNew(t, 2, "Борис", Engineering, 100),
		mustNew(t, 3, "Вера", Sales, 200),
		mustNew(t, 4, "Глеб", Sales, 50),
	}
}

func TestAdd_Duplicate(t *testing.T) {
	var es Employees
	if err := es.Add(mustNew(t, 1, "Анна", Sales, 90000)); err != nil {
		t.Fatalf("первый Add: %v", err)
	}
	if err := es.Add(mustNew(t, 1, "Другой", HR, 50000)); err == nil {
		t.Error("ожидали ошибку на дублирующийся ID, получили nil")
	}
	if len(es) != 1 {
		t.Errorf("len(es) = %d, ожидали 1", len(es))
	}
}

func TestAverageSalary(t *testing.T) {
	var empty Employees
	if got := empty.AverageSalary(); got != 0 {
		t.Errorf("пустая коллекция: got %v, ожидали 0", got)
	}
	// (150+100+200+50)/4 = 125
	if got := sample(t).AverageSalary(); math.Abs(got-125) > 1e-9 {
		t.Errorf("AverageSalary() = %v, ожидали 125", got)
	}
}

func TestFilter_NotMutatesSource(t *testing.T) {
	es := sample(t)
	before := len(es)
	_ = es.FilterByDepartment(Sales)
	if len(es) != before {
		t.Errorf("исходный срез изменился: было %d, стало %d", before, len(es))
	}
}

func TestFilterByDepartment(t *testing.T) {
	got := sample(t).FilterByDepartment(Engineering)
	if len(got) != 2 {
		t.Fatalf("len = %d, ожидали 2", len(got))
	}
	for _, e := range got {
		if e.Department != Engineering {
			t.Errorf("попал чужой отдел: %v", e)
		}
	}
}

func TestFilterAboveAverage(t *testing.T) {
	// среднее 125 => остаются 150 и 200
	got := sample(t).FilterAboveAverage()
	if len(got) != 2 {
		t.Fatalf("len = %d, ожидали 2", len(got))
	}
	for _, e := range got {
		if e.Salary <= 125 {
			t.Errorf("зарплата не выше средней: %v", e)
		}
	}
}

func TestSortBySalaryDesc(t *testing.T) {
	got := sample(t).SortBySalaryDesc()
	for i := 1; i < len(got); i++ {
		if got[i-1].Salary < got[i].Salary {
			t.Errorf("порядок нарушен на позиции %d: %v", i, got)
		}
	}
}

func TestSortByName(t *testing.T) {
	got := sample(t).SortByName()
	for i := 1; i < len(got); i++ {
		if got[i-1].Name > got[i].Name {
			t.Errorf("порядок нарушен на позиции %d: %v", i, got)
		}
	}
}

func TestDepartmentStats(t *testing.T) {
	stats := sample(t).DepartmentStats()

	eng := stats[Engineering]
	if eng.Count != 2 || eng.Sum != 250 || eng.Min != 100 || eng.Max != 150 {
		t.Errorf("Engineering: %+v", eng)
	}
	if math.Abs(eng.Avg-125) > 1e-9 {
		t.Errorf("Engineering.Avg = %v, ожидали 125", eng.Avg)
	}

	// сумма по всем отделам == общей сумме
	var total float64
	for _, s := range stats {
		total += s.Sum
	}
	if math.Abs(total-500) > 1e-9 {
		t.Errorf("сумма по отделам = %v, ожидали 500", total)
	}
}
