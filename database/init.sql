--
-- PostgreSQL database dump
--

-- Dumped from database version 13.1 (Debian 13.1-1.pgdg100+1)
-- Dumped by pg_dump version 13.1 (Debian 13.1-1.pgdg100+1)

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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: actions; Type: TABLE; Schema: public; Owner: redteam
--

CREATE TABLE public.actions (
    uuid text NOT NULL,
    actiontype text NOT NULL,
    actioncmd text NOT NULL,
    actionresponse text
);


ALTER TABLE public.actions OWNER TO redteam;

--
-- Name: agent; Type: TABLE; Schema: public; Owner: redteam
--

CREATE TABLE public.agent (
    uuid text NOT NULL,
    agentname text NOT NULL,
    agentversion text NOT NULL,
    primaryip text NOT NULL,
    hostname text NOT NULL,
    mac text NOT NULL,
    agentos text NOT NULL,
    otherips text NOT NULL
);


ALTER TABLE public.agent OWNER TO redteam;

--
-- Name: agentbeacon; Type: TABLE; Schema: public; Owner: redteam
--

CREATE TABLE public.agentbeacon (
    uuid text NOT NULL,
    registertime text NOT NULL,
    lastbeacon text NOT NULL,
    actionqueue text NOT NULL
);


ALTER TABLE public.agentbeacon OWNER TO redteam;

--
-- Data for Name: actions; Type: TABLE DATA; Schema: public; Owner: redteam
--

COPY public.actions (uuid, actiontype, actioncmd, actionresponse) FROM stdin;
\.


--
-- Data for Name: agent; Type: TABLE DATA; Schema: public; Owner: redteam
--

COPY public.agent (uuid, agentname, agentversion, primaryip, hostname, mac, agentos, otherips) FROM stdin;
\.


--
-- Data for Name: agentbeacon; Type: TABLE DATA; Schema: public; Owner: redteam
--

COPY public.agentbeacon (uuid, registertime, lastbeacon, actionqueue) FROM stdin;
\.


--
-- Name: actions actions_pkey; Type: CONSTRAINT; Schema: public; Owner: redteam
--

ALTER TABLE ONLY public.actions
    ADD CONSTRAINT actions_pkey PRIMARY KEY (uuid);


--
-- Name: agent agent_pkey; Type: CONSTRAINT; Schema: public; Owner: redteam
--

ALTER TABLE ONLY public.agent
    ADD CONSTRAINT agent_pkey PRIMARY KEY (uuid);


--
-- Name: agentbeacon agentbeacon_pkey; Type: CONSTRAINT; Schema: public; Owner: redteam
--

ALTER TABLE ONLY public.agentbeacon
    ADD CONSTRAINT agentbeacon_pkey PRIMARY KEY (uuid);


--
-- PostgreSQL database dump complete
--

