package dao

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"http/dto"
	"log"
)

var DB *sql.DB

var CreateTable = "CREATE TABLE IF NOT EXISTS user(user_id int primary key auto_increment, name varchar(20), surname varchar(20), created_at datetime default CURRENT_TIMESTAMP, updated_at datetime default CURRENT_TIMESTAMP);"

//var addUserQuery = `INSERT INTO user (name, surname, created_at, updated_at) VALUES ($1,$2,$3,$4);`
var addUserQuery = `INSERT INTO user (name, surname, created_at, updated_at) VALUES (?,?,?,?);`

var lastUser = `SELECT * FROM user WHERE user_id = LAST_INSERT_ID()`

func init() {
	var err error
	DB, err = sql.Open("mysql", "user:password@tcp(localhost:1234)/db?parseTime=true")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Fatal("Error during connection opening")
		//panic(err.Error())
	}

	log.Println("1")

	// defer the close till after the main function has finished
	// executing
	//defer DB.Close()

	creation, errCreation := DB.Query(CreateTable)

	if errCreation != nil {
		log.Fatal("Error during creation of the table")
	}

	log.Println("2")

	defer creation.Close()
}

func InsertUser(u dto.User) (dto.User, error) {
	var insertedUser dto.User

	//fmt.Println("User details")
	//fmt.Println(u.Name)
	//fmt.Println(u.CreatedAt)
	//fmt.Println(u.UpdatedAt)
	//fmt.Println(u.Surname)
	//fmt.Println("--------------")

	//err := DB.QueryRow(addUserQuery, u.Name, u.Surname, u.CreatedAt, u.UpdatedAt).Scan(&insertedUser.Id, &insertedUser.Name, &insertedUser.Surname, &insertedUser.CreatedAt, &insertedUser.UpdatedAt)
	//_, errInsert := DB.Query(addUserQuery)
	insertPrepare, errInsert := DB.Prepare(addUserQuery)
	//fmt.Println(results)
	if errInsert != nil {
		log.Printf("Unable to execute the query %s for error %s \n", addUserQuery, errInsert.Error())
		return insertedUser, errInsert
	}
	insertPrepare.Exec(u.Name, u.Surname, u.CreatedAt, u.UpdatedAt)
	errRetrieve := DB.QueryRow(lastUser).Scan(&insertedUser.Id, &insertedUser.Name, &insertedUser.Surname, &insertedUser.CreatedAt, &insertedUser.UpdatedAt)
	if errRetrieve != nil {
		log.Printf("Unable to execute the query %s for error %s \n", addUserQuery, errRetrieve.Error())
		return insertedUser, errRetrieve
	}
	log.Printf("Row inserted with success %v \n", insertedUser)
	return insertedUser, errInsert
}
