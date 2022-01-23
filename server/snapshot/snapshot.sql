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

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: common_access; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.common_access (
    name text NOT NULL,
    created_at timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL
);


ALTER TABLE public.common_access OWNER TO postgres;

--
-- Name: dnd5e_language; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.dnd5e_language (
    dnd5e_language_id bigint NOT NULL,
    created_by bigint,
    created_at timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
    version bigint DEFAULT 0 NOT NULL,
    name text NOT NULL
);


ALTER TABLE public.dnd5e_language OWNER TO postgres;

--
-- Name: dnd5e_monster; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.dnd5e_monster (
    dnd5e_monster_id bigint NOT NULL,
    created_by bigint,
    created_at timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
    version bigint DEFAULT 0 NOT NULL,
    dnd5e_world_id bigint,
    name text NOT NULL,
    user_tags text[] NOT NULL,
    system_tags text[] NOT NULL,
    monster_type text NOT NULL,
    alignment text NOT NULL,
    size_category text NOT NULL,
    milli_challenge_rating bigint NOT NULL,
    languages text[] NOT NULL,
    description text NOT NULL
);


ALTER TABLE public.dnd5e_monster OWNER TO postgres;

--
-- Name: dnd5e_monster_type; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.dnd5e_monster_type (
    created_by bigint,
    created_at timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
    version bigint DEFAULT 0 NOT NULL,
    name text NOT NULL
);


ALTER TABLE public.dnd5e_monster_type OWNER TO postgres;

--
-- Name: dnd5e_size_category; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.dnd5e_size_category (
    created_by bigint,
    created_at timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
    version bigint DEFAULT 0 NOT NULL,
    name text NOT NULL,
    space text NOT NULL
);


ALTER TABLE public.dnd5e_size_category OWNER TO postgres;

--
-- Name: dnd5e_world; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.dnd5e_world (
    dnd5e_world_id bigint NOT NULL,
    created_by bigint,
    created_at timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
    version bigint DEFAULT 0 NOT NULL,
    is_active boolean DEFAULT true NOT NULL,
    common_access text NOT NULL,
    user_tags text[] NOT NULL,
    system_tags text[] NOT NULL,
    name text NOT NULL,
    module text,
    description text NOT NULL,
    external_source_id bigint,
    external_source_key text
);


ALTER TABLE public.dnd5e_world OWNER TO postgres;

--
-- Name: dnd5e_world_assignment; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.dnd5e_world_assignment (
    created_by bigint,
    created_at timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
    version bigint DEFAULT 0 NOT NULL,
    user_id bigint NOT NULL,
    dnd5e_world_id bigint NOT NULL,
    role_action text NOT NULL
);


ALTER TABLE public.dnd5e_world_assignment OWNER TO postgres;

--
-- Name: external_source; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.external_source (
    external_source_id bigint NOT NULL,
    created_by bigint,
    created_at timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
    version bigint DEFAULT 0 NOT NULL,
    external_source_key text NOT NULL,
    external_source_version text NOT NULL,
    user_tags text[] NOT NULL,
    system_tags text[] NOT NULL,
    name text NOT NULL
);


ALTER TABLE public.external_source OWNER TO postgres;

--
-- Name: role_action; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.role_action (
    name text NOT NULL,
    role_subject text NOT NULL,
    created_at timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL
);


ALTER TABLE public.role_action OWNER TO postgres;

--
-- Name: role_subject; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.role_subject (
    name text NOT NULL,
    created_at timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL
);


ALTER TABLE public.role_subject OWNER TO postgres;

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
    version bigint DEFAULT 0 NOT NULL,
    effective_date timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
    end_date timestamp without time zone,
    is_active boolean DEFAULT true NOT NULL,
    common_access text NOT NULL,
    email text NOT NULL,
    name text NOT NULL,
    user_tags text[] NOT NULL,
    system_tags text[] NOT NULL
);


ALTER TABLE public."user" OWNER TO postgres;

--
-- Name: common_access common_access_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.common_access
    ADD CONSTRAINT common_access_pk PRIMARY KEY (name);


--
-- Name: dnd5e_language dnd5e_language_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dnd5e_language
    ADD CONSTRAINT dnd5e_language_pk PRIMARY KEY (dnd5e_language_id);


--
-- Name: dnd5e_monster dnd5e_monster_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dnd5e_monster
    ADD CONSTRAINT dnd5e_monster_pk PRIMARY KEY (dnd5e_monster_id);


--
-- Name: dnd5e_monster_type dnd5e_monster_type_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dnd5e_monster_type
    ADD CONSTRAINT dnd5e_monster_type_pk PRIMARY KEY (name);


--
-- Name: dnd5e_size_category dnd5e_size_category_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dnd5e_size_category
    ADD CONSTRAINT dnd5e_size_category_pk PRIMARY KEY (name);


--
-- Name: dnd5e_world_assignment dnd5e_world_assignment_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dnd5e_world_assignment
    ADD CONSTRAINT dnd5e_world_assignment_pk PRIMARY KEY (user_id, dnd5e_world_id);


--
-- Name: dnd5e_world dnd5e_world_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dnd5e_world
    ADD CONSTRAINT dnd5e_world_pk PRIMARY KEY (dnd5e_world_id);


