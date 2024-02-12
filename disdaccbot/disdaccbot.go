package main

import (
  "fmt"
  "flag"
  "os"
  "github.com/bwmarrin/discordgo"
  "os/signal"
  "syscall"
  "io/ioutil"
  "net/http"
  "encoding/json"
)

type Quote struct {
  Id      int
  Quote   string
  Author  string
}

var TOKEN string

func init() {
  flag.StringVar(&TOKEN, "t", "", "Bot Token")
  flag.Parse()
}

func main() {
  fmt.Println("token: " + TOKEN)
  dis, err := discordgo.New("Bot " + TOKEN)
  if err != nil {
    fmt.Println("Error creating Discord session: ", err)
    return  
  }

  dis.AddHandler(messageCreate)
  dis.Identify.Intents = discordgo.IntentsGuildMessages

  err = dis.Open()
  if err != nil {
    fmt.Println("Error opening connection: ", err)
    return
  }
  fmt.Println("Bot is now running. Press ctrl+c to exit")
  sc := make(chan os.Signal, 1)
  signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
  <-sc

  dis.Close()
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
  messageAuthor := m.Author

  if m.Author.ID == s.State.User.ID {
    return
  }

  if (m.Content == "Hello") {
    s.ChannelMessageSend(m.ChannelID, "Hello " + messageAuthor.Username)
  } else if (m.Content == "!quote") {
    resp, err := http.Get("https://dummyjson.com/quotes/random")
    if err != nil {
      fmt.Println("Error in getting dummyjson quote: ", err)
      return
    }
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      fmt.Println("Error in parsing response from dummyjson quote: ", err)
      return
    }
    var randomQuote Quote
    err = json.Unmarshal(body, &randomQuote)
    if err != nil {
      fmt.Println("Error in parsing json to quote: ", err)
      return
    }
    newMessage := "\"" + randomQuote.Quote + "\" - " + randomQuote.Author
    s.ChannelMessageSend(m.ChannelID, newMessage)
  } else if (m.Content == "!challenge") {

  }
}
