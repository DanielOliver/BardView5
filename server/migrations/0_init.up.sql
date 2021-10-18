create table "user"
(
    id             bigserial
        constraint user_pk
            primary key,
    uid            char(27)                                                       not null,
    created_by     bigint                                                         null,
    created_at     timestamp without time zone default (now() at time zone 'utc') not null,
    effective_date timestamp without time zone default (now() at time zone 'utc') not null,
    end_date       timestamp without time zone                                    null,
    is_active      boolean                                                        not null default (true),
    email          text                                                           not null,
    name           text                                                           not null,
    tags           text[]                                                         not null,

    CONSTRAINT fk_user_createdby
        FOREIGN KEY (created_by)
            REFERENCES "user" (id)
);

create unique index user_email_uindex on "user" (email);

create table "role_type"
(
    id                           bigserial
        constraint role_type_pk
            primary key,
    created_at                   timestamp without time zone default (now() at time zone 'utc') not null,
    name                         text                                                           not null,
    multiple_assignments_allowed boolean                                                        not null,
    system_managed               boolean                                                        not null
);

create unique index roletype_name_uindex on "role_type" (name);

create table "role"
(
    id             bigserial
        constraint role_pk
            primary key,
    uid            char(27)                                                       not null,
    created_by     bigint                                                         null,
    created_at     timestamp without time zone default (now() at time zone 'utc') not null,
    effective_date timestamp without time zone default (now() at time zone 'utc') not null,
    end_date       timestamp without time zone                                    null,
    is_active      boolean                                                        not null default (true),
    name           text                                                           not null,
    role_type_id   bigint,
    tags           text[]                                                         not null,

    CONSTRAINT fk_role_roletype
        FOREIGN KEY (role_type_id)
            REFERENCES role_type (id),
    CONSTRAINT fk_role_createdby
        FOREIGN KEY (created_by)
            REFERENCES "user" (id)
);

create table "role_assignment"
(
    id             bigserial
        constraint role_assignment_pk
            primary key,
    uid            char(27)                                                       not null,
    created_by     bigint                                                         null,
    created_at     timestamp without time zone default (now() at time zone 'utc') not null,
    effective_date timestamp without time zone default (now() at time zone 'utc') not null,
    end_date       timestamp without time zone                                    null,
    is_active      boolean                                                        not null default (true),
    role_id        bigint                                                         not null,
    user_id        bigint                                                         not null,
    tags           text[]                                                         not null,

    CONSTRAINT fk_roleassignment_role
        FOREIGN KEY (role_id)
            REFERENCES role (id),
    CONSTRAINT fk_roleassignment_user
        FOREIGN KEY (user_id)
            REFERENCES "user" (id),
    CONSTRAINT fk_roleassignment_createdby
        FOREIGN KEY (created_by)
            REFERENCES "user" (id)
);

create table "role_action"
(
    name       text
        constraint role_action_pk primary key,
    created_at timestamp without time zone default (now() at time zone 'utc') not null
);

create table "role_subject"
(
    name       text
        constraint role_subject_pk primary key,
    created_at timestamp without time zone default (now() at time zone 'utc') not null
);

create table "role_permission"
(
    id             bigserial
        constraint role_permission_pk
            primary key,
    uid            char(27)                                                       not null,
    created_by     bigint                                                         null,
    created_at     timestamp without time zone default (now() at time zone 'utc') not null,
    effective_date timestamp without time zone default (now() at time zone 'utc') not null,
    end_date       timestamp without time zone                                    null,
    is_active      boolean                                                        not null default (true),
    role_id        bigint                                                         not null,
    action         text                                                           not null,
    subject        text                                                           not null,
    conditions     jsonb                                                          not null,
    fields         text[]                                                         null,

    CONSTRAINT fk_rolepermission_role
        FOREIGN KEY (role_id)
            REFERENCES role (id),
    CONSTRAINT fk_rolepermission_action
        FOREIGN KEY (action)
            REFERENCES role_action (name),
    CONSTRAINT fk_rolepermission_subject
        FOREIGN KEY (subject)
            REFERENCES role_subject (name),
    CONSTRAINT fk_roleassignment_createdby
        FOREIGN KEY (created_by)
            REFERENCES "user" (id)

);

INSERT INTO role_type (name, multiple_assignments_allowed, system_managed)
SELECT 'User Role', false, true
WHERE NOT EXISTS(SELECT 1 FROM "role_type" WHERE name = 'User Role');

INSERT INTO role_type (name, multiple_assignments_allowed, system_managed)
SELECT 'User Role, Global', true, true
WHERE NOT EXISTS(SELECT 1 FROM "role_type" WHERE name = 'User Role, Global');

CREATE OR REPLACE FUNCTION get_user_role_type_id ()RETURNS bigint
AS $$ SELECT Id FROM "role_type" WHERE name = 'User Role' $$
    LANGUAGE SQL;

CREATE OR REPLACE FUNCTION get_user_role_global_type_id ()RETURNS bigint
AS $$ SELECT Id FROM "role_type" WHERE name = 'User Role, Global' $$
    LANGUAGE SQL;

create unique index role_uid_uindex
    on role (uid);

create unique index roleassignment_uid_uindex
    on role_assignment (uid);

create unique index rolepermission_uid_uindex
    on role_permission (uid);

create unique index user_uid_uindex
    on "user" (uid);

create or replace function generate_ksuid()
    returns  char(27)
    language sql
as $$
select substring(
         replace(to_char(clock_timestamp(),'yyyymmddhh24missus')
                     || (to_char(random()*1e9,'000000000')
                     ),' ',''),1,27)::char(27);
$$;

INSERT INTO role (uid, name, role_type_id, tags)
VALUES (generate_ksuid(), 'User Role, Global', get_user_role_global_type_id(), '{}');

INSERT INTO role_subject (name) VALUES ('user');
INSERT INTO role_action (name) VALUES ('manage');

INSERT INTO role_permission(role_id, action, subject, conditions, uid)
VALUES ((SELECT id FROM "role" WHERE name = 'User Role, Global'), 'manage', 'user', '{ "id:eq": "session_id" }', generate_ksuid());

create or replace function evaluate_access_user(condition jsonb, session_id bigint, user_id bigint)
    returns boolean
    language sql
as
$$
select (condition ->> 'id')::bigint = user_id
           OR ((condition ->> 'id:eq') = 'session_id'
        AND user_id = session_id)
$$;