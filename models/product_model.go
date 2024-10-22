package models

import "github.com/okidwijaya/go_wms_a/config"

type Product struct {
	name string `json:"ProductName"`
}

func AddProduct(product Product) error {
	query := "INSERT INTO master_products SET ?"
	_, err := config.DB.Exec(query, product.name)
	return err
}
