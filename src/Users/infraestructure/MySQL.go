package infraestructure

import (
	"api/src/Users/domain"
	"database/sql"
	"fmt"
)

type MySQL struct {
	DB *sql.DB
}

func NewMySQL(db *sql.DB) *MySQL {
	return &MySQL{DB: db}
}

func (mysql *MySQL) Save(name string, lastname string) error {
	_, err := mysql.DB.Exec("INSERT INTO users (name, lastname) VALUES (?, ?)", name, lastname)
	if err != nil {
		return fmt.Errorf("[MySQL] Error al guardar el ususario : %w", err)
	}
	return nil
}

func (mysql *MySQL) GetAll() ([]domain.User, error) {
	rows, err := mysql.DB.Query("SELECT * FROM users")

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User

	for rows.Next() {
		var user domain.User
		err := rows.Scan(&user.Id, &user.Name, &user.Lastname)
		if err != nil {
			return nil, err
		}
		users = append(users, user)

	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (mysql *MySQL) Delete(id int32) error {
	query := "DELETE FROM users WHERE id = ?"
	result, err := mysql.DB.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("No se encontro el usuario con ID:", id)
	}
	fmt.Println("Usuario Elimonado")
	return nil
}

func (mysql *MySQL) Update(id int32, name string, lastname string) error {
	query := "UPDATE users SET name = ?, lastname =? WHERE id = ? "

	result, err := mysql.DB.Exec(query, name, lastname, id)
	if err != nil {
		return fmt.Errorf("[MySQL] error alm verificar las filas afectadas:", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("[MySQL] Error al verificar filas afectadas: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("No se encontro el usuario")
	}

	fmt.Println("Usuario actualizado")
	return nil
}

func (mysql *MySQL) GetById(id int32) (domain.User, error) {
	var UserById domain.User

	query := "SELECT id, name, lastname FROM users WHERE id=?"
	row := mysql.DB.QueryRow(query, id)

	err := row.Scan(&UserById.Id, &UserById.Name, &UserById.Lastname)
	if err != nil {
		if err == sql.ErrNoRows {
			return UserById, fmt.Errorf("producto con id no encontrado", id)
		}
		return UserById, err
	}

	return UserById, nil
}
