package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	// "time"
	"strings"
	// "strconv"

	"image/color"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	// "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

/*--------------------------------------------------------------------------------------------
-------------------------------------- Type Struct -------------------------------------------
----------------------------------------------------------------------------------------------*/

// Artists struct
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

// Locations struct
type Locations struct {
	Index []Location `json:"index"`
}

type Location struct {
	ID         int      `json:"id"`
	LOCATIONS  []string `json:"locations"`
	DATES      string   `json:"dates"`
}

// Dates struct
type Dates struct {
	Index []Date `json:"index"`
}

type Date struct {
	ID    int      `json:"id"`
	DATES []string `json:"dates"`
}

// Realtions struct
type Relations struct {
	Index []Relation `json:"index"`
}

type Relation struct {
	ID         int                 `json:"id"`
	DATESLOCAT map[string][]string `json:"datesLocations"`
}

func main() {
	a := app.New()
	a.SetIcon(theme.FyneLogo())
	w := a.NewWindow("Groupie Tracker")
	w.Resize(fyne.NewSize(800, 600))

	Homelabel := widget.NewLabel("Page D'accueil")
	Artistslabel := widget.NewLabel("")
	Locationslabel := widget.NewLabel("")
	Dateslabel := widget.NewLabel("")
	Relationslabel := widget.NewLabel("")

	var artists []Artists
	var artistsData []string

	var locations Locations
	var dates Dates
	var relations Relations

	go func() {

		// Get data from API
		ArtistsResp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
		if err != nil {
			Artistslabel.SetText("Erreur lors du chargement des données")
			return
		}
		defer ArtistsResp.Body.Close()

		ArtistsBody, err := io.ReadAll(ArtistsResp.Body)
		if err != nil {
			Artistslabel.SetText("Erreur lors de la lecture des données")
			return
		}

		// Decode data
		err = json.Unmarshal(ArtistsBody, &artists)
		if err != nil {
			Artistslabel.SetText("Erreur lors de la décodage des données")
			return
		}

		// Recover the data in a []string
		for i := 0; i < len(artists); i++ {
			artistsData = append(artistsData, fmt.Sprintf(artists[i].NAME))
		}

	}()

	go func() {

		locationResp, err := http.Get("https://groupietrackers.herokuapp.com/api/locations")
		if err != nil {
			Locationslabel.SetText("Erreur lors du chargement des données de lieux")
			return
		}
		defer locationResp.Body.Close()

		locationBody, err := io.ReadAll(locationResp.Body)
		if err != nil {
			Locationslabel.SetText("Erreur lors de la lecture des données de lieux")
			return
		}
		
		err = json.Unmarshal(locationBody, &locations)
		if err != nil {
			Locationslabel.SetText("Erreur lors de la décodage des données de lieux")
			return
		}

	}()

	go func() {

		DatesResp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
		if err != nil {
			Dateslabel.SetText("Erreur lors du chargement des données des dates")
			return
		}
		defer DatesResp.Body.Close()

		DatesBody, err := io.ReadAll(DatesResp.Body)
		if err != nil {
			Dateslabel.SetText("Erreur lors de la lecture des données des dates")
			return
		}
		
		err = json.Unmarshal(DatesBody, &dates)
		if err != nil {
			Dateslabel.SetText("Erreur lors de la décodage des données des dates")
			return
		}
		
	}()

	go func() {
		RelationsResp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
		if err != nil {
			Relationslabel.SetText("Erreur lors du chargement des données des relations")
			return
		}
		defer RelationsResp.Body.Close()

		RelationsBody, err := io.ReadAll(RelationsResp.Body)
		if err != nil {
			Relationslabel.SetText("Erreur lors de la lecture des données des relations")
			return
		}
		
		err = json.Unmarshal(RelationsBody, &relations)
		if err != nil {
			Relationslabel.SetText("Erreur lors de la décodage des données des relations")
			return
		}
		
	}()

	returnline := widget.NewLabel("\n")

	Homelabel.SetText("Page D'accueil")
	Hometext := canvas.NewText("Bienvenue sur GROUPIE TRACKER", color.White)
	Hometext.Alignment = fyne.TextAlignCenter

	Hometext1 := canvas.NewText("Le projet Groupie Tracker a pour finalité de traiter des données à l'aide d'une API. Cette API doit être développé avec le langage GO, et la gestion des données doit etre traitée avec le langage GO.", color.White)
	Hometext1.Alignment = fyne.TextAlignCenter

	box := container.NewVBox(
		Homelabel,
		Hometext,
		returnline,
		Hometext1,
	)

	w.SetContent(box)

	// Create a menu
	menuItem1 := fyne.NewMenuItem("Home", func() {
		w.SetContent(box)
	})

	var buttons []fyne.CanvasObject

	Artistslabel.SetText("Page des artistes")
	Artiststext := canvas.NewText("Voici la liste de tout les artistes :", color.White)
	Artiststext.Alignment = fyne.TextAlignCenter

	resource, _ := fyne.LoadResourceFromURLString("")
	Artistimage := canvas.NewImageFromResource(resource)
	Artistimage.FillMode = canvas.ImageFillOriginal

	content := container.NewVBox(
		Artistslabel,
		Artiststext,
		returnline,
		Artistimage,
		Locationslabel,
		Dateslabel,
		Relationslabel,
	)

	scroll := container.NewVScroll(container.New(layout.NewVBoxLayout(), content))

	menuItem2 := fyne.NewMenuItem("Artists", func() {
		for i, artist := range artistsData {
			var artistLocations []string
			var artistDates []string
			// var artistRelations []string
			var artistLocationsVille []string
			var artistLocationsPays []string
			newLabel := widget.NewLabel("")
			buttons = append(buttons, widget.NewButton(artist, func(i int, artist string) func() {
				return func() {

					for _, loc := range locations.Index {
						if loc.ID == artists[i].ID {
							artistLocations = append(artistLocations, loc.LOCATIONS...)
							break
						}
					}

					for _, location := range artistLocations {
						splitLocation := strings.Split(location, "-")
						artistLocationsVille = append(artistLocationsVille, splitLocation[0])
						artistLocationsPays = append(artistLocationsPays, splitLocation[1])
					}

					for _, dates := range dates.Index {
						if dates.ID == artists[i].ID {
							artistDates = append(artistDates, dates.DATES...)
							break
						}
					}

					for i, date := range artistDates {
						artistDates[i] = strings.Replace(date, "*", "", -1)
					}

					// for _, relations := range relations.Index {
					// 	if relations.ID == artists[i].ID {
					// 		artistRelations = append(artistRelations, relations.DATESLOCAT[artistLocations[i]]...)
					// 		break
					// 	}
					// }

					newLabel.SetText("Membres : \n - " + strings.Join(artists[i].MEMBERS, "\n - ") + "\n" + "Date de création : " + fmt.Sprintf("%d", artists[i].CREA_DATE) + "\n" + "Premier album : " + artists[i].FIRST_ALBUM + "\n" + "Lieux : \n - Ville : " + strings.Join(artistLocationsVille, ", ") + "\n - Pays : " + strings.Join(artistLocationsPays, ", ") + "\n" + "Dates de concerts : " + strings.Join(artistDates, ", ") + "\n" + "Relations : " + artists[i].RELATION)
					resource, _ := fyne.LoadResourceFromURLString(artists[i].IMAGE)
					Artistimage.Resource = resource
					Artistimage.Refresh()
				}
			}(i, artist)))
			content.Add(buttons[i])
			content.Add(newLabel)
		}
		w.SetContent(scroll)
	})

	newMenu1 := fyne.NewMenu("Menu", menuItem1, menuItem2)

	menu := fyne.NewMainMenu(newMenu1)

	w.SetMainMenu(menu)
	w.ShowAndRun()
}
