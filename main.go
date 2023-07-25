package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

// type ResponseData struct {
// 	City string `json:"name"`
// }

const prefix string = "!gobot"
const weather string = "!weather"

func main() {

	err := godotenv.Load()

	TOKEN := os.Getenv("TOKEN")

	if err != nil {
    	log.Fatalf("Error loading .env file: %s", err)
	}

	bot := fmt.Sprintf("Bot %s", TOKEN)

	fmt.Println(bot)

	session, err := discordgo.New(bot)
	if err != nil{
		log.Fatal(err)
	}
	
	
	if err != nil{
		log.Fatal(err)
	}

	session.AddHandler(handleMessage)

	session.Identify.Intents = discordgo.IntentsAllWithoutPrivileged

	err = session.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer session.Close()

	fmt.Println("The Bot Is Online!")

	sessionConn := make(chan os.Signal, 1)
	signal.Notify(sessionConn, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sessionConn
}



func handleMessage(s *discordgo.Session, m *discordgo.MessageCreate) {
	defer func() {
		if err := recover(); err != nil{

			/* MAKE THIS BETTER, Hard Coded this
			
			Sends error that location needs to be provided, 
			but what if theres diferent error, 
			that message will be sent for all errors
			
			
			*/
			

			log.Println(err)

			fmt.Println("Provide a Location")
			embed := &discordgo.MessageEmbed{
				Title: "Error",
				Description: "Please Provide a Location! Thanks ;)",
			}
	
			s.ChannelMessageSendEmbed(m.ChannelID, embed)
			
		}
	}()

	if m.Author.ID == s.State.User.ID {
		// means that the bot itself sent the message, and there's no need for the bot to process its own messages.
		return
	}

	if m.Content == prefix + " help" {
		s.ChannelMessageSend(m.ChannelID, "Here's help")
	}

	if m.Content == "!rules" {
		s.ChannelMessageSend(m.ChannelID, "Rules...")
	}


	fmt.Println(m.Content)
	
	if strings.HasPrefix(m.Content, weather) {
		location := strings.TrimSpace(strings.TrimPrefix(m.Content, weather))
		url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=e872b5e879f04cb393a54523230505&q=%s&aqi=no", location)
		weatherData, err := getWeatherData(url)

		if err != nil {
			fmt.Println("Error:", err)
		}

		

		formatedData := parseWeather(weatherData)
		// s.ChannelMessageSend(m.ChannelID, formatedData)

		embed := &discordgo.MessageEmbed{
			Title: "Current Weather",
			Description: formatedData,
			Color: 0x00ff00,
		}

		s.ChannelMessageSendEmbed(m.ChannelID, embed)
	}



}



func parseWeather(data *WeatherData) string{
	location := fmt.Sprintf("%s, %s, %s\n", data.Location.Name, data.Location.Region, data.Location.Country)
	time := fmt.Sprintf("Local Time: %s\n", data.Location.Localtime)
	condition := data.Current.Condition.Text
	temp := fmt.Sprintf("%.1f°C\n", data.Current.TemperatureC)
	uv := fmt.Sprintf("%.1f UV\n", data.Current.UVIndex)
	wind := fmt.Sprintf("%.1f %s\n", data.Current.WindKph, data.Current.WindDirection)

	return location + time + condition + "\n" + temp + uv + wind
}