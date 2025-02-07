CREATE TABLE IF NOT EXISTS skus (
    id SERIAL PRIMARY KEY,
    seller_id VARCHAR(100) NOT NULL,
    sku_code VARCHAR(100) UNIQUE NOT NULL,
    name VARCHAR(255) NOT NULL,
    price INT NOT NULL,
    dimensions VARCHAR(100),
    fragile VARCHAR(10),
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- ✅ Add the missing comma here
    deleted_at TIMESTAMP NULL
);
