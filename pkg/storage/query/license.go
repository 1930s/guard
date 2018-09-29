package query

import domain "github.com/kamilsk/guard/pkg/service/types"

// CreateLicense TODO issue#docs
type CreateLicense struct {
	ID       *domain.ID
	Contract domain.Contract
}

// DeleteLicense TODO issue#docs
type DeleteLicense struct {
	ID domain.ID
}

// ReadLicense TODO issue#docs
type ReadLicense struct {
	ID domain.ID
}

// UpdateLicense TODO issue#docs
type UpdateLicense struct {
	ID       domain.ID
	Contract domain.Contract
}

// ---

// ExtendLicense TODO issue#docs
// TODO issue#future
type ExtendLicense struct {
	ID    domain.ID
	Patch interface{}
}

// RegisterLicense TODO issue#docs
type RegisterLicense struct {
	ID       domain.ID
	Contract domain.Contract
}