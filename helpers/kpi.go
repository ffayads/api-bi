package helpers

import (
    //"time"
    //"strconv"
	"github.com/BI/api/models"
	"github.com/BI/api/httpmodels"
)

func GetTypeWeapon() (*[]httpmodels.TypeWeaponResponse, string) {
    data := []httpmodels.TypeWeaponResponse{}
    data = models.GetTypeWeapon()
    return &data,""
}

func GetBrandWeapon() (*[]httpmodels.BrandWeaponResponse, string) {
    data := []httpmodels.BrandWeaponResponse{}
    data = models.GetBrandWeapon()
    return &data,""
}

func GetQuantityWeapon() (*[]httpmodels.QuantityWeaponResponse, string) {
    data := []httpmodels.QuantityWeaponResponse{}
    data = models.GetQuantityWeapon()
    return &data,""
}
