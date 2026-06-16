-- Creating schemas
CREATE SCHEMA data;
CREATE SCHEMA matches;

-- Enums
CREATE TYPE data.device_type AS ENUM ('phaser', 'vest', 'grenade', 'stim');
CREATE TYPE matches.match_type AS ENUM ('deathmatch', 'team_deathmatch', 'capture_the_flag');
CREATE TYPE matches.event_type AS ENUM ('shot', 'hit', 'heal', 'revive');

-- Creating data tables
CREATE TABLE data.players(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(60) NOT NULL
);

CREATE TABLE data.devices(
    id UUID PRIMARY KEY,
    owner_id UUID REFERENCES data.players(id),
    device_type data.device_type NOT NULL
);

-- Creating matches tables
CREATE TABLE matches.matches(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    started_at BIGINT NOT NULL,
    ended_at BIGINT,
    winner_id UUID REFERENCES data.players(id),
    match_type matches.match_type NOT NULL,
    match_name VARCHAR(60) NOT NULL,
    winner_team VARCHAR(60)
);

CREATE TABLE matches.players(
    match_id UUID NOT NULL REFERENCES matches.matches(id),
    player_id UUID NOT NULL REFERENCES data.players(id),
    team VARCHAR(60),
    PRIMARY KEY (match_id, player_id)
);

CREATE TABLE matches.events(
    match_id UUID NOT NULL REFERENCES matches.matches(id),
    event_time BIGINT NOT NULL,
    active_device_id UUID NOT NULL REFERENCES data.devices(id),
    passive_device_id UUID REFERENCES data.devices(id),
    event_type matches.event_type NOT NULL,

    PRIMARY KEY (match_id, event_time, active_device_id)
);
