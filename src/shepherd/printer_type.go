package main

import (
	"gorm.io/gorm"
)

type PrinterAPIConfig struct {
	Type string
	Secret string
	Hostname string
	Port int32 `gorm:"default:80"`
	gorm.Model
}

type Printer struct {
	Name string
	APIConfig PrinterAPIConfig `gorm:"foreignKey:APIConfigID;references:ID"`
	APIConfigID int `json:"-"`
	gorm.Model
}