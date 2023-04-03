--
-- PostgreSQL database dump
--

-- Dumped from database version 11.18 (Ubuntu 11.18-1.pgdg20.04+1)
-- Dumped by pg_dump version 15.1

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
-- Name: heroku_ext; Type: SCHEMA; Schema: -; Owner: vxvtteck
--

CREATE SCHEMA heroku_ext;


ALTER SCHEMA heroku_ext OWNER TO vxvtteck;

--
-- Name: main; Type: SCHEMA; Schema: -; Owner: vxvtteck
--

CREATE SCHEMA main;


ALTER SCHEMA main OWNER TO vxvtteck;

--
-- Name: public; Type: SCHEMA; Schema: -; Owner: postgres
--

-- *not* creating schema, since initdb creates it


ALTER SCHEMA public OWNER TO postgres;

--
-- Name: btree_gin; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS btree_gin WITH SCHEMA public;


--
-- Name: EXTENSION btree_gin; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION btree_gin IS 'support for indexing common datatypes in GIN';


--
-- Name: btree_gist; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS btree_gist WITH SCHEMA public;


--
-- Name: EXTENSION btree_gist; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION btree_gist IS 'support for indexing common datatypes in GiST';


--
-- Name: citext; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS citext WITH SCHEMA public;


--
-- Name: EXTENSION citext; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION citext IS 'data type for case-insensitive character strings';


--
-- Name: cube; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS cube WITH SCHEMA public;


--
-- Name: EXTENSION cube; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION cube IS 'data type for multidimensional cubes';


--
-- Name: dblink; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS dblink WITH SCHEMA public;


--
-- Name: EXTENSION dblink; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION dblink IS 'connect to other PostgreSQL databases from within a database';


--
-- Name: dict_int; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS dict_int WITH SCHEMA public;


--
-- Name: EXTENSION dict_int; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION dict_int IS 'text search dictionary template for integers';


--
-- Name: dict_xsyn; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS dict_xsyn WITH SCHEMA public;


--
-- Name: EXTENSION dict_xsyn; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION dict_xsyn IS 'text search dictionary template for extended synonym processing';


--
-- Name: earthdistance; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS earthdistance WITH SCHEMA public;


--
-- Name: EXTENSION earthdistance; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION earthdistance IS 'calculate great-circle distances on the surface of the Earth';


--
-- Name: fuzzystrmatch; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS fuzzystrmatch WITH SCHEMA public;


--
-- Name: EXTENSION fuzzystrmatch; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION fuzzystrmatch IS 'determine similarities and distance between strings';


--
-- Name: hstore; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS hstore WITH SCHEMA public;


--
-- Name: EXTENSION hstore; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION hstore IS 'data type for storing sets of (key, value) pairs';


--
-- Name: intarray; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS intarray WITH SCHEMA public;


--
-- Name: EXTENSION intarray; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION intarray IS 'functions, operators, and index support for 1-D arrays of integers';


--
-- Name: ltree; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS ltree WITH SCHEMA public;


--
-- Name: EXTENSION ltree; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION ltree IS 'data type for hierarchical tree-like structures';


--
-- Name: pg_stat_statements; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pg_stat_statements WITH SCHEMA public;


--
-- Name: EXTENSION pg_stat_statements; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pg_stat_statements IS 'track execution statistics of all SQL statements executed';


--
-- Name: pg_trgm; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pg_trgm WITH SCHEMA public;


--
-- Name: EXTENSION pg_trgm; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pg_trgm IS 'text similarity measurement and index searching based on trigrams';


--
-- Name: pgcrypto; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pgcrypto WITH SCHEMA public;


--
-- Name: EXTENSION pgcrypto; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pgcrypto IS 'cryptographic functions';


--
-- Name: pgrowlocks; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pgrowlocks WITH SCHEMA public;


--
-- Name: EXTENSION pgrowlocks; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pgrowlocks IS 'show row-level locking information';


--
-- Name: pgstattuple; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS pgstattuple WITH SCHEMA public;


