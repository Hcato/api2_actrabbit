package dependencies

import (
	"api/src/Users/application"
	"api/src/Users/infraestructure"
	"api/src/Users/infraestructure/controllers"
	"api/src/core"
	"database/sql"
	"fmt"
)

var (
	mySQL infraestructure.MySQL
	db    *sql.DB
)

func Init() {
	db, err := core.ConnectToDB()

	if err != nil {
		fmt.Println("server error")
		return
	}

	mySQL = *infraestructure.NewMySQL(db)

}
func CloseDB() {
	if db != nil {
		db.Close()
		fmt.Println("Conexión a la base de datos cerrada.")
	}
}
func GetCreateUserController() *controllers.CreateUserController {
	caseCreateProduct := application.NewCreateUser(&mySQL)
	return controllers.NewCreateProductController(caseCreateProduct)
}

func GetGetAllUserController() *controllers.GetAllUserController {
	caseGetAllUsers := application.NewGetAllProduct(&mySQL)
	return controllers.NewGetAllUserController(*caseGetAllUsers)
}

func GetDeleteUserController() *controllers.DeleteUserController {
	caseDeleteUser := application.NewDeleteUser(&mySQL)
	return controllers.NewDeleteUserController(caseDeleteUser)
}

func GetUpdateUserController() *controllers.UpdateUserController {
	caseUpdateUser := application.NewUpdateProduct(&mySQL)
	return controllers.NewUpdateUserController(caseUpdateUser)
}

func GetGetByIdUserController() *controllers.GetByIdUserController {
	caseGetByIdUser := application.NewGetByIdUser(&mySQL)
	return controllers.NewGetByIdUserController(caseGetByIdUser)
}
