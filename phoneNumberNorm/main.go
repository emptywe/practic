package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"time"
)

//var numbers = []string{"1234567890", "123 456 7891", "(123) 456 7892", "(123) 456-7893", "123-456-7894", "123-456-7890", "1234567892", "(123)456-7892"}

// ConnectDB - connecting to existing database, via given source and pq driver
func ConnectDB(source string) (*sqlx.DB, error) {
	db, err := sqlx.Connect("postgres", source)
	if err != nil {
		return nil, err
	}
	return db, err
}

// SetTable - automatically creates needed table
func SetTable(db *sqlx.DB, tableName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	sql := fmt.Sprintf(`CREATE TABLE IF NOT EXISTS %s(
    							id SERIAL PRIMARY KEY,
    							phone_number varchar NOT NULL
								)`, tableName)
	_, err := db.ExecContext(ctx, sql)
	return err
}

// SetDefaultData - sets task phone numbers
func SetDefaultData(db *sqlx.DB, tableName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	sql := fmt.Sprintf(`INSERT INTO %s (phone_number) values ('1234567890'),('123 456 7891'),('(123) 456 7892'),
										('(123) 456-7893'),('123-456-7894'),('123-456-7890'),('1234567892'),('(123)456-7892')`, tableName)
	_, err := db.ExecContext(ctx, sql)
	return err
}

// FormatPhoneNumbers - formatting phone numbers in database
func FormatPhoneNumbers(db *sqlx.DB, tableName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	sql := fmt.Sprintf(`UPDATE %s SET phone_number = REPLACE(REPLACE(REPLACE(REPLACE(phone_number, ' ', ''), '(', ''), ')', ''),'-', '')`, tableName)
	_, err := db.ExecContext(ctx, sql)
	return err
}

// DeleteDuplicates - deleting old duplicates
func DeleteDuplicates(db *sqlx.DB, tableName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	sql := fmt.Sprintf(`DELETE FROM %s T1
    							USING   numbers T2
								WHERE   T1.ctid > T2.ctid  
    							AND T1.phone_number  = T2.phone_number; `, tableName)
	_, err := db.ExecContext(ctx, sql)
	return err
}

func main() {
	var (
		username  = "username"
		dbname    = "dbname"
		password  = "password"
		sslmode   = "disable"
		tableName = "tablename"
	)
	db, err := ConnectDB(fmt.Sprintf("user=%s dbname=%s password=%s sslmode=%s", username, dbname, password, sslmode))
	if err != nil {
		panic(err)
	}
	err = SetDefaultData(db, tableName)
	if err != nil {
		log.Println(err)
		return
	}
	err = FormatPhoneNumbers(db, tableName)
	if err != nil {
		log.Println(err)
		return
	}
	err = DeleteDuplicates(db, tableName)
	if err != nil {
		log.Println(err)
		return
	}
}
