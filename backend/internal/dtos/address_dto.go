package dtos

type CreateAddressDTO struct {
	Street  string `json:"street" validate:"required"`
	Number  string `json:"number" validate:"required"`
	City    string `json:"city" validate:"required"`
	State   string `json:"state" validate:"required"`
	ZipCode string `json:"zip_code" validate:"required"`
}

type UpdateAddressDTO struct {
	Street  *string `json:"street"`
	Number  *string `json:"number"`
	City    *string `json:"city"`
	State   *string `json:"state"`
	ZipCode *string `json:"zip_code"`
}
