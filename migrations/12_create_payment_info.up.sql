CREATE TABLE IF NOT EXISTS payme_info (
    shipper_id uuid PRIMARY KEY,
    merchant_id VARCHAR(50) NOT NULL,
    login VARCHAR(50),
    key VARCHAR(100),
    token TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at INTEGER DEFAULT 0
);

CREATE TABLE IF NOT EXISTS click_info (
    shipper_id uuid PRIMARY KEY,
    merchant_id INTEGER NOT NULL,
    service_id INTEGER,
    merchant_user_id INTEGER,
    key VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at INTEGER DEFAULT 0
);