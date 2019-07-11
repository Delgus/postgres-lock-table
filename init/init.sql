--
-- PostgreSQL database dump
--

-- Dumped from database version 11.2 (Debian 11.2-1.pgdg90+1)
-- Dumped by pg_dump version 11.2

-- Started on 2019-04-19 21:48:20 UTC

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_with_oids = false;

--
-- TOC entry 196 (class 1259 OID 16386)
-- Name: books; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.books (
    isbn character(14) NOT NULL,
    title character varying(255) NOT NULL,
    author character varying(255) NOT NULL,
    price numeric(5,2) NOT NULL
);


ALTER TABLE public.books OWNER TO postgres;

--
-- TOC entry 2862 (class 0 OID 16386)
-- Dependencies: 196
-- Data for Name: books; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.books (isbn, title, author, price) FROM stdin;
978-1503261969	Emma	Jayne Austen	9.44
978-1505255607	The Time Machine	H. G. Wells	5.99
978-1503379640	The Prince	Niccolò Machiavelli	6.99
\.


--
-- TOC entry 2740 (class 2606 OID 16393)
-- Name: books books_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.books
    ADD CONSTRAINT books_pkey PRIMARY KEY (isbn);


-- Completed on 2019-04-19 21:48:20 UTC

--
-- PostgreSQL database dump complete
--