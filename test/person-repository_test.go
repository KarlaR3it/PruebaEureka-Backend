package test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/KarlaR3it/PruebaEureka-Backend/internal/models"
	"github.com/KarlaR3it/PruebaEureka-Backend/internal/repository"
)

func TestPersonRepository_Create(t *testing.T) {
	db, mock := setupTestDB(t)
	repo := repository.NewPersonRepository(db)

	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	person := &models.Person{
		Name:   "Juan Perez",
		Email:  "juan.perez@email.com",
		AreaID: 1,
	}
	err := repo.Create(context.Background(), person)
	assert.NoError(t, err)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestPersonRepository_GetAll(t *testing.T) {
	db, mock := setupTestDB(t)
	repo := repository.NewPersonRepository(db)

	rows := sqlmock.NewRows([]string{"id", "name", "email", "area_id"}).
		AddRow(1, "Juan Perez", "juan.perez@email.com", 1).
		AddRow(2, "Maria Lopez", "maria.lopez@email.com", 2)

	mock.ExpectQuery("SELECT").WillReturnRows(rows)

	// Expectativa para el Preload de Area
	areaRows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Recursos Humanos").
		AddRow(2, "Ventas")
	mock.ExpectQuery("SELECT \\* FROM `areas`").WillReturnRows(areaRows)

	results, err := repo.GetAll(context.Background())
	assert.NoError(t, err)
	assert.Len(t, results, 2)
	assert.Equal(t, "Juan Perez", results[0].Name)
	assert.Equal(t, "juan.perez@email.com", results[0].Email)
}

func TestPersonRepository_ExistsByEmail(t *testing.T) {
	db, mock := setupTestDB(t)
	repo := repository.NewPersonRepository(db)

	rows := sqlmock.NewRows([]string{"count"}).AddRow(1)
	mock.ExpectQuery("SELECT count").WillReturnRows(rows)

	exists, err := repo.ExistsByEmail(context.Background(), "juan.perez@email.com")
	assert.NoError(t, err)
	assert.True(t, exists)
	assert.NoError(t, mock.ExpectationsWereMet())
}
