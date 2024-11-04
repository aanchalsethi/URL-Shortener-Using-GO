package routes

import (
	"url-shortener-using-go/api/database"

	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
)

func ResolveURL(c *fiber.Ctx) error {
	url := c.Params("url")

	r := database.CreateClient(0)
	defer r.Close()

	value, err := r.Get(database.Ctx, url).Result()
	if err == redis.Nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "shorten url not found in the database",
		})
	} else if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "cannot connect to the database",
		})
	}

	rInr := database.CreateClient(1)
	defer rInr.Close()

	_ = rInr.Incr(database.Ctx, "counter")

	return c.Redirect(value, 301)
}
