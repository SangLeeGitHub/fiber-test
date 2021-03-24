package main

import (
	"log"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"database/sql"

	_ "github.com/lib/pq"
)
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "restaurant1"
)

type Employee struct {
	UserID   string `json:"userId"`
	UserName string `json:"userName"`
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected!")

	app := fiber.New()

	app.Get("/", func (c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})
	app.Get("/employee", func(c *fiber.Ctx) error {
		emps := make([]*Employee, 0)
		rows, err := db.Query("select userid, username from employee")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			emp := new(Employee)
			if err := rows.Scan(&emp.UserID, &emp.UserName); err != nil {
				panic(err)
			}
			emps = append(emps, emp)
		}
		if err := rows.Err(); err != nil {
			panic(err)
		}

		return c.Status(fiber.StatusOK).JSON(emps)
	})

	log.Fatal(app.Listen(":3000"))
}