package lead

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rajuuu1992/crm-gofiber-basics/database"
	"fmt"
)

type Lead struct {
	gorm.Model
	Name  string `json: "Name"`
	Phone int    `json: "Phone"`
	Email string `json: "email"`
}

func GetLeads(c *fiber.Ctx)  error {
	var leads []Lead
	database.DBConn.Find(&leads)
	c.JSON(leads)
	return nil
}

func GetLead(c *fiber.Ctx)  error {
	id := c.Params("id")
	var getLead Lead
	database.DBConn.Find(&getLead, id)
	c.JSON(getLead)
	return nil
}

func NewLead(c *fiber.Ctx) error {
	lead := new(Lead)

	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send([]byte("Can't create new error"))
		return err
	}
	database.DBConn.Create(&lead)
	c.JSON(lead)
	c.Status(200).Send([]byte(fmt.Sprintf("Created New Lead %v", lead)))
	return nil
}

func DeleteLead(c *fiber.Ctx)  error {
	id := c.Params("id")
	var lead Lead
	database.DBConn.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send([]byte("No lead found with Id"))
		return nil
	}
	database.DBConn.Delete(&lead)
	c.Send([]byte("Lead deleted successfully"))
	return nil
}
