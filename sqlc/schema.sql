--
-- PostgreSQL database dump
--

-- Dumped from database version 15.1 (Debian 15.1-1.pgdg110+1)
-- Dumped by pg_dump version 15.1 (Debian 15.1-1.pgdg110+1)

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
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: myusername
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO myusername;

--
-- Name: sessions; Type: TABLE; Schema: public; Owner: myusername
--

CREATE TABLE public.sessions (
    id bigint NOT NULL,
    user_id bigint NOT NULL,
    token character varying NOT NULL,
    created_at timestamp without time zone NOT NULL
);


ALTER TABLE public.sessions OWNER TO myusername;

--
-- Name: sessions_id_seq; Type: SEQUENCE; Schema: public; Owner: myusername
--

ALTER TABLE public.sessions ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.sessions_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: tasks; Type: TABLE; Schema: public; Owner: myusername
--

CREATE TABLE public.tasks (
    id bigint NOT NULL,
    title text NOT NULL,
    description text NOT NULL,
    result text NOT NULL,
    created_at timestamp without time zone NOT NULL,
    created_by bigint NOT NULL,
    done_at timestamp without time zone
);


ALTER TABLE public.tasks OWNER TO myusername;

--
-- Name: tasks_dependencies; Type: TABLE; Schema: public; Owner: myusername
--

CREATE TABLE public.tasks_dependencies (
    id bigint NOT NULL,
    task_id bigint NOT NULL,
    depends_on_id bigint NOT NULL,
    created_at timestamp without time zone NOT NULL
);


ALTER TABLE public.tasks_dependencies OWNER TO myusername;

--
-- Name: tasks_dependencies_id_seq; Type: SEQUENCE; Schema: public; Owner: myusername
--

ALTER TABLE public.tasks_dependencies ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.tasks_dependencies_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: tasks_id_seq; Type: SEQUENCE; Schema: public; Owner: myusername
--

ALTER TABLE public.tasks ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.tasks_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: users; Type: TABLE; Schema: public; Owner: myusername
--

CREATE TABLE public.users (
    id bigint NOT NULL,
    username character varying NOT NULL,
    password character varying,
    email character varying,
    confirmed boolean,
    created_at timestamp without time zone NOT NULL,
    root_task_id bigint
);


ALTER TABLE public.users OWNER TO myusername;

--
-- Name: users_id_seq; Type: SEQUENCE; Schema: public; Owner: myusername
--

ALTER TABLE public.users ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: myusername
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: myusername
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


--
-- Name: sessions sessions_token_key; Type: CONSTRAINT; Schema: public; Owner: myusername
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_token_key UNIQUE (token);


--
-- Name: tasks_dependencies tasks_dependencies_pkey; Type: CONSTRAINT; Schema: public; Owner: myusername
--

ALTER TABLE ONLY public.tasks_dependencies
    ADD CONSTRAINT tasks_dependencies_pkey PRIMARY KEY (id);


--
-- Name: tasks tasks_pkey; Type: CONSTRAINT; Schema: public; Owner: myusername
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: myusername
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: tasks_dependencies_depends_on_id_idx; Type: INDEX; Schema: public; Owner: myusername
--

CREATE INDEX tasks_dependencies_depends_on_id_idx ON public.tasks_dependencies USING btree (depends_on_id);


--
-- Name: tasks_dependencies_task_id_idx; Type: INDEX; Schema: public; Owner: myusername
--

CREATE INDEX tasks_dependencies_task_id_idx ON public.tasks_dependencies USING btree (task_id);


--
-- Name: tasks_search_en_idx; Type: INDEX; Schema: public; Owner: myusername
--

CREATE INDEX tasks_search_en_idx ON public.tasks USING gin (to_tsvector('english'::regconfig, title));


--
-- Name: sessions sessions_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: myusername
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: tasks tasks_created_by_fkey; Type: FK CONSTRAINT; Schema: public; Owner: myusername
--

ALTER TABLE ONLY public.tasks
    ADD CONSTRAINT tasks_created_by_fkey FOREIGN KEY (created_by) REFERENCES public.users(id);


--
-- Name: tasks_dependencies tasks_dependencies_depends_on_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: myusername
--

ALTER TABLE ONLY public.tasks_dependencies
    ADD CONSTRAINT tasks_dependencies_depends_on_id_fkey FOREIGN KEY (depends_on_id) REFERENCES public.tasks(id);


--
-- Name: tasks_dependencies tasks_dependencies_task_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: myusername
--

ALTER TABLE ONLY public.tasks_dependencies
    ADD CONSTRAINT tasks_dependencies_task_id_fkey FOREIGN KEY (task_id) REFERENCES public.tasks(id);


--
-- Name: users users_root_task_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: myusername
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_root_task_id_fkey FOREIGN KEY (root_task_id) REFERENCES public.tasks(id) ON UPDATE CASCADE ON DELETE SET NULL;


--
-- PostgreSQL database dump complete
--

