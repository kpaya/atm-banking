CREATE TABLE IF NOT EXISTS address (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  street varchar(255) NOT NULL,
  zip_code varchar(8) NOT NULL,
  city varchar(255) NOT NULL,
  state varchar(2) NOT NULL,
  country varchar(255) NOT NULL
);


CREATE TABLE IF NOT EXISTS card (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  account_id UUID NOT NULL REFERENCES account(id),
  number varchar(50) NOT NULL UNIQUE,
  cvv varchar(3) NOT NULL,
  expiration_month int NOT NULL,
  expiration_year int NOT NULL
);

CREATE TABLE IF NOT EXISTS customer (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name varchar(255) NOT NULL,
  email varchar(255) NOT NULL UNIQUE,
  phone varchar(20) NOT NULL,
  status varchar(30) NOT NULL,
  cpf varchar(11) NOT NULL UNIQUE,
  address_id UUID NOT NULL REFERENCES address(id)
);

CREATE TABLE IF NOT EXISTS account (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  customer_id UUID NOT NULL REFERENCES customer(id),
  account_number varchar(50) NOT NULL UNIQUE,
  total_balance float NOT NULL,
  avaliable_balance float NOT NULL
);

CREATE TABLE IF NOT EXISTS transactions(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  account_id UUID NOT NULL REFERENCES account(id),
  amount float NOT NULL,
  type varchar(30) NOT NULL,
  status varchar(30) NOT NULL,
  created_at timestamp NOT NULL
);