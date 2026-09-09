package employee

import "testing"

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		id      int
		empName string
		dept    Department
		salary  float64
		wantErr bool
	}{
		{"валидный", 1, "Анна", Sales, 90000, false},
		{"имя с пробелами обрезается", 2, "  Борис  ", HR, 50000, false},
		{"пустое имя", 3, "   ", Sales, 90000, true},
		{"неположительный id", 0, "Вера", HR, 50000, true},
		{"отрицательная зарплата", 4, "Глеб", HR, -1, true},
		{"неизвестный отдел", 5, "Дарья", "Ministry", 50000, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e, err := New(tt.id, tt.empName, tt.dept, tt.salary)
			if (err != nil) != tt.wantErr {
				t.Fatalf("New() error = %v, wantErr = %v", err, tt.wantErr)
			}
			if err == nil && e.Name == "" {
				t.Errorf("New() вернул пустое имя без ошибки")
			}
		})
	}
}

func TestEmployee_ToString(t *testing.T) {
	e, _ := New(1, "Анна", Sales, 92500.5)
	got := e.ToString()
	if got == "" {
		t.Fatal("ToString() пустая строка")
	}
	// String() должен совпадать с ToString() (реализация fmt.Stringer)
	if e.String() != got {
		t.Errorf("String() = %q, ToString() = %q — должны совпадать", e.String(), got)
	}
}
