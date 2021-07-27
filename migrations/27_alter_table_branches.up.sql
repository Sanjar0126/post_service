do $$
    begin
        alter table branches add column tg_chat_id varchar(20);
    exception
        when others then
            RAISE NOTICE 'Already existed';
    end $$;