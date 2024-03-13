package common

type CustomerAddressData struct {
	Street       string `json:"street"`
	Number       string `json:"number"`
	ZipCode      string `json:"zip_code"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	Complement   string `json:"complement"`
}

type CustomerData struct {
	Name      string              `json:"name"`
	Email     string              `json:"email"`
	Phone     string              `json:"phone"`
	BirthDate string              `json:"birth_date"`
	Document  string              `json:"document"`
	Address   CustomerAddressData `json:"address"`
}
