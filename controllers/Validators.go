package controllers

type CreateAuthorValidator struct {
	Name string `form:"name" json:"name" binding:"required"`
	Introduction string `form:"intro" json:"intro"`
}

type UpdateAuthorValidator struct {
	 Name string `form:"name" json:"name"`
	 Introduction string `form:"intro" json:"intro"`
}

type CreateBookValidator struct {
	Title string `form:"name" json:"name" binding:"required"`
	Description string `form:"description" json:"description"`
	Author uint `form:"author" json"author" binding:"required"`
}

type UpdateBookValidator struct {
	Title string `form:"name" json:"name"`
	Description string `form:"description" json:"description"`
}
