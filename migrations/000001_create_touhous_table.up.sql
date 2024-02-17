CREATE TABLE IF NOT EXISTS touhous (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    name text NOT NULL,
    species text NOT NULL,
    abilities text [] NOT NULL,
    version integer NOT NULL DEFAULT 1
)