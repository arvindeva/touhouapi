--
-- PostgreSQL database dump
--

-- Dumped from database version 14.10 (Ubuntu 14.10-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.10 (Ubuntu 14.10-0ubuntu0.22.04.1)

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


ALTER TABLE public.touhous_id_seq OWNER TO admin;

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
3	f
\.


--
-- Data for Name: touhous; Type: TABLE DATA; Schema: public; Owner: admin
--

COPY public.touhous (id, created_at, name, species, abilities, version) FROM stdin;
2	2024-02-17 11:56:22+01	Marisa Kirisame	Human	{"Using Magic"}	1
4	2024-02-17 16:40:40+01	Sanae Kochiya	Human	{"Causing miracles to occur"}	4
3	2024-02-17 12:04:14+01	Sakuya Izayoi	Human	{"Space-time manipulation"}	4
5	2024-02-17 19:51:34+01	Reimu Hakurei	Human	{Flying,"Aura Manipulation","Youkai Extermination"}	1
6	2024-02-17 19:51:34+01	Marisa Kirisame	Human	{"Using Magic"}	1
7	2024-02-17 19:51:34+01	Sanae Kochiya	Human	{"Divine Miracles","Control Over Wind and Rain"}	1
8	2024-02-17 19:51:34+01	Yukari Yakumo	Youkai	{"Manipulation of Boundaries"}	1
9	2024-02-17 19:51:34+01	Sakuya Izayoi	Youkai	{"Time Manipulation","Precision and Speed"}	1
10	2024-02-17 19:51:34+01	Remilia Scarlet	Vampire	{"Control Over Fate","Superhuman Strength"}	1
11	2024-02-17 19:51:34+01	Flandre Scarlet	Vampire	{"Destructive Power",Immortality}	1
12	2024-02-17 19:51:34+01	Patchouli Knowledge	Youkai	{"Master of Magic","Library of Forbidden Tomes"}	1
13	2024-02-17 19:51:34+01	Hong Meiling	Youkai	{"Martial Arts","Chi Manipulation"}	1
14	2024-02-17 19:51:34+01	Koakuma	Youkai	{"Knowledgeable Assistant"}	1
15	2024-02-17 19:51:34+01	Iku Nagae	Youkai	{"Control Over the Seas","Weather Manipulation"}	1
16	2024-02-17 19:51:34+01	Tenshi Hinanawi	Celestial	{"Control Over Earthquakes"}	1
17	2024-02-17 19:51:34+01	Reisen Udongein Inaba	Youkai	{"Illusion Manipulation","Lunatic Red Eyes"}	1
18	2024-02-17 19:51:34+01	Kaguya Houraisan	Youkai	{"Eternal Youth","Manipulation of Eternity"}	1
19	2024-02-17 19:51:34+01	Mokou Fujiwara	Human	{Immortality,"Manipulation of Fire"}	1
20	2024-02-17 19:51:34+01	Nitori Kawashiro	Kappa	{"Technological Expertise","Water Manipulation"}	1
21	2024-02-17 19:51:34+01	Aya Shameimaru	Tengu	{"Control Over Wind","Reporting Skills"}	1
22	2024-02-17 19:51:34+01	Kanako Yasaka	God	{"Control Over Mountains",Harvest}	1
23	2024-02-17 19:51:34+01	Suwako Moriya	God	{"Frog Deity","Earthquake Control"}	1
24	2024-02-17 19:51:34+01	Byakuren Hijiri	Youkai	{"Sealing Techniques","Buddhist Magic"}	1
25	2024-02-17 19:51:34+01	Miko Toyosatomimi	Human	{"Divine Spirit Possession"}	1
26	2024-02-17 19:51:34+01	Mamizou Futatsuiwa	Youkai	{Transformation,Disguise}	1
27	2024-02-17 19:51:34+01	Kokoro Hata	Youkai	{"Emotion Manipulation"}	1
28	2024-02-17 19:51:34+01	Mononobe no Futo	Human	{Onbashira,"Scroll of Origin"}	1
29	2024-02-17 19:51:34+01	Toyosatomimi no Miko	Human	{"Divine Spirit Possession"}	1
30	2024-02-17 19:51:34+01	Minamitsu Murasa	Ghost	{"Control Over Shipwrecks"}	1
31	2024-02-17 19:51:34+01	Shinmyoumaru Sukuna	Youkai	{"Size Manipulation","Mallet of Luck"}	1
32	2024-02-17 19:51:34+01	Raiko Horikawa	Youkai	{Drumming,"Rhythm Manipulation"}	1
33	2024-02-17 19:51:34+01	Kagerou Imaizumi	Werewolf	{"Lunar Manipulation",Shapeshifting}	1
34	2024-02-17 19:51:34+01	Wakasagihime	Mermaid	{"Aquatic Abilities"}	1
35	2024-02-17 19:51:34+01	Seiran	Moon Rabbit	{"Lunatic Bullet","Inaba Boxing"}	1
36	2024-02-17 19:51:34+01	Reisen II	Moon Rabbit	{"Illusion Manipulation","Lunatic Red Eyes"}	1
37	2024-02-17 19:51:34+01	Ringo	Moon Rabbit	{"Space Manipulation","UFO Summoning"}	1
38	2024-02-17 19:51:34+01	Doremy Sweet	Dream Youkai	{"Dream Manipulation"}	1
39	2024-02-17 19:51:34+01	Toki	Dream Youkai	{"Time Manipulation","Dream Manipulation"}	1
40	2024-02-17 19:51:34+01	Yumemi Okazaki	Human	{"Space Manipulation","Scientific Knowledge"}	1
41	2024-02-17 19:51:34+01	Rikako Asakura	Human	{"Scientific Knowledge",Alchemy}	1
42	2024-02-17 19:51:34+01	Chiyuri Kitashirakawa	Human	{"Time Travel","Space Manipulation"}	1
43	2024-02-17 19:51:34+01	Yumeko	Human	{"Space Manipulation","Dimensional Travel"}	1
44	2024-02-17 19:51:34+01	Yumemi	Human	{"Space Manipulation","Scientific Knowledge"}	1
45	2024-02-17 19:51:34+01	Irusu	Human	{Teleportation,Invisibility}	1
1	2024-02-17 11:56:05+01	Reimu Hakurei	Human	{"Ability to Float","Aura Manipulation"}	5
47	2024-02-18 09:20:08+01	Shinmy	Inchling	{"Be small","Get bigger"}	1
\.


--
-- Name: touhous_id_seq; Type: SEQUENCE SET; Schema: public; Owner: admin
--

SELECT pg_catalog.setval('public.touhous_id_seq', 47, true);


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
-- Name: touhous_name_idx; Type: INDEX; Schema: public; Owner: admin
--

CREATE INDEX touhous_name_idx ON public.touhous USING gin (to_tsvector('simple'::regconfig, name));


--
-- PostgreSQL database dump complete
--

