do $$
    begin
        alter table branch_users drop column platform_id;
        alter table customers drop column tg_chat_id;
    exception
        when others then
            RAISE NOTICE 'Already dropped';
    end $$;