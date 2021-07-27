do $$
    begin
        alter table branches add column iiko_terminal_id uuid;
        alter table branches add column fare_id uuid;
    exception
        when others then
            RAISE NOTICE 'Already existed';
    end $$;