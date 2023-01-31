package main

import (
	"encoding/json"
	"fmt"
	"image"
	"io/ioutil"
	"net/http"

	// "time"
	"strings"
	// "strconv"

	"fyne.io/fyne"
	"fyne.io/fyne/app"

	// "fyne.io/fyne/container"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/widget"
)

/*--------------------------------------------------------------------------------------------
-------------------------------------- Type Struct -------------------------------------------
----------------------------------------------------------------------------------------------*/

// Artists struct match artists.json var type and logic
type Artists struct {
	ID           int      `json:"id"`
	IMAGE        string   `json:"image"`
	NAME         string   `json:"name"`
	MEMBERS      []string `json:"members"`
	CREA_DATE    int      `json:"creationDate"`
	FIRST_ALBUM  string   `json:"firstAlbum"`
	LOCATIONS    string   `json:"locations"`
	CONCERT_DATE string   `json:"concertDates"`
	RELATION     string   `json:"relations"`
}

// Locations struct locations.json
type Locations struct {
	ID        int      `json:"id"`
	LOCATIONS []string `json:"locations"`
	DATES     string   `json:"dates"`
}

// Dates struct match dates.json var type and logic
type Dates struct {
	ID    int      `json:"id"`
	DATES []string `json:"dates"`
}

// Realtions struct match relations.json var type and logic
type Relations struct {
	ID         int                 `json:"id"`
	DATESLOCAT map[string][]string `json:"datesLocations"`
}

func main() {
	a := app.New()
	a.SetIcon(theme.FyneLogo())
	w := a.NewWindow("Groupie Tracker")
	w.Resize(fyne.NewSize(800, 600))

	label := widget.NewLabel("Page D'accueil")

	go func() {

		// Get data from API
		resp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
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

		// Decode data
		var artists []Artists
		err = json.Unmarshal(body, &artists)
		if err != nil {
			label.SetText("Erreur lors de la décodage des données")
			return
		}

		// Recover the data in a []string
		var artistsData []string
		for i := 0; i < len(artists); i++ {
			artistsData = append(artistsData, fmt.Sprintf(artists[i].NAME))
		}

		// Create a menu
		menuItem1 := fyne.NewMenuItem("Home", func() {
			label.SetText("Page D'accueil")
		})

		var buttons []fyne.CanvasObject
		container := fyne.NewContainerWithLayout(layout.NewVBoxLayout())

		menuItem2 := fyne.NewMenuItem("Artists", func() {
			label.SetText("Liste des artistes")
			for i, artist := range artistsData {
				newLabel := widget.NewLabel("")
				buttons = append(buttons, widget.NewButton(artist, func(i int, artist string) func() {
				  return func() {
					newLabel.SetText("Membres : \n - " + strings.Join(artists[i].MEMBERS, "\n - ") + "\n" + "Date de création : " + fmt.Sprintf("%d", artists[i].CREA_DATE) + "\n" + "Premier album : " + artists[i].FIRST_ALBUM + "\n" + "Lieux : " + artists[i].LOCATIONS + "\n" + "Dates de concerts : " + artists[i].CONCERT_DATE + "\n" + "Relations : " + artists[i].RELATION)
				  }
				}(i, artist)))
				container.Add(buttons[i])
				container.Add(newLabel)
			}
		})

		newMenu1 := fyne.NewMenu("Menu", menuItem1, menuItem2)

		menu := fyne.NewMainMenu(newMenu1)

		w.SetMainMenu(menu)
		scroll := widget.NewScrollContainer(fyne.NewContainerWithLayout(layout.NewVBoxLayout(), label ,fyne.NewContainerWithLayout(layout.NewGridLayout(len(buttons)+1), container)))
		w.SetContent(scroll)


	}()
	w.ShowAndRun()
}
