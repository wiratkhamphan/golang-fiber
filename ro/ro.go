package ro

import (
	"database/sql"
	"lek/bd"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

var db *sql.DB

type album struct {
	D_id     int    `json:"d_id"`
	Name_id  int    `json:"name_id"`
	Fname    string `json:"fname"`
	Username string `json:"username"`
}

func Router() {
	router := fiber.New()

	router.Use(cors.New())

	router.Get("/Get", Get)
	router.Post("/Post", Post)

	router.Listen(":8080") // Use Listen instead of Run for Fiber v2
}

func Get(c *fiber.Ctx) error {
	db, err := bd.BdconMysql()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close() // Close the database connection when done

	var res []album
	rows, err := db.Query("SELECT * FROM employees_id")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var temp album
		err := rows.Scan(&temp.D_id, &temp.Name_id, &temp.Fname, &temp.Username)
		if err != nil {
			log.Fatal(err)
		}
		res = append(res, temp)
	}

	return c.JSON(res)
}

func Post(c *fiber.Ctx) error {
	db, err := bd.BdconMysql()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Example of how to handle incoming JSON data
	var data album
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	// Example query to insert data into the database
	_, err = db.Exec("INSERT INTO employees_id (d_id, name_id, Fname, Username) VALUES (?, ?, ?, ?)",
		data.D_id, data.Name_id, data.Fname, data.Username)
	if err != nil {
		log.Fatal(err)
		return err
	}

	return c.SendString("Data inserted successfully")
}
