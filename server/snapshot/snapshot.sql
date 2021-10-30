--
-- PostgreSQL database dump
--

-- Dumped from database version 13.4 (Debian 13.4-4.pgdg110+1)
-- Dumped by pg_dump version 13.4 (Debian 13.4-4.pgdg110+1)

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
-- Name: evaluate_access_user(jsonb, bigint, bigint); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.evaluate_access_user(condition jsonb, session_id bigint, user_id bigint) RETURNS boolean
    LANGUAGE sql
    AS $$
select (condition -> 'id' ->> 'value')::bigint = user_id
           OR ((condition -> 'id' ->> 'field') = 'session_id'
        AND user_id = session_id)
$$;


ALTER FUNCTION public.evaluate_access_user(condition jsonb, session_id bigint, user_id bigint) OWNER TO postgres;

--
-- Name: generate_ksuid(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.generate_ksuid() RETURNS character
    LANGUAGE sql
    AS $$
select substring(
               replace(to_char(clock_timestamp(), 'yyyymmddhh24missus')
                           || (to_char(random() * 1e9, '000000000')
                           ), ' ', ''), 1, 27)::char(27);
$$;


ALTER FUNCTION public.generate_ksuid() OWNER TO postgres;

--
-- Name: get_user_role_global_type_id(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_user_role_global_type_id() RETURNS bigint
    LANGUAGE sql
    AS $$
SELECT role_type_id
FROM "role_type"
WHERE name = 'User Role, Global'
$$;


ALTER FUNCTION public.get_user_role_global_type_id() OWNER TO postgres;

--
-- Name: get_user_role_type_id(); Type: FUNCTION; Schema: public; Owner: postgres
--

CREATE FUNCTION public.get_user_role_type_id() RETURNS bigint
    LANGUAGE sql
    AS $$
SELECT role_type_id
FROM "role_type"
WHERE name = 'User Role'
$$;


ALTER FUNCTION public.get_user_role_type_id() OWNER TO postgres;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: role; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.role (
    role_id bigint NOT NULL,
    created_by bigint,
    created_at timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
    effective_date timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
    end_date timestamp without time zone,
    is_active boolean DEFAULT true NOT NULL,
    name text NOT NULL,
    role_type_id bigint,
    tags text[] NOT NULL
);


ALTER TABLE public.role OWNER TO postgres;

--
-- Name: role_action; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.role_action (
    name text NOT NULL,
    created_at timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL
);


ALTER TABLE public.role_action OWNER TO postgres;

--
-- Name: role_assignment; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.role_assignment (
    role_assignment_id bigint NOT NULL,
    created_by bigint,
    created_at timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
    effective_date timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
    end_date timestamp without time zone,
    is_active boolean DEFAULT true NOT NULL,
    role_id bigint NOT NULL,
    user_id bigint NOT NULL,
    tags text[] NOT NULL
);


ALTER TABLE public.role_assignment OWNER TO postgres;

--
-- Name: role_permission; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.role_permission (
    role_permission_id bigint NOT NULL,
    created_by bigint,
    created_at timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
    effective_date timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
    end_date timestamp without time zone,
    is_active boolean DEFAULT true NOT NULL,
    role_id bigint NOT NULL,
    action text NOT NULL,
    subject text NOT NULL,
    subject_id bigint,
    conditions jsonb NOT NULL,
    fields text[]
);


ALTER TABLE public.role_permission OWNER TO postgres;

--
-- Name: role_subject; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.role_subject (
    name text NOT NULL,
    created_at timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL
);


ALTER TABLE public.role_subject OWNER TO postgres;

--
-- Name: role_type; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.role_type (
    role_type_id bigint NOT NULL,
    created_at timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
    name text NOT NULL,
    multiple_assignments_allowed boolean NOT NULL,
    system_managed boolean NOT NULL
);


ALTER TABLE public.role_type OWNER TO postgres;

--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO postgres;

--
-- Name: user; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."user" (
    user_id bigint NOT NULL,
    uuid uuid NOT NULL,
    created_by bigint,
    created_at timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
    effective_date timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
    end_date timestamp without time zone,
    is_active boolean DEFAULT true NOT NULL,
    email text NOT NULL,
    name text NOT NULL,
    tags text[] NOT NULL,
    common_access text DEFAULT 'private'::text NOT NULL
);


ALTER TABLE public."user" OWNER TO postgres;

--
-- Name: role_action role_action_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_action
    ADD CONSTRAINT role_action_pk PRIMARY KEY (name);


--
-- Name: role_assignment role_assignment_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_assignment
    ADD CONSTRAINT role_assignment_pk PRIMARY KEY (role_assignment_id);


--
-- Name: role_permission role_permission_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_permission
    ADD CONSTRAINT role_permission_pk PRIMARY KEY (role_permission_id);


--
-- Name: role role_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role
    ADD CONSTRAINT role_pk PRIMARY KEY (role_id);


--
-- Name: role_subject role_subject_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_subject
    ADD CONSTRAINT role_subject_pk PRIMARY KEY (name);


--
-- Name: role_type role_type_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_type
    ADD CONSTRAINT role_type_pk PRIMARY KEY (role_type_id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: user user_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pk PRIMARY KEY (user_id);


--
-- Name: roletype_name_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX roletype_name_uindex ON public.role_type USING btree (name);


--
-- Name: user_email_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX user_email_uindex ON public."user" USING btree (email);


--
-- Name: user_uuid_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX user_uuid_uindex ON public."user" USING btree (uuid);


--
-- Name: role fk_role_createdby; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role
    ADD CONSTRAINT fk_role_createdby FOREIGN KEY (created_by) REFERENCES public."user"(user_id);


--
-- Name: role fk_role_roletype; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role
    ADD CONSTRAINT fk_role_roletype FOREIGN KEY (role_type_id) REFERENCES public.role_type(role_type_id);


--
-- Name: role_assignment fk_roleassignment_createdby; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_assignment
    ADD CONSTRAINT fk_roleassignment_createdby FOREIGN KEY (created_by) REFERENCES public."user"(user_id);


--
-- Name: role_permission fk_roleassignment_createdby; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_permission
    ADD CONSTRAINT fk_roleassignment_createdby FOREIGN KEY (created_by) REFERENCES public."user"(user_id);


--
-- Name: role_assignment fk_roleassignment_role; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_assignment
    ADD CONSTRAINT fk_roleassignment_role FOREIGN KEY (role_id) REFERENCES public.role(role_id);


--
-- Name: role_assignment fk_roleassignment_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_assignment
    ADD CONSTRAINT fk_roleassignment_user FOREIGN KEY (user_id) REFERENCES public."user"(user_id);


--
-- Name: role_permission fk_rolepermission_action; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_permission
    ADD CONSTRAINT fk_rolepermission_action FOREIGN KEY (action) REFERENCES public.role_action(name);


--
-- Name: role_permission fk_rolepermission_role; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_permission
    ADD CONSTRAINT fk_rolepermission_role FOREIGN KEY (role_id) REFERENCES public.role(role_id);


--
-- Name: role_permission fk_rolepermission_subject; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_permission
    ADD CONSTRAINT fk_rolepermission_subject FOREIGN KEY (subject) REFERENCES public.role_subject(name);


--
-- Name: user fk_user_createdby; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT fk_user_createdby FOREIGN KEY (created_by) REFERENCES public."user"(user_id);


--
-- PostgreSQL database dump complete
--

