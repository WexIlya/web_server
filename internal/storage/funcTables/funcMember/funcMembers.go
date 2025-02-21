package storage

import (
	"encoding/json"
	"fmt"
	"net/http"
	connect "services/internal"
	storage "services/internal/storage"

	_ "github.com/lib/pq"
)

func FindAllMembers() string {
	//чтение строк таблицы
	rows, err := connect.Connectdb.Query("SELECT * FROM squad")
	result := ""
	if err != nil {
		fmt.Printf("Ошибка при запросе на чтение таблицы: %v", err)
	}
	defer rows.Close()
	var (
		id              int
		name            string
		power           int
		hp              int
		id_starter_pack int
		mp              int
	)
	for rows.Next() {
		if err := rows.Scan(&id, &name, &power, &hp, &id_starter_pack, &mp); err != nil {
			fmt.Printf("Строку с id %v не удалось прочесть!", id)
		}
		str := fmt.Sprintf("id = %v name = %s power = %v hp = %v id starter pack = %v mp = %v\n", id, name, power, hp, id_starter_pack, mp)
		result += str
	}

	return result
}

func InsertMember(w http.ResponseWriter, r *http.Request) {
	//добавление строки в таблицу
	var sm storage.SquadMembers
	err := json.NewDecoder(r.Body).Decode(&sm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	result, err := connect.Connectdb.Exec("INSERT INTO squad (id, name, power, healthpoint, id_starter_pack, manapoint) VALUES ($1, $2, $3, $4, $5, $6)",
		sm.Id,
		sm.Name,
		sm.Power,
		sm.HP,
		sm.IdStarterPack,
		sm.MP,
	)
	if err != nil {
		fmt.Printf("Ошибка при добавлении строки: %v", err)
	} else {
		fmt.Printf("Строка успешно добавлена! %v", result)
	}
}

func DeleteMemberByID(w http.ResponseWriter, r *http.Request) {
	//Удаление строки из таблицы по ID
	var sm storage.ID
	err := json.NewDecoder(r.Body).Decode(&sm)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	result, err := connect.Connectdb.Exec("DELETE FROM squad WHERE id = $1", sm.Id)
	if err != nil {
		fmt.Printf("Неудалось выполнить запрос на удаление строки по id: %v", err)
	} else {
		fmt.Printf("Строка с id %v успешно удалена!", sm.Id)
	}
	_ = result
}
