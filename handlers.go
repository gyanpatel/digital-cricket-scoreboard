///////////////////////////////////////////////////////////////////
//      (c) 2021 Fujitsu Services                                //
//       By: GyanPatel                                           //
//      Ref: https://pol-jira.atlassian.net/browse/BMP-4421      //
//     Date: 27-Jan-2021                                         //
//  Version: v01.001                                             //
///////////////////////////////////////////////////////////////////
package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
)

func handleScore(w http.ResponseWriter, r *http.Request) {
	currentScore := readScoreFile()
	t := template.Must(template.ParseFS(templates, "templates/scoreview.html", "templates/_menu.html", "templates/_sidenav.html", "templates/_footer.html"))
	err := t.Execute(w, currentScore)
	if err != nil {
		log.Println("ERROR:handleScore Error occured Parsing - templates/login ", err)
	}
}

func handleScoreBoard(w http.ResponseWriter, r *http.Request) {
	currentScore := readScoreFile()
	t := template.Must(template.ParseFS(templates, "templates/scoreboardview.html"))
	err := t.Execute(w, currentScore)
	if err != nil {
		log.Println("ERROR:handleScoreBoard Error occured Parsing - templates/login ", err)
	}
}

func handleSaveScore(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("ERROR:handleSaveScore Error occured r.ParseForm() ", err)
	}
	params := r.PostForm
	matchID := params.Get("matchid") //+ " On " + time.Now().Format("02-Jan-2006")
	fiTeam := params.Get("fiteam")
	fiScore := params.Get("fiscore")
	fiWicketsct := params.Get("fiwicketsct")
	fiOvers := params.Get("fiovers")
	siTeam := params.Get("siteam")
	siScore := params.Get("siscore")
	siWicketsct := params.Get("siwicketsct")
	siOvers := params.Get("siovers")
	bat1 := params.Get("bat1")
	bat2 := params.Get("bat2")
	bat1name := params.Get("bat1name")
	bat2name := params.Get("bat2name")
	bat1score := params.Get("bat1score")
	bat2score := params.Get("bat2score")
	bgcolor := params.Get("bgcolor")

	updatedScore := ScoreDetails{
		Match:      matchID,
		Fiteamname: fiTeam,
		FiScore:    fiScore,
		FiWickets:  fiWicketsct,
		FiOvers:    fiOvers,
		Siteamname: siTeam,
		SiScore:    siScore,
		SiWickets:  siWicketsct,
		SiOvers:    siOvers,
		Bat1:       bat1,
		Bat2:       bat2,
		Bat1Name:   bat1name,
		Bat2Name:   bat2name,
		Bat1score:  bat1score,
		Bat2score:  bat2score,
		BgColor:    bgcolor,
	}

	file, err := json.MarshalIndent(updatedScore, "", " ")
	if err != nil {
		log.Fatal("ERROR : UNABLE TO PREPARE THE SCORE JSON", err)
	}
	err = ioutil.WriteFile("templates/ScoreBoard.json", file, 0644)
	if err != nil {
		log.Fatal("ERROR : UNABLE TO SAVE THE SCORE FILE", err)
	}
	http.Redirect(w, r, "/scoreview", http.StatusSeeOther)

}

func readScoreFile() ScoreDetails {
	var scoreDetails ScoreDetails
	ScoreFile, err := ioutil.ReadFile("templates/ScoreBoard.json")
	if err != nil {
		log.Fatal("FATAL : UNABLE TO READ THE SCORE FILE", err)
	}
	err = json.Unmarshal(ScoreFile, &scoreDetails)
	if err != nil {
		log.Fatal("FATAL : UNABLE TO PARSE THE SCORE FILE", err)
	}
	return scoreDetails
}
