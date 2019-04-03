package model

type Invoice struct {
	BillingAddress string `gorm:"column:BillingAddress" json:"BillingAddress"`
	BillingCity    string `gorm:"column:BillingCity" json:"BillingCity"`
	BillingCountry string `gorm:"column:BillingCountry"`
}

func (Invoice) TableName() string {
	return "Invoice"
}

type Employee struct {
	FirstName string `gorm:"column:FirstName"`
	LastName  string `gorm:"column:LastName"`
	Title     string `gorm:"column:Title"`
}

func (Employee) TableName() string {
	return "Employee"
}
