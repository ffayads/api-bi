package helpers

import (
    //"time"
    //"strconv"
	"github.com/BI/api/models"
	"github.com/BI/api/httpmodels"
)

func GetSearchMonthly() (*[]httpmodels.SearchResponse, string) {
    search := []httpmodels.SearchResponse{}
    search = models.GetSearchMonthly()
    return &search,""
}

func GetSearchYearly() (*[]httpmodels.SearchYearlyResponse, string) {
    search := []httpmodels.SearchYearlyResponse{}
    search = models.GetSearchYearly()
    return &search,""
}

func GetAtention() (*[]httpmodels.AtentionResponse, string) {
    search := []httpmodels.AtentionResponse{}
    search = models.GetAtention()
    return &search,""
}

func GetAffiliationByNationality() (*[]httpmodels.AffiliationResponse, string) {
    search := []httpmodels.AffiliationResponse{}
    search = models.GetAffiliation()
    return &search,""
}
