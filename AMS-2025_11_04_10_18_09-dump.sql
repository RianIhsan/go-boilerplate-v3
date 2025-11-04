--
-- PostgreSQL database dump
--

\restrict Q2Zn8Y9XULufo2p0sjwaOiGKhiqlppVpjtUUwoctqhTfXPFvGavrYTNto418hyS

-- Dumped from database version 14.19 (Ubuntu 14.19-0ubuntu0.22.04.1)
-- Dumped by pg_dump version 14.19 (Ubuntu 14.19-0ubuntu0.22.04.1)

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
-- Name: accesses; Type: TABLE; Schema: public; Owner: sentuhxams
--

CREATE TABLE public.accesses (
    id bigint NOT NULL,
    name text,
    link text,
    priority bigint,
    access_id bigint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.accesses OWNER TO sentuhxams;

--
-- Name: accesses_id_seq; Type: SEQUENCE; Schema: public; Owner: sentuhxams
--

CREATE SEQUENCE public.accesses_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.accesses_id_seq OWNER TO sentuhxams;

--
-- Name: accesses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentuhxams
--

ALTER SEQUENCE public.accesses_id_seq OWNED BY public.accesses.id;


--
-- Name: permissions; Type: TABLE; Schema: public; Owner: sentuhxams
--

CREATE TABLE public.permissions (
    id bigint NOT NULL,
    name character varying(100),
    path character varying(100),
    method character varying(100),
    access_id bigint,
    type character varying(100),
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.permissions OWNER TO sentuhxams;

--
-- Name: permissions_id_seq; Type: SEQUENCE; Schema: public; Owner: sentuhxams
--

CREATE SEQUENCE public.permissions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.permissions_id_seq OWNER TO sentuhxams;

--
-- Name: permissions_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentuhxams
--

ALTER SEQUENCE public.permissions_id_seq OWNED BY public.permissions.id;


--
-- Name: role_accesses; Type: TABLE; Schema: public; Owner: sentuhxams
--

CREATE TABLE public.role_accesses (
    access_id bigint NOT NULL,
    role_id bigint NOT NULL
);


ALTER TABLE public.role_accesses OWNER TO sentuhxams;

--
-- Name: role_permissions; Type: TABLE; Schema: public; Owner: sentuhxams
--

CREATE TABLE public.role_permissions (
    role_id bigint NOT NULL,
    permission_id bigint NOT NULL,
    access_id bigint
);


ALTER TABLE public.role_permissions OWNER TO sentuhxams;

--
-- Name: roles; Type: TABLE; Schema: public; Owner: sentuhxams
--

CREATE TABLE public.roles (
    id bigint NOT NULL,
    name text,
    description text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone
);


ALTER TABLE public.roles OWNER TO sentuhxams;

--
-- Name: roles_id_seq; Type: SEQUENCE; Schema: public; Owner: sentuhxams
--

CREATE SEQUENCE public.roles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.roles_id_seq OWNER TO sentuhxams;

--
-- Name: roles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentuhxams
--

ALTER SEQUENCE public.roles_id_seq OWNED BY public.roles.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: sentuhxams
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    avatar character varying(255),
    name character varying(100) NOT NULL,
    email character varying(100) NOT NULL,
    password character varying(100) NOT NULL,
    nfc_tag character varying(100),
    role_id bigint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    username character varying(100)
);


ALTER TABLE public.users OWNER TO sentuhxams;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: sentuhxams
--

CREATE SEQUENCE public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_id_seq OWNER TO sentuhxams;

--
-- Name: users_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: sentuhxams
--

ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;


--
-- Name: accesses id; Type: DEFAULT; Schema: public; Owner: sentuhxams
--

ALTER TABLE ONLY public.accesses ALTER COLUMN id SET DEFAULT nextval('public.accesses_id_seq'::regclass);


--
-- Name: permissions id; Type: DEFAULT; Schema: public; Owner: sentuhxams
--

ALTER TABLE ONLY public.permissions ALTER COLUMN id SET DEFAULT nextval('public.permissions_id_seq'::regclass);


--
-- Name: roles id; Type: DEFAULT; Schema: public; Owner: sentuhxams
--

ALTER TABLE ONLY public.roles ALTER COLUMN id SET DEFAULT nextval('public.roles_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: sentuhxams
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);


--
-- Data for Name: accesses; Type: TABLE DATA; Schema: public; Owner: sentuhxams
--

COPY public.accesses (id, name, link, priority, access_id, created_at, updated_at) FROM stdin;
1	User	user	1	\N	2025-10-30 08:19:00.452486+00	2025-10-30 08:19:00.452486+00
2	Access	access	2	\N	2025-11-03 06:28:23.369863+00	2025-11-03 06:28:23.369863+00
3	Role	role	3	\N	2025-11-03 06:40:28.444404+00	2025-11-03 06:40:28.444404+00
\.


--
-- Data for Name: permissions; Type: TABLE DATA; Schema: public; Owner: sentuhxams
--

COPY public.permissions (id, name, path, method, access_id, type, created_at, updated_at) FROM stdin;
1	User	/users	READ	1	page	2025-10-30 08:20:23.664433+00	2025-10-30 08:20:23.664433+00
2	Get List User	/api/v1/users	GET	1	api	2025-10-30 08:20:38.879324+00	2025-10-30 08:20:38.879324+00
3	Get User By ID	/api/v1/users/*	GET	1	api	2025-10-30 08:21:09.545225+00	2025-10-30 08:21:09.545225+00
4	Self Update User	/api/v1/users/protected	PUT	1	api	2025-10-30 08:21:56.859776+00	2025-10-30 09:23:12.625251+00
5	Update Avatar	/api/v1/users/avatar	PUT	1	api	2025-10-30 09:59:07.725214+00	2025-10-30 09:59:07.725214+00
6	Update User By SuperAdmin	/api/v1/users/*	PUT	1	api	2025-10-31 09:29:59.399907+00	2025-10-31 09:29:59.399907+00
7	Delete User By SuperAdmin	/api/v1/users/*	DELETE	1	api	2025-10-31 09:30:22.132497+00	2025-10-31 09:30:22.132497+00
8	Access	/access	READ	2	page	2025-11-03 06:29:28.507554+00	2025-11-03 06:29:28.507554+00
10	Get list Access	/api/v1/access	GET	2	api	2025-11-03 06:30:20.335955+00	2025-11-03 06:30:20.335955+00
11	Register Access	/api/v1/access	POST	2	api	2025-11-03 06:31:35.577264+00	2025-11-03 06:31:35.577264+00
12	Get Access by ID	/api/v1/access/*	GET	2	api	2025-11-03 06:32:22.254828+00	2025-11-03 06:32:22.254828+00
13	Update access	/api/v1/access/*	PUT	2	api	2025-11-03 06:32:57.730064+00	2025-11-03 06:32:57.730064+00
14	Delete access	/api/v1/access/*	DELETE	2	api	2025-11-03 06:33:06.409231+00	2025-11-03 06:33:06.409231+00
15	Role	/role	READ	3	page	2025-11-03 06:41:48.440791+00	2025-11-03 06:41:48.440791+00
16	Register Role	/api/v1/role	POST	3	api	2025-11-03 06:42:46.56164+00	2025-11-03 06:42:46.56164+00
17	Get list Role	/api/v1/role	GET	3	api	2025-11-03 06:43:01.333854+00	2025-11-03 06:43:01.333854+00
18	Modify Role Permission	/api/v1/role/permission	PUT	3	api	2025-11-03 06:44:28.317506+00	2025-11-03 06:44:28.317506+00
19	Get Role By ID	/api/v1/role/*	GET	3	api	2025-11-03 06:44:50.033112+00	2025-11-03 06:44:50.033112+00
20	Update Role	/api/v1/role/*	PUT	3	api	2025-11-03 06:45:07.524422+00	2025-11-03 06:45:07.524422+00
21	Delete Role	/api/v1/role/*	DELETE	3	api	2025-11-03 06:45:19.949566+00	2025-11-03 06:45:19.949566+00
\.


--
-- Data for Name: role_accesses; Type: TABLE DATA; Schema: public; Owner: sentuhxams
--

COPY public.role_accesses (access_id, role_id) FROM stdin;
1	1
2	1
3	1
\.


--
-- Data for Name: role_permissions; Type: TABLE DATA; Schema: public; Owner: sentuhxams
--

COPY public.role_permissions (role_id, permission_id, access_id) FROM stdin;
1	1	1
1	2	1
1	3	1
1	4	1
1	5	1
1	6	1
1	7	1
1	8	2
1	10	2
1	11	2
1	12	2
1	13	2
1	14	2
1	15	3
1	16	3
1	17	3
1	18	3
1	19	3
1	20	3
1	21	3
\.


--
-- Data for Name: roles; Type: TABLE DATA; Schema: public; Owner: sentuhxams
--

COPY public.roles (id, name, description, created_at, updated_at, deleted_at) FROM stdin;
1	SUPER ADMIN		2025-10-30 07:46:04.619771+00	2025-10-30 07:46:04.619771+00	\N
2	Admin		2025-11-03 03:14:46.420101+00	2025-11-03 03:14:46.420101+00	\N
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: sentuhxams
--

COPY public.users (id, avatar, name, email, password, nfc_tag, role_id, created_at, updated_at, deleted_at, username) FROM stdin;
2	http://170.64.164.121:9000/ams-bucket/default-avatar.png	Haikal	haikal@sentuh.id	$2a$10$.IKEs86juBItx5DPAObj4OLNY0QcvTclIj/Kx3spmofAlFHUPgusu	\N	1	2025-11-03 03:49:05.31228+00	2025-11-03 03:49:05.31228+00	\N	\N
3	http://170.64.164.121:9000/ams-bucket/default-avatar.png	Haikal2	haikal2@sentuh.id	$2a$10$5vqFJLArDLNLXv/0rKp8Cu4QMCgage/86hAIuOZaOLTa6HpKcz3da	\N	1	2025-11-03 03:49:38.313422+00	2025-11-03 03:49:38.313422+00	\N	\N
4	http://170.64.164.121:9000/ams-bucket/default-avatar.png	Haikal3	haikal3@sentuh.id	$2a$10$Hy4qek2o.wdnxvABZQX69uI1v6Zxjrg3cAKiMblWZqeIK9fEQe/VC	\N	1	2025-11-03 04:04:18.631776+00	2025-11-03 04:04:18.631776+00	\N	
5	http://170.64.164.121:9000/ams-bucket/default-avatar.png	Haikal5	haikal5@sentuh.id	$2a$10$kzXQQNcYoZZjg1SlpLazUuHQihPxe9H9BNYGCLU1VAaOtQmWUX0rO	\N	1	2025-11-03 04:24:25.100728+00	2025-11-03 04:24:25.100728+00	\N	haikal5
15	http://170.64.164.121:9000/ams-bucket/1762225267-1761818417-police-girl-cute-retro-anime-aesthetic-xr3f0iq1jd5jsdf0.webp	asd11 update	asdupdate@gmail.com	$2a$10$qBfo/pbRpNE8j5axlXZ7FO3DgXQhvfwuRLqNhssNbFJdwfL.F1w3y	\N	2	2025-11-03 09:48:10.426043+00	2025-11-04 03:01:07.861169+00	\N	asd11
1	http://170.64.164.121:9000/ams-bucket/default-avatar.png	SUPERADMIN	superadmin@sentuh.id	$2a$10$cpNUaTI.ApW4jZjQrO89M.DboO8j9/pxUvyMm6xjCGz8nryBUv3ae	\N	1	2025-11-03 03:48:46.067872+00	2025-11-04 03:02:54.467905+00	\N	\N
\.


--
-- Name: accesses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentuhxams
--

SELECT pg_catalog.setval('public.accesses_id_seq', 3, true);


--
-- Name: permissions_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentuhxams
--

SELECT pg_catalog.setval('public.permissions_id_seq', 21, true);


--
-- Name: roles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentuhxams
--

SELECT pg_catalog.setval('public.roles_id_seq', 2, true);


--
-- Name: users_id_seq; Type: SEQUENCE SET; Schema: public; Owner: sentuhxams
--

SELECT pg_catalog.setval('public.users_id_seq', 15, true);


--
-- Name: accesses accesses_pkey; Type: CONSTRAINT; Schema: public; Owner: sentuhxams
--

ALTER TABLE ONLY public.accesses
    ADD CONSTRAINT accesses_pkey PRIMARY KEY (id);


--
-- Name: permissions permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: sentuhxams
--

ALTER TABLE ONLY public.permissions
    ADD CONSTRAINT permissions_pkey PRIMARY KEY (id);


--
-- Name: role_accesses role_accesses_pkey; Type: CONSTRAINT; Schema: public; Owner: sentuhxams
--

ALTER TABLE ONLY public.role_accesses
    ADD CONSTRAINT role_accesses_pkey PRIMARY KEY (access_id, role_id);


--
-- Name: role_permissions role_permissions_pkey; Type: CONSTRAINT; Schema: public; Owner: sentuhxams
--

ALTER TABLE ONLY public.role_permissions
    ADD CONSTRAINT role_permissions_pkey PRIMARY KEY (role_id, permission_id);


--
-- Name: roles roles_pkey; Type: CONSTRAINT; Schema: public; Owner: sentuhxams
--

ALTER TABLE ONLY public.roles
    ADD CONSTRAINT roles_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: sentuhxams
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: idx_roles_deleted_at; Type: INDEX; Schema: public; Owner: sentuhxams
--

CREATE INDEX idx_roles_deleted_at ON public.roles USING btree (deleted_at);


--
-- Name: idx_users_deleted_at; Type: INDEX; Schema: public; Owner: sentuhxams
--

CREATE INDEX idx_users_deleted_at ON public.users USING btree (deleted_at);


--
-- Name: uni_users_email; Type: INDEX; Schema: public; Owner: sentuhxams
--

CREATE UNIQUE INDEX uni_users_email ON public.users USING btree (email);


--
-- Name: permissions fk_accesses_permissions; Type: FK CONSTRAINT; Schema: public; Owner: sentuhxams
--

ALTER TABLE ONLY public.permissions
    ADD CONSTRAINT fk_accesses_permissions FOREIGN KEY (access_id) REFERENCES public.accesses(id);


--
-- Name: accesses fk_accesses_sub_access; Type: FK CONSTRAINT; Schema: public; Owner: sentuhxams
--

ALTER TABLE ONLY public.accesses
    ADD CONSTRAINT fk_accesses_sub_access FOREIGN KEY (access_id) REFERENCES public.accesses(id);


--
-- Name: accesses fk_permissions_access; Type: FK CONSTRAINT; Schema: public; Owner: sentuhxams
--

ALTER TABLE ONLY public.accesses
    ADD CONSTRAINT fk_permissions_access FOREIGN KEY (access_id) REFERENCES public.permissions(id) ON UPDATE CASCADE ON DELETE CASCADE;


--
-- Name: role_accesses fk_role_accesses_access; Type: FK CONSTRAINT; Schema: public; Owner: sentuhxams
--

ALTER TABLE ONLY public.role_accesses
    ADD CONSTRAINT fk_role_accesses_access FOREIGN KEY (access_id) REFERENCES public.accesses(id);


--
-- Name: role_accesses fk_role_accesses_role; Type: FK CONSTRAINT; Schema: public; Owner: sentuhxams
--

ALTER TABLE ONLY public.role_accesses
    ADD CONSTRAINT fk_role_accesses_role FOREIGN KEY (role_id) REFERENCES public.roles(id);


--
-- Name: users fk_users_role; Type: FK CONSTRAINT; Schema: public; Owner: sentuhxams
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT fk_users_role FOREIGN KEY (role_id) REFERENCES public.roles(id);


--
-- PostgreSQL database dump complete
--

\unrestrict Q2Zn8Y9XULufo2p0sjwaOiGKhiqlppVpjtUUwoctqhTfXPFvGavrYTNto418hyS

