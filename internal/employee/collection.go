package employee

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
	"text/tabwriter"
)

// Employees — коллекция сотрудников. Это именованный тип на основе среза,
// поэтому на него можно навешивать методы (Add, Filter*, Sort* и т. д.).
type Employees []Employee

// NewEmployees собирает коллекцию из набора сотрудников,
// добавляя их по одному через Add (с проверкой уникальности ID).
func NewEmployees(list ...Employee) (Employees, error) {
	var es Employees
	for _, e := range list {
		if err := es.Add(e); err != nil {
			return nil, err
		}
	}
	return es, nil
}

// Add добавляет сотрудника в коллекцию.
// Возвращает ошибку, если сотрудник с таким ID уже есть.
// Получатель — указатель (*Employees), потому что метод изменяет сам срез.
func (es *Employees) Add(e Employee) error {
	for _, existing := range *es {
		if existing.ID == e.ID {
			return fmt.Errorf("employees: duplicate id %d", e.ID)
		}
	}
	*es = append(*es, e)
	return nil
}

// AverageSalary возвращает среднюю зарплату по коллекции.
// Для пустой коллекции возвращает 0 (защита от деления на ноль).
func (es Employees) AverageSalary() float64 {
	if len(es) == 0 {
		return 0
	}
	var sum float64
	for _, e := range es {
		sum += e.Salary
	}
	return sum / float64(len(es))
}

// Filter возвращает новый срез с элементами, для которых pred вернул true.
// Исходная коллекция не изменяется.
func (es Employees) Filter(pred func(Employee) bool) Employees {
	result := make(Employees, 0, len(es))
	for _, e := range es {
		if pred(e) {
			result = append(result, e)
		}
	}
	return result
}

// FilterByDepartment оставляет сотрудников заданного отдела.
func (es Employees) FilterByDepartment(d Department) Employees {
	return es.Filter(func(e Employee) bool { return e.Department == d })
}

// FilterAboveAverage оставляет сотрудников с зарплатой строго выше средней.
func (es Employees) FilterAboveAverage() Employees {
	avg := es.AverageSalary()
	return es.Filter(func(e Employee) bool { return e.Salary > avg })
}

// SortBy возвращает отсортированную КОПИЮ коллекции.
// cmp(a, b) должен вернуть <0, 0 или >0 (контракт slices.SortFunc).
func (es Employees) SortBy(cmpFunc func(a, b Employee) int) Employees {
	result := slices.Clone(es)
	slices.SortFunc(result, cmpFunc)
	return result
}

// SortBySalaryDesc сортирует по убыванию зарплаты.
func (es Employees) SortBySalaryDesc() Employees {
	return es.SortBy(func(a, b Employee) int {
		return cmp.Compare(b.Salary, a.Salary) // b и a местами => по убыванию
	})
}

// SortByName сортирует по имени в алфавитном порядке.
func (es Employees) SortByName() Employees {
	return es.SortBy(func(a, b Employee) int {
		return strings.Compare(a.Name, b.Name)
	})
}

// Table возвращает список сотрудников в виде выровненной таблицы.
// Выравнивание колонок делает text/tabwriter.
func (es Employees) Table() string {
	var sb strings.Builder
	w := tabwriter.NewWriter(&sb, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ID\tИМЯ\tОТДЕЛ\tЗАРПЛАТА")
	fmt.Fprintln(w, "--\t---\t-----\t--------")
	for _, e := range es {
		fmt.Fprintf(w, "%d\t%s\t%s\t%.2f ₽\n", e.ID, e.Name, e.Department, e.Salary)
	}
	w.Flush()
	return sb.String()
}
