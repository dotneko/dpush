package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/dotneko/dpush/config"
	"github.com/dotneko/dpush/hook"
)

const (
	defaultCfgFile = "dpush.yaml"
	envKey         = "DPUSH_CONFIG"
)

var usage = `
dpush is a Go commandline tool to send a message to a Discord webhook
Usage:
    dpush [webhook alias] [command] [arguments]
Commands:
    msg [message]                Send a message
    pre [message]                Send a pre-formatted code message
    embed [title] [description]  Send an embed
`[1:]

func getEnv(key string) (string, error) {
	val, ok := os.LookupEnv(key)
	if !ok {
		return "", fmt.Errorf("%s not set", key)
	}
	return val, nil
}

func exit(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, filepath.Base(os.Args[0])+": "+format+"\n", a...)
	fmt.Print("\n" + usage)
	os.Exit(1)
}

func help() {
	fmt.Print(usage)
	os.Exit(0)
}

func main() {
	// Check if environment variable set for location of config file
	cfgFile, err := getEnv(envKey)
	if err != nil {
		cfgFile = defaultCfgFile
	}
	err = config.ReadConfig(cfgFile)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	if len(os.Args) < 3 {
		help()
	}

	// Show help if -h[elp] appears anywhere before we do anything else.
	for _, f := range os.Args[1:] {
		switch f {
		case "help", "-h", "-help", "--help":
			help()
		}
	}

	webhook, cmd, args := os.Args[1], os.Args[2], os.Args[3:]
	botname, url, err := config.GetWebhook(webhook)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	switch cmd {
	default:
		exit("unknown command: %q", cmd)
	case "embed":
		hook.Embed(botname, url, args...)
	case "msg":
		hook.Message(botname, url, args...)
	case "pre":
		hook.Pre(botname, url, args...)
	}
}
