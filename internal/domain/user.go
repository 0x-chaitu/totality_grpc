package domain

type User struct {
	Id      int64   `bson:"id"`
	Name    string  `bson:"fname"`
	City    string  `bson:"city"`
	Phone   uint64  `bson:"phone" json:"phone"`
	Height  float32 `bson:"height" json:"height"`
	Married bool    `bson:"Married" json:"Married"`
}
