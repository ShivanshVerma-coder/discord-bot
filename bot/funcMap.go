package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

var commands = map[string]string{
	"joshi bsdk":                "bol bsdk",
	"joshi dance":               "tu dance kar bsdk",
	"joshi fuckoff":             "gali kahe diye bsdk",
	"kaise ho":                  "Ab kya bolu guyzz!!",
	"joshi op":                  "Thanx guyzz!!",
	"lol":                       "Has kya raha hai bsdk",
	"joshi ggl ko gali do":      "ggl gandu bsdk ",
	"bye":                       "Bye guyzzz",
	"joshi joke sunao":          "abhi nahi badme",
	"joshi ajke contests batao": "Custom",
	"today's contests":          "Custom",
}

func FuncMap(s *discordgo.Session, m *discordgo.MessageCreate) {
	content := strings.ToLower(m.Content)
	if content == "joshi commands" {
		SendCommandsHandler(s, m)
		return
	}

	if _, ok := commands[content]; ok { // if the message is in the map
		if commands[content] != "Custom" {
			SendResponse(s, m, commands[content])
		} else {
			//Custom responses
			if content == "joshi ajke contests batao" || content == "today's contests" {
				SendContests(s, m)
			}
		}
	} else {
		if strings.Contains(content, "joshi") {
			SendResponse(s, m, "Pata nahi bhai")
		}
		return //returns if no match found
	}

}
