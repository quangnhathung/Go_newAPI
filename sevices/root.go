package sevices

import (
	"gorm.io/gorm"
	"go_API/internal/model"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Root(db *gorm.DB) {
	app := fiber.New()
	app.Use(cors.New()) // Accept request from React

	app.Get("/api/v1/vocabs", func(c *fiber.Ctx) error {
		categoryID := c.Query("cate")
		
		if categoryID == "" {
			var vocabs []model.Vocab
			if err := db.Find(&vocabs).Error; err != nil {
				return c.Status(500).JSON(fiber.Map{"error": "Không thể lấy dữ liệu"})
			}
			return c.JSON(vocabs)
		}

		var category model.Category
		if err := db.Preload("Vocabs").First(&category, categoryID).Error; err != nil {
			return c.Status(404).JSON(fiber.Map{"error": "Category not found"})
		}

		return c.JSON(category.Vocabs)
	})

	app.Listen(":3000")
}
