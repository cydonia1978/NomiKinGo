package NomiKin

import (
    "log"
    "os"
)

var Version = "v0.3.0"

type NomiKin struct {
    ApiKey      string
    CompanionId string
    TooLongMessage string
}

var NomiUrlComponents map[string][]string
var KinUrlComponents map[string][]string

func (nomi *NomiKin) Init(companionType string) {
    log.Println("NomiKin Init - version: " + Version)

    KinUrlComponents = make(map[string][]string)
    KinUrlComponents["SendMessage"] = []string {"https://api.kindroid.ai/v1/send-message"}
    KinUrlComponents["ChatBreak"] = []string {"https://api.kindroid.ai/v1/chat-break"}
    KinUrlComponents["DiscordBot"] = []string {"https://api.kindroid.ai/v1/discord-bot"}

    NomiUrlComponents = make(map[string][]string)
    NomiUrlComponents["SendMessage"] = []string {"https://api.nomi.ai/v1/nomis", "chat"}
    NomiUrlComponents["RoomCreate"] = []string {"https://api.nomi.ai/v1/rooms"}
    NomiUrlComponents["RoomReply"] = []string {"https://api.nomi.ai/v1/rooms", "chat/request"}
    NomiUrlComponents["RoomSend"] = []string {"https://api.nomi.ai/v1/rooms", "chat"}
    // load our custom TOO_LONG_MESSAGE (with a default fallback)
    if v := os.Getenv("TOO_LONG_MESSAGE"); v != "" {
        nomi.TooLongMessage = v
    } else {
        nomi.TooLongMessage = "Sorry, your message was too long."
   }
}

