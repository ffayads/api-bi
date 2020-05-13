package models

import (
    "fmt"
    "time"
    "github.com/jinzhu/gorm"
	"github.com/BI/api/httpmodels"
	"github.com/BI/api/db"
)

type Registros struct {
    gorm.Model
	TipoIdentificacion  string
	Identificacion      string
	Llegada             time.Time
	Atendido            time.Time
	Salida              time.Time
}

func GetSearchMonthly() []httpmodels.SearchResponse {
    var total float64
    var xx float64
	search := []httpmodels.Search{}
	response := []httpmodels.SearchResponse{}
    db.DB.Table("registros").Select("YEAR(llegada) as year, MONTH(llegada) as month, DAY(llegada) day, count(id) as count").Group("YEAR(llegada), MONTH(llegada), DAY(llegada)").Scan(&search)
    for _, elem := range search {
        total += elem.Count
    }
    for _, elem := range search {
	    data := httpmodels.SearchResponse{
            Year:       elem.Year,
            Month:      elem.Month,
            Day:        elem.Day,
            Count:      elem.Count,
            Percentage: (elem.Count/total)*100,
        }
        xx += (elem.Count/total)*100
        response = append(response, data)
    }
    fmt.Println(xx)
	return response
}

func GetSearchYearly() []httpmodels.SearchYearlyResponse {
    var total float64
    var xx float64
	search := []httpmodels.SearchYearly{}
	response := []httpmodels.SearchYearlyResponse{}
    db.DB.Table("registros").Select("YEAR(llegada) as year, MONTH(llegada) as month, count(id) as count").Group("YEAR(llegada), MONTH(llegada)").Scan(&search)
    for _, elem := range search {
        total += elem.Count
    }
    for _, elem := range search {
	    data := httpmodels.SearchYearlyResponse{
            Year:       elem.Year,
            Month:      elem.Month,
            Count:      elem.Count,
            Percentage: (elem.Count/total)*100,
        }
        xx += (elem.Count/total)*100
        response = append(response, data)
    }
    fmt.Println(xx)
	return response
}

func GetAtention() []httpmodels.AtentionResponse {
	search := []httpmodels.DateTime{}
	array := []httpmodels.Minutes{}
	response := []httpmodels.AtentionResponse{}
    db.DB.Table("registros").Select("DATE(llegada) as date").Group("DATE(llegada)").Scan(&search)
    for _, elem := range search {
        var total int
        var count int
        db.DB.Table("registros").Select("((HOUR(TIMEDIFF(atendido,llegada))*60)+(MINUTE(TIMEDIFF(atendido,llegada)))) as minutes").Where("DATE(llegada) = ?",elem.Date.Format("2006-01-02")).Scan(&array)
        for _, elemM := range array {
            count++
            total += elemM.Minutes
        }
        fmt.Println(total)
        fmt.Println(count)
	    data := httpmodels.AtentionResponse{
            Date:       elem.Date,
            Average:   (total/count),
        }
        response = append(response, data)
    }
	return response
}
