do $$
    begin
        alter table branches drop column tg_chat_id;
    exception
        when others then
            RAISE NOTICE 'Already existed';
    end $$;