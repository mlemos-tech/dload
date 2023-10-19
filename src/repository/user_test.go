package repository_test

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mikaellemos.com.br/dload/src/model"
	"mikaellemos.com.br/dload/src/repository"
	"regexp"
	"testing"
	"time"
)

var (
	id, _ = uuid.NewRandom()
)

type RepositorySuite struct {
	suite.Suite
	conn *sql.DB
	DB   *gorm.DB
	mock sqlmock.Sqlmock

	repo *repository.Repository
	user *model.User
}

func (rs *RepositorySuite) SetupSuite() {
	var err error

	rs.conn, rs.mock, err = sqlmock.New()
	assert.NoError(rs.T(), err)

	fmt.Println("Config setup")

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 rs.conn,
		PreferSimpleProtocol: true,
	})

	rs.DB, err = gorm.Open(dialector, &gorm.Config{})
	assert.NoError(rs.T(), err)

	rs.repo = repository.NewRepositoryDB(rs.DB)
	assert.IsType(rs.T(), &repository.Repository{}, rs.repo)

	rs.user = &model.User{
		ID:        id,
		Email:     "mikaellemos033@gmail.com",
		Birthday:  "1997-05-27",
		Name:      "Mikael Lemos",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func (rs *RepositorySuite) AfterTest(_, _ string) {
	assert.NoError(rs.T(), rs.mock.ExpectationsWereMet())
}

func (rs *RepositorySuite) TestInsert() {
	rs.mock.ExpectBegin()
	rs.mock.ExpectExec(regexp.QuoteMeta(`INSERT INTO "users" ("id","name","birthday","email","created_at","updated_at") VALUES ($1,$2,$3,$4,$5,$6)`)).
		WithArgs(
			id,
			rs.user.Name,
			rs.user.Birthday,
			rs.user.Email,
			rs.user.CreatedAt,
			rs.user.UpdatedAt).
		WillReturnResult(sqlmock.NewResult(1, 1))

	rs.mock.ExpectCommit()
	u, err := rs.repo.Insert(rs.user)

	assert.NoError(rs.T(), err)
	assert.Equal(rs.T(), rs.user.Email, u.Email)
}

func (rs *RepositorySuite) TestDelete() {
	rs.mock.ExpectBegin()
	rs.mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM "users" WHERE "users"."id" = $1`)).
		WithArgs(rs.user.ID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	rs.mock.ExpectCommit()
	err := rs.repo.Delete(rs.user.ID)
	assert.NoError(rs.T(), err)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(RepositorySuite))
}
