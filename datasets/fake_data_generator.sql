CREATE OR replace FUNCTION random_string(length INTEGER) returns text
AS
$$
DECLARE
    chars text[] := '{A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z}';
    result text := '';
    i INTEGER := 0;
BEGIN
    IF length < 0 THEN
        RAISE
            EXCEPTION
            'Given length cannot be less than 0';
    END IF;
    FOR i IN 1..length
        LOOP
            result := result
                || chars[1+random()*(array_length(chars, 1)-1)];
        END LOOP;
    RETURN result;
END;
$$ LANGUAGE plpgsql;

INSERT INTO verification
(
    name,
    created_at,
    updated_at,
    is_removed
)
SELECT random_string( (random() * 4 + 5)::int4),
       now() - '2 years'::interval * random(),
       now() - '2 years'::interval * random(),
       (round(random())::INT)::BOOLEAN
FROM   generate_series(1, 10000);


do
$do$
    declare
        i int;
        type_str event_type;
        status_str event_status;
    begin
        for  i in 1..10000
            loop
                type_str := (select (array['CREATED', 'UPDATED', 'REMOVED'])[floor(random() * 3 + 1)]);
                status_str := (select (array['DEFERRED', 'PROCESSED'])[floor(random() * 2 + 1)]);
                insert into verification_events (
                    verification_id,
                    type,
                    status,
                    payload,
                    updated_at
                ) values (i,
                          type_str,
                          status_str,
                          '{"verification_id": 1, "name": "SQL"}',
                          now() - '2 years'::interval * random());
            end loop;
    end;
$do$