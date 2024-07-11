CREATE TABLE daily_stats (
    chain_id SERIAL PRIMARY KEY,
    name VARCHAR(255) UNIQUE NOT NULL,
    staked_buds BIGINT NOT NULL,
    lost_buds BIGINT NOT NULL
);

CREATE TABLE user_data (
    user_address VARCHAR(255) NOT NULL PRIMARY KEY,
    user_name VARCHAR(255) UNIQUE NOT NULL,
    buds_won BIGINT,
    games_won BIGINT
);
