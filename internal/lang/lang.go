package lang

import (
	"encoding/json"
	"go-fiber-svelte/internal/config"
	"os"
	"strings"
)

var translations map[string]string

func Init() {
	locale := config.APP_Locale
	if locale == "" {
		locale = "en"
	}
	file, err := os.ReadFile("internal/lang/locales/" + locale + ".json")
	if err != nil {
		translations = make(map[string]string)
		return
	}
	json.Unmarshal(file, &translations)
}

func T(key string, args ...map[string]string) string {
	msg, ok := translations[key]
	if !ok {
		return key
	}
	if len(args) > 0 {
		for k, v := range args[0] {
			msg = strings.ReplaceAll(msg, ":"+k, v)
		}
	}
	return msg
}
