do $$
    begin
        alter table branch_users add column platform_id uuid;
        alter table customers add column tg_chat_id varchar(100);
    exception
        when others then
            RAISE NOTICE 'Already existed';
    end $$;