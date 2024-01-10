package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

var mg MongoInstance

const dbName = "fiber-hrms"
const mongoURI = "mongodb://localhost:27017/" + dbName

//	ID is different in JSON and MongoDB, so define additional bson
//
// For Others (name, age, salary), mongodb column name is same as in json
type Employee struct {
	ID     string `json : "id, omitempty" bson:"_id,omitempty"`
	Name   string `json : "name"`
	Age    int    `json : "age"`
	Salary int    `json : "salary"`
}

func Connect() error {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		fmt.Printf("Can't crete new client mongo %v", err)
		return err
	}
	// mongo db timeout to avoid blocking call
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		fmt.Printf("Can't Connect client mongo %v", err)
		return err
	}

	db := client.Database(dbName)

	mg = MongoInstance{Client: client, Db: db}

	return nil
}

func setupRoutes(app *fiber.App) {
	app.Get("/employee", GetAllEmployees)
	app.Get("/employee/:id", GetEmployee)
	app.Post("/employee/", NewEmployee)
	app.Put("/employee/:id", UpdateEmployee)
	app.Delete("employee/:id", DeleteEmployee)
}

func GetAllEmployees(c *fiber.Ctx) error {

	query := bson.D{{}}
	cursor, err := mg.Db.Collection("employees").Find(c.Context(), query)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}
	var employees []Employee = make([]Employee, 0)
	if err := cursor.All(c.Context(), &employees); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	return c.JSON(employees)
}

func GetEmployee(c *fiber.Ctx) error {
	empId, err := primitive.ObjectIDFromHex(c.Params("id"))
	query := bson.D{{Key: "_id", Value: empId}}

	var employee Employee
	err = mg.Db.Collection("employees").FindOne(c.Context(), query).Decode(&employee)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.SendStatus(404)
		}
		return c.Status(500).SendString(err.Error())
	}

	return c.Status(200).JSON(employee)

}

func NewEmployee(c *fiber.Ctx) error {
	collection := mg.Db.Collection("employees")

	employee := new(Employee)

	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	employee.ID = ""

	insertionRes, err := collection.InsertOne(c.Context(), employee)

	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: insertionRes.InsertedID}}
	createdRec := collection.FindOne(c.Context(), filter)

	createdEmployee := &Employee{}
	createdRec.Decode(createdEmployee)
	return c.Status(201).JSON(createdEmployee)
}

func UpdateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	// query := bson.D{{Key : id}}
	// cursor, err := collection := mg.Db.Collection("employees").Find(c.Context, query)
	// if err != nil {
	// 	return c.Status(500).SendString(err.Error())
	// }

	empId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.SendStatus(400)
	}

	employee := new(Employee)

	if err := c.BodyParser(employee); err != nil {
		return c.Status(400).SendString(err.Error())
	}

	query := bson.D{{Key: "_id", Value: empId}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: employee.Name}, {Key: "age", Value: employee.Age}, {Key: "salary", Value: employee.Salary}}}}
	err = mg.Db.Collection("employees").FindOneAndUpdate(c.Context(), query, update).Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(400).SendString(err.Error())
		}
		return c.SendStatus(500)
	}
	employee.ID = id

	return c.Status(200).JSON(employee)

}

func DeleteEmployee(c *fiber.Ctx) error {
	eId, err := primitive.ObjectIDFromHex(c.Params("id"))

	if err != nil {
		return c.SendStatus(400)
	}
	query := bson.D{{Key: "_id", Value: eId}}

	result, err := mg.Db.Collection("employees").DeleteOne(c.Context(), &query)
	if err != nil {
		return c.SendStatus(500)
	}

	if result.DeletedCount < 1 {
		return c.SendStatus(404)
	}

	return c.Status(200).SendString("record Deleted")
}

func main() {
	app := fiber.New()
	setupRoutes(app)

	Connect()

	if err := app.Listen(":7777"); err != nil {
		fmt.Printf("Error while starting server: %v", err)
		panic(err)
	}

}
