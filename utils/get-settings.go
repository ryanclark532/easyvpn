package utils

import (
	"easyvpn/settings/settings_dtos"
	"encoding/json"
	"io"
	"os"
)

func GetSettings() (*settings_dtos.Settings, error) {
	js, err := os.Open(`./settings.json`)
	if err != nil {
		return nil, err
	}
	defer js.Close()
	jsonBytes, err := io.ReadAll(js)
	if err != nil {
		return nil, err
	}
	var settings settings_dtos.Settings
	err = json.Unmarshal(jsonBytes, &settings)
	return &settings, err
}
