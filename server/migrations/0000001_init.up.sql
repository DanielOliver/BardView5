create table "common_access"
(
    name       text
        constraint common_access_pk primary key,
    created_at timestamp without time zone default (now() at time zone 'utc') not null
);

INSERT INTO common_access (name)
VALUES ('private'),
       ('anyuser'),
       ('public');

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

INSERT INTO role_subject (name)
VALUES ('dnd5esetting');

create table "role_type"
(
    name       text
        constraint role_type_pk primary key,
    created_at timestamp without time zone default (now() at time zone 'utc') not null
);

INSERT INTO role_type (name)
VALUES ('innate'),
       ('custom');

create table "role_action"
(
    name         text                                                           not null,
    role_subject text                                                           not null,
    created_at   timestamp without time zone default (now() at time zone 'utc') not null,

    constraint role_action_pk primary key (name, role_subject),

    CONSTRAINT fk_role_action_subject
        FOREIGN KEY (role_subject)
            REFERENCES "role_subject" (name)
);

INSERT INTO role_action (name, role_subject)
VALUES ('manage', 'dnd5esetting'),
       ('owner', 'dnd5esetting'),
       ('view', 'dnd5esetting');

create table "role"
(
    role_id          bigint
        constraint role_pk
            primary key,
    name             text                                                           not null,
    created_by       bigint                                                         null,
    created_at       timestamp without time zone default (now() at time zone 'utc') not null,
    role_type        text                                                           not null,
    role_subject     text                                                           not null,
    -- If populated, FK to a specific record.
    scope_id         bigint                                                         null,
    -- FK to role_action
    capabilities     text[]                                                         not null,
    assign_on_create boolean                     default (false)                    not null,
    assign_on_add    boolean                     default (false)                    not null,

    CONSTRAINT fk_role_createdby
        FOREIGN KEY (created_by)
            REFERENCES "user" (user_id),
    CONSTRAINT fk_role_type_fk
        FOREIGN KEY (role_type)
            REFERENCES role_type (name),
    CONSTRAINT fk_role_action_subject
        FOREIGN KEY (role_subject)
            REFERENCES "role_subject" (name)
);

INSERT INTO role (role_id, role_type, role_subject, scope_id, capabilities, name, assign_on_create, assign_on_add)
VALUES (10, 'innate', 'dnd5esetting', null, '{ "view", "manage", "owner" }', 'Owner', true, false),
       (20, 'innate', 'dnd5esetting', null, '{ "view", "manage", "owner" }', 'Admin', false, false),
       (30, 'innate', 'dnd5esetting', null, '{ "view" }', 'Viewer', false, true);

create table "role_assignment"
(
    created_by   bigint null,
    created_at   timestamp without time zone default (now() at time zone 'utc') not null,
    version      bigint not null             default (0),
    user_id      bigint not null,
    role_id      bigint not null,
    -- FK to a specific record, such as a setting. Defaults to this if not specified on Role.
    -- COALESCE(role.scope_id, role_assignment.scope_id)
    scope_id     bigint not null,

    CONSTRAINT role_assignment_pk
        PRIMARY KEY (user_id, scope_id, role_id),

    CONSTRAINT fk_role_assignment_user
        FOREIGN KEY (user_id)
            REFERENCES "user" (user_id),
    CONSTRAINT fk_role_assignment_role
        FOREIGN KEY (role_id)
            REFERENCES "role" (role_id)
);

create table "dnd5e_setting"
(
    dnd5e_setting_id bigint
        constraint dnd5e_setting_pk
            primary key,
    created_by       bigint  null,
    created_at       timestamp without time zone default (now() at time zone 'utc') not null,
    version          bigint  not null            default (0),
    is_active        boolean not null            default (true),
    common_access    text    not null,
    user_tags        text[]  not null,
    system_tags      text[]  not null,
    name             text    not null,
    module           text    null,
    description      text    not null,

    CONSTRAINT fk_dnd5e_setting_createdby
        FOREIGN KEY (created_by)
            REFERENCES "user" (user_id),
    CONSTRAINT fk_dnd5e_setting_commonaccess
        FOREIGN KEY (common_access)
            REFERENCES common_access (name)
);

create table "dnd5e_monster_type"
(
    created_by bigint null,
    created_at timestamp without time zone default (now() at time zone 'utc') not null,
    version    bigint not null             default (0),
    name       text   not null
        constraint dnd5e_monster_type_pk
            primary key,

    CONSTRAINT fk_dnd5e_monster_type_createdby
        FOREIGN KEY (created_by)
            REFERENCES "user" (user_id)
);

create table "dnd5e_size_category"
(
    created_by bigint null,
    created_at timestamp without time zone default (now() at time zone 'utc') not null,
    version    bigint not null             default (0),
    name       text   not null
        constraint dnd5e_size_category_pk
            primary key,
    space      text   not null,

    CONSTRAINT fk_dnd5e_size_category_createdby
        FOREIGN KEY (created_by)
            REFERENCES "user" (user_id)
);

INSERT INTO "dnd5e_size_category" (name, space)
VALUES ('Tiny', '2½ by 2½ ft.'),
       ('Small', '5 by 5 ft.'),
       ('Medium', '5 by 5 ft.'),
       ('Large', '10 by 10 ft.'),
       ('Huge', '15 by 15 ft.'),
       ('Gargantuan', '20 by 20 ft. or larger');

create table "dnd5e_monster"
(
    dnd5e_monster_id       bigint
        constraint dnd5e_monster_pk
            primary key,
    created_by             bigint  null,
    created_at             timestamp without time zone default (now() at time zone 'utc') not null,
    version                bigint  not null            default (0),
    dnd5e_setting_id       bigint  not null,
    name                   text    not null,
    sources                text[]  not null,
    user_tags              text[]  not null,
    languages              text[]  not null,
    environments           text[]  not null,
    is_legendary           boolean not null            default (false),
    is_unique              boolean not null            default (false),
    monster_type           text    null,
    alignment              text    null,
    size_category          text    null,
    milli_challenge_rating bigint  null,
    armor_class            int     null,
    hit_points             int     null,
    description            text    null,
    str_score              int     null,
    dex_score              int     null,
    int_score              int     null,
    wis_score              int     null,
    con_score              int     null,
    cha_score              int     null,

    CONSTRAINT fk_dnd5e_monster_createdby
        FOREIGN KEY (created_by)
            REFERENCES "user" (user_id),
    CONSTRAINT fk_dnd5e_monster_setting
        FOREIGN KEY (dnd5e_setting_id)
            REFERENCES "dnd5e_setting" (dnd5e_setting_id),
    CONSTRAINT fk_dnd5e_monster_dnd5e_size_category
        FOREIGN KEY (size_category)
            REFERENCES "dnd5e_size_category" (name),
    CONSTRAINT fk_dnd5e_monster_type
        FOREIGN KEY (monster_type)
            REFERENCES "dnd5e_monster_type" (name)
);

create index dnd5e_monster_setting on dnd5e_monster (dnd5e_setting_id, dnd5e_monster_id);

create index dnd5e_monster_setting_name on dnd5e_monster (dnd5e_setting_id, name) include (dnd5e_monster_id);

create table "dnd5e_language"
(
    dnd5e_language_id bigint
        constraint dnd5e_language_pk
            primary key,
    created_by        bigint null,
    created_at        timestamp without time zone default (now() at time zone 'utc') not null,
    version           bigint not null             default (0),
    name              text   not null,

    CONSTRAINT fk_dnd5e_language_createdby
        FOREIGN KEY (created_by)
            REFERENCES "user" (user_id)
);
