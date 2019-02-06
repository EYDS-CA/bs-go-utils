package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/FreshworksStudio/bs-go-utils/api"
	"github.com/FreshworksStudio/bs-go-utils/apiEntity"
	"github.com/FreshworksStudio/bs-go-utils/game"
	"github.com/FreshworksStudio/bs-go-utils/lib"
)

func main() {
	http.HandleFunc("/", Index)
	http.HandleFunc("/start", Start)
	http.HandleFunc("/move", Move)
	http.HandleFunc("/end", End)
	http.HandleFunc("/ping", Ping)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	// Add filename into logging messages
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	log.Printf("Running server on port %s...\n", port)
	http.ListenAndServe(":"+port, lib.LoggingHandler(http.DefaultServeMux))
}

func Index(res http.ResponseWriter, req *http.Request) {
	res.WriteHeader(http.StatusOK)
	res.Write([]byte("Battlesnake documentation can be found at <a href=\"https://docs.battlesnake.io\">https://docs.battlesnake.io</a>."))
}

func Start(res http.ResponseWriter, req *http.Request) {
	decoded := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &decoded)
	if err != nil {
		log.Printf("Bad start request: %v", err)
	}

	lib.Respond(res, api.StartResponse{
		Color: "#75CEDD",
	})
}

func Move(res http.ResponseWriter, req *http.Request) {
	decoded := api.SnakeRequest{}
	err := api.DecodeSnakeRequest(req, &decoded)
	if err != nil {
		log.Printf("Bad move request: %v", err)
	}

	manager := game.InitializeBoard(&decoded)
	log.Printf("Bad move request: %v", manager.OurHead)
	path, err := manager.FindPath(manager.OurHead, apiEntity.Coord{1, 3})
	if err != nil {
		fmt.Printf("ERROR")
	}
	for i := 0; i < len(path)-1; i++ {
		fmt.Printf("%v\n", path[i])
	}
	fmt.Printf("%v\n", lib.DirectionFromCoords(path[0], path[1]))
	lib.Respond(res, api.MoveResponse{Move: apiEntity.Down})
}

func End(res http.ResponseWriter, req *http.Request) {
	lib.Respond(res, api.EmptyResponse{})
}

func Ping(res http.ResponseWriter, req *http.Request) {
	lib.Respond(res, api.EmptyResponse{})
}
