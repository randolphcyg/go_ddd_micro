package main

import (
	"encoding/json"
	"log/slog"
	"math/rand"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Res struct {
	AvailableHotels []struct {
		Name               string `json:"name"`
		PriceInUSDPerNight int    `json:"priceInUSDPerNight"`
	} `json:"availableHotels"`
}

func main() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	minIdx := 1
	maxIdx := 10

	sampleRes := Res{AvailableHotels: []struct {
		Name               string `json:"name"`
		PriceInUSDPerNight int    `json:"priceInUSDPerNight"`
	}{
		{
			Name:               "some hotel",
			PriceInUSDPerNight: 300,
		},
		{
			Name:               "some other hotel",
			PriceInUSDPerNight: 30,
		},
		{
			Name:               "some third hotel",
			PriceInUSDPerNight: 90,
		},
		{
			Name:               "some fourth hotel",
			PriceInUSDPerNight: 80,
		},
	}}

	b, err := json.Marshal(sampleRes)
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.Path("/partnerships").HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		ran := rand.Intn(maxIdx - minIdx + 1)
		if ran > 7 {
			writer.WriteHeader(http.StatusInternalServerError)
			return
		}
		writer.WriteHeader(http.StatusOK)
		_, _ = writer.Write(b)
	})

	slog.Info("partnerships service running on:localhost:3031")
	if err := http.ListenAndServe(":3031", r); err != nil {
		panic(err)
	}
}
