package tools

import (
	"ConsultantBack/internal/settings"
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDb() (*sql.DB, error) {
	fmt.Println("Start connectin to db")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		settings.DbHost, settings.DbPort, settings.DbUser, settings.DbPassword, settings.DbName)

	db, err := sql.Open("postgres", dsn)

	fmt.Println("Connected to db")

	if err != nil {
		//defer db.Close()
		fmt.Printf("Error:%s\n", err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		//defer db.Close()
		fmt.Printf("Error:%s\n", err)
		return nil, err
	}
	fmt.Println("End connectin to db")
	return db, nil
}