--
-- Name: EXTENSION pgstattuple; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION pgstattuple IS 'show tuple-level statistics';


--
-- Name: tablefunc; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS tablefunc WITH SCHEMA public;


--
-- Name: EXTENSION tablefunc; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION tablefunc IS 'functions that manipulate whole tables, including crosstab';


--
-- Name: unaccent; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS unaccent WITH SCHEMA public;


--
-- Name: EXTENSION unaccent; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION unaccent IS 'text search dictionary that removes accents';


--
-- Name: uuid-ossp; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA public;


--
-- Name: EXTENSION "uuid-ossp"; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION "uuid-ossp" IS 'generate universally unique identifiers (UUIDs)';


--
-- Name: xml2; Type: EXTENSION; Schema: -; Owner: -
--

CREATE EXTENSION IF NOT EXISTS xml2 WITH SCHEMA public;


--
-- Name: EXTENSION xml2; Type: COMMENT; Schema: -; Owner: 
--

COMMENT ON EXTENSION xml2 IS 'XPath querying and XSLT';


--
-- Name: notify_bkng_changes_func(); Type: FUNCTION; Schema: public; Owner: vxvtteck
--

CREATE FUNCTION public.notify_bkng_changes_func() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
			BEGIN
			
			IF TG_OP='INSERT' or OLD.dte != NEW.dte or NEW.canceled_at is not null THEN
			PERFORM pg_notify('bkng_changes_chan', 'changed');
			END IF;
			
			RETURN NEW;
			END;
			$$;


ALTER FUNCTION public.notify_bkng_changes_func() OWNER TO vxvtteck;

SET default_tablespace = '';

--
-- Name: bkng; Type: TABLE; Schema: main; Owner: vxvtteck
--

CREATE TABLE main.bkng (
    bkng_id integer NOT NULL,
    status text DEFAULT 'booked'::text NOT NULL,
    dte timestamp without time zone NOT NULL,
    gme_id integer NOT NULL,
    nme text NOT NULL,
    mob_num text NOT NULL,
    email_addr text NOT NULL,
    tmstmp timestamp without time zone DEFAULT CURRENT_TIMESTAMP NOT NULL,
    canceled_at timestamp without time zone,
    notes text,
    bked_by integer,
    lst_edted_by integer
);


ALTER TABLE main.bkng OWNER TO vxvtteck;

--
-- Name: bkng_bkng_id_seq; Type: SEQUENCE; Schema: main; Owner: vxvtteck
--

CREATE SEQUENCE main.bkng_bkng_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE main.bkng_bkng_id_seq OWNER TO vxvtteck;

--
-- Name: bkng_bkng_id_seq; Type: SEQUENCE OWNED BY; Schema: main; Owner: vxvtteck
--

ALTER SEQUENCE main.bkng_bkng_id_seq OWNED BY main.bkng.bkng_id;


--
-- Name: gme; Type: TABLE; Schema: main; Owner: vxvtteck
--

CREATE TABLE main.gme (
    gme_id integer NOT NULL,
    status text DEFAULT 'inactive'::text NOT NULL,
    img_url text NOT NULL,
    map_url text NOT NULL,
    plrs text NOT NULL,
    age_rng text NOT NULL,
    nme text NOT NULL,
    descr text NOT NULL,
    addr text NOT NULL,
    dur integer NOT NULL
);


ALTER TABLE main.gme OWNER TO vxvtteck;

--
-- Name: gme_gme_id_seq; Type: SEQUENCE; Schema: main; Owner: vxvtteck
--

CREATE SEQUENCE main.gme_gme_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE main.gme_gme_id_seq OWNER TO vxvtteck;

--
-- Name: gme_gme_id_seq; Type: SEQUENCE OWNED BY; Schema: main; Owner: vxvtteck
--

ALTER SEQUENCE main.gme_gme_id_seq OWNED BY main.gme.gme_id;


--
-- Name: gme_lcl; Type: TABLE; Schema: main; Owner: vxvtteck
--

