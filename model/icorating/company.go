package model

type ICORatingCompany struct {
	Title            string
	Markets          []Market
}

type Market struct {
	Name string
	Max string
	Min string
	Volume string
}
