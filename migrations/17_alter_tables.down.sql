do $$
    begin
        alter table customers drop column bot_language;
    exception
        when others then
            RAISE NOTICE 'Already existed';
    end $$;s