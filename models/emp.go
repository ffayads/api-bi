package models

import (
    "github.com/jinzhu/gorm"
	"github.com/BI/api/httpmodels"
	"github.com/BI/api/db"
)

type Emp struct {
    gorm.Model
	Reports             Reports             `gorm:"foreignkey:ReportsID"`
	ReportsID           uint
	TypeWeapons         TypeWeapons         `gorm:"foreignkey:TypeWeaponsID"`
	TypeWeaponsID       uint
	BrandWeapons        BrandWeapons        `gorm:"foreignkey:BrandWeaponsID"`
	BrandWeaponsID      uint
	InLetX              int                 `gorm:"not null"`
	InLetY              int                 `gorm:"not null"`
	InLetZ              int                 `gorm:"not null"`
	InLetSize           int                 `gorm:"not null"`
	OutLetX             int
	OutLetY             int
	OutLetZ             int
	OutLetSize          int                 `gorm:"not null"`
	Description         string              `gorm:"type:text"`
	ModelWeapon         string
	Serial              string
	Industrial          bool
	Caliber             string
	InternalNumber      int                 `gorm:"type:bigint"`
	Performance         string
	BarrelWeapons       BarrelWeapons       `gorm:"foreignkey:BarrelWeaponsID"`
	BarrelWeaponsID     uint
	SizeBarrel          string
	Size                string
	Capacity            int
	Empcol              string
	TypePerformances    TypePerformances    `gorm:"foreignkey:TypePerformancesID"`
	TypePerformancesID  uint
	Details             string
	WeaponStatus        string              `gorm:"not null"`
	Status              bool                `gorm:"not null"`
}

func GetTypeWeapon() []httpmodels.TypeWeaponResponse {
	query := []httpmodels.Month{}
	response := []httpmodels.TypeWeaponResponse{}
    db.DB.Table("emp").Select("MONTH(emp.created_at) as month").Group("MONTH(emp.created_at)").Scan(&query)
    for _, elem := range query {
	    typeWeapons := []httpmodels.TypeWeapon{}
        var count float64
	    subquery := []httpmodels.TypeWeaponQuery{}
        db.DB.Table("emp, type_weapons").Select("count(type_weapons_id) as count, type_weapons.description").Where("month(emp.created_at) = ? AND emp.type_weapons_id = type_weapons.id",elem.Month).Group("emp.type_weapons_id").Scan(&subquery)
        for _, elemm := range subquery {
            count += elemm.Count
        }
        for _, elemm := range subquery {
	        typeWeapon := httpmodels.TypeWeapon{
                Description: elemm.Description,
                Count: elemm.Count,
                Percentage: (elemm.Count*100)/count,
            }
            typeWeapons = append(typeWeapons, typeWeapon)
        }
	    month := httpmodels.TypeWeaponResponse{
            Month: elem.Month,
        }
        month.TypeWeapon = append(month.TypeWeapon, typeWeapons)
        response = append(response, month)
    }
	return response
}

func GetBrandWeapon() []httpmodels.BrandWeaponResponse {
	query := []httpmodels.Month{}
	response := []httpmodels.BrandWeaponResponse{}
    db.DB.Table("emp").Select("MONTH(emp.created_at) as month").Group("MONTH(emp.created_at)").Scan(&query)
    for _, elem := range query {
	    brandWeapons := []httpmodels.BrandWeapon{}
        var count float64
	    subquery := []httpmodels.BrandWeaponQuery{}
        db.DB.Table("emp, brand_weapons").Select("count(brand_weapons_id) as count, brand_weapons.description").Where("month(emp.created_at) = ? AND emp.brand_weapons_id = brand_weapons.id",elem.Month).Group("emp.brand_weapons_id").Scan(&subquery)
        for _, elemm := range subquery {
            count += elemm.Count
        }
        for _, elemm := range subquery {
	        brandWeapon := httpmodels.BrandWeapon{
                Description: elemm.Description,
                Count: elemm.Count,
                Percentage: (elemm.Count*100)/count,
            }
            brandWeapons = append(brandWeapons, brandWeapon)
        }
	    month := httpmodels.BrandWeaponResponse{
            Month: elem.Month,
        }
        month.BrandWeapon = append(month.BrandWeapon, brandWeapons)
        response = append(response, month)
    }
	return response
}

func GetQuantityWeapon() []httpmodels.QuantityWeaponResponse {
	response := []httpmodels.QuantityWeaponResponse{}
    db.DB.Table("emp").Select("MONTH(created_at) as month, count(created_at) as count").Group("MONTH(created_at)").Scan(&response)
	return response
}

func (Emp) TableName() string {
    return "emp"
}
