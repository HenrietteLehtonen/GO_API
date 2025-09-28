package controllers

import (
	"context"

	"github.com/henriettelehtonen/GO_API/configs"
	"github.com/henriettelehtonen/GO_API/models"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var uhCollection *mongo.Collection = configs.GetCollection(configs.DB, "uhanalaisuus")
var validateUhanalaisuus = validator.New()

// Get all uhanalaisuus
func GetUhanalaisuus(c *fiber.Ctx) error {
	var uhanalaisuus []models.Uhanalaisuus

	result, err := uhCollection.Find(context.Background(), bson.M{})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error while fetching uhanalaisuus"})
	}
	defer result.Close(context.Background())
	for result.Next(context.Background()) {
		var singleUhanalaisuus models.Uhanalaisuus
		if err = result.Decode(&singleUhanalaisuus); err != nil {
			return c.Status(500).JSON("Error while fetching uhanalaisuus")
		}
		uhanalaisuus = append(uhanalaisuus, singleUhanalaisuus)
	}
	if len(uhanalaisuus) == 0 {
		return c.Status(200).JSON(fiber.Map{"message": "No uhanalaisuus found", "data": []models.Uhanalaisuus{}})
	}
	return c.Status(200).JSON(uhanalaisuus)
	// tai 
	// return c.Status(http.StatusOK).JSON(responses.BirdResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"Birds": birds}})
}