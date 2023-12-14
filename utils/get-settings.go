package utils

import (
	"easyvpn/settings/settings_dtos"
	"encoding/json"
	"os"
)

var Settings *settings_dtos.Settings

func GetSettings() error {
	js, err := os.ReadFile(`./settings.json`)
	if err != nil {
		return err
	}
	var settings settings_dtos.Settings
	err = json.Unmarshal(js, &settings)
	Settings = &settings
	return err
}
