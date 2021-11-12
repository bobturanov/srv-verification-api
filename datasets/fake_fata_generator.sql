Create or replace function random_string(length integer) returns text as
$$
declare
    chars text[] := '{A,B,C,D,E,F,G,H,I,J,K,L,M,N,O,P,Q,R,S,T,U,V,W,X,Y,Z}';
    result text := '';
    i integer := 0;
begin
    if length < 0 then
        raise exception 'Given length cannot be less than 0';
    end if;
    for i in 1..length loop
            result := result || chars[1+random()*(array_length(chars, 1)-1)];
        end loop;
    return result;
end;
$$ language plpgsql;

insert into verification (name, created_at, updated_at, is_removed)
select
    random_string( (random() * 4 + 5)::int4),
    now() - '2 years'::interval * random(),
    now() - '2 years'::interval * random(),
    (round(random())::int)::boolean
from generate_series(1, 10000);