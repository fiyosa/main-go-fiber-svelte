package logger_repository

import (
	"os"
	"sort"
	"strings"

	"go-fiber-svelte/internal/helper"

	"github.com/gofiber/fiber/v2"
)

func LogListRepository(c *fiber.Ctx) error {
	dir := "./logs"
	entries, err := os.ReadDir(dir)
	if err != nil {
		return c.JSON(helper.Res.SuccessData([]string{}, "Log files retrieved successfully"))
	}

	var files []string
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".log") {
			continue
		}
		files = append(files, e.Name())
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i] > files[j]
	})

	return c.JSON(helper.Res.SuccessData(files, "Log files retrieved successfully"))
}
