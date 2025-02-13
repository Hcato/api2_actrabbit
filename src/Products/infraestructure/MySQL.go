package infraestructure

import (
	"api/src/Products/domain"
	"database/sql"
	"fmt"
)

type MySQL struct {
	DB *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{DB: db}
}

func (mysql *MySQL) Save(name string, price float32) error {
	_, err := mysql.DB.Exec("INSERT INTO products (name, price) VALUES (?, ?)", name, price)
	if err != nil {
		return fmt.Errorf("[MySQL] Error al guardar el producto : %w", err)
	}
	return nil
}

func (mysql *MySQL) GetAll() ([]domain.Product, error) {
	rows, err := mysql.DB.Query("SELECT * FROM products")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products []domain.Product

	for rows.Next() {
		var product domain.Product
		err := rows.Scan(&product.Id, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return products, nil
}

func (mysql *MySQL) Delete(id int32) error {
	query := "DELETE FROM products WHERE id = ?"
	result, err := mysql.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("No se encontro el Producto con ID:", id)
	}
	fmt.Println("Ticket Elimonado")
	return nil
}

func (mysql *MySQL) Update(id int32, name string, price float32) error {
	query := "UPDATE products SET name = ?, price =? WHERE id = ? "

	result, err := mysql.DB.Exec(query, name, price, id)
	if err != nil {
		return fmt.Errorf("[MySQL] error alm verificar las filas afectadas:", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("[MySQL] Error al verificar filas afectadas: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("No se encontro el producto")
	}

	fmt.Println("Producto actualizado")
	return nil
}
func (mysql *MySQL) GetById(id int32) (domain.Product, error) {
	var UserById domain.Product

	query := "SELECT id, name, price FROM products WHERE id=?"
	row := mysql.DB.QueryRow(query, id)

	err := row.Scan(&UserById.Id, &UserById.Name, &UserById.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return UserById, fmt.Errorf("producto con id no encontrado", id)
		}
		return UserById, err
	}

	return UserById, nil
}
