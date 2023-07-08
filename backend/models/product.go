package models

type Review struct {
	Rating    int    `json:"rating" min:"1" max:"5"`
	Comment   string `json:"comment"`
	ID        int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	ProductID uint   `json:"product_id"`
	UserID    uint   `json:"user_id"`
}

type Product struct {
	Name          string `json:"name"`
	ImageURL      string `json:"image_url"`
	TotalRating   int64  `json:"total_rating"`
	Total_reviews int64  `json:"total_reviews"`
	ID            int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Price         int64  `json:"price"`
}