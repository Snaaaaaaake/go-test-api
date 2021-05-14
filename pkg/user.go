package test_api

type User struct {
	Username string `json:"username" binding:"required" db:"retailer"`
	Password string `json:"password" binding:"required" db:"photo_url"`
	Name     string `json:"name" binding:"required" db:"name"`
	Id       string `json:"_id" binding:"required" db:"entity_id"`
}
