--
-- PostgreSQL database dump
--

-- Dumped from database version 10.0
-- Dumped by pg_dump version 10.0

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner: 
--

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';


SET search_path = public, pg_catalog;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: schema_migration; Type: TABLE; Schema: public; Owner: sujiths
--

CREATE TABLE schema_migration (
    version character varying(255) NOT NULL
);


ALTER TABLE schema_migration OWNER TO sujiths;

--
-- Name: users; Type: TABLE; Schema: public; Owner: sujiths
--

CREATE TABLE users (
    id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE users OWNER TO sujiths;

--
-- Name: whitelisted_clients; Type: TABLE; Schema: public; Owner: sujiths
--

CREATE TABLE whitelisted_clients (
    id uuid NOT NULL,
    client_id character varying(255) NOT NULL,
    pass_key character varying(255) NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL
);


ALTER TABLE whitelisted_clients OWNER TO sujiths;

--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: sujiths
--

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: whitelisted_clients whitelisted_clients_pkey; Type: CONSTRAINT; Schema: public; Owner: sujiths
--

ALTER TABLE ONLY whitelisted_clients
    ADD CONSTRAINT whitelisted_clients_pkey PRIMARY KEY (id);


--
-- Name: schema_migration_version_idx; Type: INDEX; Schema: public; Owner: sujiths
--

CREATE UNIQUE INDEX schema_migration_version_idx ON schema_migration USING btree (version);


--
-- PostgreSQL database dump complete
--

