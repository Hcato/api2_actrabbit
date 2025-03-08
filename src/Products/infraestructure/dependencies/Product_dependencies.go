package dependencies

import (
	"api/src/Products/application"
	"api/src/Products/infraestructure"
	"api/src/Products/infraestructure/controllers"
	"api/src/core"
	"database/sql"
	"fmt"
)

var (
	mySQL    infraestructure.MySQL
	db       *sql.DB
	consumer *infraestructure.RabbitMQConsumer
)

func Init() {
	var err error
	db, err = core.ConnectToDB()
	if err != nil {
		fmt.Println("server error")
		return
	}

	mySQL = *infraestructure.NewMySQL(db)

	// Inicializar casos de uso
	createProductUseCase := application.NewCreateProduct(&mySQL)
	getByIdProductUseCase := application.NewGetByIdProduct(&mySQL)
	getAllProductUseCase := application.NewGetAllProduct(&mySQL)
	updateProductUseCase := application.NewUpdateProduct(&mySQL)
	deleteProductUseCase := application.NewDeleteProduct(&mySQL)

	// Inicializar RabbitMQConsumer con todos los casos de uso
	consumer, err = infraestructure.NewRabbitMQConsumer(
		"amqp://cato:5678@3.233.111.240/", // URL de RabbitMQ
		"product",                         // Nombre de la cola
		"products",                        // Mensaje de respuesta
		createProductUseCase,              // Caso de uso de creación
		getAllProductUseCase,              // Caso de uso de obtener todos
		getByIdProductUseCase,             // Caso de uso de obtener por ID
		updateProductUseCase,              // Caso de uso de actualización
		deleteProductUseCase,              // Caso de uso de eliminación
	)

	if err != nil {
		fmt.Println("[RabbitMQ] Error al conectar con la cola:", err)
		return
	}

	// Iniciar la escucha de mensajes
	consumer.Start()
}

func CloseDB() {
	if db != nil {
		db.Close()
		fmt.Println("Conexión a la base de datos cerrada.")
	}

	if consumer != nil {
		consumer.Close()
		fmt.Println("Conexión a RabbitMQ cerrada.")
	}
}

func GetCreateProductController() *controllers.CreateProductController {
	caseCreateProduct := application.NewCreateProduct(&mySQL)
	return controllers.NewCreateProductController(caseCreateProduct)
}

func GetGetAllProductController() *controllers.GetAllProductController {
	caseGetAllProduct := application.NewGetAllProduct(&mySQL)
	return controllers.NewGetAllProductController(*caseGetAllProduct)
}

func GetDeleteProductController() *controllers.DeleteProductController {
	caseDeleteProduct := application.NewDeleteProduct(&mySQL)
	return controllers.NewDeleteProductController(caseDeleteProduct)
}

func GetUpdateProductController() *controllers.UpdateProductController {
	caseUpdateProduct := application.NewUpdateProduct(&mySQL)
	return controllers.NewUpdateProductController(caseUpdateProduct)
}

func GetByIdProductController() *controllers.GetByIdProductController {
	caseGetByIdProduct := application.NewGetByIdProduct(&mySQL)
	return controllers.NewGetByIdProductController(caseGetByIdProduct)
}
