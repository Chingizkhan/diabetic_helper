create table if not exists sugar_levels(
    user_id     int,
    "value"     varchar(4),
    created_at  timestamp(6),
    updated_at  timestamp(6)
);

-- create index if not exists "sugar_level_name_value" on sugar_levels(user_name, "value");