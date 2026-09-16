# Лабораторная работа №1 — «Типы данных в Go»

Программный прототип: учёт сотрудников (`Employee`) с фильтрацией, сортировкой
и статистикой по отделам.

## Требования

- Go 1.21+ (используются пакеты `slices`, `cmp`).

## Запуск

```bash
go run ./cmd/lab1                        # встроенный набор / testdata/employees.json
go run ./cmd/lab1 -data path/to/file.json
```

## Тесты

```bash
go test ./...
go test -v -cover ./internal/employee
```

## Проверки качества

```bash
gofmt -l .
go vet ./...
```

## Структура

| Путь | Назначение |
|------|-----------|
| `internal/employee/employee.go`   | типы `Department`, `Grade`, `Money`, `Employee`; конструктор `New`; `ToString`/`String` |
| `internal/employee/collection.go` | тип `Employees`; `Add`, `AverageSalary`, `Filter*`, `Sort*`, `Table` |
| `internal/employee/stats.go`      | `Stats`, `DepartmentStats`, `DepartmentStatsString` |
| `internal/employee/loader.go`     | `LoadJSON` — чтение данных из JSON |
| `cmd/lab1/main.go`                 | демонстрационный сценарий |
| `testdata/employees.json`         | тестовый набор данных |

Подробный разбор кода для защиты — в [EXPLANATION.md](EXPLANATION.md).
Распределение работ и этапы — в [PLAN.md](PLAN.md).
