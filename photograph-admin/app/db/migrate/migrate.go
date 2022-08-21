package migrate

import (
	"photograph/app/db"
	"photograph/app/models"

	"gorm.io/gorm"
)

func main() {
	con := db.Init()
	defer db.Close(con)
	migrate(con)
}

func migrate(con *gorm.DB) {
	con.AutoMigrate(&models.Photo{})
}
