package main

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// TO-DO: These should all be hooked into the generated protobuf with gorm functionality. These are (bad) wrappers
//        to support db operations without a bunch of (un)marshaling and other modifications because of this bad idea.

type PrintRequest struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name      string    `gorm:"uniqueIndex"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (p *PrintRequest) BeforeCreate(tx *gorm.DB) (err error) {
	p.ID = uuid.New()
	return
}
