package main

import (
	"context"
	"fmt"
	"log"

	"github.com/PON/app/controllers"
	_ "github.com/PON/app/docs"
	"github.com/PON/app/ent"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/PON/app/ent/userstatus"

	"github.com/PON/app/ent/roomstatus"
	"github.com/PON/app/ent/roomtype"
)

//-------------------------------------------------------------------

// Users defines the struct for the Users
type Users struct {
	User []User
}

// User defines the struct for the User
type User struct {
	Name       string
	Email      string
	Userstatus int
}

// Roomages defines the struct for the Roomages
type Roomages struct {
	Roomage []Roomage
}

// Roomage defines the struct for the Roomage
type Roomage struct {
	RoomAge int
	Text    string
}

// RoomTypes defines the struct for the RoomTypes
type RoomTypes struct {
	RoomType []RoomType
}

// RoomType defines the struct for the RoomType
type RoomType struct {
	Type string
	COST int
}

// Rooms defines the struct for the Rooms
type Rooms struct {
	Room []Room
}

// Room defines the struct for the Room
type Room struct {
	Type   int
	Status int
	RoomName string
}



//-------------------------------------------------------------------
// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())
	client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")

	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewUserController(v1, client)
	controllers.NewRoomController(v1, client)
	controllers.NewInsuranceController(v1, client)
	controllers.NewRoomStatusController(v1, client)
	controllers.NewRoomTypeController(v1, client)
	controllers.NewRoomageController(v1, client)
	controllers.NewUserStatusController(v1, client)
	controllers.NewRentController(v1, client)
	//-------------------------------------------------------------------
	
	
	// Set userstatus Data
	usersta := []string{"Admin", "Resident", "Staff"}
	for _, r := range usersta {
		client.UserStatus.
			Create().
			SetUserStatus(r).
			Save(context.Background())
	}
	
	// Set Users Data
	users := Users{
		User: []User{
			User{"Khatadet khianchainat", "B6103866@gmail.com", 1},
			User{"nara haru", "c101@example.com", 2},
			User{"morani rode", "c102@example.com", 2},
			User{"faratell yova", "c103@example.com", 2},
			User{"pulla visan", "c104@example.com", 2},
			User{"omaha yad", "c104@example.com", 2},
		},
	}

	for _, p := range users.User {

		u, err := client.UserStatus.
			Query().
			Where(userstatus.IDEQ(int(p.Userstatus))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.User.
			Create().
			SetUserEmail(p.Email).
			SetNAME(p.Name).
			SetUserUserstatus(u).
			Save(context.Background())
	}

	// Set RoomTypes Data
	Roomage := Roomages{
		Roomage: []Roomage{
			Roomage{365, "ปี"},
			Roomage{30, "เดือน"},
			Roomage{1, "วัน"},
		},
	}

	for _, u := range Roomage.Roomage {
		client.Roomage.
			Create().
			SetRoomAge(u.RoomAge).
			SetText(u.Text).
			Save(context.Background())
	}

	// Set Resolutions Data
	Insurance := []int{2000, 3600, 4000, 7000, 18000}
	for _, r := range Insurance {
		client.Insurance.
			Create().
			SetInsurance(r).
			Save(context.Background())
	}
	// Set RoomTypes Data
	roomtypes := RoomTypes{
		RoomType: []RoomType{
			RoomType{"ห้อง แอร์", 5000},
			RoomType{"ห้อง พัดลม", 2500},
			RoomType{"ห้อง Suite", 19000},
		},
	}

	for _, u := range roomtypes.RoomType {
		client.RoomType.
			Create().
			SetRoomType(u.Type).
			SetCOST(u.COST).
			Save(context.Background())
	}

	

	// Set roomstatus Data
	Roomstatus := []string{"Available", "Unavailable"}
	for _, r := range Roomstatus {
		client.RoomStatus.
			Create().
			SetRoomStatus(r).
			Save(context.Background())
	}

	// Set room Data
	room := Rooms{
		Room: []Room{
			Room{1, 1,"101"},
			Room{1, 1,"102"},
			Room{1, 1,"103"},
		},
	}
	for _, y := range room.Room {

		RoomType, err := client.RoomType.
			Query().
			Where(roomtype.IDEQ(int(y.Type))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		RoomStatus, err := client.RoomStatus.
			Query().
			Where(roomstatus.IDEQ(int(y.Status))).
			Only(context.Background())

		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.Room.
			Create().
			SetRoomRoomstatus(RoomStatus).
			SetRoomRoomtype(RoomType).
			SetRoomName(y.RoomName).
			Save(context.Background())
	}

	//-------------------------------------------------------------------
	
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
