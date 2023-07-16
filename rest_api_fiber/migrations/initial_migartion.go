package migrations

import (
	"github.com/andreichristian/rest/internal/db"
	"github.com/andreichristian/rest/internal/models"
)

func Migrate(){


	 db.DB.AutoMigrate(&models.Author{}, &models.Book{} )

}

