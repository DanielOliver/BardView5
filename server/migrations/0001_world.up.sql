create table "dnd5e_world"
(
    dnd5e_world_id      bigint
        constraint dnd5e_world_pk
            primary key,
    created_by          bigint  null,
    created_at          timestamp without time zone default (now() at time zone 'utc') not null,
    version             bigint  not null            default (0),
    is_active           boolean not null            default (true),
    common_access       text    not null,
    user_tags           text[]  not null,
    system_tags         text[]  not null,
    derived_from_world  bigint  null,
    name                text    not null,
    module              text    null,
    description         text    not null,
    external_source_id  bigint  null,
    external_source_key text    null,

    CONSTRAINT fk_dnd5e_world_createdby
        FOREIGN KEY (created_by)
            REFERENCES "user" (user_id),
    CONSTRAINT fk_dnd5e_world_commonaccess
        FOREIGN KEY (common_access)
            REFERENCES common_access (name),
    CONSTRAINT fk_dnd5e_world_derived_from
        FOREIGN KEY (derived_from_world)
            REFERENCES "dnd5e_world" (dnd5e_world_id),

    CONSTRAINT fk_dnd5e_world_derived_from_external_source
        FOREIGN KEY (external_source_id)
            REFERENCES "external_source" (external_source_id)

);

create table "dnd5e_world_assignment"
(
    created_by     bigint null,
    created_at     timestamp without time zone default (now() at time zone 'utc') not null,
    version        bigint not null             default (0),
    user_id        bigint not null,
    dnd5e_world_id bigint not null,
    role_action    text   not null,

    CONSTRAINT dnd5e_world_assignment_pk
        PRIMARY KEY (user_id, dnd5e_world_id),

    CONSTRAINT fk_dnd5e_world_assignment_user
        FOREIGN KEY (user_id)
            REFERENCES "user" (user_id),
    CONSTRAINT fk_dnd5e_world_assignment_world
        FOREIGN KEY (dnd5e_world_id)
            REFERENCES "dnd5e_world" (dnd5e_world_id),
    CONSTRAINT fk_dnd5e_world_assignment_role_action
        FOREIGN KEY (role_action)
            REFERENCES "role_action" (name)
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

create table "dnd5e_monster"
(

    dnd5e_monster_id       bigint
        constraint dnd5e_monster_pk
            primary key,
    created_by             bigint null,
    created_at             timestamp without time zone default (now() at time zone 'utc') not null,
    version                bigint not null             default (0),
    dnd5e_world_id         bigint null,
    name                   text   not null,
    user_tags              text[] not null,
    system_tags            text[] not null,
    monster_type           text   not null,
    alignment              text   not null,
    size_category          text   not null,
    milli_challenge_rating bigint not null,
    languages              text[] not null,
    description            text   not null,


    CONSTRAINT fk_dnd5e_monster_createdby
        FOREIGN KEY (created_by)
            REFERENCES "user" (user_id),
    CONSTRAINT fk_dnd5e_monster_world
        FOREIGN KEY (dnd5e_world_id)
            REFERENCES "dnd5e_world" (dnd5e_world_id),
    CONSTRAINT fk_dnd5e_monster_dnd5e_size_category
        FOREIGN KEY (size_category)
            REFERENCES "dnd5e_size_category" (name),
    CONSTRAINT fk_dnd5e_monster_type
        FOREIGN KEY (monster_type)
            REFERENCES "dnd5e_monster_type" (name)
);

create index dnd5e_monster_world on dnd5e_monster (dnd5e_world_id, dnd5e_monster_id);

create index dnd5e_monster_world_name on dnd5e_monster (dnd5e_world_id, name) include (dnd5e_monster_id);

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

INSERT INTO "dnd5e_size_category" (name, space)
VALUES ('Tiny', '2½ by 2½ ft.'),
       ('Small', '5 by 5 ft.'),
       ('Medium', '5 by 5 ft.'),
       ('Large', '10 by 10 ft.'),
       ('Huge', '15 by 15 ft.'),
       ('Gargantuan', '20 by 20 ft. or larger');

