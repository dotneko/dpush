package config

import (
	"errors"
	"fmt"
	"log"
	"os"
	"path"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var (
	w []WebhookT
)

type WebhookT struct {
	Alias   string `yaml:"alias"`
	Botname string `yaml:"botname"`
	URL     string `yaml:"url"`
}
type ConfigT struct {
	Default  string     `yaml:"default_webhook"`
	Webhooks []WebhookT `yaml:"webhooks"`
}

func ReadConfig(fpath string) error {
	var config ConfigT

	dir := filepath.Dir(fpath)
	filename := filepath.Base(fpath)
	home, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
		return err
	}
	fileExists := false
	configFilepath := path.Join(dir, filename)
	if _, err := os.Stat(configFilepath); err == nil {
		fileExists = true
	} else if errors.Is(err, os.ErrNotExist) {
		// Try home directory
		configFilepath = path.Join(home, filename)
		if _, err := os.Stat(configFilepath); err == nil {
			fileExists = true
		}
	}
	if !fileExists {
		return fmt.Errorf("config file not found")
	}

	configFile, err := os.ReadFile(configFilepath)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		log.Fatal(err)
		return err
	}
	w = config.Webhooks

	return nil
}

func GetWebhook(alias string) (string, string, error) {
	for _, webhook := range w {
		if alias == webhook.Alias {
			return webhook.Botname, webhook.URL, nil
		}
	}
	return "", "", fmt.Errorf("webhook for %s not found", alias)
}
