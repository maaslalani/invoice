package utils

import "time"

type Invoice struct {
	Id    string `json:"id" yaml:"id"`
	Title string `json:"title" yaml:"title"`

	Logo string `json:"logo" yaml:"logo"`
	From string `json:"from" yaml:"from"`
	To   string `json:"to" yaml:"to"`
	Date string `json:"date" yaml:"date"`
	Due  string `json:"due" yaml:"due"`

	Items      []string  `json:"items" yaml:"items"`
	Quantities []int     `json:"quantities" yaml:"quantities"`
	Rates      []float64 `json:"rates" yaml:"rates"`

	Tax      float64 `json:"tax" yaml:"tax"`
	Discount float64 `json:"discount" yaml:"discount"`
	Currency string  `json:"currency" yaml:"currency"`

	Note string `json:"note" yaml:"note"`
}

func DefaultInvoice() Invoice {
	return Invoice{
		Id:         time.Now().Format("20060102"),
		Title:      "INVOICE",
		Rates:      []float64{25},
		Quantities: []int{2},
		Items:      []string{"Paper Cranes"},
		From:       "Project Folded, Inc.",
		To:         "Untitled Corporation, Inc.",
		Date:       time.Now().Format("Jan 02, 2006"),
		Due:        time.Now().AddDate(0, 0, 14).Format("Jan 02, 2006"),
		Tax:        0,
		Discount:   0,
		Currency:   "USD",
	}
}
