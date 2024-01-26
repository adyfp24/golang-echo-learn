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
	db.Exec("ALTER TABLE books ADD CONSTRAINT fk_author_id FOREIGN KEY (author_id) REFERENCES authors(id) ON DELETE CASCADE ON UPDATE CASCADE")
	if err != nil {
		log.Print(err)
	}

	fmt.Println("database berhasil di migrasi")
}
