package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

// type ResponseData struct {
// 	City string `json:"name"`
// }

const prefix string = "!gobot"
const weather string = "!weather"
var city string

func main() {


	session, err := discordgo.New("Bot MTEzMjc5OTE2ODI5ODA5ODgwMA.G0Xb_U.cbqXVj_71r94x0RfPjomLYzPUFZQu6x-gOACQY")
	if err != nil{
		log.Fatal(err)
	}
	
	
	if err != nil{
		log.Fatal(err)
	}

	session.AddHandler(func(s *discordgo.Session, m *discordgo.MessageCreate)  {
		if m.Author.ID == s.State.User.ID {
			return
		}



		fmt.Println(m.Content)
		
		if strings.HasPrefix(m.Content, weather) {
			location := strings.TrimSpace(strings.TrimPrefix(m.Content, weather))
			url := fmt.Sprintf("http://api.weatherapi.com/v1/current.json?key=e872b5e879f04cb393a54523230505&q=%s&aqi=no", location)
			weatherData, err := getWeatherData(url)
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			formatedData := parseWeather(weatherData)
			s.ChannelMessageSend(m.ChannelID, formatedData)
		}






	})

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



func parseWeather(data *WeatherData) string{
	condition := data.Current.Condition.Text
	temp := fmt.Sprintf("%.1f°C\n", data.Current.TemperatureC)
	uv := fmt.Sprintf("%.1f UV\n", data.Current.UVIndex)
	wind := fmt.Sprintf("%.1f %s\n", data.Current.WindKph, data.Current.WindDirection)
	location := fmt.Sprintf("%s, %s, %s\n", data.Location.Name, data.Location.Region, data.Location.Country)
	time := fmt.Sprintf("Local Time: %s\n", data.Location.Localtime)

	return condition + "\n" + temp + uv + wind + location + time
}