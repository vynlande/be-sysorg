create table example(
    id serial not null,
    uuid uuid not null DEFAULT gen_random_uuid(),
    name text not null,
    age int,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    created_by int,
    updated_by int,

    primary key(id),
    unique(uuid)
);