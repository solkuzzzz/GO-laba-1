package main

import (
	"flag"
	"fmt"
	"os"

	"unik/lab1/internal/employee"
)

func main() {
	// Флаг -data позволяет подставить свой JSON; по умолчанию — набор из testdata.
	dataPath := flag.String("data", "testdata/employees.json", "путь к JSON со списком сотрудников")
	flag.Parse()

	staff, err := loadStaff(*dataPath)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ошибка загрузки данных:", err)
		os.Exit(1)
	}

	section("1. Все сотрудники")
	fmt.Print(staff.Table())

	section("2. Добавление записей")
	demoAdd(&staff)

	section("3. Фильтр по отделу: Engineering")
	fmt.Print(staff.FilterByDepartment(employee.Engineering).Table())

	section(fmt.Sprintf("4. Зарплата выше средней (%.2f ₽)", staff.AverageSalary()))
	fmt.Print(staff.FilterAboveAverage().SortBySalaryDesc().Table())

	section("5. Сортировка по имени")
	fmt.Print(staff.SortByName().Table())

	section("6. Статистика по отделам")
	fmt.Print(staff.DepartmentStatsString())
}

// loadStaff берёт данные из файла, а если файла нет — из встроенного набора.
func loadStaff(path string) (employee.Employees, error) {
	if _, err := os.Stat(path); err == nil {
		return employee.LoadJSON(path)
	}
	fmt.Printf("(%s не найден, использую встроенный набор)\n", path)
	return employee.NewEmployees(
		mustNew(1, "Анна Петрова", employee.Engineering, 145000),
		mustNew(2, "Борис Смирнов", employee.Engineering, 120000),
		mustNew(3, "Вера Кузнецова", employee.Sales, 98000),
		mustNew(4, "Глеб Соколов", employee.Sales, 87000),
		mustNew(5, "Дарья Морозова", employee.HR, 76000),
		mustNew(6, "Егор Волков", employee.Marketing, 91000),
	)
}

// demoAdd показывает и успешное добавление, и ошибку дубликата ID.
func demoAdd(staff *employee.Employees) {
	newbie, _ := employee.New(100, "Ксения Орлова", employee.HR, 70000)
	if err := staff.Add(newbie); err != nil {
		fmt.Println("не добавлено:", err)
	} else {
		fmt.Println("добавлен:", newbie)
	}

	dup, _ := employee.New(1, "Дубликат", employee.Sales, 50000)
	if err := staff.Add(dup); err != nil {
		fmt.Println("ожидаемая ошибка:", err)
	}
}

func mustNew(id int, name string, d employee.Department, salary float64) employee.Employee {
	e, err := employee.New(id, name, d, salary)
	if err != nil {
		panic(err)
	}
	return e
}

func section(title string) {
	fmt.Printf("\n=== %s ===\n", title)
}
