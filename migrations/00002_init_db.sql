-- +goose Up
CREATE TYPE event_type AS ENUM ('CREATED', 'UPDATED', 'REMOVED');
CREATE TYPE event_status AS ENUM ('DEFERRED', 'PROCESSED');

CREATE TABLE verification_events (
                                event_id BIGSERIAL PRIMARY KEY,
                                verification_id BIGSERIAL NOT NULL,
                                type event_type NOT NULL,
                                status event_status NOT NULL,
                                payload JSONB NOT NULL,
                                updated_at TIMESTAMP NOT NULL,
                                CONSTRAINT fk_verification FOREIGN KEY (verification_id) REFERENCES verification (id) ON DELETE CASCADE

);

-- +goose Down
DROP TYPE event_status;
DROP TYPE event_type;
DROP TABLE verification_events;
