package entity

import "time"

type Fighter struct {
	ID          string    `json:"id"`          //格闘家ID
	Name        string    `json:"name"`        //格闘家名
	Type        string    `json:"type"`        //格闘技の種類 例）総合格闘技
	Results     string    `json:"results"`     //戦績
	Description string    `json:"description"` //選手の説明
	Backbone    string    `json:"backbone"`    //選手のバックボーン
	Age         int       `json:"age"`         //年齢
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}
