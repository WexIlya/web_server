package storage

type SquadMembers struct {
	Id            int
	Name          string
	Power         int
	HP            int
	IdStarterPack int
	MP            int
}

type StarterPack struct {
	Id      int
	Ability string
	Item    string
	Tool    string
	Weapon  string
}

func InitMember(id int, name string, power int, hp int, idsp int, mp int) SquadMembers {
	// функция для заполнения полей структуры
	sm := SquadMembers{
		Id:            id,
		Name:          name,
		Power:         power,
		HP:            hp,
		IdStarterPack: idsp,
		MP:            mp}
	return sm
}

func InitPack(id int, ability string, item string, tool string, weapon string) StarterPack {
	// функция для заполнения полей структуры
	sp := StarterPack{
		Id:      id,
		Ability: ability,
		Item:    item,
		Tool:    tool,
		Weapon:  weapon}
	return sp
}