CREATE TABLE main.gme_lcl (
    gme_lcl_id integer NOT NULL,
    gme_id integer NOT NULL,
    lcl text NOT NULL,
    nme text NOT NULL,
    descr text NOT NULL,
    addr text NOT NULL
);


ALTER TABLE main.gme_lcl OWNER TO vxvtteck;

--
-- Name: gme_lcl_gme_lcl_id_seq; Type: SEQUENCE; Schema: main; Owner: vxvtteck
--

CREATE SEQUENCE main.gme_lcl_gme_lcl_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE main.gme_lcl_gme_lcl_id_seq OWNER TO vxvtteck;

--
-- Name: gme_lcl_gme_lcl_id_seq; Type: SEQUENCE OWNED BY; Schema: main; Owner: vxvtteck
--

ALTER SEQUENCE main.gme_lcl_gme_lcl_id_seq OWNED BY main.gme_lcl.gme_lcl_id;


--
-- Name: schdl; Type: TABLE; Schema: main; Owner: vxvtteck
--

CREATE TABLE main.schdl (
    schdl_id integer NOT NULL,
    gme_id integer NOT NULL,
    schedule jsonb DEFAULT '{}'::jsonb NOT NULL,
    active_by date
);


ALTER TABLE main.schdl OWNER TO vxvtteck;

--
-- Name: schdl_schdl_id_seq; Type: SEQUENCE; Schema: main; Owner: vxvtteck
--

CREATE SEQUENCE main.schdl_schdl_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE main.schdl_schdl_id_seq OWNER TO vxvtteck;

--
-- Name: schdl_schdl_id_seq; Type: SEQUENCE OWNED BY; Schema: main; Owner: vxvtteck
--

ALTER SEQUENCE main.schdl_schdl_id_seq OWNED BY main.schdl.schdl_id;


--
-- Name: usr; Type: TABLE; Schema: main; Owner: vxvtteck
--

CREATE TABLE main.usr (
    usr_id integer NOT NULL,
    usrnme text NOT NULL,
    psswrd text NOT NULL,
    salt text NOT NULL
);


ALTER TABLE main.usr OWNER TO vxvtteck;

--
-- Name: usr_usr_id_seq; Type: SEQUENCE; Schema: main; Owner: vxvtteck
--

CREATE SEQUENCE main.usr_usr_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE main.usr_usr_id_seq OWNER TO vxvtteck;

--
-- Name: usr_usr_id_seq; Type: SEQUENCE OWNED BY; Schema: main; Owner: vxvtteck
--

ALTER SEQUENCE main.usr_usr_id_seq OWNED BY main.usr.usr_id;


--
-- Name: bkng bkng_id; Type: DEFAULT; Schema: main; Owner: vxvtteck
--

ALTER TABLE ONLY main.bkng ALTER COLUMN bkng_id SET DEFAULT nextval('main.bkng_bkng_id_seq'::regclass);


--
-- Name: gme gme_id; Type: DEFAULT; Schema: main; Owner: vxvtteck
--

ALTER TABLE ONLY main.gme ALTER COLUMN gme_id SET DEFAULT nextval('main.gme_gme_id_seq'::regclass);


--
-- Name: gme_lcl gme_lcl_id; Type: DEFAULT; Schema: main; Owner: vxvtteck
--

ALTER TABLE ONLY main.gme_lcl ALTER COLUMN gme_lcl_id SET DEFAULT nextval('main.gme_lcl_gme_lcl_id_seq'::regclass);


--
-- Name: schdl schdl_id; Type: DEFAULT; Schema: main; Owner: vxvtteck
--

ALTER TABLE ONLY main.schdl ALTER COLUMN schdl_id SET DEFAULT nextval('main.schdl_schdl_id_seq'::regclass);


--
-- Name: usr usr_id; Type: DEFAULT; Schema: main; Owner: vxvtteck
--

ALTER TABLE ONLY main.usr ALTER COLUMN usr_id SET DEFAULT nextval('main.usr_usr_id_seq'::regclass);


--
-- Data for Name: bkng; Type: TABLE DATA; Schema: main; Owner: vxvtteck
--

