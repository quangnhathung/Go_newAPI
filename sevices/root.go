package sevices

import (
	"gorm.io/gorm"
	"go_API/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Root(db *gorm.DB){
	app := fiber.New()
	app.Use(cors.New())//accept request from react 
	defer app.Listen(":3000")

	app.Get("/api/v1/vocabs", func(c *fiber.Ctx) error {
		categoryID := c.Query("cate")
		if categoryID == "" {
			return c.Status(400).JSON(fiber.Map{"error": "cate parameter is required"})
		}

		var category model.Category
		if err := db.Preload("Vocabs.Categories").First(&category, categoryID).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Category not found"})
		}
		return c.JSON(category.Vocabs)
	})
}
