-- CREATE DATABASE cryptoc WITH OWNER 'cryptoc' ENCODING 'UTF8';

CREATE TYPE crypto_currency AS ENUM ('ETH', 'BTC', 'XMR');
CREATE TYPE payment_status AS ENUM ('unconfirmed', 'weakConfirmation', 'confirmed', 'weakCancelled', 'cancelled');

CREATE TABLE users(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    email TEXT NOT NULL,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    password TEXT NOT NULL,
    registration_date TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ,
    updated_at TIMESTAMPTZ,
    is_admin BOOLEAN NOT NULL DEFAULT FALSE,
    CONSTRAINT user_name_cnst CHECK (char_length(first_name) <= 256 AND char_length(last_name) <= 256),
    CONSTRAINT user_email_cnst CHECK (char_length(email) <= 256)
);

CREATE UNIQUE INDEX users_email_idx ON users (email);

CREATE TABLE pay_per_view_events(
    "id" BIGSERIAL NOT NULL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "event_type" TEXT NOT NULL,
    "start" TIMESTAMPTZ,
    "end" TIMESTAMPTZ,
    "price_ETH" BIGINT,
    "price_BTC" BIGINT,
    "price_XMR" BIGINT,
    "eth_contract_addr" TEXT,
    "created_at" TIMESTAMPTZ,
    "updated_at" TIMESTAMPTZ
);

CREATE TABLE payments(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    user_id BIGINT,
    pay_per_view_event_id BIGINT NOT NULL,
    currency crypto_currency NOT NULL,
    currency_payment_id TEXT,
    amount BIGINT,
    wallet_address TEXT,
    status payment_status NOT NULL DEFAULT 'unconfirmed',
    block_hash TEXT,
    block_number_hex TEXT,
    tx_hash TEXT,
    tx_number_hex TEXT,
    cancelled_block_hash TEXT,
    cancelled_block_number_hex TEXT,
    cancelled_tx_hash TEXT,
    cancelled_tx_number_hex TEXT,
    cancelled_at TIMESTAMPTZ,
    CONSTRAINT payments_user_fk FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE INDEX payments_wallet_address_currency_idx ON payments (currency, wallet_address);

CREATE TABLE smart_contracts(
    "id" BIGSERIAL NOT NULL PRIMARY KEY,
    "pay_per_view_event_id" BIGINT NOT NULL,
    "address" TEXT NOT NULL,
    CONSTRAINT smart_contracts_event_fk FOREIGN KEY (pay_per_view_event_id) REFERENCES pay_per_view_events (id)
);

CREATE TABLE smart_contract_events(
    id BIGSERIAL NOT NULL PRIMARY KEY,
    smart_contract_id BIGINT NOT NULL,
    data TEXT,
    CONSTRAINT smart_contract_events_contract_fk FOREIGN KEY (smart_contract_id) REFERENCES smart_contracts (id)
);