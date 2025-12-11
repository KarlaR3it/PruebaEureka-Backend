package test

import (
	"context"
	"testing"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/KarlaR3it/PruebaEureka-Backend/internal/models"
	"github.com/KarlaR3it/PruebaEureka-Backend/internal/repository"
)

func setupTestDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
    db, mock, err := sqlmock.New()
    assert.NoError(t, err)

    // Espera la consulta interna de GORM
    mock.ExpectQuery("SELECT VERSION()").WillReturnRows(sqlmock.NewRows([]string{"VERSION()"}).AddRow("8.0"))

    gormDB, err := gorm.Open(mysql.New(mysql.Config{
        Conn: db,
    }), &gorm.Config{})
    assert.NoError(t, err)

    return gormDB, mock
}

func TestAreaRepository_Create(t *testing.T) {
	db, mock := setupTestDB(t)
	repo := repository.NewAreaRepository(db)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	area := &models.Area{Name: "Test"}
	err := repo.Create(context.Background(), area)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAreaRepository_GetAll(t *testing.T) {
	db, mock := setupTestDB(t)
	repo := repository.NewAreaRepository(db)

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Recursos Humanos").
		AddRow(2, "Ventas")

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	results, err := repo.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Len(t, results, 2)
}

func TestAreaRepository_GetAreaCounts(t *testing.T) {
	db, mock := setupTestDB(t)
	repo := repository.NewAreaRepository(db)

	rows := sqlmock.NewRows([]string{"name", "persons_quantity"}).
		AddRow("Recursos Humanos", 5).
		AddRow("Ventas", 3)

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	results, err := repo.GetAreaCounts(context.Background())
	assert.NoError(t, err)
	assert.Len(t, results, 2)
	assert.Equal(t, "Recursos Humanos", results[0].Name)
	assert.EqualValues(t, 5, results[0].PersonsQuantity)
}
