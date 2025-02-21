package storage

import (
	"encoding/json"
	"fmt"
	"net/http"
	connect "services/internal"
	"services/internal/storage"
)

func FindAllPack() string {
	//чтение строк таблицы
	rows, err := connect.Connectdb.Query("Select * FROM starter_packs")
	result := ""
	if err != nil {
		fmt.Printf("Ошибка при запросе на чтение таблицы: %v", err)
	}
	defer rows.Close()
	var (
		id      int
		ability string
		item    string
		tool    string
		weapon  string
	)
	for rows.Next() {
		if err := rows.Scan(&id, &ability, &item, &tool, &weapon); err != nil {
			fmt.Printf("Строку с id %v не удалось прочесть!", id)
		}
		str := fmt.Sprintf("id = %v, ability = %s, item = %s, tool = %s, weapon = %s\n", id, ability, item, tool, weapon)
		result += str
	}
	return result
}

func InsertPack(w http.ResponseWriter, r *http.Request) {
	//добавление строки в таблицу
	var sp storage.StarterPack
	err := json.NewDecoder(r.Body).Decode(&sp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	result, err := connect.Connectdb.Exec("INSERT INTO starter_packs (id, ability, item, tool, weapon) VALUES ($1, $2, $3, $4, $5)",
		sp.Id,
		sp.Ability,
		sp.Item,
		sp.Tool,
		sp.Weapon,
	)
	if err != nil {
		fmt.Printf("Ошибка при добавлении строки: %v", err)
	} else {
		fmt.Printf("Строка успешно добавлена! %v", result)
	}
}

func DeletePackByID(w http.ResponseWriter, r *http.Request) {
	//Удаление строки из таблицы по ID
	var sp storage.ID
	err := json.NewDecoder(r.Body).Decode(&sp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := connect.Connectdb.Exec("DELETE FROM starter_packs WHERE id = $1", sp.Id)
	if err != nil {
		fmt.Printf("Неудалось выполнить запрос на удаление строки по id: %v", err)
	} else {
		fmt.Printf("Строка с id %v успешно удалена!", sp.Id)
	}
	_ = result
}
