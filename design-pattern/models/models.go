package models

import (
	"database/sql"
	"time"
)

// var db *sql.DB
var repo Repository

type Models struct {
	DogBreed DogBreed
}

func New(conn *sql.DB) *Models {
	if conn != nil {
		repo = newMysqlRepository(conn)
	} else {
		repo = newTestRepository(conn)
	}

	return &Models{
		DogBreed: DogBreed{},
	}
}

type DogBreed struct {
	ID               int    `json:"id"`
	Breed            string `json:"breed"`
	Details          string `json:"details"`
	AlternateNames   string `json:"alternate_names"`
	AverageWeight    int    `json:"average_weight"`
	Lifespan         int    `json:"lifespan"`
	WeightHighLbs    int    `json:"weight_high_lbs"`
	GeographicOrigin string `json:"geographic_origin"`
	WeightLowLbs     int    `json:"weight_low_lbs"`
}

type CatBreed struct {
	ID               int    `json:"id"`
	Breed            string `json:"breed"`
	AlternateNames   string `json:"alternate_names"`
	Lifespan         int    `json:"lifespan"`
	AverageWeight    int    `json:"average_weight"`
	Details          string `json:"details"`
	GeographicOrigin string `json:"geographic_origin"`
	WeightHighLbs    int    `json:"weight_high_lbs"`
	WeightLowLbs     int    `json:"weight_low_lbs"`
}

func (d *DogBreed) All() ([]*DogBreed, error) {
	return repo.AllDogBreeds()
}

type Dog struct {
	Breed            DogBreed  `json:"breed"`
	Breeder          Breeder   `json:"breeder"`
	BreederID        int       `json:"breeder_id"`
	BreedID          int       `json:"breed_id"`
	Color            string    `json:"color"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	Description      string    `json:"description"`
	DogName          string    `json:"dog_name"`
	ID               int       `json:"id"`
	SpayedOrNeutered int       `json:"spayed_or_neutered"`
	Weight           int       `json:"weight"`
}

type Cat struct {
	Breed            CatBreed  `json:"breed"`
	Breeder          Breeder   `json:"breeder"`
	BreederID        int       `json:"breeder_id"`
	BreedID          int       `json:"breed_id"`
	CatName          string    `json:"cat_name"`
	Color            string    `json:"color"`
	DateOfBirth      time.Time `json:"date_of_birth"`
	Description      string    `json:"description"`
	ID               int       `json:"id"`
	SpayedOrNeutered int       `json:"spayed_or_neutered"`
	Weight           int       `json:"weight"`
}

type Breeder struct {
	Active      int         `json:"active"`
	Address     string      `json:"address"`
	BreederName string      `json:"breeder_name"`
	CatBreeds   []*CatBreed `json:"cat_breeds"`
	City        string      `json:"city"`
	Country     string      `json:"country"`
	DogBreeds   []*DogBreed `json:"dog_breeds"`
	Email       string      `json:"email"`
	ID          int         `json:"id"`
	Phone       string      `json:"phone"`
	State       string      `json:"prov_state"`
	Zip         string      `json:"zip"`
}

type Pet struct {
	Breed       string `json:"breed"`
	Description string `json:"description"`
	LifeSpan    int    `json:"lifespan"`
	MaxWeight   int    `json:"max_weight"`
	MinWeight   int    `json:"min_weight"`
	Species     string `json:"species"`
}
