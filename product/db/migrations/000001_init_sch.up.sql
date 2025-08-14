CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    image VARCHAR(255) NOT NULL,
    price FLOAT NOT NULL CHECK (price > 0),
    description TEXT NOT NULL,
    quantity INT NOT NULL CHECK (quantity >= 0),
    stars FLOAT NOT NULL CHECK (stars >= 0 AND stars <= 5) DEFAULT 0,
    created_at TIMESTAMP NOT NULL DEFAULT NOW()
);