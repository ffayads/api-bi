package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/BI/api/helpers"
)

func GetKpiSearchs(c *gin.Context){
    search, errMsg := helpers.GetSearchMonthly()
    if search == nil {
        c.JSON(200, gin.H{
			"status": false,
			"err":    "",
			"msg":    errMsg,
			"data": gin.H{
			},
		})
    } else {
        c.JSON(200, gin.H{
			"status": true,
			"err":    "",
			"msg":    errMsg,
			"data": gin.H{
                "Search":  search,
			},
		})
    }
    return
}

func GetKpiSearchsYearly(c *gin.Context){
    search, errMsg := helpers.GetSearchYearly()
    if search == nil {
        c.JSON(200, gin.H{
			"status": false,
			"err":    "",
			"msg":    errMsg,
			"data": gin.H{
			},
		})
    } else {
        c.JSON(200, gin.H{
			"status": true,
			"err":    "",
			"msg":    errMsg,
			"data": gin.H{
                "Search":  search,
			},
		})
    }
    return
}

func GetAffiliationByNationality(c *gin.Context){
    affiliation, errMsg := helpers.GetAffiliationByNationality()
    if affiliation == nil {
        c.JSON(200, gin.H{
			"status": false,
			"err":    "",
			"msg":    errMsg,
			"data": gin.H{
			},
		})
    } else {
        c.JSON(200, gin.H{
			"status": true,
			"err":    "",
			"msg":    errMsg,
			"data": gin.H{
                "affiliation":  affiliation,
			},
		})
    }
    return
}

func GetAtention(c *gin.Context){
    atention, errMsg := helpers.GetAtention()
    if atention == nil {
        c.JSON(200, gin.H{
			"status": false,
			"err":    "",
			"msg":    errMsg,
			"data": gin.H{
			},
		})
    } else {
        c.JSON(200, gin.H{
			"status": true,
			"err":    "",
			"msg":    errMsg,
			"data": gin.H{
                "atention":  atention,
			},
		})
    }
    return
}
