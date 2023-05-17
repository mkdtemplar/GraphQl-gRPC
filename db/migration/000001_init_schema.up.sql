create table if not exists public."UserRepo"
(
    id   bigserial constraint id primary key,
    name varchar(100)
);

alter table public."UserRepo" owner to graphql;

create table if not exists public.cars
(
    id       bigserial
    constraint cars_id
    primary key,
    car_name varchar(100),
    model    varchar(100),
    user_id  integer
    constraint cars_user_id_fk
    references public."UserRepo"
    );

alter table public.cars
    owner to graphql;