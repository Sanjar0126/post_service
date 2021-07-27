do $$
    begin
        alter table system_users alter column phone type varchar;
    exception
        when duplicate_column then
            RAISE NOTICE 'Already existed';
    end $$;
