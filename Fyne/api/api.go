package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/widget"
)

type WeatherData struct {
	Latitude  float64            `json:"latitude"`
	Longitude float64            `json:"longitude"`
	Hourly    map[string]float64 `json:"hourly"`
}

func main() {
	a := app.New()
	w := a.NewWindow("Météo")

	label := widget.NewLabel("Chargement...")

	go func() {
		resp, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=52.52&longitude=13.41&hourly=temperature_2m")
		if err != nil {
			label.SetText("Erreur lors du chargement des données")
			return
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			label.SetText("Erreur lors de la lecture des données")
			return
		}

		weatherData := WeatherData{}
		err = json.Unmarshal(body, &weatherData)
		if err != nil {
			label.SetText("Erreur lors de la décodage des données")
			return
		}

		label.SetText(fmt.Sprintf("La température à Latitude: %.2f° et Longitude: %.2f° est de %.2f°C", weatherData.Latitude, weatherData.Longitude, weatherData.Hourly["temperature_2m"]))
	}()

	w.SetContent(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), label))

	w.ShowAndRun()
}
