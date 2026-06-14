-- Creating schemas
CREATE SCHEMA data;
CREATE SCHEMA matches;

-- Creating data tables
CREATE TABLE data.players(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
);

CREATE TABLE data.devices(
    id UUID PRIMARY KEY,
    owner_id UUID REFERENCES players(id),
    device_type VARCHAR(255) NOT NULL,
);

-- Creating matches tables
CREATE TABLE matches.matches(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
);

CREATE TABLE matches.players(
    match_id UUID NOT NULL REFERENCES matches(id),
    player_id UUID NOT NULL REFERENCES players(id),
    team VARCHAR(63),
    PRIMARY KEY (match_id, player_id)
);

CREATE TABLE matches.events(
    match_id UUID NOT NULL REFERENCES matches(id),
    event_time BIGINT NOT NULL,
    event_type VARCHAR(255) NOT NULL,
    active_device_id UUID NOT NULL REFERENCES devices(id),
    passive_device_id UUID REFERENCES devices(id),

    PRIMARY KEY (match_id, event_time, active_device_id)
);
