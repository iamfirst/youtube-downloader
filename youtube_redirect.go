package main

import (
	"io"
	"net/http"
	"os"
	"time"
)

type Relay struct{}

var client = http.Client{}

func LogEvent(entry string) {
	file, err := os.OpenFile("log.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil { panic(err) }
	file.WriteString(entry)
	err = file.Close()
	if err != nil { panic(err) }
}

func (relay Relay) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	relativeUrl := request.URL
	urlString := "https://r2---sn-jvhja5g3-hoar.googlevideo.com/" + relativeUrl.String()
	logString := request.RemoteAddr + urlString + time.Now().String()
	LogEvent(logString)
	response, _ := client.Get(urlString)
	//writer.Write([]byte("sss"))
	io.Copy(writer, response.Body)
	response.Body.Close()
}

func CheckLog() {
	_, err := os.Lstat("log.txt")
	if err != nil {
		os.Create("log.txt")
	}
}

func main() {
	CheckLog()
	relay := Relay{}
	http.Handle("/", relay)
	http.ListenAndServe(":4000", nil)
}
