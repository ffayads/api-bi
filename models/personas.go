package models

import (
    "github.com/jinzhu/gorm"
    "time"
	"github.com/BI/api/db"
	"github.com/BI/api/httpmodels"
)

type Personas struct {
    gorm.Model
	Identificacion      string
	TipoIdentificacion  string
	Apellidos           string
	Nombres             string
	Genero              string
	FechaNacimiento     time.Time   `gorm:"type:date;"`
	Nacionalidad        string
}

func GetAffiliation() []httpmodels.AffiliationResponse {
	response := []httpmodels.AffiliationResponse{}
    db.DB.Table("personas").Select("nacionalidad as nationality, count(nacionalidad) as count").Group("nacionalidad").Scan(&response)
	return response
}
