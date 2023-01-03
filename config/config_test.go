package config_test

import (
	"testing"

	"github.com/dotneko/dpush/config"
)

func TestReadConfig(t *testing.T) {
	var exampleFilepath = "../dpush_example.yaml"
	var testConfig1 = config.WebhookT{
		"webhook1", "someBot1",
		"https://discord.com/api/webhooks/999999999999999999/zzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzzz",
	}
	var testConfig2 = config.WebhookT{
		"webhook2", "someBot2",
		"https://discord.com/api/webhooks/111111111111111111/aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
	}

	err := config.ReadConfig(exampleFilepath)
	if err != nil {
		t.Fatal(err)
		return
	}
	botname1, url1, err := config.GetWebhook(testConfig1.Alias)
	if err != nil {
		t.Fatal(err)
	}
	if botname1 != testConfig1.Botname {
		t.Errorf("%s => Expected %s, got %s", testConfig1.Alias, botname1, testConfig1.Botname)
	}
	if url1 != testConfig1.URL {
		t.Errorf("%s => Expected %s, got %s", testConfig1.Alias, url1, testConfig1.URL)
	}
	botname2, url2, err := config.GetWebhook(testConfig2.Alias)
	if err != nil {
		t.Fatal(err)
	}
	if botname2 != testConfig2.Botname {
		t.Errorf("%s => Expected %s, got %s", testConfig2.Alias, botname2, testConfig2.Botname)
	}
	if url2 != testConfig2.URL {
		t.Errorf("%s => Expected %s, got %s", testConfig2.Alias, url2, testConfig2.URL)
	}
}