COPY main.bkng (bkng_id, status, dte, gme_id, nme, mob_num, email_addr, tmstmp, canceled_at, notes, bked_by, lst_edted_by) FROM stdin;
48	canceled	2022-09-29 19:00:00	1	ioulios	+306975645865	ioulios@email.com	2022-09-23 11:27:02.220301	2022-09-26 18:58:38.201466	\N	\N	\N
52	booked	2022-09-27 14:00:00	2	Ioulios Tsiko	+306975645865	ioulios.tsiko@gmail.com	2022-09-27 09:33:54.242757	\N	\N	-1	\N
53	booked	2022-09-29 19:15:00	2	Ioulios Tsiko	+306975645865	ioulios.tsiko@gmail.com	2022-09-27 09:45:22.544097	\N	\N	-1	\N
54	booked	2022-09-29 14:00:00	2	Ioulios Tsiko	+306975645865	ioulios.tsiko@gmail.com	2022-09-27 09:47:15.539244	\N	\N	-1	\N
55	booked	2022-09-27 19:15:00	2	Ioulios Tsiko	+306975645865	ioulios.tsiko@gmail.com	2022-09-27 09:56:10.343438	\N	\N	-1	\N
56	booked	2022-09-30 14:00:00	2	Ioulios Tsiko	+306975645865	ioulios.tsiko@gmail.com	2022-09-27 09:56:30.888465	\N	\N	-1	\N
28	booked	2022-07-28 14:00:00	1	ioulios	+306975645865	ioulis@email.com	2022-04-05 10:24:21.862661	\N	\N	\N	\N
31	booked	2022-06-28 14:00:00	1	ioulios	+306975645865	ioulis@email.com	2022-04-07 11:36:43.363868	\N	\N	\N	\N
57	booked	2022-09-30 19:00:00	1	Ioulios Tsiko	+306975645865	ioulios.tsiko@gmail.com	2022-09-27 09:58:30.949979	\N	\N	-1	\N
58	booked	2022-09-27 14:00:00	1	Ioulios Tsiko	+306975645865	ioulios.tsiko@gmail.com	2022-09-27 09:59:51.260581	\N	tes τεστ 123 '	-1	\N
59	booked	2022-09-30 14:00:00	1	Ioulios Tsiko	+306975645865	ioulios.tsiko@gmail.com	2022-09-27 10:46:05.609366	\N	test	-1	\N
20	booked	2022-06-30 14:00:00	1	ioulios	+306975645865	ioulis@email.com	2022-04-04 14:57:55.136072	\N	\N	\N	\N
22	booked	2022-06-30 19:00:00	2	ioulios	+306975645865	ioulis@email.com	2022-04-04 14:58:40.813914	\N	\N	\N	\N
50	canceled	2022-09-26 14:00:00	1	iouliosa	+306975645865	ioulis@email.com	2022-09-26 17:42:53.969542	2022-09-26 18:57:05.854178	\N	-1	\N
51	booked	2022-09-26 17:00:00	1	iouliosa	+306975645865	ioulis@email.com	2022-09-26 19:04:34.731	\N	test	\N	1
2	booked	2022-08-30 14:00:00	1	ioulios	+306975645865	ioulios@email.com	2022-04-04 14:23:19.033811	\N	\N	\N	\N
60	booked	2022-09-26 14:00:00	1	iouliosa	+306975645865	ioulis@email.com	2022-09-27 10:47:25.401628	2022-09-27 11:07:08.0254	\N	1	\N
61	booked	2022-09-26 14:00:00	1	iouliosa	+306975645865	ioulis@email.com	2022-09-27 10:47:25.401628	\N	\N	1	\N
40	booked	2022-09-21 14:00:00	2	Ioulios Tsiko	+306975645865	ioulios.tsiko@gmail.com	2022-09-20 12:46:54.029074	\N	\N	\N	\N
41	booked	2022-09-20 19:15:00	2	Ioulios Tsiko	+306975645865	ioulios.tsiko@gmail.com	2022-09-20 12:49:12.307808	\N	\N	\N	\N
45	booked	2022-09-20 14:00:00	2	Ioulios Tsiko	+306975645865	ioulios.tsiko@gmail.com	2022-09-20 13:55:19.928095	\N	\N	\N	\N
25	booked	2022-07-29 14:00:00	1	ioulios	+306975645865	ioulis@email.com	2022-04-05 09:54:08.472114	\N	\N	\N	\N
44	booked	2022-09-22 14:00:00	1	Dim11	+306972782751	dim1@gmail.com	2022-09-20 13:27:11.615863	\N	\N	\N	\N
35	booked	2022-08-28 14:00:00	2	iouliosa	+306975645865	ioulis@email.com	2022-08-17 17:52:18.646659	\N	\N	\N	\N
46	booked	2022-09-10 14:00:00	1	ioulios	+306975645865	ioulis@email.com	2022-09-21 14:49:40.41655	\N	\N	\N	\N
1	canceled	2022-09-24 23:00:00	1	ioulios	+306975645865	ioulis@email.com	2022-04-04 14:10:21.79407	2022-09-21 11:16:17.095017	\N	\N	\N
62	booked	2023-01-09 14:00:00	2	Ioulios Tsiko	+306975645865	ioulios.tsiko@gmail.com	2022-12-27 11:46:51.744793	\N	test τεστ 123	-1	\N
63	canceled	2022-12-31 14:00:00	1	Ioulios Tsikos	+306975645865	ioulios.tsiko@gmail.com	2022-12-27 12:14:20.009273	2022-12-27 12:14:50.515939	\N	1	1
\.


