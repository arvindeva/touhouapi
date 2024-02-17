--
-- PostgreSQL database dump
--

-- Dumped from database version 16.2 (Postgres.app)
-- Dumped by pg_dump version 16.2 (Postgres.app)

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

--
-- Name: citext; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;


--
-- Name: EXTENSION citext; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION citext IS 'data type for case-insensitive character strings';


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO admin;

--
-- Name: touhous; Type: TABLE; Schema: public; Owner: admin
--

CREATE TABLE public.touhous (
    id bigint NOT NULL,
    created_at timestamp(0) with time zone DEFAULT now() NOT NULL,
    name text NOT NULL,
    species text NOT NULL,
    abilities text[] NOT NULL,
    version integer DEFAULT 1 NOT NULL,
    CONSTRAINT abilities_length_check CHECK (((array_length(abilities, 1) >= 1) AND (array_length(abilities, 1) <= 10)))
);


ALTER TABLE public.touhous OWNER TO admin;

--
-- Name: touhous_id_seq; Type: SEQUENCE; Schema: public; Owner: admin
--

CREATE SEQUENCE public.touhous_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.touhous_id_seq OWNER TO admin;

--
-- Name: touhous_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: admin
--

ALTER SEQUENCE public.touhous_id_seq OWNED BY public.touhous.id;


--
-- Name: touhous id; Type: DEFAULT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.touhous ALTER COLUMN id SET DEFAULT nextval('public.touhous_id_seq'::regclass);


--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.schema_migrations (version, dirty) FROM stdin;
2	f
\.


--
-- Data for Name: touhous; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.touhous (id, created_at, name, species, abilities, version) FROM stdin;
1	2024-02-17 11:56:05+01	Reimu Hakurei	Human	{"Ability to Float","Aura Manipulation"}	1
2	2024-02-17 11:56:22+01	Marisa Kirisame	Human	{"Using Magic"}	1
3	2024-02-17 12:04:14+01	Sakuya Izayoi	Human	{"Space-time manipulation"}	1
\.


--
-- Name: touhous_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.touhous_id_seq', 4, false);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: touhous touhous_pkey; Type: CONSTRAINT; Schema: public; Owner: admin
--

ALTER TABLE ONLY public.touhous
    ADD CONSTRAINT touhous_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

