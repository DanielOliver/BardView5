create table "world"
(
    world_id           bigint
        constraint world_pk
            primary key,
    created_by         bigint  null,
    created_at         timestamp without time zone default (now() at time zone 'utc') not null,
    version            bigint  not null            default (0),
    is_active          boolean not null            default (true),
    common_access      text    not null,
    user_tags          text[]  not null,
    system_tags        text[]  not null,
    derived_from_world bigint  null,
    name               text    not null,

    CONSTRAINT fk_world_createdby
        FOREIGN KEY (created_by)
            REFERENCES "user" (user_id),
    CONSTRAINT fk_world_commonaccess
        FOREIGN KEY (common_access)
            REFERENCES common_access (name),
    CONSTRAINT fk_world_derived_from
        FOREIGN KEY (derived_from_world)
            REFERENCES "world" (world_id)

);

create table "monster"
(

    monster_id     bigint
        constraint monster_pk
            primary key,
    created_by     bigint null,
    created_at     timestamp without time zone default (now() at time zone 'utc') not null,
    version        bigint not null             default (0),
    first_world_id bigint null,
    name           text   not null,

    CONSTRAINT fk_monster_createdby
        FOREIGN KEY (created_by)
            REFERENCES "user" (user_id),
    CONSTRAINT monster_world
        FOREIGN KEY (first_world_id)
            REFERENCES "world" (world_id)
);

create table "world_monster"
(

    world_monster_id bigint
        constraint world_monster_pk
            primary key,
    created_by       bigint  null,
    created_at       timestamp without time zone default (now() at time zone 'utc') not null,
    version          bigint  not null            default (0),
    user_tags        text[]  not null,
    system_tags      text[]  not null,
    world_id         bigint  not null,
    monster_id       bigint  not null,
    original_world   boolean not null            default (false),

    CONSTRAINT fk_world_monster_createdby
        FOREIGN KEY (created_by)
            REFERENCES "user" (user_id),
    CONSTRAINT fk_world_monster_world
        FOREIGN KEY (world_id)
            REFERENCES "world" (world_id),
    CONSTRAINT fk_world_monster_monster
        FOREIGN KEY (monster_id)
            REFERENCES "monster" (monster_id)
);

