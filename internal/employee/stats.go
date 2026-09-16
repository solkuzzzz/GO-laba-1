package employee

import (
	"fmt"
	"slices"
	"strings"
	"text/tabwriter"
)

// Stats — агрегаты по одному отделу.
type Stats struct {
	Count int
	Sum   float64
	Avg   float64
	Min   float64
	Max   float64
}

// DepartmentStats группирует сотрудников по отделам и считает агрегаты.
// Ключ map — отдел, значение — статистика по нему.
func (es Employees) DepartmentStats() map[Department]Stats {
	acc := make(map[Department]Stats)

	for _, e := range es {
		s, ok := acc[e.Department]
		if !ok {
			// первый сотрудник отдела: инициализируем Min/Max его зарплатой
			s = Stats{Min: e.Salary, Max: e.Salary}
		}
		s.Count++
		s.Sum += e.Salary
		if e.Salary < s.Min {
			s.Min = e.Salary
		}
		if e.Salary > s.Max {
			s.Max = e.Salary
		}
		acc[e.Department] = s
	}

	// среднее считаем отдельным проходом, когда известны Count и Sum
	for d, s := range acc {
		s.Avg = s.Sum / float64(s.Count)
		acc[d] = s
	}
	return acc
}

// DepartmentStatsString форматирует статистику по отделам в таблицу.
// Отделы выводятся в алфавитном порядке (обход map недетерминирован).
func (es Employees) DepartmentStatsString() string {
	stats := es.DepartmentStats()

	depts := make([]string, 0, len(stats))
	for d := range stats {
		depts = append(depts, string(d))
	}
	slices.Sort(depts)

	var sb strings.Builder
	w := tabwriter.NewWriter(&sb, 0, 0, 2, ' ', 0)
	fmt.Fprintln(w, "ОТДЕЛ\tКОЛ-ВО\tСУММА\tСРЕДНЕЕ\tМИН\tМАКС")
	for _, d := range depts {
		s := stats[Department(d)]
		fmt.Fprintf(w, "%s\t%d\t%.2f\t%.2f\t%.2f\t%.2f\n",
			d, s.Count, s.Sum, s.Avg, s.Min, s.Max)
	}
	w.Flush()
	return sb.String()
}
