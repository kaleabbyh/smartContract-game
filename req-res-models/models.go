package models

import "cicada/web-service-gin/models"

type Pool struct {
	models.Model
	Gameid     string   `gorm:"not null" json:"gameid"`
	GameStatus string   `gorm:"default:'ongoing'" json:"gameStatus"`
	Poolli     []Poolli `gorm:"foreignKey:PoolID" json:"Poolli,omitempty"`
}

type Poolli struct {
	models.Model
	Value  float32  `json:"Value"`
	Owners []string `gorm:"type:text[]" json:"Owners"`
	PoolID string   `gorm:"not null" json:"PoolID"`
	Pool   Pool     `gorm:"foreignKey:PoolID" json:"-"`
}

type PlayerScore struct {
	models.Model
	Gameid string `gorm:"not null" json:"gameid"`
	Name   string `gorm:"not null" json:"name"`
	Score  int    `gorm:"not null" json:"score"`
}
type PlayersScore struct {
	Gameid       string        `json:"gameid"`
	PlayersScore []PlayerScore `json:"playersScore"`
}

type PlayerScoreResponse struct {
	Gameid       string        `json:"gameid"`
	PlayersScore []PlayerScore `json:"playersScore"`
}

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Player struct {
	Name   string
	Amount float32
}

type Players struct {
	Players []Player
}

type PoolResponse struct {
	ID     string
	Value  float32
	Owners []string
}
type CreatePoolResponse struct {
	PoolID  string
	GameID  string
	Poollis []PoolResponse
}

type PoolToSort struct {
	PoolID  string `json:"poolID"`
	Pooliis []Poollis
}
type Poollis struct {
	ID     string   `json:"id"`
	PoolID string   `json:"poolID,omitempty"`
	Value  float32  `json:"value"`
	Owners []string `json:"owners"`
}

type Winner struct {
	ID              string   `json:"id"`
	Name            string   `json:"name"`
	NotsharedAmount float32  `json:"not_shared_amount"`
	SharedAmount    float32  `json:"shared_amount"`
	TotalAmount     float32  `json:"total_amount"`
	PoolsWonID      []string `json:"pools_won_id"`
}
type AllWinners struct {
	Winnners []Winner
}

type GameRequest struct {
	PlayerScore []struct {
		ID    string `json:"id"`
		Name  string `json:"name"`
		Score int    `json:"score"`
	} `json:"playerScore"`

	Players []struct {
		ID     string  `json:"id"`
		Name   string  `json:"name"`
		Amount float32 `json:"amount"`
	} `json:"players"`
}

type EscrowHoF struct {
	models.Model
	ContractAddress string   `gorm:"not null" json:"contractAddress"`
	Status          string   `gorm:"not null" json:"status"`
	UserId          string   `gorm:"not null" json:"userId"`
	User            User     `gorm:"not null" json:"user"`
	HoFID           string   `gorm:"not null" json:"HoFID"`
	Transaction     []string `gorm:"type:text[]" json:"transaction"`
}
type User struct {
	models.Model
	// add other fields here
}
