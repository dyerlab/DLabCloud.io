package postgres

import (
	"github.com/dyerlab/DLabCloud.io/pkg/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)



type SnippetModel struct {
	DB *gorm.DB
}



func (s *SnippetModel) AutoMigrate() {
	s.DB.AutoMigrate( &models.Snippet{} )

	var snippets []models.Snippet
	s.DB.Find(&snippets)
	if len(snippets) == 0 {
		d1 := models.Snippet{Title: "Clear workspace",Contents: "rm(list=ls())"}
		s.DB.Save(&d1)
		d2 := models.Snippet{Title: "Load in the arapat data set", Contents: "library(gstudio); data(arapat)"}
		s.DB.Save(&d2)
	}

}



func (s *SnippetModel) Insert( title, contents string ) (uint, error) {
	d1 := models.Snippet{Title: title, Contents: contents}
	s.DB.Debug().Create( &d1 )
	return d1.ID, s.DB.Error
}

func (s *SnippetModel) Get(id int) (*models.Snippet, error ) {

	var snip models.Snippet

	res := s.DB.First(&snip, id)

	if res.RecordNotFound() {
		return nil, s.DB.Error
	}
	
	return &snip, nil 
}

func (m *SnippetModel) Latest() ([]*models.Manuscript, error) {
	return nil, nil 
}
