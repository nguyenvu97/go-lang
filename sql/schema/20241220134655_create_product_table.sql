-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS product (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_url VARCHAR(255) not null,
    product_category varchar(50) not null,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    status Int not null default 1,
    quantity INT DEFAULT 0,
    product_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
    );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
Drop TABLE If  EXISTS product
-- +goose StatementEnd
