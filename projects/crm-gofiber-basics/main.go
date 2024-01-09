package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	"github.com/rajuuu1992/crm-gofiber-basics/database"
	"github.com/rajuuu1992/crm-gofiber-basics/lead"
	"fmt"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/lead", lead.GetLeads)
	app.Get("/api/v1/lead/:id", lead.GetLead)
	app.Post("/api/v1/lead", lead.NewLead)
	app.Delete("/api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {

	var err error
	database.DBConn, err = gorm.Open("sqlite3", "leads.db")
	if err != nil {
		panic(err)
	}
	fmt.Println("DB Connect success")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated")
}
func main() {
	app := fiber.New()
	initDatabase()
	setupRoutes(app)
	if err := app.Listen(":7777"); err != nil {
		fmt.Printf("Error while starting server: %v", err)
		panic(err)
	}

	defer database.DBConn.Close()

}
