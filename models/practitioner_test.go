package models

import (
	"testing"
)

func TestToPublicPractitioner(t *testing.T) {
	p := &PublicPractitioner{
		ID:             1,
		UID:            "123",
		FirstName:      "John",
		SecondName:     "Doe",
		LastName:       "Smith",
		SecondLastName: "Jones",
		FullName:       "John Doe Smith Jones",
		RegistryID:     "ABC123",
		Image:          "https://example.com/image.jpg",
		Gender:         "male",
		Bio:            "A brief bio",
	}

	result := p.ToPublicPractitioner()

	if result.ID != p.ID {
		t.Errorf("Expected ID to be %d, got %d", p.ID, result.ID)
	}

	if result.UID != p.UID {
		t.Errorf("Expected UID to be %s, got %s", p.UID, result.UID)
	}

	if result.FirstName != p.FirstName {
		t.Errorf("Expected FirstName to be %s, got %s", p.FirstName, result.FirstName)
	}

	// Similarly, add assertions for all other fields
}

func TestToPublicPractitionerAggregated(t *testing.T) {
	// Similarly, write test cases for ToPublicPractitionerAggregated method
}
