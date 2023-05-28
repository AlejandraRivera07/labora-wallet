CREATE DATABASE labora-wallet
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    CONNECTION LIMIT = -1
    IS_TEMPLATE = False;

CREATE TABLE IF NOT EXISTS public.wallet
(
    id integer SERIAL PRIMARY KEY,
    customer_id integer NOT NULL,
    country_id varchar(10) NOT NULL,
    create_date date NOT NULL,
    CONSTRAINT wallet_pkey PRIMARY KEY (id)
)


ALTER TABLE IF EXISTS public.wallet
    OWNER to postgres;

CREATE TABLE IF NOT EXISTS public.logger
(
    id integer SERIAL PRIMARY KEY,
    create_date date NOT NULL,
    status_creation varchar(10) NOT NULL,
    customer_id integer NOT NULL,
    country_id varchar(10) NOT NULL,
    codigo varchar(10) NOT NULL,
    CONSTRAINT logger_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.log
    OWNER to postgres;


























CREATE TABLE solicitud (
	id_persona int PRIMARY KEY,
	dni_solicitud VARCHAR(250),
	fecha_solicitud DATE,
	pais varchar(50),
	estado varchar(50),
	codigo int
)

CREATE TABLE wallet (
	id int PRIMARY KEY REFERENCES solicitud(id_persona),
	dni VARCHAR(250),
	pais_id VARCHAR(50),
	creacion DATE
)