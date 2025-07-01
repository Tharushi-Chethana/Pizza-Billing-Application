package models

type Item struct {
	ItemID    uint    `gorm:"primaryKey" json:"id"`
	//gorm:"primaryKey" tells a database tool called GORM, This field is the primary key in the database table.
	//When converting this struct to JSON (like sending data over the internet), call this field ‘id’.
	Name  string  `json:"name"`
	Type  string  `json:"type"`  // pizza, topping, beverage, desert
	Price float64 `json:"price"`
}