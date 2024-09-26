CREATE TABLE IF NOT EXISTS checks (
                                      id                    SERIAL PRIMARY KEY,
                                      brand_name            VARCHAR(255),
    terminal_id           VARCHAR(255),
    amount                DECIMAL(10, 2),
    amount_sbp            DECIMAL(10, 2),
    amount_card           DECIMAL(10, 2),
    amount_cash           DECIMAL(10, 2),
    amount_prepayment_sum DECIMAL(10, 2),
    amount_postpay_sum    DECIMAL(10, 2),
    amount_providing_sum  DECIMAL(10, 2),
    seller                VARCHAR(255),
    fnsn                  VARCHAR(255),
    phone                 VARCHAR(255),
    email                 VARCHAR(255),
    customer              VARCHAR(255),
    customer_inn          VARCHAR(12),
    offline_check_id      VARCHAR(255),
    calculation_type      INT,
    tax_system_type       INT,
    check_type            VARCHAR(50),
    seller_inn            VARCHAR(12),
    status                INT
    );

CREATE INDEX IF NOT EXISTS idx_checks_customer_inn ON checks(customer_inn);

CREATE TABLE IF NOT EXISTS cloud_fiscal (
                                            id                     SERIAL PRIMARY KEY,
                                            created_at             TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                            updated_at             TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                            deleted_at             TIMESTAMP,
                                            check_id               INT,
                                            fiscal_id              VARCHAR(255),
    status                 VARCHAR(50) NOT NULL,
    operation              VARCHAR(50) NOT NULL,
    fiscal_provider_type   VARCHAR(50) NOT NULL,
    fiscal_provider_status VARCHAR(50) NOT NULL,
    error_code             INT,
    error_message          TEXT,
    error_type             VARCHAR(255),
    check_url              TEXT,
    FOREIGN KEY (check_id) REFERENCES checks(id)
    );

CREATE INDEX IF NOT EXISTS idx_cloud_fiscal_check_id ON cloud_fiscal(check_id);

CREATE TABLE IF NOT EXISTS check_product (
                                             id                        SERIAL PRIMARY KEY,
                                             check_id                  INT,
                                             trans_kind                VARCHAR(50),
    total_price               DECIMAL(10, 2),
    item_count                INT,
    name                      VARCHAR(255),
    goods_attribute           VARCHAR(255),
    unit_of_measurement       VARCHAR(50),
    additional_attribute      VARCHAR(255),
    manufacturer_country_code VARCHAR(3),
    customs_declaration_number VARCHAR(255),
    supplier_inn              VARCHAR(12),
    supplier_info             TEXT,
    agent_type                VARCHAR(50),
    agent_info                TEXT,
    tax_id                    INT,
    FOREIGN KEY (check_id) REFERENCES checks(id) ON DELETE CASCADE ON UPDATE CASCADE
    );

CREATE INDEX IF NOT EXISTS idx_check_product_check_id ON check_product(check_id);

CREATE TABLE IF NOT EXISTS cloud_fiscal_cred (
                                                 id                     SERIAL PRIMARY KEY,
                                                 created_at             TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                                 updated_at             TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                                 deleted_at             TIMESTAMP,
                                                 fiscal_service         VARCHAR(50),
    group_code             VARCHAR(255),
    login                  VARCHAR(255),
    password               VARCHAR(255),
    company_name           VARCHAR(255),
    company_inn            VARCHAR(12),
    company_sno            VARCHAR(50),
    company_payment_address VARCHAR(255),
    company_email          VARCHAR(255)
    );

CREATE INDEX IF NOT EXISTS idx_cloud_fiscal_cred_company_inn ON cloud_fiscal_cred(company_inn);

CREATE TABLE IF NOT EXISTS terminals (
                                         id                      SERIAL PRIMARY KEY,
                                         created_at              TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                         updated_at              TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                         deleted_at              TIMESTAMP,
                                         serial_number           VARCHAR(255) UNIQUE NOT NULL,
    terminal_id             VARCHAR(255) UNIQUE NOT NULL,
    is_cloud_fiscal         BOOLEAN NOT NULL DEFAULT FALSE,
    cloud_fiscal_cred_id    INT,
    FOREIGN KEY (cloud_fiscal_cred_id) REFERENCES cloud_fiscal_cred(id) ON DELETE CASCADE ON UPDATE CASCADE
    );

CREATE INDEX IF NOT EXISTS idx_terminals_cloud_fiscal_cred_id ON terminals(cloud_fiscal_cred_id);

CREATE TABLE IF NOT EXISTS tokens (
                                      id SERIAL PRIMARY KEY,
                                      service VARCHAR(50) NOT NULL,
    login VARCHAR(255) UNIQUE NOT NULL,
    token VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP NOT NULL,
    error_code INT,
    error_message TEXT,
    error_type VARCHAR(255)
    );
CREATE INDEX IF NOT EXISTS idx_tokens_login ON tokens(login);