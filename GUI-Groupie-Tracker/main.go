package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	// "time"
	"strconv"
	"strings"

	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	// "fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/dialog"
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
	// Locationslabel := widget.NewLabel("")
	// Dateslabel := widget.NewLabel("")
	// Relationslabel := widget.NewLabel("")

	var artists []Artists
	var artistsData []string

	var locations Locations
	var dates Dates
	var relations Relations

	go func() {

		// Get data from API
		ArtistsResp, err := http.Get("https://groupietrackers.herokuapp.com/api/artists")
		if err != nil {
			// Artistslabel.SetText("Erreur lors du chargement des données")
			errs := errors.New("erreur lors du chargement des données")
			dialog.ShowError(errs, w)
			return
		}
		defer ArtistsResp.Body.Close()

		ArtistsBody, err := io.ReadAll(ArtistsResp.Body)
		if err != nil {
			// Artistslabel.SetText("Erreur lors de la lecture des données")
			errs := errors.New("erreur lors de la lecture des données")
			dialog.ShowError(errs, w)
			return
		}

		// Decode data
		err = json.Unmarshal(ArtistsBody, &artists)
		if err != nil {
			// Artistslabel.SetText("Erreur lors de la décodage des données")
			errs := errors.New("erreur lors de la décodage des données")
			dialog.ShowError(errs, w)
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
			// Locationslabel.SetText("Erreur lors du chargement des données de lieux")
			errs := errors.New("erreur lors du chargement des données de lieux")
			dialog.ShowError(errs, w)
			return
		}
		defer locationResp.Body.Close()

		locationBody, err := io.ReadAll(locationResp.Body)
		if err != nil {
			// Locationslabel.SetText("Erreur lors de la lecture des données de lieux")
			errs := errors.New("erreur lors de la lecture des données de lieux")
			dialog.ShowError(errs, w)
			return
		}
		
		err = json.Unmarshal(locationBody, &locations)
		if err != nil {
			// Locationslabel.SetText("Erreur lors de la décodage des données de lieux")
			errs := errors.New("erreur lors de la décodage des données de lieux")
			dialog.ShowError(errs, w)
			return
		}

	}()

	go func() {

		DatesResp, err := http.Get("https://groupietrackers.herokuapp.com/api/dates")
		if err != nil {
			// Dateslabel.SetText("Erreur lors du chargement des données des dates")
			errs := errors.New("erreur lors du chargement des données des dates")
			dialog.ShowError(errs, w)
			return
		}
		defer DatesResp.Body.Close()

		DatesBody, err := io.ReadAll(DatesResp.Body)
		if err != nil {
			// Dateslabel.SetText("Erreur lors de la lecture des données des dates")
			errs := errors.New("erreur lors de la lecture des données des dates")
			dialog.ShowError(errs, w)
			return
		}
		
		err = json.Unmarshal(DatesBody, &dates)
		if err != nil {
			// Dateslabel.SetText("Erreur lors de la décodage des données des dates")
			errs := errors.New("erreur lors de la décodage des données des dates")
			dialog.ShowError(errs, w)
			return
		}
		
	}()

	go func() {
		RelationsResp, err := http.Get("https://groupietrackers.herokuapp.com/api/relation")
		if err != nil {
			// Relationslabel.SetText("Erreur lors du chargement des données des relations")
			errs := errors.New("erreur lors du chargement des données des relations")
			dialog.ShowError(errs, w)
			return
		}
		defer RelationsResp.Body.Close()

		RelationsBody, err := io.ReadAll(RelationsResp.Body)
		if err != nil {
			// Relationslabel.SetText("Erreur lors de la lecture des données des relations")
			errs := errors.New("erreur lors de la lecture des données des relations")
			dialog.ShowError(errs, w)
			return
		}
		
		err = json.Unmarshal(RelationsBody, &relations)
		if err != nil {
			// Relationslabel.SetText("Erreur lors de la décodage des données des relations")
			errs := errors.New("erreur lors de la décodage des données des relations")
			dialog.ShowError(errs, w)
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

	content := container.NewVBox(
		Artistslabel,
		Artiststext,
		returnline,
		// Locationslabel,
		// Dateslabel,
		// Relationslabel,
	)

	scroll := container.NewVScroll(container.New(layout.NewVBoxLayout(), content))

	menuItem2 := fyne.NewMenuItem("Artists", func() {

		for i, artist := range artistsData {

			var artistLocations []string
			var artistDates []string
			// var artistRelations []string
			var artistLocationsVille []string
			var artistLocationsPays []string
			
			resource, _ := fyne.LoadResourceFromURLString("")
			Artistimage := canvas.NewImageFromResource(resource)
			Artistimage.FillMode = canvas.ImageFillOriginal
			newLabel := widget.NewLabel("")
			
			buttons = append(buttons, widget.NewButton(artist, func(i int, artist string) func() {
				return func() {

					artistLocations = nil
					artistDates = nil
					artistLocationsVille = nil
					artistLocationsPays = nil

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
					resource, _ = fyne.LoadResourceFromURLString(artists[i].IMAGE)
					Artistimage.Resource = resource
					Artistimage.Refresh()
				}
			}(i, artist)))
			content.Add(buttons[i])
			content.Add(Artistimage)
			content.Add(newLabel)
		}
		w.SetContent(scroll)
	})

	annoncelabel := widget.NewLabel("Vous pouvez rechercher un artiste en écrivant son nom dans la barre de recherche ci-dessous : \n*soit le nom complet de l'artiste \n*soit le début du nom \n*soit le nom complet du membre \n*soit le début du nom")
	searchArtist := widget.NewEntry()

	filterlabel := widget.NewLabel("")
	otherartistlabel := widget.NewLabel("")
	newfilter := widget.NewLabel("")

	var filteredArtists []string

	f := 1000.0
	data := binding.BindFloat(&f)
	slider := widget.NewSliderWithData(1950, 2020, data)
	createlabel := widget.NewLabelWithData(
		binding.FloatToStringWithFormat(data, "%.0f"),
	)

	slider.Hidden = true
	createlabel.Hidden = true

	checkcheckbox := false
	checkcheckcheckbox := false
	checkcheckcheckcheckbox := false

	var membersCheckboxes [7]*widget.Check

	for i := 0; i < 7; i++ {
		membersCheckboxes[i] = widget.NewCheck("", func(checked bool) {
			// Mettre en œuvre la logique pour filtrer les artistes en fonction du nombre de membres sélectionné
		})
		membersCheckboxes[i].Text = strconv.Itoa(i + 1) + " membres"
	}

	members := container.NewHBox(
		membersCheckboxes[0],
		membersCheckboxes[1],
		membersCheckboxes[2],
		membersCheckboxes[3],
		membersCheckboxes[4],
		membersCheckboxes[5],
		membersCheckboxes[6],
	)

	members.Hidden = true

	checkbox3 := widget.NewCheck("Filtrer par nombre de membres", func(plot bool) {
		if plot {
			checkcheckcheckcheckbox = true
			dialog.ShowInformation("Filtrer par nombre de membres", "Veuillez sélectionner le nombre de membres que vous souhaitez filtrer", w)
			members.Show()
		} else {
			checkcheckcheckcheckbox = false
			members.Hide()
		}
	})

	var checkedMembers []int

	filterbutton := widget.NewButton("Filtrer", func() {

		newfilter.SetText("")
		otherartistlabel.SetText("")

		if checkcheckbox{
			createdate, _ := strconv.Atoi(createlabel.Text)
			filteredArtists = filterArtistsDate(artists, createdate)
		}else if checkcheckcheckbox {
			filteredArtists = filterArtistsMembers(artists, searchArtist.Text)
		}else if checkcheckcheckcheckbox {
			checkedMembers = nil
			for i, checkbox := range membersCheckboxes {
				if checkbox.Checked {
					checkedMembers = append(checkedMembers, i + 1)
				}
			}
			filteredArtists = filterArtistsLenMembers(artists, checkedMembers)
		}else {
			filteredArtists = filterArtists(artistsData, searchArtist.Text)
		}
		filterlabel.SetText(strings.Join(filteredArtists, ", "))

		if len(filterlabel.Text) <= 654 {

			if len(filteredArtists) <= 0 {
				otherartistlabel.SetText("Ecrivez le nom complet de l'artiste ou selectionner une autre date pour voir ses informations si ce n'est pas le bon artiste.")
			} else {
				otherartistlabel.SetText("Possible résultat de votre recherche : " + strings.Join(filteredArtists, ", ") )
			}

			for i := range artistsData {

				var artistLocations []string
				var artistDates []string
				var artistLocationsVille []string
				var artistLocationsPays []string

				if strings.EqualFold(filterlabel.Text,artists[i].NAME) {
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

					newfilter.SetText("Nom du groupe : " + filterlabel.Text + "\n" +"Membres : \n - " + strings.Join(artists[i].MEMBERS, "\n - ") + "\n" + "Date de création : " + fmt.Sprintf("%d", artists[i].CREA_DATE) + "\n" + "Premier album : " + artists[i].FIRST_ALBUM + "\n" + "Lieux : \n - Ville : " + strings.Join(artistLocationsVille, ", ") + "\n - Pays : " + strings.Join(artistLocationsPays, ", ") + "\n" + "Dates de concerts : " + strings.Join(artistDates, ", ") + "\n" + "Relations : " + artists[i].RELATION)
				}
			}
		} else {
			newfilter.SetText("Veuillez écrire quelque chose dans la barre de recherche")
		}
	})

	annoncelabel2 := widget.NewLabel("Vous pouvez filtrer les artistes par date de création en utilisant le curseur ci-dessous :")

	checkbox := widget.NewCheck("Filtrer par date de création", func(plot bool) {
		if plot {
			checkcheckbox = true
			dialog.ShowInformation("Filtrer par date de création", "Veuillez sélectionner la date de création que vous souhaitez filtrer", w)
			createlabel.Show()
			slider.Show()
		} else {
			checkcheckbox = false
			createlabel.Hide()
			slider.Hide()
		}
	})

	checkbox2 := widget.NewCheck("Filtrer par nom de membre", func(plot bool) {
		if plot {
			checkcheckcheckbox = true
			dialog.ShowInformation("Filtrer par nom de membre", "Veuillez écrire le nom du membre que vous souhaitez filtrer", w)
		} else {
			checkcheckcheckbox = false
		}
	})

	search := container.NewVBox(
		annoncelabel,
		searchArtist,
		checkbox2,
		annoncelabel2,
		checkbox,
		createlabel,
		slider,
		checkbox3,
		members,
		filterbutton,
		otherartistlabel,
		newfilter,
	)

	menuItem3 := fyne.NewMenuItem("Recherche", func() {
		w.SetContent(search)
	})

	newMenu1 := fyne.NewMenu("Menu", menuItem1, menuItem2,menuItem3)

	menu := fyne.NewMainMenu(newMenu1)

	w.SetMainMenu(menu)
	w.ShowAndRun()
}

func filterArtists(artists []string, searchTerm string) []string {
	filteredArtists := []string{}
	for _, artist := range artists {
	  if strings.Contains(strings.ToLower(artist), strings.ToLower(searchTerm)) {
		filteredArtists = append(filteredArtists, artist)
	  }
	}
	return filteredArtists
}

func filterArtistsDate(artists []Artists, searchTerm int) []string {
	filteredArtists := []string{}
	for i := range artists {
		if artists[i].CREA_DATE == searchTerm {
			filteredArtists = append(filteredArtists, artists[i].NAME)
		}
	}
	return filteredArtists
}

func filterArtistsMembers(artists []Artists, searchTerm string) []string {
	filteredArtists := []string{}
	for i := range artists {
		for j := range artists[i].MEMBERS {
			if strings.Contains(strings.ToLower(artists[i].MEMBERS[j]), strings.ToLower(searchTerm)) {
				filteredArtists = append(filteredArtists, artists[i].NAME)
				break
			}
		}
	}
	return filteredArtists
}

func filterArtistsLenMembers(artists []Artists, searchTerm []int) []string {
	filteredArtists := []string{}
	for i := range artists {
		for _, nb := range searchTerm {
			if len(artists[i].MEMBERS) == nb {
				filteredArtists = append(filteredArtists, artists[i].NAME)
				break
			}
		}
	}
	return filteredArtists
}