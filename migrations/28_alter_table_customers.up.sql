do $$
    begin
        alter table customers add column platform_id uuid;
        alter table customers add column date_of_birth date;
    exception
        when others then
            RAISE NOTICE 'Already existed';
    end $$;