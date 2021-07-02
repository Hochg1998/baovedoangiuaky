package repo

import (
	"fmt"

	"github.com/TechMaster/golang/15GoMySQL/model"
)

func initCategory() {
	categories := map[int]string{
		1: "cats",
		2: "dogs",
		3: "catsfood",
		4: "dogsfood",
		5: "catsitem",
		6: "dogsitem",
	}
	fmt.Println(categories)

	for _, category := range categories {
		createCategory := model.Category{Name: category}
		result := Db.Create(&createCategory)
		if result.Error != nil {
			panic(result.Error)
		}
	}
}
