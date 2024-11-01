package main

import "time"

type XmlFile struct {
	ID        uint `gorm:"primaryKey"`
	File      string
	Active    bool
	CreatedAt time.Time
	UpdatedAt time.Time
}
