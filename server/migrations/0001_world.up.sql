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

create table "dnd5e_setting_assignment"
(
    created_by       bigint null,
    created_at       timestamp without time zone default (now() at time zone 'utc') not null,
    version          bigint not null             default (0),
    user_id          bigint not null,
    dnd5e_setting_id bigint not null,
    role_action      text   not null,

    CONSTRAINT dnd5e_setting_assignment_pk
        PRIMARY KEY (user_id, dnd5e_setting_id),

    CONSTRAINT fk_dnd5e_setting_assignment_user
        FOREIGN KEY (user_id)
            REFERENCES "user" (user_id),
    CONSTRAINT fk_dnd5e_setting_assignment_setting
        FOREIGN KEY (dnd5e_setting_id)
            REFERENCES "dnd5e_setting" (dnd5e_setting_id),
    CONSTRAINT fk_dnd5e_setting_assignment_role_action
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

INSERT INTO "dnd5e_size_category" (name, space)
VALUES ('Tiny', '2½ by 2½ ft.'),
       ('Small', '5 by 5 ft.'),
       ('Medium', '5 by 5 ft.'),
       ('Large', '10 by 10 ft.'),
       ('Huge', '15 by 15 ft.'),
       ('Gargantuan', '20 by 20 ft. or larger');