--
-- Data for Name: gme; Type: TABLE DATA; Schema: main; Owner: vxvtteck
--

COPY main.gme (gme_id, status, img_url, map_url, plrs, age_rng, nme, descr, addr, dur) FROM stdin;
2	active	https://res.cloudinary.com/hwrkhvisl/image/upload/v1586116498/Paradox%20Project/iizdo9dvtx21dqmvrpwp.jpg?fbclid=IwAR0f2SG_pt8_iFH3liW-xHJFz3jt8PKNq4nUmd_2hVd0XRFkouFzhy5f4WU	https://www.google.com/maps/place/Paradox+Project/@37.9600721,23.7055711,17z/data=!3m1!4b1!4m5!3m4!1s0x14a1bcf8e3a8f505:0xd72de356a1596eff!8m2!3d37.9600721!4d23.7077598	3-7	15+	Bookstore	The first escape house in Europe edit	Charokopou 93, Kallithea, 2st Floor	180
87	inactive	https://res.cloudinary.com/hwrkhvisl/image/upload/v1586116498/Paradox%20Project/iizdo9dvtx21dqmvrpwp.jpg?fbclid=IwAR0f2SG_pt8_iFH3liW-xHJFz3jt8PKNq4nUmd_2hVd0XRFkouFzhy5f4WU	https://www.google.com/maps/place/Paradox+Project/@37.9600721,23.7055711,17z/data=!3m1!4b1!4m5!3m4!1s0x14a1bcf8e3a8f505:0xd72de356a1596eff!8m2!3d37.9600721!4d23.7077598	3-7	15+	Academy	The first escape house in Europe edit	Charokopou 93, Kallithea, 2st Floor	180
1	active	https://res.cloudinary.com/hwrkhvisl/image/upload/v1586116498/Paradox%20Project/iizdo9dvtx21dqmvrpwp.jpg?fbclid=IwAR0f2SG_pt8_iFH3liW-xHJFz3jt8PKNq4nUmd_2hVd0XRFkouFzhy5f4WU	https://www.google.com/maps/place/Paradox+Project/@37.9600721,23.7055711,17z/data=!3m1!4b1!4m5!3m4!1s0x14a1bcf8e3a8f505:0xd72de356a1596eff!8m2!3d37.9600721!4d23.7077598	3-7	15+	Mansion	The first escape house in Europe !	Charokopou 93, Kallithea, 2st Floor	180
88	active	https://www.washingtonpost.com/wp-apps/imrs.php?src=https://arc-anglerfish-washpost-prod-washpost.s3.amazonaws.com/public/CAGNYXUKG4I6VHP5TEHZ3TDR7Q.jpg&w=1200	https://www.google.com/maps/place/Acropolis+Museum/@37.9684499,23.7285227,15z/data=!4m2!3m1!1s0x0:0xb00fb46a2c010a3c?sa=X&ved=2ahUKEwjS_4zm45n8AhV3iv0HHVa6CxIQ_BJ6BQiFARAI	3-7	18+	Museum	Solve the mysteries of the museum	Athinon 10, athina 	100
\.


