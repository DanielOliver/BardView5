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
    dnd5e_setting_id bigint NOT NULL,
    name text NOT NULL,
    sources text[] NOT NULL,
    user_tags text[] NOT NULL,
    languages text[] NOT NULL,
    environments text[] NOT NULL,
    is_legendary boolean DEFAULT false NOT NULL,
    is_unique boolean DEFAULT false NOT NULL,
    monster_type text,
    alignment text,
    size_category text,
    milli_challenge_rating bigint,
    armor_class integer,
    hit_points integer,
    description text,
    str_score integer,
    dex_score integer,
    int_score integer,
    wis_score integer,
    con_score integer,
    cha_score integer
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
-- Name: dnd5e_setting; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.dnd5e_setting (
    dnd5e_setting_id bigint NOT NULL,
    created_by bigint,
    created_at timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
    version bigint DEFAULT 0 NOT NULL,
    is_active boolean DEFAULT true NOT NULL,
    common_access text NOT NULL,
    user_tags text[] NOT NULL,
    system_tags text[] NOT NULL,
    name text NOT NULL,
    module text,
    description text NOT NULL
);


ALTER TABLE public.dnd5e_setting OWNER TO postgres;

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
-- Name: role; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.role (
    role_id bigint NOT NULL,
    name text NOT NULL,
    created_by bigint,
    created_at timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
    role_type text NOT NULL,
    role_subject text NOT NULL,
    scope_id bigint,
    capabilities text[] NOT NULL,
    assign_on_create boolean DEFAULT false NOT NULL,
    assign_on_add boolean DEFAULT false NOT NULL
);


ALTER TABLE public.role OWNER TO postgres;

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
-- Name: role_assignment; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.role_assignment (
    created_by bigint,
    created_at timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL,
    version bigint DEFAULT 0 NOT NULL,
    user_id bigint NOT NULL,
    role_id bigint NOT NULL,
    scope_id bigint NOT NULL
);


ALTER TABLE public.role_assignment OWNER TO postgres;

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
    name text NOT NULL,
    created_at timestamp without time zone DEFAULT timezone('utc'::text, now()) NOT NULL
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
-- Name: dnd5e_setting dnd5e_setting_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dnd5e_setting
    ADD CONSTRAINT dnd5e_setting_pk PRIMARY KEY (dnd5e_setting_id);


--
-- Name: dnd5e_size_category dnd5e_size_category_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dnd5e_size_category
    ADD CONSTRAINT dnd5e_size_category_pk PRIMARY KEY (name);


--
-- Name: role_action role_action_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_action
    ADD CONSTRAINT role_action_pk PRIMARY KEY (name, role_subject);


--
-- Name: role_assignment role_assignment_pk; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_assignment
    ADD CONSTRAINT role_assignment_pk PRIMARY KEY (user_id, scope_id, role_id);


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
    ADD CONSTRAINT role_type_pk PRIMARY KEY (name);


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
-- Name: dnd5e_monster_setting; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX dnd5e_monster_setting ON public.dnd5e_monster USING btree (dnd5e_setting_id, dnd5e_monster_id);


--
-- Name: dnd5e_monster_setting_name; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX dnd5e_monster_setting_name ON public.dnd5e_monster USING btree (dnd5e_setting_id, name) INCLUDE (dnd5e_monster_id);


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
-- Name: dnd5e_monster fk_dnd5e_monster_setting; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dnd5e_monster
    ADD CONSTRAINT fk_dnd5e_monster_setting FOREIGN KEY (dnd5e_setting_id) REFERENCES public.dnd5e_setting(dnd5e_setting_id);


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
-- Name: dnd5e_setting fk_dnd5e_setting_commonaccess; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dnd5e_setting
    ADD CONSTRAINT fk_dnd5e_setting_commonaccess FOREIGN KEY (common_access) REFERENCES public.common_access(name);


--
-- Name: dnd5e_setting fk_dnd5e_setting_createdby; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dnd5e_setting
    ADD CONSTRAINT fk_dnd5e_setting_createdby FOREIGN KEY (created_by) REFERENCES public."user"(user_id);


--
-- Name: dnd5e_size_category fk_dnd5e_size_category_createdby; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.dnd5e_size_category
    ADD CONSTRAINT fk_dnd5e_size_category_createdby FOREIGN KEY (created_by) REFERENCES public."user"(user_id);


--
-- Name: role_action fk_role_action_subject; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_action
    ADD CONSTRAINT fk_role_action_subject FOREIGN KEY (role_subject) REFERENCES public.role_subject(name);


--
-- Name: role fk_role_action_subject; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role
    ADD CONSTRAINT fk_role_action_subject FOREIGN KEY (role_subject) REFERENCES public.role_subject(name);


--
-- Name: role_assignment fk_role_assignment_role; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_assignment
    ADD CONSTRAINT fk_role_assignment_role FOREIGN KEY (role_id) REFERENCES public.role(role_id);


--
-- Name: role_assignment fk_role_assignment_user; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role_assignment
    ADD CONSTRAINT fk_role_assignment_user FOREIGN KEY (user_id) REFERENCES public."user"(user_id);


--
-- Name: role fk_role_createdby; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role
    ADD CONSTRAINT fk_role_createdby FOREIGN KEY (created_by) REFERENCES public."user"(user_id);


--
-- Name: role fk_role_type_fk; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.role
    ADD CONSTRAINT fk_role_type_fk FOREIGN KEY (role_type) REFERENCES public.role_type(name);


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

