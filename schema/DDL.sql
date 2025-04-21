create table if not exists public.asset_class
(
    name                          varchar(255),
    expected_return_in_percentage double precision,
    id                            bigint default nextval('asset_class_id_seq'::regclass) not null
    primary key
    );

alter table public.asset_class
    owner to myuser;

create table if not exists public.allocation_type
(
    name        varchar(255)                                               not null,
    description text,
    id          bigint default nextval('allocation_type_id_seq'::regclass) not null
    primary key,
    min_age     integer,
    max_age     integer
    );

alter table public.allocation_type
    owner to myuser;

create table if not exists public.allocation_type_config
(
    id                       bigserial
    primary key,
    allocation_type_id       bigint           not null
    references public.allocation_type,
    asset_class_id           bigint           not null
    references public.asset_class,
    allocation_in_percentage double precision not null
);

alter table public.allocation_type_config
    owner to myuser;

create table if not exists public.cashflow
(
    id        bigserial
    primary key,
    name      varchar(255)     not null,
    amount    double precision not null,
    is_inflow boolean          not null
    );

alter table public.cashflow
    owner to myuser;

create table if not exists public.liabilities
(
    id           bigserial
    primary key,
    name         varchar(255)     not null,
    amount       double precision not null,
    due_date     date,
    is_long_term boolean
    );

alter table public.liabilities
    owner to myuser;

create table if not exists public.goals
(
    id                     bigserial
    primary key,
    name                   varchar(255)     not null,
    description            text,
    years_left             integer          not null,
    inflation_percentage   double precision default 0.0,
    today_amount           double precision not null,
    allocated_amount       double precision default 0.0,
    sip_step_up_percentage double precision default 0.0
    );

alter table public.goals
    owner to myuser;

create table if not exists public.asset_sub_category
(
    id             bigint       not null
    primary key,
    asset_class_id bigint       not null
    references public.asset_class,
    name           varchar(255) not null,
    priority_order integer      not null
    );

alter table public.asset_sub_category
    owner to myuser;

create table if not exists public.investments
(
    id                    bigserial
    primary key,
    asset_id              bigint           not null
    references public.asset_class,
    amount                double precision not null,
    type                  varchar(10)      not null
    constraint investments_type_check
    check ((type)::text = ANY ((ARRAY ['liquid'::character varying, 'Illiquid'::character varying])::text[])),
    name                  varchar(255)     not null,
    asset_sub_category_id integer
    constraint fk_asset_sub_category
    references public.asset_sub_category
    );

alter table public.investments
    owner to myuser;

