package schemas

import (
	"time"

	"gorm.io/gorm"
)

// modelo de dados gorm.Model, uma estrutura base fornecida pela GORM. Contêm pacotes uteis para rastrear a criação, atualização e exclusão.
type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}

// estrutura já formata que o cliente recebe ao interagir com as vagas na API
type OpeningResponse struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deteledAt,omitempty"`
	Role      string    `json:"role"`
	Company   string    `json:"company"`
	Location  string    `json:"location"`
	Remote    bool      `json:"remote"`
	Link      string    `json:"link"`
	Salary    int64     `json:"salary"`
}
