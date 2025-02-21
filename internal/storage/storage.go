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

type ID struct {
	Id int
}
