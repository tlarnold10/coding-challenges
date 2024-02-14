package main

import (
  "fmt"
  "flag"
  "os"
  "strings"
  "github.com/bwmarrin/discordgo"
  "os/signal"
  "syscall"
  "io/ioutil"
  "net/http"
  "encoding/json"
  "math/rand"
  "golang.org/x/net/html"
)

type Quote struct {
  Id      int
  Quote   string
  Author  string
}

type Challenge struct {
  Name  string `json:"name"`
  Url   string `json:"url"`
}

type Challenges struct {
  Challenges []Challenge `json:"challenges"`
}

var TOKEN string

func init() {
  flag.StringVar(&TOKEN, "t", "", "Bot Token")
  flag.Parse()
}

// If the user just types "Hello"
func sayHi(s *discordgo.Session, m *discordgo.MessageCreate, author string) {
  s.ChannelMessageSend(m.ChannelID, "Hello " + author)
}

// If the user types "!quote"
func randomQuote(s *discordgo.Session, m *discordgo.MessageCreate) {
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
}

// if the user types "!challenge"
func randomChallenge(s *discordgo.Session, m *discordgo.MessageCreate) {
  jsonFile, err := os.Open("challenges.json")
  if err != nil {
    fmt.Println("Error reading challenges.json file")
    return
  }
  var challenges Challenges
  challengesData, _ := ioutil.ReadAll(jsonFile)
  err = json.Unmarshal(challengesData, &challenges)
  if err != nil {
    fmt.Println("Error in parsing json to a []Challenge")
    return
  }
  max := len(challenges.Challenges)
  random := rand.Intn(max) - 1
  randomPick := challenges.Challenges[random]
  newMessage := randomPick.Name + ": " + randomPick.Url
  s.ChannelMessageSend(m.ChannelID, newMessage)
}

// if the user types "!list"
func getAllChallenges(s *discordgo.Session, m *discordgo.MessageCreate) {
  jsonFile, err := os.Open("challenges.json")
  if err != nil {
    fmt.Println("Error reading challenges.json file")
    return
  }
  var challenges Challenges
  challengesData, _ := ioutil.ReadAll(jsonFile)
  err = json.Unmarshal(challengesData, &challenges)
  if err != nil {
    fmt.Println("Error in parsing json to a []Challenge")
    return
  }
  newMessage := ""
  max := len(challenges.Challenges)
  for i := 0; i < max; i++ {
    if (i == max - 1) {
      newMessage = newMessage + challenges.Challenges[i].Name + ": " + challenges.Challenges[i].Url
    } else {
      newMessage = newMessage + challenges.Challenges[i].Name + ": " + challenges.Challenges[i].Url + "\n"
    }
  }
  s.ChannelMessageSend(m.ChannelID, newMessage)
}

// if the user types "!add"
func addChallenge(s *discordgo.Session, m *discordgo.MessageCreate, url string) {
  resp, err := http.Get(url)
  if err != nil {
    newMessage := "Unable to add: " + url + " please check it is a valid Coding Challenge url"
    s.ChannelMessageSend(m.ChannelID, newMessage)
    return
  }
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    fmt.Println("Error in parsing response from dummyjson quote: ", err)
    return
  }
  doc, err := html.Parse(strings.NewReader(string(body)))
  if err != nil {
    fmt.Println("Error in parsing body: ", err)
    return
  }
  nextPass := false
  newMessage := ""
  var f func(*html.Node, bool)
  f = func(n *html.Node, nextPass bool) {
    if n.Type == html.ElementNode && n.Data == "title" {
      nextPass = true
    } else if (nextPass) {
      newMessage = "Added: " + n.Data + ": " + url
      nextPass = false
    }
    for c := n.FirstChild; c != nil; c = c.NextSibling {
      f(c, nextPass)
    }
  }
  f(doc, nextPass)
  s.ChannelMessageSend(m.ChannelID, newMessage)
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

// listens for messages from discord and responds
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
  if m.Author.ID == s.State.User.ID {
    return
  }

  if (m.Content == "Hello") {
    sayHi(s, m, m.Author.Username) 
  } else if (m.Content == "!quote") {
    randomQuote(s, m)
  } else if (m.Content == "!challenge") {
    randomChallenge(s, m)
  } else if (m.Content == "!list") {
    getAllChallenges(s, m)
  } else if (m.Content[:4] == "!add") {
    if (len(strings.Split(m.Content, " ")) < 2) {
      missingUrl := "Please provide a url with the \"!add\" command"
      fmt.Println(missingUrl)
      s.ChannelMessageSend(m.ChannelID, missingUrl)
      return
    }
    newUrl := strings.Split(m.Content, " ")[1]
    addChallenge(s, m, newUrl)
  }
}
