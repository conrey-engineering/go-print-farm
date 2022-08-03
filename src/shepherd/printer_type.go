package main

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

// TO-DO: These should all be hooked into the generated protobuf with gorm functionality. These are (bad) wrappers
//        to support db operations without a bunch of (un)marshaling and other modifications because of this bad idea.

type PrinterAPIConfig struct {
	ID        uuid.UUID `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Type      string
	Secret    string
	Hostname  string
	Port      int32 `gorm:"default:80"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type Printer struct {
	ID          uuid.UUID        `gorm:"primary_key;type:uuid;default:uuid_generate_v4()"`
	Name        string           `gorm:"uniqueIndex"`
	APIConfig   PrinterAPIConfig `gorm:"foreignKey:APIConfigID;references:ID"`
	APIConfigID string           `json:"-"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	// DeletedAt gorm.DeletedAt `gorm:"index"`
}

func (printer *Printer) BeforeCreate(tx *gorm.DB) (err error) {
	printer.ID = uuid.New()
	return
}

func (papi *PrinterAPIConfig) BeforeCreate(tx *gorm.DB) (err error) {
	papi.ID = uuid.New()
	return
}
