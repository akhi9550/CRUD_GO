package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var id int
var name, domain string

func main() {
	var choice int
	db := connectPostgresDB()
	for {
		fmt.Println("\nChoose\n1.Insert\n2.Read\n3.Update\n4.Delete\n5.Exit")
		fmt.Println("Enter Your Choice:")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			Insert(db)
		case 2:
			Read(db)
		case 3:
			Update(db)
		case 4:
			Delete(db)
		case 5:
			os.Exit(0)
		}
	}
}

// db connection
func connectPostgresDB() *sql.DB {
	connectTo := "user=postgres dbname=databasename password='*******' host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connectTo)
	if err != nil {
		fmt.Println(err)
	}
	return db
}
func Insert(db *sql.DB) {
	fmt.Println("Enter the id:")
	fmt.Scan(&id)
	fmt.Println("Enter the name:")
	fmt.Scan(&name)
	fmt.Println("Enter the domain:")
	fmt.Scan(&domain)
	insertIntoPostgres(db, id, name, domain)
}
func insertIntoPostgres(db *sql.DB, id int, name, domain string) {
	_, err := db.Exec("INSERT INTO student(id,name,domain)values($1,$2,$3)", id, name, domain)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value Inserted!!!")
	}
}
func Read(db *sql.DB) {
	rows, err := db.Query("SELECT * FROM student")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("id   name   domain")
		for rows.Next() {
			rows.Scan(&id, &name, &domain)
			fmt.Printf("%d - %s - %s \n", id, name, domain)
		}
	}
}
func Update(db *sql.DB) {
	fmt.Println("Enter the id:")
	fmt.Scan(&id)
	fmt.Println("Enter new name:")
	fmt.Scan(&name)
	fmt.Println("Enter new domain:")
	fmt.Scan(&domain)
	_, err := db.Query("UPDATE student SET name=$1,domain=$2 where id=$3", name, domain, id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Updated!!!")
	}
}
func Delete(db *sql.DB) {
	fmt.Println("Enter the id:")
	fmt.Scan(&id)
	_, err := db.Query("DELETE FROM student WHERE id=$1", id)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Deleted!!!")
	}
}

//postgress:records-student
