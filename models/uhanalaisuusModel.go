package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Uhanalaisuus struct {
    ID                primitive.ObjectID `bson:"_id" json:"-"`  // "-" piilottaa kentän JSON:sta
    UhanalaisuusID    int                `bson:"uhanalaisuus_id" json:"uhanalaisuus_id"`
    Nimi              string             `bson:"nimi" json:"nimi"`
    Lyhenne           string             `bson:"lyhenne" json:"lyhenne"`
}
