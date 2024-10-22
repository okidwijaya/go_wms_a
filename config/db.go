package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func DbConn() {
	dsn := "cloudsan_userroot:2024okicloud@tcp(cloudsand.my.id:3306)/cloudsan_wms_a?charset=utf8mb4&parseTime=True&loc=Local"

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error Database Connection", err)
	}
	// defer DB.Close()

	pingErr := DB.Ping()
	if pingErr != nil {
		log.Fatal("Error Ping Connection", pingErr)
	}

	fmt.Println("Database Connection Established Successfully!")
}

// router := gin.Default()

// router.GET("/ping", func(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{"Message": "Ping Connection Successfully"})
// })

// router.Run(":8080")
