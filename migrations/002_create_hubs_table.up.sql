CREATE TABLE  hubs (
    hub_id VARCHAR(100) NOT NULL UNIQUE,
    tenant_id VARCHAR(100) NOT NULL,
    manager_email VARCHAR(255) NOT NULL,
    contact_no VARCHAR(50) NOT NULL,
    hub_name VARCHAR(255) NOT NULL,
    location VARCHAR(255),
    PRIMARY KEY (hub_id)
);