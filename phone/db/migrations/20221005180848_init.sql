-- +goose Up

CREATE TABLE phone_numbers (
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    phone text NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO phone_numbers (phone) VALUES
('1234567890'),
('123 456 7891'),
('(123) 456 7892'),
('(123) 456-7893'),
('123-456-7894'),
('123-456-7890'),
('1234567892'),
('(123)456-7892');

-- +goose Down

DROP TABLE phone_numbers;
