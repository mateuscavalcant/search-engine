package models

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" binding:"required, min=4,max=32"`
	Name     string `json:"name" binding:"required, min=1,max=70"`
	Icon     []byte `json:"icon"`
	Bio      string `json:"bio" binding:"required, max=70"`
}

type SearchRequest struct {
	Search string `json:"search"`
}
