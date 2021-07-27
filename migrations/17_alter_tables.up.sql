do $$
    begin
        alter table customers add column bot_language varchar(2) check(bot_language in('uz','ru','en')) default 'ru';
    exception
        when others then
            RAISE NOTICE 'Already existed';
    end $$;