--
-- Data for Name: gme_lcl; Type: TABLE DATA; Schema: main; Owner: vxvtteck
--

COPY main.gme_lcl (gme_lcl_id, gme_id, lcl, nme, descr, addr) FROM stdin;
1	1	en	Mansion	The first escape house in Europe	Charokopou 93, Kallithea, 2st Floor
\.


--
-- Data for Name: schdl; Type: TABLE DATA; Schema: main; Owner: vxvtteck
--

COPY main.schdl (schdl_id, gme_id, schedule, active_by) FROM stdin;
1	1	{"0": ["14:00", "19:00"], "1": ["14:00", "19:00"], "2": ["14:00", "19:00"], "3": ["14:00", "19:00"], "4": ["14:00", "19:00"], "5": ["14:00", "19:00"], "6": ["14:00", "19:00", "23:00"]}	2022-12-27
3	2	{"0": ["14:00", "18:00"], "1": ["14:00", "18:00"], "2": ["14:00", "18:00"], "3": ["14:00", "19:00"], "4": ["14:00", "19:00"], "5": ["14:00", "19:00"], "6": ["14:00", "19:00", "23:00"]}	2022-12-27
4	2	{"1": ["17:15"], "3": ["16:00"], "5": ["15:15"]}	2023-01-17
5	1	{"0": ["14:00"], "1": ["13:15"], "2": ["14:00", "19:45"], "3": ["15:15"], "4": ["12:30", "20:00"], "6": ["17:15"]}	2023-01-17
6	88	{"1": ["11:00", "13:00"], "2": ["16:30"], "3": ["12:15"], "4": ["14:15"]}	2022-12-27
7	88	{"2": ["17:45", "13:45"], "3": ["14:15"], "5": ["17:45"], "6": ["14:30"]}	2023-01-17
\.


--
-- Data for Name: usr; Type: TABLE DATA; Schema: main; Owner: vxvtteck
--

COPY main.usr (usr_id, usrnme, psswrd, salt) FROM stdin;
1	ioulios	c8b2505b76926abdc733523caa9f439142f66aa7293a7baaac0aed41a191eef6	salt
2	dimrks	2bff1032fa51b78a95877cf40d48f522c0a282990048b336e005b9b25ccf6ca6	salt
-1	selfbooked	system reserved user	
\.


--
-- Name: bkng_bkng_id_seq; Type: SEQUENCE SET; Schema: main; Owner: vxvtteck
--

SELECT pg_catalog.setval('main.bkng_bkng_id_seq', 63, true);


--
-- Name: gme_gme_id_seq; Type: SEQUENCE SET; Schema: main; Owner: vxvtteck
--

SELECT pg_catalog.setval('main.gme_gme_id_seq', 88, true);


--
-- Name: gme_lcl_gme_lcl_id_seq; Type: SEQUENCE SET; Schema: main; Owner: vxvtteck
--

SELECT pg_catalog.setval('main.gme_lcl_gme_lcl_id_seq', 1, true);


--
-- Name: schdl_schdl_id_seq; Type: SEQUENCE SET; Schema: main; Owner: vxvtteck
--

SELECT pg_catalog.setval('main.schdl_schdl_id_seq', 7, true);


--
-- Name: usr_usr_id_seq; Type: SEQUENCE SET; Schema: main; Owner: vxvtteck
--

SELECT pg_catalog.setval('main.usr_usr_id_seq', 1, true);


--
-- Name: bkng bkng_pkey; Type: CONSTRAINT; Schema: main; Owner: vxvtteck
--

