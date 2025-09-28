package models

// omitempty poistaa tyhjän kentän jsonista
// validate määrittelee pakolliset kentät

import "go.mongodb.org/mongo-driver/bson/primitive"

type Bird struct {
    ID               primitive.ObjectID `bson:"_id,omitempty" json:"id"`
    Nimi             string             `bson:"nimi" json:"nimi" validate:"required"`
    TieteellinenNimi string             `bson:"tieteellinenNimi" json:"tieteellinenNimi" validate:"required"`
    // Koko             Koko               `bson:"koko,omitempty" json:"koko"`
    SiipivaliCm     int                `bson:"siipivali_cm,omitempty" json:"siipivali_cm"`
    Varitys          string             `bson:"varitys,omitempty" json:"varitys,omitempty"`
    Elinymparisto    []string           `bson:"elinymparisto,omitempty" json:"elinymparisto,omitempty"`
    Ravinto          []string           `bson:"ravinto,omitempty" json:"ravinto,omitempty"`
    Uhanalaisuus     int                `bson:"uhanalaisuus_id" json:"uhanalaisuus_id"`
}

// testi
// type Koko struct {
//     PituusCm    int `bson:"pituus_cm,omitempty" json:"pituus_cm"`
//     SiipivaliCm int `bson:"siipivali_cm,omitempty" json:"siipivali_cm"`
//     PainoG      int `bson:"paino_g,omitempty" json:"paino_g"`
// }
