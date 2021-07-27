do $$
    begin
        alter table customers drop column platform_id;
        alter table customers drop column date_of_birth;
    exception
        when others then
            RAISE NOTICE 'Already existed';
    end $$;