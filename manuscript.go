package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// Manuscript is the base type for a journal article
type Manuscript struct {
	gorm.Model
	Authors  string `json:"authors"`
	Year     int    `json:"year"`
	Title    string `json:"title"`
	Journal  string `json:"journal"`
	Volume   int    `json:"volume"`
	Pages    string `json:"pages"`
	Doi      string `json:"doi"`
	Abstract string `json:"abstract"`
}

// MigrateManuscript opens the db and makes a fresh copy
func MigrateManuscript() {

	db, err := gorm.Open(dbName, dbConnect)
	if err != nil {
		fmt.Println(err.Error())
		panic("Failed to connect to the db")
	}
	defer db.Close()

	db.AutoMigrate(&Manuscript{})

	var manuscripts []Manuscript
	db.Find(&manuscripts)
	if len(manuscripts) == 0 {

		db.Create(&Manuscript{
			Authors:  "Dyer RJ",
			Year:     2007,
			Title:    "The Evolution of Genetic Topologies",
			Journal:  "Theoretical Population Biology",
			Volume:   71,
			Pages:    "71-79",
			Doi:      "http://dx.doi.org/10.1016/j.tpb.2006.07.001",
			Abstract: "This manuscript explores the simultaneous evolution of population genetic parameters and topological features within a population graph through a series of Monte Carlo simulations. I show that node centrality and graph breadth are significantly correlated to population genetic parameters F<sub>ST</sub> and M (r 1⁄4 0:95; r 1⁄4 0:98, respectively), which are commonly used in quantifying among population genetic structure and isolation by distance. Next, the topological consequences of migration patterns are examined by contrasting N-island and stepping stone models of gene movement. Finally, I show how variation in migration rate influences the rate of formation of specific topological features with particular emphasis to the phase transition that occurs when populations begin to become fixed due to restricted movement of genes among populations. I close by discussing the utility of this method for the analysis of intraspecific genetic variation."})

		db.Create(&Manuscript{
			Authors:  "Dyer RJ",
			Year:     2015,
			Title:    "Is there such a thing as landscape genetics? ",
			Journal:  "Molecular Ecology",
			Volume:   24,
			Pages:    "3518-3528",
			Doi:      "http://dx.doi.org/10.1111/mec.13249",
			Abstract: "For a scientific discipline to be interdisciplinary, it must satisfy two conditions; it must consist of contributions from at least two existing disciplines, and it must be able to provide insights, through this interaction, that neither progenitor discipline could address. In this study, I examine the complete body of peer-reviewed literature self-identified as landscape genetics (LG) using the statistical approaches of text mining and natural language processing. The goal here was to quantify the kinds of questions being addressed in LG studies, the ways in which questions are evaluated mechanistically, and how they are differentiated from the progenitor disciplines of landscape ecology and population genetics. I then circumscribe the main factions within published LG studies examining the extent to which emergent questions are being addressed and highlighting a deep bifurcation between existing individual- and population-based approaches. I close by providing some suggestions on where theoretical and analytical work is needed if LGs is to serve as a real bridge connecting evolution and ecology <em>sensu lato</em>."})
	}

}

// ManuscriptData holds data for the template
type ManuscriptData struct {
	Title       string
	Manuscripts []Manuscript
}

// AllManuscripts lists all manuscripts in the database.
func AllManuscripts(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open(dbName, dbConnect)
	if err != nil {
		panic("Could not open database")
	}
	defer db.Close()

	var manuscripts []Manuscript
	db.Find(&manuscripts)

	data := ManuscriptData{
		Title:       "Manuscripts",
		Manuscripts: manuscripts,
	}

	fmt.Printf("Serving up %d manuscripts", len(manuscripts))

	tmpl := template.Must(template.ParseFiles("manuscripts.html"))
	tmpl.Execute(w, data)
	//json.NewEncoder(w).Encode(manuscripts)
}
