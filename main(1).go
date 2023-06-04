package main

import (
	"context"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/golang-jwt/jwt/v5"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	type User struct {
		Name     string `json:"username" xml:"username" form:"username"`
		Password string `json:"password" xml:"password" form:"password"`
		Role     string `json:"role" xml:"role" form:"role"`
	}

	var sign_key = []byte("hello jwt")
	//mongodb连接
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("user").Collection("user")

	app := fiber.New()
	//一句话解决跨域问题
	app.Use(cors.New())

	app.Post("/login", func(c *fiber.Ctx) error {
		user := new(User)
		if err := c.BodyParser(user); err != nil {
			return err
		}
		query := bson.M{"username": user.Name, "password": user.Password, "role": user.Role}
		var result bson.M
		err := collection.FindOne(context.TODO(), query).Decode(&result)
		if err != nil {
			return c.JSON(fiber.Map{
				"code": 1001,
				"msg":  "deny",
			})
		} else {
			newtoken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
				"username": user.Name,
				"role":     user.Role,
			})
			tokenString, err := newtoken.SignedString(sign_key)
			if err != nil {
				return err
			} else {
				return c.JSON(fiber.Map{
					"code":  1000,
					"msg":   "allow",
					"token": tokenString,
				})
			}
		}
	})

	app.Get("/app", func(c *fiber.Ctx) error {
		return c.SendString("app is good")
	})

	app.Post("/test", func(c *fiber.Ctx) error {
		user := new(User)
		fmt.Println(c.Path())
		if err := c.BodyParser(user); err != nil {
			return err
		}
		fmt.Println(user.Name)
		return c.SendString("username: " + user.Name)
	})

	app.Listen(":8080")
}
