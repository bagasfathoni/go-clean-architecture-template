package model

import "time"

type Vendor struct {
	Name          string     `json:"name" mapstructure:"name"`
	Address       string     `json:"address" mapstructure:"address"`
	Phone         string     `json:"phone" mapstructure:"phone"`
	BankName      string     `json:"bankName" mapstructure:"bank_name"`
	AccountNumber string     `json:"accountNumber" mapstructure:"account_number"`
	AccountName   string     `json:"accountName" mapstructure:"account_name"`
	CreatedAt     *time.Time `json:"createdAt" mapstructure:"created_at"`
	UpdatedAt     *time.Time `json:"updatedAt" mapstructure:"updated_at"`
	DeletedAt     *time.Time `json:"deletedAt" mapstructure:"deleted_at"`
}
