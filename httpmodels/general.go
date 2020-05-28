package httpmodels

import (
)

type Respose struct {
  Status    bool
  Data      string
  Msg       string
  Error     error
}

type Login struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Role     string `form:"role" json:"role" binding:"required"`
}

type Month struct {
	Month       int
}

type TypeWeaponResponse struct {
	Month       int
	TypeWeapon  [][]TypeWeapon
}

type TypeWeapon struct {
	Description     string
	Count           float64
    Percentage      float64
}

type TypeWeaponQuery struct {
	Count           float64
	Description     string
}

type BrandWeaponResponse struct {
	Month       int
	BrandWeapon  [][]BrandWeapon
}

type BrandWeapon struct {
	Description     string
	Count           float64
    Percentage      float64
}

type BrandWeaponQuery struct {
	Count           float64
	Description     string
}

type QuantityWeaponResponse struct {
	Month       int
	Count       int
}
