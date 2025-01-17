package CRUD

import (
	"context"

	_ "github.com/lib/pq"
)

const CreateAccountSQL = `
INSERT INTO "Account" 
("Owner","Balance","Currency") 
VALUES($1,$2,$3) 
RETURNING "ID","Owner","Balance","Currency","Createdat"`

type CreateAccountParam struct {
	Owner    string
	Balance  int64
	Currency string
}

const GetAccountSQL = `
SELECT 
"ID","Owner","Balance","Currency","Createdat" 
FROM "Account" WHERE "ID" = $1`

type GetAccountParam struct {
	ID int64
}

const ListAccountSQL = `
SELECT "ID","Owner","Balance","Currency","Createdat" 
FROM "Account"
ORDER BY "ID"
LIMIT $1
OFFSET $2`

type ListAccountParam struct {
	Offset int32
	Limit  int32
}

const UpdateAccountSQL = `
UPDATE "Account"
SET "Balance" = $1
WHERE "ID" = $2`

type UpdateAccountParam struct {
	Balance int64
	ID      int64
}

const DeleteAccountSQL = `
DELETE FROM "Account"
WHERE "ID" = $1`

type DeleteAccountParam struct {
	ID int64
}

const CountOfAccountSQL = `
SELECT
COUNT(*)
FROM "Account"
`

const GetAccountForUpdateSQL = `
SELECT 
"ID","Owner","Balance","Currency","Createdat" 
FROM "Account" WHERE "ID" = $1
FOR NO KEY UPDATE`

type GetAccountForUpdateParam struct {
	ID int64
}

func (q *Queries) GetAccountForUpdate(ctx context.Context, arg GetAccountForUpdateParam) (Account, error) {
	tmp := q.db.QueryRowContext(ctx, GetAccountForUpdateSQL, arg.ID)
	var A Account
	err := tmp.Scan(
		&A.ID,
		&A.Owner,
		&A.Balance,
		&A.Currency,
		&A.Createdat,
	)

	return A, err
}

func (q *Queries) CountOfAccount(ctx context.Context) int64 {
	row := q.db.QueryRowContext(ctx, CountOfAccountSQL)

	var count int64

	err := row.Scan(&count)

	if err != nil {
		return count
	}
	return count

}
func (q *Queries) CreateAccount(ctx context.Context, arg CreateAccountParam) (Account, error) {
	tmp := q.db.QueryRowContext(ctx, CreateAccountSQL, arg.Owner, arg.Balance, arg.Currency)
	var A Account
	err := tmp.Scan(
		&A.ID,
		&A.Owner,
		&A.Balance,
		&A.Currency,
		&A.Createdat,
	)

	return A, err
}

func (q *Queries) GetAccount(ctx context.Context, arg GetAccountParam) (Account, error) {
	tmp := q.db.QueryRowContext(ctx, GetAccountSQL, arg.ID)
	var A Account
	err := tmp.Scan(
		&A.ID,
		&A.Owner,
		&A.Balance,
		&A.Currency,
		&A.Createdat,
	)

	return A, err
}
func (q *Queries) ListAccount(ctx context.Context, arg ListAccountParam) ([]Account, error) {
	tmps, err := q.db.QueryContext(ctx, ListAccountSQL, arg.Limit, arg.Offset)

	if err != nil {
		return nil, err
	}

	defer tmps.Close()

	A_arr := []Account{}

	for tmps.Next() {
		var A Account
		if err := tmps.Scan(
			&A.ID,
			&A.Owner,
			&A.Balance,
			&A.Currency,
			&A.Createdat,
		); err != nil {
			return nil, err
		}
		A_arr = append(A_arr, A)
	}

	//Avoid next enumeration.
	if err := tmps.Close(); err != nil {
		return nil, err
	}
	if err := tmps.Err(); err != nil {
		return nil, err
	}
	return A_arr, nil
}

func (q *Queries) UpdateAccount(ctx context.Context, arg UpdateAccountParam) error {
	_, err := q.db.ExecContext(ctx, UpdateAccountSQL, arg.Balance, arg.ID)
	return err
}

func (q *Queries) DeleteAccount(ctx context.Context, arg DeleteAccountParam) error {
	_, err := q.db.ExecContext(ctx, DeleteAccountSQL, arg.ID)
	return err
}
