package dto

type Credentials struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CredentialsResponse struct {
	Prefix string `json:"prefix"`
	Token  string `json:"token"`
}

type AlbumRequestForm struct {
	Name   string `json:"name" binding:"required"`
	Photos []uint `json:"photoId,omitempty"`
}

type AuthorRequestform struct {
	Name   string `json:"name" binding:"required"`
	Photos []uint `json:"photoId,omitempty"`
}

type PhotoMetadataRequestform struct {
	Height      int64  `json:"height" binding:"required"`
	Width       int64  `json:"width" binding:"required"`
	Orientation string `json:"orientation" binding:"required"`
	Compressed  int64  `json:"compressed" binding:"required"`
	Comment     string `json:"comment" binding:"required"`
	PhotoId     uint   `json:"photoId" binding:"required"`
}

type PhotoRequestform struct {
	Description string `json:"description" binding:"required"`
	FileName    string `json:"fileName" binding:"required"`
	IsPublished bool   `json:"isPublished" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Views       int64  `json:"views" binding:"required"`
	AuthorId    uint   `json:"authorId" binding:"required"`
}

type UserRequestform struct {
	Name      string  `json:"name" binding:"required"`
	Email     *string `json:"email" binding:"required"`
	Age       uint8   `json:"age" binding:"required"`
	FirstName string  `json:"firstName" binding:"required"`
	LastName  string  `json:"lastName" binding:"required"`
	Username  string  `json:"username" binding:"required"`
	Password  string  `json:"password" binding:"required"`
}
