package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	tmio "github.com/Johnnycyan/go-tmio-sdk"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide the port number")
	}
	port := os.Args[1]
	http.HandleFunc("/", getLeaderboardRank)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func getLeaderboardRank(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprint(w, "User not found")
		}
	}()
	username := r.URL.Query().Get("username")
	if username == "" {
		http.Error(w, "username is required", http.StatusBadRequest)
		return
	}
	campaign := r.URL.Query().Get("campaign")
	if campaign == "" {
		http.Error(w, "campaign is required", http.StatusBadRequest)
		return
	}

	rank, err := tmio.GetPlayerCampaignRank(username, campaign)
	if err != nil {
		log.Println(err)
		http.Error(w, "Player not found in top 500", http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, rank)
}
