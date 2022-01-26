create table "common_access"
(
    name       text
        constraint common_access_pk primary key,
    created_at timestamp without time zone default (now() at time zone 'utc') not null
);

create table "user"
(
    user_id        bigint
        constraint user_pk
            primary key,
    uuid           uuid                        not null,
    created_by     bigint                      null,
    created_at     timestamp without time zone          default (now() at time zone 'utc') not null,
    version        bigint                      not null default (0),
    effective_date timestamp without time zone          default (now() at time zone 'utc') not null,
    end_date       timestamp without time zone null,
    is_active      boolean                     not null default (true),
    common_access  text                        not null,
    email          text                        not null,
    name           text                        not null,
    user_tags      text[]                      not null,
    system_tags    text[]                      not null,

    CONSTRAINT fk_user_createdby
        FOREIGN KEY (created_by)
            REFERENCES "user" (user_id),
    CONSTRAINT fk_user_commonaccess
        FOREIGN KEY (common_access)
            REFERENCES common_access (name)
);

create unique index user_email_uindex on "user" (email);

create unique index user_uuid_uindex on "user" (uuid);

create table "role_subject"
(
    name       text
        constraint role_subject_pk primary key,
    created_at timestamp without time zone default (now() at time zone 'utc') not null
);

create table "role_action"
(
    name         text
        constraint role_action_pk primary key,
    role_subject text                                                           not null,
    created_at   timestamp without time zone default (now() at time zone 'utc') not null,

    CONSTRAINT fk_role_action_subject
        FOREIGN KEY (role_subject)
            REFERENCES "role_subject" (name)
);

INSERT INTO role_subject (name)
VALUES ('dnd5esetting');

INSERT INTO role_action (name, role_subject)
VALUES ('manage', 'dnd5esetting'),
       ('owner', 'dnd5esetting'),
       ('view', 'dnd5esetting');

INSERT INTO common_access (name)
VALUES ('private'),
       ('anyuser'),
       ('public');

create table "external_source"
(
    external_source_id      bigint
        constraint external_source_pk
            primary key,
    created_by              bigint null,
    created_at              timestamp without time zone default (now() at time zone 'utc') not null,
    version                 bigint not null             default (0),
    external_source_key     text   not null,
    external_source_version text   not null,
    user_tags               text[] not null,
    system_tags             text[] not null,
    name                    text   not null,

    CONSTRAINT fk_external_source_createdby
        FOREIGN KEY (created_by)
            REFERENCES "user" (user_id)
);
