# dpush

Discord webhook messaging for the command line

## Credits

Thank you to [@gtuk](https://github.com/gtuk) for his excellent and simple to use golang library:
[discordwebhook](https://github.com/gtuk/discordwebhook)

## Prerequisites

- Go v1.18+
- Discord Channel Webhook URL

To create a webhook, refer to Discord's documentation here:

https://support.discord.com/hc/en-us/articles/228383668-Intro-to-Webhooks

## Installation

Clone this repo

```
git clone https://github.com/dotneko/dpush.git
```

Change into the `dpush` directory then run `go build .`

## Configuration

`dpush` will check for a configuration file `dpush.yaml` in the current directory or home directory by default.

A custom location/filename can be set using the environmental variable `DPUSH_CONFIG=/path/to/my_config.yaml`

```
webhooks:
  - alias: 'webhook1'
    botname: 'someBot1'
    url: 'https://discord.com/api/webhooks/999999999999999999/zzz...'
  - alias: 'webhook2'
    botname: 'someBot2'
    url: 'https://discord.com/api/webhooks/111111111111111111/aaa...'
```
## Usage

```
Usage:
    dpush [webhook alias] [command] [arguments]
Commands:
    msg [message]                Send a message
    pre [message]                Send a pre-formatted code message
    embed [title] [description]  Send an embed
```

## Examples

### Simple message

Send a simple message to a pre-configured `webhook1`

```
dpush webhook1 msg Some message here...
```

### Pre-formatted code

Send pre-formatted code

```
dpush webhook1 pre Testing 123
```

### Simple embed

Send a simple embedded message consisting of *title* and *description*

Note:

- Use quotation marks for longer titles/description containing spaces
- Newlines will be added between each word unless enclosed in quotes

```
dpush webhook1 embed Title "Some description..."
```

```
dpush webhook1 embed "A longer title" "Some description or contents"
```

Discord markdown formatting should technically work, but remember to use a backslash for quotes and backticks

```
dpush webhook1 embed "Testing title" "Line 1" "Line 2" \
    "\`Preformated line 3\`" "**Bolded text** on line 4"
```