--
-- Name: external_source external_source_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.external_source
    ADD CONSTRAINT external_source_pk PRIMARY KEY (external_source_id);


--
-- Name: role_action role_action_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_action
    ADD CONSTRAINT role_action_pk PRIMARY KEY (name);


--
-- Name: role_subject role_subject_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_subject
    ADD CONSTRAINT role_subject_pk PRIMARY KEY (name);


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
-- Name: dnd5e_monster_world; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX dnd5e_monster_world ON public.dnd5e_monster USING btree (dnd5e_world_id, dnd5e_monster_id);


--
-- Name: dnd5e_monster_world_name; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX dnd5e_monster_world_name ON public.dnd5e_monster USING btree (dnd5e_world_id, name) INCLUDE (dnd5e_monster_id);


--
-- Name: user_email_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX user_email_uindex ON public."user" USING btree (email);


--
-- Name: user_uuid_uindex; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX user_uuid_uindex ON public."user" USING btree (uuid);


--
-- Name: dnd5e_language fk_dnd5e_language_createdby; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dnd5e_language
    ADD CONSTRAINT fk_dnd5e_language_createdby FOREIGN KEY (created_by) REFERENCES public."user"(user_id);


--
-- Name: dnd5e_monster fk_dnd5e_monster_createdby; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dnd5e_monster
    ADD CONSTRAINT fk_dnd5e_monster_createdby FOREIGN KEY (created_by) REFERENCES public."user"(user_id);


--
-- Name: dnd5e_monster fk_dnd5e_monster_dnd5e_size_category; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dnd5e_monster
    ADD CONSTRAINT fk_dnd5e_monster_dnd5e_size_category FOREIGN KEY (size_category) REFERENCES public.dnd5e_size_category(name);


--
-- Name: dnd5e_monster fk_dnd5e_monster_type; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dnd5e_monster
    ADD CONSTRAINT fk_dnd5e_monster_type FOREIGN KEY (monster_type) REFERENCES public.dnd5e_monster_type(name);


--
-- Name: dnd5e_monster_type fk_dnd5e_monster_type_createdby; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dnd5e_monster_type
    ADD CONSTRAINT fk_dnd5e_monster_type_createdby FOREIGN KEY (created_by) REFERENCES public."user"(user_id);


--
-- Name: dnd5e_monster fk_dnd5e_monster_world; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dnd5e_monster
    ADD CONSTRAINT fk_dnd5e_monster_world FOREIGN KEY (dnd5e_world_id) REFERENCES public.dnd5e_world(dnd5e_world_id);


--
-- Name: dnd5e_size_category fk_dnd5e_size_category_createdby; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dnd5e_size_category
    ADD CONSTRAINT fk_dnd5e_size_category_createdby FOREIGN KEY (created_by) REFERENCES public."user"(user_id);


--
-- Name: dnd5e_world_assignment fk_dnd5e_world_assignment_role_action; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dnd5e_world_assignment
    ADD CONSTRAINT fk_dnd5e_world_assignment_role_action FOREIGN KEY (role_action) REFERENCES public.role_action(name);


--
-- Name: dnd5e_world_assignment fk_dnd5e_world_assignment_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dnd5e_world_assignment
    ADD CONSTRAINT fk_dnd5e_world_assignment_user FOREIGN KEY (user_id) REFERENCES public."user"(user_id);


--
-- Name: dnd5e_world_assignment fk_dnd5e_world_assignment_world; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dnd5e_world_assignment
    ADD CONSTRAINT fk_dnd5e_world_assignment_world FOREIGN KEY (dnd5e_world_id) REFERENCES public.dnd5e_world(dnd5e_world_id);


--
-- Name: dnd5e_world fk_dnd5e_world_commonaccess; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dnd5e_world
    ADD CONSTRAINT fk_dnd5e_world_commonaccess FOREIGN KEY (common_access) REFERENCES public.common_access(name);


--
-- Name: dnd5e_world fk_dnd5e_world_createdby; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dnd5e_world
    ADD CONSTRAINT fk_dnd5e_world_createdby FOREIGN KEY (created_by) REFERENCES public."user"(user_id);


--
-- Name: dnd5e_world fk_dnd5e_world_derived_from_external_source; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dnd5e_world
    ADD CONSTRAINT fk_dnd5e_world_derived_from_external_source FOREIGN KEY (external_source_id) REFERENCES public.external_source(external_source_id);


--
-- Name: external_source fk_external_source_createdby; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.external_source
    ADD CONSTRAINT fk_external_source_createdby FOREIGN KEY (created_by) REFERENCES public."user"(user_id);


--
-- Name: role_action fk_role_action_subject; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_action
    ADD CONSTRAINT fk_role_action_subject FOREIGN KEY (role_subject) REFERENCES public.role_subject(name);


--
-- Name: user fk_user_commonaccess; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT fk_user_commonaccess FOREIGN KEY (common_access) REFERENCES public.common_access(name);


--
-- Name: user fk_user_createdby; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT fk_user_createdby FOREIGN KEY (created_by) REFERENCES public."user"(user_id);


--
-- PostgreSQL database dump complete
--

