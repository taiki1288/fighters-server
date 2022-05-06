package entity

import "time"

type User struct {
	ID               string    //UserID
	Name             string    //名前
	SelfIntroduction string    //自己紹介
	Age              int       //年齢
	likeFighters     []Fighter //好きな格闘家
	CreatedAt        time.Time
	UpdatedAt        time.Time
}
