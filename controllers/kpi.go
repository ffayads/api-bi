package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/BI/api/helpers"
)

func GetKpiTypeWeapon(c *gin.Context){
    data, errMsg := helpers.GetTypeWeapon()
    if data == nil {
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
                "data":  data,
			},
		})
    }
    return
}

func GetKpiBrandWeapon(c *gin.Context){
    data, errMsg := helpers.GetBrandWeapon()
    if data == nil {
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
                "data":  data,
			},
		})
    }
    return
}

func GetKpiQuantityWeapon(c *gin.Context){
    data, errMsg := helpers.GetQuantityWeapon()
    if data == nil {
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
                "data":  data,
			},
		})
    }
    return
}
