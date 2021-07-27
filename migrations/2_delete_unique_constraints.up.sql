do $$
    begin
        alter table branches drop constraint branches_username_key;
        alter table branches drop constraint branches_name_key;
    exception
        when duplicate_column then
            RAISE NOTICE 'Already existed';
    end $$;