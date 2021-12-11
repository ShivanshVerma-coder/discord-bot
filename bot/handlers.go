package bot

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/bwmarrin/discordgo"
)

func SendResponse(s *discordgo.Session, m *discordgo.MessageCreate, response string) {
	s.ChannelMessageSend(m.ChannelID, response)
}

func SendCommandsHandler(s *discordgo.Session, m *discordgo.MessageCreate) {
	var response string = "Commands are:\n"
	index := 1
	for key, _ := range commands {
		response += strconv.Itoa(index) + ")" + "\t" + key + "\n"
		index++
	}
	s.ChannelMessageSend(m.ChannelID, response)
	time.Sleep(time.Second * 1)
	s.ChannelMessageSend(m.ChannelID, "sahi se type karna varna reply nahi dunga")
}

type Contest struct {
	Name        string `json:"name"`
	Url         string `json:"url"`
	Start_time  string `json:"start_time"`
	End_time    string `json:"end_time"`
	Duration    string `json:"duration"`
	Site        string `json:"site"`
	In_24_hours string `json:"in_24_hours"`
	Status      string `json:"status"`
}

func SendContests(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "Kya matlab ab mujhse ye sab bhi karaoge")
	time.Sleep(time.Second * 1)
	s.ChannelMessageSend(m.ChannelID, "Acha Ruko dekhra hu")
	time.Sleep(time.Second * 1)
	var contests []Contest
	resp, err := http.Get("https://kontests.net/api/v1/all")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&contests)
	if err != nil {
		fmt.Println(err)
	}
	var index int = 0
	for _, contest := range contests {
		if contest.In_24_hours == "Yes" {
			fst, err := time.Parse(time.RFC3339, contest.Start_time)

			string_fst := fmt.Sprintf("%02d:%02d", fst.Hour(), fst.Minute())
			if err != nil {
				fmt.Println(err)
			}

			s.ChannelMessageSend(m.ChannelID, strconv.Itoa(index+1)+") "+contest.Name+" starts at "+string_fst+"\n"+contest.Url)
			index++
		}
	}
	time.Sleep(time.Second * 1)
	s.ChannelMessageSend(m.ChannelID, "\nDena hai nahi par contests jan ne hai tumko")
}
