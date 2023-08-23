-- name: CreateAddress :one
INSERT INTO address (
  street, zip_code, city, state, country
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id;

-- name: UpdateAddress :one
UPDATE address
SET street = $1, zip_code = $2, city = $3, state = $4, country = $5
WHERE id = $6
RETURNING id;

-- name: GetAddress :one
SELECT * FROM address WHERE id = $1 LIMIT 1;

-- name: ListAddresses :many
SELECT * FROM address;

-- name: CreateCustomer :one
INSERT INTO customer (
  name, email, phone, status, cpf, address_id
) VALUES (
  $1, $2, $3, $4, $5, $6
)
RETURNING id;

-- name: UpdateCustomer :one
UPDATE customer
SET name = $1, email = $2, phone = $3, status = $4, cpf = $5, address_id = $6
WHERE id = $7
RETURNING id;

-- name: GetCustomer :one
SELECT * FROM customer
WHERE id = $1 LIMIT 1;

-- name: ListCustomers :many
SELECT * FROM customer;

-- name: CreateAccount :one
INSERT INTO account (
  customer_id, account_number, total_balance, avaliable_balance
) VALUES (
  $1, $2, $3, $4
)
RETURNING id;

-- name: UpdateAccount :one
UPDATE account
SET customer_id = $1, account_number = $2, total_balance = $3, avaliable_balance = $4
WHERE id = $5
RETURNING id;

-- name: GetAccount :one
SELECT * FROM account
WHERE id = $1 LIMIT 1;

-- name: ListAccounts :many
SELECT * FROM account;

-- name: CreateCard :one
INSERT INTO card (
  account_id, number, cvv, expiration_month, expiration_year
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING id;

-- name: UpdateCard :one
UPDATE card
SET account_id = $1, number = $2, cvv = $3, expiration_month = $4, expiration_year = $5
WHERE id = $6
RETURNING id;

-- name: GetCard :one
SELECT * FROM card
WHERE id = $1 LIMIT 1;

-- name: ListCards :many
SELECT * FROM card;

-- name: CreateTransactions :one
INSERT INTO transactions (
  account_id, amount, type, status
) VALUES (
  $1, $2, $3, $4
)
RETURNING id;

-- name: UpdateTransactions :one
UPDATE transactions
SET account_id = $1, amount = $2, type = $3, status = $4
WHERE id = $5
RETURNING id;

-- name: GetTransactions :one
SELECT * FROM transactions
WHERE id = $1 LIMIT 1;

-- name: ListTransactionss :many
SELECT * FROM transactions;