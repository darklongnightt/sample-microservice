package db

import (
	"fmt"
	"time"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
)

// Product migration will create default table name: products
type Product struct {
	tableName struct{} `sql:"product_items"`
	ID        int      `sql:"id,pk"`
	Name      string   `sql:"name,unique"`
	Desc      string   `sql:"desc"`
	Image     string   `sql:"image"`
	Price     float64  `sql:"price,type:real"`
	IsActive  bool     `sql:"is_active"`
	Features  struct { // struct,array,map -> jsonb
		Name string
		Desc string
	} `sql:"features,type:jsonb"`
	CreatedAt time.Time `sql:"created_at"`
	UpdatedAt time.Time `sql:"updated_at"`
}

// CreateProductTable ...
func CreateProductTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	if err := db.CreateTable(&Product{}, opts); err != nil {
		return fmt.Errorf("error while creating table for Product\nreason: %v", err)
	}

	fmt.Println("Table for Product created successfully")
	return nil
}
