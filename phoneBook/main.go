package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "modernc.org/sqlite"
)

type Contact struct {
	id     int
	name   string
	number string
}

func main() {
	db, err := sql.Open("sqlite", "./database.db")
	if err != nil {
		log.Fatal(err)
	}

	initTables(db)
	for {
		if err := showContacts(db); err != nil {
			log.Fatal(err)
		}
		if err := execCommand(db); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

func initTables(db *sql.DB) error {
	const sql = `
		CREATE TABLE IF NOT EXISTS contacts (
			id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			number INTEGER NOT NULL
		);`
	_, err := db.Exec(sql)
	return err
}

func showContacts(db *sql.DB) error {
	rows, err := db.Query("SELECT * FROM contacts;")
	if err != nil {
		return err
	}
	fmt.Printf("--------------------------------------------------\n")
	fmt.Printf("| id   | name               | number             |\n")
	fmt.Printf("--------------------------------------------------\n")
	for rows.Next() {
		var c Contact
		if err := rows.Scan(&c.id, &c.name, &c.number); err != nil {
			return err
		}
		fmt.Printf("|%6d|%19s |%19s |\n", c.id, c.name, c.number)
	}
	fmt.Printf("--------------------------------------------------\n")
	return nil
}

func scanTextContent(contentName string) string {
	sc := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s >> ", contentName)
		sc.Scan()
		if len(sc.Text()) != 0 {
			return sc.Text()
		}
		fmt.Fprintln(os.Stderr, "invalid value")
	}
}

func scanIntegerContent(contentName string) int {
	for {
		res, err := strconv.Atoi(scanTextContent(contentName))
		if err == nil {
			return res
		}
		fmt.Fprintln(os.Stderr, "invalid value")
	}
}

func execCommand(db *sql.DB) error {
	sc := bufio.NewScanner(os.Stdin)

	fmt.Print(">> ")
	sc.Scan()

	cmd := sc.Text()
	err := error(nil)
	switch cmd {
	case "add":
		err = commandAdd(db)
	case "delete":
		err = commandDelete(db)
	case "":
		// nop
	default:
		commandHelp()
	}
	return err
}

func commandDelete(db *sql.DB) error {
	id := scanIntegerContent("id")
	return deleteContact(db, id)
}

func deleteContact(db *sql.DB, id int) error {
	const sql = "DELETE FROM contacts WHERE id=?"
	_, err := db.Exec(sql, id)
	return err
}

func commandAdd(db *sql.DB) error {
	name := scanTextContent("name")
	number := scanTextContent("phone number")
	return addContact(db, name, number)
}

func addContact(db *sql.DB, name string, number string) error {
	const sql = "INSERT INTO contacts (name, number) values (?, ?)"
	_, err := db.Exec(sql, name, number)
	return err
}

func commandHelp() {
	fmt.Println("commands:")
	fmt.Println("     add: add a new contact")
	fmt.Println("  delete: delete a contact")
}
