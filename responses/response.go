package responses

import "github.com/gofiber/fiber/v2"

// Response testi

type BirdResponse struct {
    Status  int        `json:"status" bson:"status"`
    Message string     `json:"message" bson:"message"`
    Data    *fiber.Map `json:"data" bson:"data"`
}