package pets

import "design-pattern/models"

func NewPet(species string) *models.Pet {
	pet := models.Pet{
		Species:     species,
		Breed:       "",
		MaxWeight:   0,
		MinWeight:   0,
		Description: "no description entered yet",
		LifeSpan:    0,
	}

	return &pet
}
