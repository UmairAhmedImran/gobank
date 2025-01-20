package storage

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/UmairAhmedImran/gobank/internal/account"
)


type Storage interface {
	CreateAccount(*account.Account) error
	DeleteAccount(int) error
	UpdateAccount(*account.Account) error
	GetAccountById(int) (*account.Account, error)
}

type PostgresStore struct {
	db 	*sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {

	connStr := "user=postgres dbname=postgres password=mygobank port=5435 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil

}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account(
		id 		serial primary key,
		first_name	varchar(50), 
		last_name	varchar(50),
		number		serial,
		balnace		serial,
		created_at	timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(*account.Account) error{
	return nil
}
func (s *PostgresStore) DeleteAccount(id int) error{
	return nil
}
func (s *PostgresStore) UpdateAccount(*account.Account) error{
	return nil
}
func (s *PostgresStore) GetAccountById(id int) (*account.Account, error){
	return nil, nil
}



