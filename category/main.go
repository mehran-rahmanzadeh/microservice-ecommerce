package main

import (
	"context"
	"fmt"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/domain/model"
	"github.com/mehran-rahmanzadeh/microservice-ecommerce/repository"
	"github.com/spf13/viper"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func main() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	log.Println("Config loaded successfully")

	// init DB
	// TODO: change sqlite to PostgreSQL
	db, err := gorm.Open(sqlite.Open(viper.GetString("db.name")), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database")
	}
	log.Println("Connected to db successfully")

	err = db.AutoMigrate(model.Category{})
	if err != nil {
		panic("failed to migrate Category")
	}
	log.Println("Database migrated successfully")

	// manual test
	db.Delete(&model.Category{}, "deleted_at is null")
	var parent model.Category
	db.Create(&model.Category{
		Title:       "Parent Category",
		Description: "Set of best products",
		Image:       "mobile.png",
	}).Scan(&parent)
	db.Create(&model.Category{
		Title:       "Child Category",
		Description: "Child",
		Image:       "a.png",
		ParentID:    &parent.ID,
	}).Scan(&parent)
	db.Create(&model.Category{
		Title:       "Child of child",
		Description: "Child of child",
		Image:       "cc.png",
		ParentID:    &parent.ID,
	})

	categoryRepository := repository.NewCategoryRepository(db)
	fmt.Println("------ FetchCategories ------")
	categories, err := categoryRepository.FetchCategories(context.Background())
	fmt.Println(categories)
	fmt.Println("-----------------------------")

	fmt.Println("------ GetCategoryByID ------")
	category, err := categoryRepository.GetCategoryByID(context.Background(), parent.ID)
	fmt.Println(category)
	fmt.Println("-----------------------------")

	fmt.Println("------ GetChildrenCategories ------")
	categories, err = categoryRepository.GetChildrenCategories(context.Background(), parent.ID)
	fmt.Println(categories)
	fmt.Println("-----------------------------")

	fmt.Println("------ GetParentCategories ------")
	categories, err = categoryRepository.GetParentCategories(context.Background())
	fmt.Println(categories)
	fmt.Println("-----------------------------")

	fmt.Println("------ SearchCategories ------")
	categories, err = categoryRepository.SearchCategories(context.Background(), "parent")
	fmt.Println(categories)
	fmt.Println("-----------------------------")

}
