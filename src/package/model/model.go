package model

type CryptoCurrency struct {
	Id    int    `bson:"id,omitempty" json:"id,omitempty"`
	Name  string `bson:"name,omitempty" json:"name,omitempty" validate:"required"`
	Token string `bson:"token,omitempty" json:"token,omitempty" validate:"required"`
	Votes int    `bson:"votes,omitempty" json:"votes,omitempty"`
}
