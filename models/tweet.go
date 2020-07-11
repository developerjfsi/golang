package models

/*Tweet es el formato o estructura que tendrá nuestro Tweet en la BD */
type Tweet struct {
	Mensaje string `bson:"mensaje" json:"mensaje,omitempty"`
}
