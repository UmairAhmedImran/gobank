package storage

import (
	"fmt"
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/UmairAhmedImran/gobank/internal/account"
)


type Storage interface {
	CreateAccount(*account.Account) error
	DeleteAccount(int) error
	UpdateAccount(*account.Account) error
	GetAccounts() ([]*account.Account, error)
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


func (s *PostgresStore) DropTable(tableName string) error {
	query := fmt.Sprintf("DROP TABLE IF EXISTS %s", tableName)
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	fmt.Printf("Table %s has been dropped.\n", tableName)
	return nil
}

func (s *PostgresStore) createAccountTable() error {
	query := `CREATE TABLE IF NOT EXISTS account(
		id 		serial primary key,
		first_name	varchar(50), 
		last_name	varchar(50),
		number		serial,
		balance		serial,
		created_at	timestamp
	)`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(acc *account.Account) error{
	query :=`INSERT INTO account 
		(first_name, last_name, number, balance, created_at)	
		VALUES ($1, $2, $3, $4, $5)
		`
	resp, err := s.db.Exec(
	query,
	acc.FirstName,
	acc.LastName,
	acc.Number,
	acc.Balance,
	acc.CreatedAt,
	)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)

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

func (s *PostgresStore) GetAccounts() ([]*account.Account, error){
	rows, err := s.db.Query("SELECT * FROM account")
	if err != nil {
		return nil, err
	}

	accounts := []*account.Account{}
	for rows.Next() {
		account := new(account.Account)
		err := rows.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.Balance,
		&account.CreatedAt)
	
		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	
	}
	
	return accounts, nil
}


