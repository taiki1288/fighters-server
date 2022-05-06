package entity

import "time"

type User struct {
	ID               string    `json:"id"`               //UserID
	Name             string    `json:"name"`             //名前
	SelfIntroduction string    `json:"selfintroduction"` //自己紹介
	Age              int       `json:"age"`              //年齢
	LikeFighters     []Fighter `json:"likefighters"`     //好きな格闘家
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
}
