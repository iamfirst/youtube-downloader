package main

import (
	"./goydl"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

type Relay struct{}

var client = http.Client{}

func (relay Relay) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	relativeUrl := request.URL
	urlString := GetRealVideoUrl(relativeUrl.String())
	if urlString == "" {
		fmt.Println("forbidden")
		return
	}

	logString := request.RemoteAddr + urlString + time.Now().String()
	LogEvent(logString)
	response, _ := client.Get(urlString)
	io.Copy(writer, response.Body)
	response.Body.Close()
}

func LogEvent(entry string) {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil { panic(err) }
	file.WriteString(entry)
	err = file.Close()
	if err != nil { panic(err) }
}

func CheckLog() {
	_, err := os.Lstat("log.txt")
	if err != nil {
		os.Create("log.txt")
	}
}

func GetRealVideoUrl(url string) (VideoUrl string) {
	url = strings.ReplaceAll(url, "/https:/", "https://")
	fmt.Println(url)
	youtubeDl := goydl.NewYoutubeDl()
	youtubeDl.VideoURL = url
	info, err := youtubeDl.GetInfo()

	if err != nil {
		fmt.Println(err)
	}

	len := len(info.Formats)
	if len < 2 {
		fmt.Println("json info is too short")
		return
	}

	for k, v := range info.Formats {
		if v.Ext == "mp4" && v.FormatNote == "360p" && v.Acodec == "mp4a.40.2"{
			VideoUrl = v.URL
			fmt.Println(k, v)
		}
	}

	return VideoUrl
}

func main() {
	CheckLog()
	relay := Relay{}
	http.Handle("/", relay)
	http.ListenAndServe(":4000", nil)
}
