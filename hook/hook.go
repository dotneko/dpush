package hook

import (
	"fmt"
	"strings"

	"github.com/gtuk/discordwebhook"
)

func Embed(botname, url string, args ...string) error {
	if len(args) < 2 {
		return fmt.Errorf("must specify at least [title] and [description]")
	}
	title := args[0]
	description := strings.Join(args[1:], "\n")
	embed := discordwebhook.Embed{
		Title:       &title,
		Description: &description,
	}

	message := discordwebhook.Message{
		Username: &botname,
		Embeds:   &[]discordwebhook.Embed{embed},
	}

	err := discordwebhook.SendMessage(url, message)
	return err
}

func Message(botname, url string, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("must specify at least one message")
	}
	content := strings.Join(args, " ")

	message := discordwebhook.Message{
		Username: &botname,
		Content:  &content,
	}

	err := discordwebhook.SendMessage(url, message)
	return err
}

func Pre(botname, url string, args ...string) error {
	if len(args) < 1 {
		return fmt.Errorf("must specify at least one message")
	}
	content := "`" + strings.Join(args, " ") + "`"
	message := discordwebhook.Message{
		Username: &botname,
		Content:  &content,
	}

	err := discordwebhook.SendMessage(url, message)
	return err
}