ALTER TABLE ONLY main.bkng
    ADD CONSTRAINT bkng_pkey PRIMARY KEY (bkng_id);


--
-- Name: gme_lcl gme_lcl_pkey; Type: CONSTRAINT; Schema: main; Owner: vxvtteck
--

ALTER TABLE ONLY main.gme_lcl
    ADD CONSTRAINT gme_lcl_pkey PRIMARY KEY (gme_lcl_id);


--
-- Name: gme gme_pkey; Type: CONSTRAINT; Schema: main; Owner: vxvtteck
--

ALTER TABLE ONLY main.gme
    ADD CONSTRAINT gme_pkey PRIMARY KEY (gme_id);


--
-- Name: schdl schdl_pkey; Type: CONSTRAINT; Schema: main; Owner: vxvtteck
--

ALTER TABLE ONLY main.schdl
    ADD CONSTRAINT schdl_pkey PRIMARY KEY (schdl_id);


--
-- Name: usr usr_pkey; Type: CONSTRAINT; Schema: main; Owner: vxvtteck
--

ALTER TABLE ONLY main.usr
    ADD CONSTRAINT usr_pkey PRIMARY KEY (usr_id);


--
-- Name: one_bkng; Type: INDEX; Schema: main; Owner: vxvtteck
--

CREATE UNIQUE INDEX one_bkng ON main.bkng USING btree (dte, gme_id) WHERE (canceled_at IS NULL);


--
-- Name: bkng bkng_changed_trig; Type: TRIGGER; Schema: main; Owner: vxvtteck
--

CREATE TRIGGER bkng_changed_trig AFTER INSERT OR UPDATE ON main.bkng FOR EACH ROW EXECUTE PROCEDURE public.notify_bkng_changes_func();


--
-- Name: bkng bkng_gme_id_fkey; Type: FK CONSTRAINT; Schema: main; Owner: vxvtteck
--

ALTER TABLE ONLY main.bkng
    ADD CONSTRAINT bkng_gme_id_fkey FOREIGN KEY (gme_id) REFERENCES main.gme(gme_id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- Name: bkng bkng_lst_edted_by_fkey; Type: FK CONSTRAINT; Schema: main; Owner: vxvtteck
--

ALTER TABLE ONLY main.bkng
    ADD CONSTRAINT bkng_lst_edted_by_fkey FOREIGN KEY (lst_edted_by) REFERENCES main.usr(usr_id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- Name: bkng bkng_src_fkey; Type: FK CONSTRAINT; Schema: main; Owner: vxvtteck
--

ALTER TABLE ONLY main.bkng
    ADD CONSTRAINT bkng_src_fkey FOREIGN KEY (bked_by) REFERENCES main.usr(usr_id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- Name: gme_lcl gme_lcl_gme_id_fkey; Type: FK CONSTRAINT; Schema: main; Owner: vxvtteck
--

ALTER TABLE ONLY main.gme_lcl
    ADD CONSTRAINT gme_lcl_gme_id_fkey FOREIGN KEY (gme_id) REFERENCES main.gme(gme_id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- Name: schdl schdl_gme_id_fkey; Type: FK CONSTRAINT; Schema: main; Owner: vxvtteck
--

ALTER TABLE ONLY main.schdl
    ADD CONSTRAINT schdl_gme_id_fkey FOREIGN KEY (gme_id) REFERENCES main.gme(gme_id) ON UPDATE CASCADE ON DELETE RESTRICT;


--
-- Name: SCHEMA heroku_ext; Type: ACL; Schema: -; Owner: vxvtteck
--

REVOKE ALL ON SCHEMA heroku_ext FROM vxvtteck;
GRANT CREATE ON SCHEMA heroku_ext TO vxvtteck;
GRANT USAGE ON SCHEMA heroku_ext TO vxvtteck WITH GRANT OPTION;


--
-- Name: SCHEMA public; Type: ACL; Schema: -; Owner: postgres
--

REVOKE USAGE ON SCHEMA public FROM PUBLIC;
GRANT ALL ON SCHEMA public TO PUBLIC;


--
-- PostgreSQL database dump complete
--

