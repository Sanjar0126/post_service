ALTER TABLE payme_info DROP CONSTRAINT payme_info_pkey;
ALTER TABLE click_info DROP CONSTRAINT click_info_pkey;

ALTER TABLE payme_info add column branch_id uuid default null;
ALTER TABLE click_info add column branch_id uuid;

ALTER TABLE payme_info add unique (shipper_id,branch_id,deleted_at);
ALTER TABLE click_info add unique (shipper_id,branch_id,deleted_at);
