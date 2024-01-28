package migrations

import(
	"fmt"
	"github.com/adyfp24/golang-fiber-learn/database"
	"github.com/adyfp24/golang-fiber-learn/models"
	"log"
)

func RunMigration(){
	db, err := database.InitDB()
	err = db.AutoMigrate(&models.Author{}, &models.Book{})
	if err != nil {
		log.Print(err)
	}

	fmt.Println("database berhasil di migrasi")
}
