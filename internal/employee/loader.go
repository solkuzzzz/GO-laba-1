package employee

import (
	"encoding/json"
	"fmt"
	"os"
)

// jsonEmployee — промежуточная структура для разбора JSON.
// Теги `json:"..."` связывают поля структуры с ключами в файле.
type jsonEmployee struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Department string  `json:"department"`
	Salary     float64 `json:"salary"`
}

// LoadJSON читает список сотрудников из JSON-файла.
// Каждая запись проходит через New (валидацию) и Add (уникальность ID).
func LoadJSON(path string) (Employees, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("load json: %w", err)
	}

	var raw []jsonEmployee
	if err := json.Unmarshal(data, &raw); err != nil {
		return nil, fmt.Errorf("parse json: %w", err)
	}

	var es Employees
	for i, r := range raw {
		e, err := New(r.ID, r.Name, Department(r.Department), r.Salary)
		if err != nil {
			return nil, fmt.Errorf("record %d: %w", i, err)
		}
		if err := es.Add(e); err != nil {
			return nil, fmt.Errorf("record %d: %w", i, err)
		}
	}
	return es, nil
}
