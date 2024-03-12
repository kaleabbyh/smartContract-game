package models

import "github.com/lib/pq"

type PlayerScore struct {
	Model
	Gameid string `gorm:"not null" json:"gameid"`
	Name   string `gorm:"not null" json:"name"`
	Score  int    `gorm:"not null" json:"score"`
}

type Pool struct {
	Model
	Gameid      string   `gorm:"not null" json:"gameid"`
	GameStatus  string   `gorm:"default:'ongoing'" json:"gameStatus"`
	Poolli      []Poolli `gorm:"foreignKey:PoolID" json:"Poolli,omitempty"`
	PlayerScore []Poolli `gorm:"foreignKey:PoolID" json:"PlayerScore,omitempty"`
}

type Poolli struct {
	Model
	Value  float32        `json:"Value"`
	Owners pq.StringArray `gorm:"type:text[]" json:"Owners"`
	PoolID string         `gorm:"not null" json:"PoolID"`
	Pool   Pool           `gorm:"foreignKey:PoolID" json:"-"`
}

type AdminWallet struct {
	Model
	Wallet string `gorm:"not null;unique" json:"wallet"`
}

type User struct {
	Model
	// add other fields here
}

type EscrowHire struct {
	Id               int            `gorm:"primary_key" json:"id,omitempty"`
	Contract_address string         `gorm:"not null" json:"contract_address"`
	Transaction      pq.StringArray `gorm:"type:text[]" json:"transaction"`
	State            string         `gorm:"column:state"`
	UserId           string         `gorm:"column:userId"`
	User             User           `gorm:"not null" json:"user"`
	Chat_id          string         `gorm:"not null" json:"chat_id"`
}

func (EscrowHire) TableName() string {
	return "Escrow_Hire"
}

type EscrowHoF struct {
	Id               int            `gorm:"primary_key" json:"id,omitempty"`
	Contract_address string         `gorm:"not null" json:"contract_address"`
	State            string         `gorm:"not null" json:"state"`
	User             User           `gorm:"not null" json:"user"`
	UserId           string         `gorm:"column:userId" json:"userId"`
	Transaction      pq.StringArray `gorm:"type:text[]" json:"transaction"`
	Hof_id           int            `gorm:"column:hof_id" json:"hof_id"`
}

func (EscrowHoF) TableName() string {
	return "Escrow_HoF"
}
