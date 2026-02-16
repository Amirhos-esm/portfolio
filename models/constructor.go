package models

import (
	"time"
)

func NewProject() *Project {

	return &Project{
		Tags:          make([]string, 0),
		GalleryImages: make([]string, 0),
		TechStack:     make(map[string][]string),
		Features:      make([]*ProjectFeature, 0),
	}
}

func NewEducation() *Education {
	return &Education{}
}
func NewExperience() *Experience {
	return &Experience{}
}

func NewMessage(message, email, fullname string) *Message {
	return &Message{
		Message:  message,
		Email:    email,
		Fullname: fullname,
		Createat: time.Now(),
	}
}
