package models

// Smarthphone
// Se escribe Smartphone (CamelCase) porque va a ser exportada
// si se escribe con minuscula solo es posible verla
// dentro del paquete.
type Smartphone struct {
	Id            int64
	Name          string
	Price         int
	CountryOrigin string
	Os            string
}

type CreateSmartphoneCMD struct {
	Name          string `json:"name"`
	Price         int    `json:"price"`
	CountryOrigin string `json:"country_origin"`
	Os            string `json:"os"`
}
