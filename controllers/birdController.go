package controllers

import (
	"context"
	"net/http"
	"time"

	"github.com/henriettelehtonen/GO_API/configs"
	"github.com/henriettelehtonen/GO_API/models"
	"github.com/henriettelehtonen/GO_API/responses"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Yhteys böörd collectioniin
var birdCollection *mongo.Collection = configs.GetCollection(configs.DB, "birds")
var validate = validator.New()


/// CRUDS ***********************************

// CREATE BIRD
func CreateBird(c *fiber.Ctx) error {

	var bird models.Bird
	if err := c.BodyParser(&bird); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid req"})
		// TESTI käyttämään response.go
		// return c.Status(http.StatusBadRequest).JSON(responses.BirdResponse{Status: http.StatusBadRequest, Message: "Invalid req", Data: &fiber.Map{"Bird": err.Error()}})
	}
	// validointi
	if err := validate.Struct(bird); err != nil {
		errs := err.(validator.ValidationErrors)
		validationErrors := make(map[string]string)
		for _, e := range errs {
			validationErrors[e.Field()] = "Field is required!"
		}
		return c.Status(400).JSON(fiber.Map{"validation_errors": validationErrors})
	}

	bird.ID = primitive.NewObjectID()
	_, err := birdCollection.InsertOne(context.Background(), bird)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error while crreating bird"})
	}
	return c.Status(201).JSON(bird)
	// tai käytä response.go
	// return c.Status(http.StatusCreated).JSON(responses.BirdResponse{Status: http.StatusCreated, Message: "Succes!", Data: &fiber.Map{"Bird": bird}})
}

// GET ALL BIRDS
func GetBirds(c *fiber.Ctx) error {
	var birds []models.Bird

	result, err := birdCollection.Find(context.Background(), bson.M{})

	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": "Error while fetching birds"})
	}
	defer result.Close(context.Background())
	for result.Next(context.Background()) {
		var singleBird models.Bird
		if err = result.Decode(&singleBird); err != nil {
			return c.Status(500).JSON("Error while fetching birds")
		}
		birds = append(birds, singleBird)
	}
	return c.Status(200).JSON(birds)
	// tai 
	// return c.Status(http.StatusOK).JSON(responses.BirdResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"Birds": birds}})
}

// GET SIINGLE BIRD
func GetSingleBirdByID(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second) // testi
	defer cancel()

	birdID := c.Params("id")
	var bird models.Bird

	objId, _ := primitive.ObjectIDFromHex(birdID)

	err := birdCollection.FindOne(ctx, bson.M{"_id": objId}).Decode(&bird)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.BirdResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"Bird": err.Error()}})
	}

	return c.Status(http.StatusOK).JSON(responses.BirdResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"Bird": bird}})
}


// UPDATE BIRD -> käytetty patchia
func UpdateBirdByID(c *fiber.Ctx) error {
    idParam := c.Params("id")
    objID, err := primitive.ObjectIDFromHex(idParam)
    if err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
    }

    var updateData map[string]interface{}
    if err := c.BodyParser(&updateData); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Invalid request body"})
    }

    filter := bson.M{"_id": objID}
    update := bson.M{"$set": updateData}

    result, err := birdCollection.UpdateOne(context.Background(), filter, update)
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Update failed"})
    }

    if result.MatchedCount == 0 {
        return c.Status(404).JSON(fiber.Map{"error": "Bird not found"})
    }

	// patchattu lintu
    var updatedBird models.Bird
    if err := birdCollection.FindOne(context.Background(), filter).Decode(&updatedBird); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": "Could not fetch updated bird"})
    }

    return c.Status(200).JSON(updatedBird)
}


// DELETE BIRD BY ID
func DeleteBirdByID(c *fiber.Ctx) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	birdID := c.Params("id")
	objId, _ := primitive.ObjectIDFromHex(birdID)

	result, err := birdCollection.DeleteOne(ctx, bson.M{"_id": objId})
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(responses.BirdResponse{Status: http.StatusInternalServerError, Message: "error", Data: &fiber.Map{"Bird": err.Error()}})
	}

	if result.DeletedCount < 1 {
		return c.Status(http.StatusNotFound).JSON(
			responses.BirdResponse{Status: http.StatusNotFound, Message: "Bird not found!", Data:nil},
		)
	}

	return c.Status(http.StatusOK).JSON(
		responses.BirdResponse{Status: http.StatusOK, Message: "success", Data: &fiber.Map{"Bird": "Bird deleted successfully!"}},
	)
	// tai pelkkä
	// return c.Status(200).JSON(fiber.Map{"message": "Bird deleted successfully!"})
}

