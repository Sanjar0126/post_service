do $$
    begin
        alter table shippers add column crm varchar(10) check(crm in('jowi', 'iiko', 'none')) default 'none';
    exception
        when others then
            RAISE NOTICE 'Already existed';
    end $$;