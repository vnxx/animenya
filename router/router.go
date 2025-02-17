package router

import (
	"animenya.site/handler"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, handler *handler.Handler) {
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("pong")
	})

	anime := app.Group("/anime")
	anime.Get("/", handler.LatestAnimeEpisode)
	anime.Get("/search", handler.SearchAnime)
	anime.Get("/:anime_id", handler.Anime)
	anime.Get("/:anime_id/cover", handler.AnimeCover)
	anime.Get("/:anime_id/episode/:episode_id", handler.Episode)
}
