package services

import (
	"fmt"
)

type InfoService interface {
	GetInfo() string
}

type Availability interface {
	VerifyAvailability(requested int, available int) bool
}

type SupplierService interface {
	InfoService
	Availability
}

type Supplier struct {
	CNPJ    string
	Contact string
	City    string
}

func (s Supplier) GetInfo() string {
	return fmt.Sprintf("CNPJ: %s | Contact: %s | City: %s", s.CNPJ, s.Contact, s.City)
}

func (s Supplier) VerifyAvailability(requested int, available int) bool {
	return requested <= available
}
