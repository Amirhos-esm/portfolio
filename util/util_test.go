package util

import (
	"fmt"
	"testing"
	"time"
)

// --- Your types ---
type Education struct {
	ID          uint       `json:"id"`
	Degree      string     `json:"degree"`
	School      string     `json:"school"`
	Location    string     `json:"location"`
	StartDate   time.Time  `json:"startDate"`
	EndDate     *time.Time `json:"endDate,omitempty"`
	Description string     `json:"description"`
}

type UpdateEducationInput struct {
	Degree      *string    `json:"degree,omitempty"`
	School      *string    `json:"school,omitempty"`
	StartDate   *time.Time `json:"startDate,omitempty"`
	Location    *string    `json:"location,omitempty"`
	EndDate     *time.Time `json:"endDate,omitempty"`
	Description *string    `json:"description,omitempty"`
}

func TestPatchStruct(t *testing.T) {
	start := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	timePatch := time.Now()
	newDegree := "Master of Science"
	newLocation := "Berlin"
	newDesc := "Updated description"

	dst := &Education{
		ID:          1,
		Degree:      "Bachelor of Arts",
		School:      "Old School",
		Location:    "Munich",
		StartDate:   start,
		EndDate:     &end,
		Description: "Old description",
	}

	src := &UpdateEducationInput{
		Degree:      &newDegree,
		Location:    &newLocation,
		Description: &newDesc,
		StartDate:   Ptr(timePatch),
	}

	err := PatchStruct(dst, src)
	if err != nil {
		t.Fatalf("PatchStruct failed: %v", err)
	}

	// ✅ Assertions
	if dst.Degree != newDegree {
		t.Errorf("Degree not updated, got %s", dst.Degree)
	}
	if dst.Location != newLocation {
		t.Errorf("Location not updated, got %s", dst.Location)
	}
	if dst.Description != newDesc {
		t.Errorf("Description not updated, got %s", dst.Description)
	}
	if dst.School != "Old School" {
		t.Errorf("School should remain unchanged, got %s", dst.School)
	}
	if !dst.StartDate.Equal(timePatch) {
		t.Errorf("StartDate should remain unchanged, got %v", dst.StartDate)
	}
	if dst.EndDate == nil || !dst.EndDate.Equal(end) {
		t.Errorf("EndDate should remain unchanged, got %v", dst.EndDate)
	}
	fmt.Println(dst)
}



// --- Test ---
func TestPatchStruct2(t *testing.T) {
	// Original data
	start := time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC)
	end := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)

	// New patch data
	timePatch := time.Now().Truncate(time.Second) // truncate for safe equality
	endPatch := time.Now().Add(24 * time.Hour).Truncate(time.Second)
	newDegree := "Master of Science"
	newLocation := "Berlin"
	newDesc := "Updated description"
	newSchool := "New School"

	dst := &Education{
		ID:          1,
		Degree:      "Bachelor of Arts",
		School:      "Old School",
		Location:    "Munich",
		StartDate:   start,
		EndDate:     &end,
		Description: "Old description",
	}

	src := &UpdateEducationInput{
		Degree:      Ptr(newDegree),
		School:      Ptr(newSchool),       // let's patch school too
		Location:    Ptr(newLocation),
		Description: Ptr(newDesc),
		StartDate:   Ptr(timePatch),
		EndDate:     Ptr(endPatch),
	}

	fmt.Println("Before patch:", dst)

	err := PatchStruct(dst, src)
	if err != nil {
		t.Fatalf("PatchStruct failed: %v", err)
	}

	fmt.Println("After patch:", dst)

	// --- Assertions with subtests ---
	tests := []struct {
		name     string
		got      interface{}
		expected interface{}
	}{
		{"Degree", dst.Degree, newDegree},
		{"School", dst.School, newSchool},
		{"Location", dst.Location, newLocation},
		{"Description", dst.Description, newDesc},
		{"StartDate", dst.StartDate, timePatch},
		{"EndDate", *dst.EndDate, endPatch},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.got != tt.expected {
				t.Errorf("%s not updated correctly: got %v, want %v", tt.name, tt.got, tt.expected)
			}
		})
	}
}