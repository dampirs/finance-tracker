package expences

import "time"

type Expence struct {
	Id          int       `json:"id"`
	Time        time.Time `json:"time"`
	Category    string    `json:"category"`
	Description string    `json:"description"`
	Amount      int       `json:"amount"`
}

func NewExpence(description string, category string, amount int) *Expence {
	return &Expence{Id: len(storage), Time: time.Now(), Category: category, Description: description, Amount: amount}
}
