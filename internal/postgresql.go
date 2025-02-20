package storage

import (
	"database/sql"
	"fmt"
	storage "services/internal/storage"

	_ "github.com/lib/pq"
)

func FindAllMembers(db *sql.DB) string {
	//чтение строк таблицы
	rows, err := db.Query("SELECT * FROM squad")
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

func FindAllPack(db *sql.DB) string {
	//чтение строк таблицы
	rows, err := db.Query("Select * FROM starter_packs")
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

func InsertMember(db *sql.DB) {
	//добавление строки в таблицу
	sm := storage.InitMember(2, "Ivan", 60, 800, 1, 80)
	result, err := db.Exec("INSERT INTO squad (id, name, power, healthpoint, id_starter_pack, manapoint) VALUES ($1, $2, $3, $4, $5, $6)",
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

func InsertPack(db *sql.DB) {
	//добавление строки в таблицу
	sp := storage.InitPack(3, "Explosion", "Fetil", "Pool", "1231")
	result, err := db.Exec("INSERT INTO starter_packs (id, ability, item, tool, weapon) VALUES ($1, $2, $3, $4, $5)",
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

func DeleteMemberByID(db *sql.DB, id int) {
	//Удаление строки из таблицы по ID
	result, err := db.Exec("DELETE FROM squad WHERE id = $1", id)
	if err != nil {
		fmt.Printf("Неудалось выполнить запрос на удаление строки по id: %v", err)
	} else {
		fmt.Printf("Строка с id %v успешно удалена!", id)
	}
	_ = result
}

func DeletePackByID(db *sql.DB, id int) {
	//Удаление строки из таблицы по ID
	result, err := db.Exec("DELETE FROM starter_packs WHERE id = $1", id)
	if err != nil {
		fmt.Printf("Неудалось выполнить запрос на удаление строки по id: %v", err)
	} else {
		fmt.Printf("Строка с id %v успешно удалена!", id)
	}
	_ = result
}
