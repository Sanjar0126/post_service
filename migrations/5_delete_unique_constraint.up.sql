do $$
    begin
        alter table branches drop constraint branches_phone_key;
    exception
        when duplicate_column then
            RAISE NOTICE 'Already dropped';
    end $$;