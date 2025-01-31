package routes

import (
	"fallback/internal/embeded"
	"net/http"

	"github.com/gofiber/fiber/v2/middleware/filesystem"
)

var staticsHandler = filesystem.New(filesystem.Config{
	Root:       http.FS(embeded.EmbedDirStatic),
	PathPrefix: "statics",
	MaxAge:     60 * 60,
})
