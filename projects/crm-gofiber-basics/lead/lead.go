package lead

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"github.com/rajuuu1992/crm-gofiber-basics/database"
)

type Lead struct {
	gorm.Model
	Name  string `json: "Name"`
	Phone int    `json: "Phone"`
	Email string `json: "email"`
}

func GetLeads(c *fiber.Ctx) error {
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func GetLead(c *fiber.Ctx) error {
	res := utils.CopyString(c.Params("id"))

	var getLead Lead
	db.Find(&getLead, id)
	c.JSON(lead)
}

func NewLead(c *fiber.Ctx) error {
	params := c.Params("id")

	lead := new(Lead)

	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	database.DBConn.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) error {
	params := c.Params("id")

	// var lead Lead
	// db.Where("ID=?", ID).Delete(lead)

	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send("No lead found with Id")
		return
	}
	db.Delete(&lead)
	c.Send("Lead deleted successfully")
}
