CREATE TABLE daily_stats (
    chain_id serial PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    staked_buds bigint NOT NULL,
    lost_buds bigint NOT NULL,
);

CREATE TABLE user_data (
    user_address NOT NULL PRIMARY KEY,
    user_name VARCHAR(255) NOT NULL,
    buds_won bigint,
    games_won bigint,
);