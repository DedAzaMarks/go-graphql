package repository

import "github.com/google/uuid"

type Role string

const (
	Civilian Role = "civilian"
	Mafia         = "mafia"
)

type Player struct {
	Username  string
	Role      Role
	IsAlive   bool
	DaysAlive int
}

type Game struct {
	Id         uuid.UUID
	Scoreboard []Player
	Comments   []string
}

type Repository struct {
	Games   []Game
	Players []Player
}
