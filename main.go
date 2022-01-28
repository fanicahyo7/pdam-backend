package main

import (
	"log"
	"os"
	"pdam/handler"
	"pdam/repository"
	"pdam/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	//cek env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbusername := os.Getenv("DB_USERNAME")
	if dbusername == "" {
		log.Fatal("Username is required!")
	}

	dbpassword := os.Getenv("DB_PASSWORD")

	dbhost := os.Getenv("DB_HOST")
	if dbusername == "" {
		log.Fatal("Host is required!")
	}

	dbport := os.Getenv("DB_PORT")
	if dbusername == "" {
		log.Fatal("Port is required!")
	}

	dbdatabase := os.Getenv("DB_DATABASE")
	if dbusername == "" {
		log.Fatal("Database is required!")
	}

	dsn := dbusername + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbdatabase + "?charset=utf8mb4&parseTime=True&loc=Local"

	//koneksi ke database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	// rand.Seed(time.Now().UTC().UnixNano())
	// fmt.Println(rand.Int())
	// fmt.Println(rand.Int())
	// fmt.Println(rand.Int())

	grouprepository := repository.NewGroupRepository(db)

	// var input model.Group
	// input = model.Group{}
	// input.Kode = "111222"
	// input.Nama = "dari repo"
	// input.Tarif1 = 10000
	// input.Tarif2 = 15000

	// anu, err := grouprepository.FindGroupByName("repo")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(anu)

	groupservice := service.NewService(grouprepository)

	// input := helper.Inputgroup{}
	// input.Kode = "111222"
	// input.Nama = "dari repo"
	// input.Tarif1 = 10000
	// input.Tarif2 = 15000

	// anu, err := groupservice.GetGroups()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(anu)

	groupHandler := handler.NewGroupHandler(groupservice)

	router := gin.Default()
	api := router.Group("/api/v1")
	api.GET("/group/:id", groupHandler.GetGroupByID)
	groupResource := api.Group("/groups/")
	{
		groupResource.GET("/", groupHandler.GetGroups)
		groupResource.GET("/:nama", groupHandler.GetGroupByNama)
		groupResource.POST("/", groupHandler.CreateGroup)
		groupResource.POST("/edit/", groupHandler.UpdateGroup)
		groupResource.GET("/delete/:kode", groupHandler.DeleteGroup)
	}

	router.Run(":8081")
}
