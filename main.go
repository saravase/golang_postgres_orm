package main

import (
	plant "golang_postgresql_model/plant"
	"log"
	"time"

	pg "github.com/go-pg/pg"

	db "golang_postgresql_model/database"
)

func main() {
	log.Printf("Hi primz...")
	dbRef := db.Connect()
	//savePlantItem(dbRef)
	//updatePlantItem(dbRef)
	//deletePlantItem(dbRef)
	deletePlantItemUsingTransaction(dbRef)
	//getPlantByID(dbRef)
	//getPlantByIDAndName(dbRef)
	//getPlantOrderByName(dbRef)
}

func savePlantItem(db *pg.DB) {
	newPI1 := &plant.PlantItem{
		Name:        "White Rose",
		Category:    "Flower",
		Description: "Beautiful flower",
		Price:       150.00,
		Avatar:      "white_rose.png",
		Features: struct {
			Name string
			Desc string
			Imp  int
		}{
			Name: "Size",
			Desc: "Small, Big",
			Imp:  1,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive:  true,
	}

	newPI2 := &plant.PlantItem{
		Name:        "Red Rose",
		Category:    "Flower",
		Description: "Beautiful flower",
		Price:       150.00,
		Avatar:      "red_rose.png",
		Features: struct {
			Name string
			Desc string
			Imp  int
		}{
			Name: "Size",
			Desc: "Small, Big",
			Imp:  1,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsActive:  true,
	}

	plantItems := []*plant.PlantItem{
		newPI1, newPI2,
	}
	//newPI.Save(db)
	newPI1.SaveMultiple(db, plantItems)
}

func deletePlantItem(db *pg.DB) {
	plantItem := &plant.PlantItem{
		Name: "White Rose",
	}
	plantItem.DeletePlantItem(db)

}

func deletePlantItemUsingTransaction(db *pg.DB) {
	plantItem := &plant.PlantItem{
		ID: 7,
	}
	plantItem.DeletePlantItemUsingTransaction(db)
}

func updatePlantItem(db *pg.DB) {
	plantItem := &plant.PlantItem{
		ID:        1,
		Price:     150.0,
		UpdatedAt: time.Now(),
	}
	plantItem.UpdatePlantItem(db)

}

func getPlantByID(db *pg.DB) {
	plantItem := &plant.PlantItem{
		ID: 1,
	}
	plantItem.GetPlantByID(db)

}

func getPlantByIDAndName(db *pg.DB) {
	plantItem := &plant.PlantItem{
		ID:   1,
		Name: "Rose",
	}
	plantItem.GetPlantByIDAndName(db)

}

func getPlantOrderByName(db *pg.DB) {
	plantItem := &plant.PlantItem{
		ID: 1,
	}
	plantItem.GetPlantOrderByName(db)

}
