package types

import (
	"time"

	domain "github.com/kamilsk/guard/pkg/service/types"
)

// License TODO issue#docs
type License struct {
	ID        domain.ID       `db:"id"`
	Contract  domain.Contract `db:"contract"`
	CreatedAt time.Time       `db:"created_at"`
	UpdatedAt *time.Time      `db:"updated_at"`
	DeletedAt *time.Time      `db:"deleted_at"`
}
