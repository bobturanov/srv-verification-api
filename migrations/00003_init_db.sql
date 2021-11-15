-- +goose Up
CREATE OR replace FUNCTION create_partition(n            INTEGER,
                                            max_el_limit INTEGER) returns void AS
-- +goose StatementBegin
$$
DECLARE
    i INTEGER := 0;
    parition_name text;
    l_limit INTEGER := 0;
    r_limit INTEGER;
    r_limit_srt text;
    l_limit_srt text;
BEGIN
    IF n < 0 THEN
        RAISE
            EXCEPTION
            'Given n cannot be less than 0';
    END IF;
    FOR i IN 0..n-1
        LOOP
            r_limit := l_limit + max_el_limit;
            parition_name := format( 'verification_%s', 1 + i );
            l_limit_srt := format( '%s', l_limit );
            r_limit_srt := format( '%s', r_limit );
            execute 'create table ' || parition_name || ' ( like verification including all ) inherits (verification)';
            execute 'alter table ' || parition_name || ' add constraint partition_check check (id >= '|| l_limit_srt || ' and id < ' || r_limit_srt || ')';
            l_limit := r_limit;
        END LOOP;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

SELECT create_partition(10, 10000);

CREATE OR replace FUNCTION partition_for_verification() returns TRIGGER AS
-- +goose StatementBegin
$$
DECLARE
    v_parition_name text;
BEGIN
    v_parition_name := format( 'verification_%s', 1 + NEW.id / 10000);
    execute 'INSERT INTO ' || v_parition_name || ' VALUES ( ($1).* )' USING NEW;
    RETURN NULL;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

CREATE TRIGGER partition_verification BEFORE
    INSERT
    ON verification FOR EACH ROW EXECUTE PROCEDURE partition_for_verification();

-- +goose Down
DROP TRIGGER partition_verification ON verification;

CREATE OR replace FUNCTION delete_partition(n INTEGER) returns void AS
-- +goose StatementBegin
$$
DECLARE
    i INTEGER := 0;
    parition_name text;
BEGIN
    IF n < 0 THEN
        RAISE
            EXCEPTION
            'Given n cannot be less than 0';
    END IF;
    FOR i IN 0..n-1
        LOOP
            parition_name := format( 'verification_%s', 1 + i );
            EXECUTE format('drop table %s', parition_name);
        END LOOP;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd

SELECT delete_partition(10);
