CREATE TABLE inventories (
    inv_id VARCHAR(100) PRIMARY KEY,
    hub_id VARCHAR(100) NOT NULL,
    sku_id VARCHAR(100) NOT NULL,
    quantity INT NOT NULL
    
);
