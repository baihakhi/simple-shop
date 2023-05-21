DROP TABLE IF EXISTS transactions;
DROP TABLE IF EXISTS carts;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS locations;

DO $$
DECLARE
    t text;
BEGIN
    FOR t IN
        SELECT routine_name FROM information_schema.routines
        WHERE routine_name LIKE '%timestamp_func'
        AND routine_type = 'FUNCTION' AND routine_schema = 'public'
   LOOP
     EXECUTE 'DROP FUNCTION IF EXISTS ' || t;
   END LOOP;
END;
$$ language 'plpgsql';
