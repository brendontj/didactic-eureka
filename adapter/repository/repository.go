package repository

import (
	"cloud.google.com/go/civil"
	"context"
	"database/sql"
	"fmt"
	"github.com/brendontj/didactic-eureka/core/entity"
	"github.com/brendontj/didactic-eureka/infrastructure/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

type Adapter struct {
	*postgres.DB
}

func NewAdapter(db *postgres.DB) *Adapter {
	return &Adapter{DB: db}
}

func (a *Adapter) FindAll(ctx context.Context) ([]entity.Customer, error) {
	var customers []entity.Customer
	query := `
		SELECT 
			a.id, a.version, a.name, a.email, a.phone, a.birth_date, a.document, b.street, b.number, b.complement,
			b.neighborhood, b.city, b.state, b.country, b.zip_code, a.created_at, a.updated_at 
		FROM customer.customer as a
		INNER JOIN customer.address as b 
			ON a.id = b.customer_id`

	rows, err := a.DB.Conn.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		c, err := a.scanRowInCustomer(rows)
		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return customers, nil
}

func (a *Adapter) FindById(ctx context.Context, id uuid.UUID) (entity.Customer, error) {
	var c entity.Customer
	query := `
		SELECT 
			a.id, a.version, a.name, a.email, a.phone, a.birth_date, a.document, b.street, b.number, b.complement,
			b.neighborhood, b.city, b.state, b.country, b.zip_code, a.created_at, a.updated_at 
		FROM customer.customer as a
		INNER JOIN customer.address as b 
			ON a.id = b.customer_id
		WHERE a.id = $1`

	row := a.DB.Conn.QueryRow(ctx, query, id.String())
	c, err := a.scanRowInCustomer(row)
	if err != nil {
		return entity.Customer{}, err
	}

	return c, nil
}

func (a *Adapter) Save(ctx context.Context, customer entity.Customer) (err error) {
	customerQuery := `INSERT INTO customer.customer (id, version, name, email, phone, birth_date, document, created_at, updated_at) 
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	addressQuery := `INSERT INTO customer.address (id, customer_id, street, number, complement, neighborhood, city, state, country, zip_code)
			VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	tx, err := a.DB.Conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	defer func() {
		_ = a.CommitOrRollback(ctx, tx, err)
	}()

	_, err = tx.Exec(
		ctx,
		customerQuery,
		customer.ID.String(),
		customer.Version.String(),
		customer.Name,
		customer.Email,
		customer.Phone,
		customer.BirthDate.String(),
		customer.Document,
		customer.CreatedAt,
		customer.UpdatedAt)
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		ctx,
		addressQuery,
		uuid.New().String(),
		customer.ID.String(),
		customer.Address.Street,
		customer.Address.Number,
		customer.Address.Complement,
		customer.Address.Neighborhood,
		customer.Address.City,
		customer.Address.State,
		customer.Address.Country,
		customer.Address.ZipCode)
	if err != nil {
		return err
	}

	return nil
}

func (a *Adapter) Update(ctx context.Context, customer entity.Customer, currentCustomerVersion uuid.UUID) (err error) {
	customerQuery := `UPDATE customer.customer
	SET version = $1,
		name = $2,
		email = $3,
		phone = $4,
		birth_date = $5,
		document = $6,
		created_at = $7,
		updated_at = $8
	WHERE id = $9 and version = $10`

	addressQuery := `UPDATE customer.address 
	SET street = $1,
		number = $2,
		complement = $3,
		neighborhood = $4,
		city = $5,
		state = $6,
		country = $7,
		zip_code = $8
	WHERE customer_id = $9`

	tx, err := a.DB.Conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	defer func() {
		_ = a.CommitOrRollback(ctx, tx, err)
	}()

	_, err = tx.Exec(
		ctx,
		customerQuery,
		customer.Version.String(),
		customer.Name,
		customer.Email,
		customer.Phone,
		customer.BirthDate.String(),
		customer.Document,
		customer.CreatedAt,
		customer.UpdatedAt,
		customer.ID.String(),
		currentCustomerVersion.String())
	if err != nil {
		return err
	}

	_, err = tx.Exec(
		ctx,
		addressQuery,
		customer.Address.Street,
		customer.Address.Number,
		customer.Address.Complement,
		customer.Address.Neighborhood,
		customer.Address.City,
		customer.Address.State,
		customer.Address.Country,
		customer.Address.ZipCode,
		customer.ID.String())
	if err != nil {
		return err
	}

	return nil
}

func (a *Adapter) Delete(ctx context.Context, id uuid.UUID) (err error) {
	customerQuery := `DELETE FROM customer.customer WHERE id = $1`
	addressQuery := `DELETE FROM customer.address WHERE customer_id = $1`

	tx, err := a.DB.Conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}

	defer func() {
		_ = a.CommitOrRollback(ctx, tx, err)
	}()

	_, err = tx.Exec(ctx, customerQuery, id.String())
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, addressQuery, id.String())
	if err != nil {
		return err
	}

	return nil
}

func (a *Adapter) CommitOrRollback(ctx context.Context, tx pgx.Tx, err error) error {
	if err != nil {
		if rollbackErr := tx.Rollback(ctx); rollbackErr != nil {
			return rollbackErr
		}
		return err
	}
	return tx.Commit(ctx)
}

func (a *Adapter) scanRowInCustomer(row pgx.Row) (c entity.Customer, err error) {
	var (
		customerId      sql.NullString
		customerVersion sql.NullString
		birthDate       pgtype.Date
	)
	if err := row.Scan(&customerId, &customerVersion, &c.Name, &c.Email, &c.Phone, &birthDate, &c.Document,
		&c.Address.Street, &c.Address.Number, &c.Address.Complement, &c.Address.Neighborhood, &c.Address.City,
		&c.Address.State, &c.Address.Country, &c.Address.ZipCode, &c.CreatedAt, &c.UpdatedAt); err != nil {
		return entity.Customer{}, err
	}
	if !customerId.Valid {
		return entity.Customer{}, fmt.Errorf("customer id is null")
	}

	c.ID, err = uuid.Parse(customerId.String)
	if err != nil {
		return entity.Customer{}, err
	}
	if !customerVersion.Valid {
		return entity.Customer{}, fmt.Errorf("customer version is null")
	}

	c.Version, err = uuid.Parse(customerVersion.String)
	if err != nil {
		return entity.Customer{}, err
	}

	c.BirthDate, err = civil.ParseDate(birthDate.Time.Format("2006-01-02"))
	if err != nil {
		return entity.Customer{}, err
	}

	return c, nil
}
