package models

type User struct {
	ID               uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name             string `json:"name"`
	ProfileImage_URL string `json:"profile_image_url"`
	Password         string `json:"password"`
	Email            string `json:"email"`
	Cart             []byte `json:"cart"`
	Orders           []byte `json:"orders"`
}

type SignUps struct {
	ID               uint   `json:"id" gorm:"primaryKey;autoIncrement"`
	Name             string `json:"name"`
	ProfileImage_URL string `json:"profile_image_url"`
	Password         string `json:"password"`
	Email            string `json:"email"`
	VerifyToken      string `json:"verify_token"`
}
