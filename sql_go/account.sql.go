package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "fighting"
	dbname   = "bank"
)

type account struct {
	owner    string
	balance  int
	currency string
}

func main() {
	// 连接数据库
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Unable to connect to the database: %v\n", err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatalf("Unable to reach the database: %v\n", err)
	}

	fmt.Println("Successfully connected!")

	//createAccount(db, `HongWei`, 100, `NT`)
	var Account account = getAccount(db, `HongWei`)

	fmt.Printf("Your Name: %s\nYour balance: %d\nCurrency\n: %s", Account.owner, Account.balance, Account.currency)

}

func createAccount(db *sql.DB, owner string, balance int, currency string) {
	sqlstatement := `INSERT INTO account (owner,balance,currency) VALUES($1,$2,$3) RETURNING id`
	id := 0

	err := db.QueryRow(sqlstatement, owner, balance, currency).Scan(&id)

	if err != nil {
		log.Fatalf("Unable to create the account!Infor:%v\n", err)
	}

	fmt.Printf("Create Successfully!Your ID is %d", id)
}

func getAccount(db *sql.DB, owner string) account {
	sqlstatement := `SELECT owner,balance,currency from account WHERE owner = $1`

	var tmp account = account{}

	err := db.QueryRow(sqlstatement, owner).Scan(&tmp.owner, &tmp.balance, &tmp.currency)

	if err != nil {
		log.Fatalf("Unable to get the account!Infor:%v\n", err)
	}

	return tmp
}

func updateAccount(db *sql.DB, owner string, balance int) {

}

func deleteAccount(db *sql.DB, owner string) {

}
