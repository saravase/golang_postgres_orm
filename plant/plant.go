package plant

import (
	"log"
	"time"

	pg "github.com/go-pg/pg"
	orm "github.com/go-pg/pg/orm"
)

type PlantItem struct {
	tableName   struct{} `sql:"plant_info_collection"`
	ID          int      `sql:"id,pk"`
	Name        string   `sql:"name,unique"`
	Category    string   `sql:"category"`
	Description string   `sql:"description"`
	Price       float64  `sql:"price,type:real"`
	Avatar      string   `sql:"avatar"`
	Features    struct {
		Name string
		Desc string
		Imp  int
	} `sql:"features,type:jsonb"`
	CreatedAt time.Time `sql:"created_at"`
	UpdatedAt time.Time `sql:"updated_at"`
	IsActive  bool      `sql:"is_active"`
}

//Crate plantInfo table
func CreatePlantInfoTable(db *pg.DB) error {
	options := &orm.CreateTableOptions{
		IfNotExists: true,
	}

	createErr := db.CreateTable(&PlantItem{}, options)
	if createErr != nil {
		log.Printf("Error while creating plantInfo table. Reason: %v\n", createErr)
		return createErr
	}

	log.Printf("Table PlantInfo created successfully.\n")
	return nil

}

//Insert plantItem into plantInfo table
func (plantItem *PlantItem) Save(db *pg.DB) error {
	insertError := db.Insert(plantItem)

	if insertError != nil {
		log.Printf("Error, While Inserting new plantItem into plantInfo table. Reason %v\n", insertError)
		return insertError
	}

	log.Printf("PlantItem %s inserted into PlantInfo table successfully.\n", plantItem.Name)
	return nil
}

//Insert and Return the inserted plantItem into plantInfo table
func (plantItem *PlantItem) SaveAndReturn(db *pg.DB) (*PlantItem, error) {
	insertResult, insertError := db.Model(plantItem).Returning("*").Insert()

	if insertError != nil {
		log.Printf("Error, While Inserting new plantItem into plantInfo table. Reason %v\n", insertError)
		return nil, insertError
	}
	log.Printf("PlantItem %s inserted into PlantInfo table successfully.\n", plantItem.Name)
	log.Printf("No of affected row: %v.\n", insertResult.RowsAffected)
	return plantItem, nil
}

//Bulk plantItems insert into plantInfo table
func (plantItem *PlantItem) SaveMultiple(db *pg.DB, plantItems []*PlantItem) error {
	_, insertError := db.Model(&plantItems).Insert()

	if insertError != nil {
		log.Printf("Error, While Inserting bulk plantItem into plantInfo table. Reason %v\n", insertError)
		return insertError
	}

	log.Printf("Bulk PlantItem inserted into PlantInfo table successfully.\n")
	return nil
}

//Update the plantItem into plantInfo table
func (plantItem *PlantItem) UpdatePlantItem(db *pg.DB) error {
	_, updateError := db.Model(plantItem).Set("price = ?price").Set("updated_at = ?updated_at").Where("id = ?id").Update()

	if updateError != nil {
		log.Printf("Error, While Updating plantItem into plantInfo table. Reason %v\n", updateError)
		return updateError
	}

	log.Printf("PlantItem %d updated into PlantInfo table successfully.\n", plantItem.ID)
	return nil

}

//Delete the plantItem into plantInfo table
func (plantItem *PlantItem) DeletePlantItem(db *pg.DB) error {
	_, deleteError := db.Model(plantItem).Where("name = ?name").Delete()

	if deleteError != nil {
		log.Printf("Error, While Deleting plantItem into plantInfo table. Reason %v\n", deleteError)
		return deleteError
	}

	log.Printf("PlantItem %s deleted into PlantInfo table successfully.\n", plantItem.Name)
	return nil
}

//Delete the plantItem into plantInfo table using Transaction
func (plantItem *PlantItem) DeletePlantItemUsingTransaction(db *pg.DB) error {

	transaction, transactionError := db.Begin()

	if transactionError != nil {
		log.Printf("Error, While openning the transaction. Reason %v\n", transactionError)
		return transactionError
	}

	_, deleteError := transaction.Model(plantItem).Where("id = ?id").Delete()

	if deleteError != nil {
		log.Printf("Error, While Deleting plantItem into plantInfo table. Reason %v\n", deleteError)
		transaction.Rollback()
		return deleteError
	}

	transaction.Commit()
	log.Printf("PlantItem %s deleted into PlantInfo table successfully.\n", plantItem.Name)
	return nil
}

//Select the plantItem based on plantID
func (plantItem *PlantItem) GetPlantByID(db *pg.DB) error {
	selectError := db.Select(plantItem)

	if selectError != nil {
		log.Printf("Error, While Selecting PlantItem from PlantInfo table through PlantID. Reason %v\n", selectError)
		return selectError
	}

	log.Printf("PlantItem retrived successfully from PlantInfo table through PlantID. PlantItem :%v\n", *plantItem)
	return nil

}

//Select the plantItem based on plantID and plantName
func (plantItem *PlantItem) GetPlantByIDAndName(db *pg.DB) error {
	selectError := db.Model(plantItem).
		Column("name", "price").
		Where("id = ?id", plantItem.ID).
		Where("name = ?0", plantItem.Name).
		Select()

	if selectError != nil {
		log.Printf("Error, While Selecting PlantItem from PlantInfo table through PlantID and PlantName. Reason %v\n", selectError)
		return selectError
	}

	log.Printf("PlantItem retrived successfully from PlantInfo table through PlantID and PlantName. PlantItem :%v\n", *plantItem)
	return nil

}

//Select the plantItem order by name based on plantID and plantName
func (plantItem *PlantItem) GetPlantOrderByName(db *pg.DB) error {
	var items []PlantItem

	selectError := db.Model(&items).Column("name").
		Where("id = ?0", plantItem.ID).
		WhereOr("id = ?0", 2).
		Offset(1).
		Limit(2).
		Order("name asc").
		Select()
	if selectError != nil {
		log.Printf("Error, While Selecting PlantItem from PlantInfo table through PlantID and order by PlantName. Reason %v\n", selectError)
		return selectError
	}

	log.Printf("PlantItem retrived successfully from PlantInfo table through PlantID and order by PlantName. PlantItem :%v\n", items)
	return nil

}
