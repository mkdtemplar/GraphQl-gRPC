create table if not exists public.users
(
    id   uuid constraint id primary key,
    name varchar(100)
);

alter table public.users owner to graphql;

create table if not exists public.cars
(
    id       uuid
    constraint cars_id
    primary key,
    car_name varchar(100),
    model    varchar(100),
    user_id  uuid
    constraint cars_user_id_fk
    references public.users
    );

alter table public.cars
    owner to graphql;