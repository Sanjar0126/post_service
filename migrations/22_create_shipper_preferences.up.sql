CREATE TABLE IF NOT EXISTS shipper_preferences(
    shipper_id uuid,
    courier_accepts_first boolean default false,
    check_courier_action_radius boolean default false,
    courier_action_radius integer default -1,
    max_delivery_time integer default 60,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (shipper_id)
);

CREATE OR REPLACE FUNCTION create_shipper_preferences()
  RETURNS TRIGGER 
  LANGUAGE PLPGSQL
  AS
$$
BEGIN
    INSERT INTO shipper_preferences(shipper_id) VALUES (NEW.id);
    RETURN NEW;
END;
$$;

CREATE TRIGGER create_shipper_preferences BEFORE INSERT ON shippers
    FOR EACH ROW EXECUTE PROCEDURE create_shipper_preferences();

do
$$
declare
    f record;
begin
    for f in select id from shippers where deleted_at=0
    loop
    insert into shipper_preferences(shipper_id) values (f.id);
    end loop;
end;
$$;
