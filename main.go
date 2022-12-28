package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/oscareduardolp6/random-spanish-words-api-go/words"
	"github.com/rs/cors"
)

// Check if the request has a number of words requested, else returns 1
func getSelectedNumberOfWordsFromRequest(request *http.Request) int {
	queryParams := request.URL.Query()
	numberOfWordsText := queryParams.Get("num")
	DEFAULT_NUMBER_OF_WORDS := 1
	MAX_NUMBER := 1082
	if numberOfWordsText == "" {
		return DEFAULT_NUMBER_OF_WORDS
	}
	numberOfWords, err := strconv.Atoi(numberOfWordsText)
	if err != nil {
		return DEFAULT_NUMBER_OF_WORDS
	}
	if numberOfWords >= 1083 {
		return MAX_NUMBER
	}
	return numberOfWords
}

// Returns a map of the console args passed
// ONLY named args are included
func getNamedArgs() map[string]string {
	argsWithoutProgramName := os.Args[1:]
	argsMap := make(map[string]string)
	separator := "="

	for _, argument := range argsWithoutProgramName {
		if strings.Contains(argument, separator) {
			parts := strings.Split(argument, separator)
			key := parts[0]
			value := parts[1]
			argsMap[key] = value
		}
	}
	return argsMap
}

// Returns the port to be used for the server, if an args port is passed
// (PORT=5030) then that port is used, else check if a PORT enviroment variable
// is set, if not, then the default value (8080) is used
func getPort() string {
	const DEFAULT_PORT = "8080"
	namedArguments := getNamedArgs()
	port, existsPort := namedArguments["PORT"]
	if existsPort {
		return port
	}
	port = os.Getenv("PORT")
	if port != "" {
		return port
	}
	return DEFAULT_PORT
}

// Handler of getting a random word from the "dictionary" and return it in the response
func getRandomWord(writer http.ResponseWriter, request *http.Request) {
	numberOfWords := getSelectedNumberOfWordsFromRequest(request)
	wordsArray := make([]string, numberOfWords)

	for index := 0; index < numberOfWords; index++ {
		randomIndex := rand.Intn(len(words.SpanishWords))
		randomWord := words.SpanishWords[randomIndex]
		wordsArray[index] = randomWord
	}
	jsonResult, err := json.Marshal(wordsArray)
	if err != nil {
		panic(err)
	}
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusAccepted)
	writer.Write(jsonResult)
}

// Starts the Server API
func main() {
	initServer()
}

// Initialize a muxServer with default cors, the server automatically
// gets the port from console args(Ex: PORT=5030), enviroment variable or the default (8080),
// that is the order of priority, if a port if passed in args, that port is used over
// enviroment variable port or default port
func initServer() {
	muxServer := http.NewServeMux()
	muxServer.HandleFunc("/", getRandomWord)
	port := ":" + getPort()
	fmt.Printf("Iniciando servidor en el puerto%s", port)
	corsHandler := cors.Default().Handler(muxServer)
	log.Fatal(http.ListenAndServe(port, corsHandler))
}
