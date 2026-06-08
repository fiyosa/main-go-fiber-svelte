package logger_repository

import (
	"bufio"
	"encoding/json"
	"os"
	"strings"

	"go-fiber-svelte/internal/helper"

	"github.com/gofiber/fiber/v2"
)

type logEntry struct {
	Level   string `json:"level"`
	Time    string `json:"time"`
	Message string `json:"message"`
}

type logDetail struct {
	Name  string     `json:"name"`
	Total int        `json:"total"`
	Logs  []logEntry `json:"logs"`
}

func LogDetailRepository(c *fiber.Ctx) error {
	filename := c.Params("filename")
	if filename == "" || strings.Contains(filename, "..") {
		return c.Status(fiber.StatusBadRequest).JSON(helper.Res.Error("Invalid filename", nil))
	}

	file, err := os.Open("./logs/" + filename)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(helper.Res.Error("Log file not found", nil))
	}
	defer file.Close()

	var entries []logEntry
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var entry logEntry
		if err := json.Unmarshal([]byte(line), &entry); err != nil {
			continue
		}
		entries = append(entries, entry)
	}

	if err := scanner.Err(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(helper.Res.Error("Failed to read log file", nil))
	}

	for i, j := 0, len(entries)-1; i < j; i, j = i+1, j-1 {
		entries[i], entries[j] = entries[j], entries[i]
	}

	return c.JSON(helper.Res.SuccessData(logDetail{
		Name:  filename,
		Total: len(entries),
		Logs:  entries,
	}, "Log detail retrieved successfully"))
}
