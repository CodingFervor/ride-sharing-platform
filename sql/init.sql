CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY, phone VARCHAR(20) UNIQUE NOT NULL, password VARCHAR(255) NOT NULL,
    name VARCHAR(100), email VARCHAR(100), avatar VARCHAR(500),
    role VARCHAR(20) DEFAULT 'passenger' CHECK (role IN ('passenger','driver','admin')),
    status VARCHAR(20) DEFAULT 'active', created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE drivers (
    id BIGSERIAL PRIMARY KEY, user_id BIGINT UNIQUE NOT NULL REFERENCES users(id),
    license_no VARCHAR(50) NOT NULL, license_photo VARCHAR(500),
    rating DECIMAL(3,2) DEFAULT 5.00, total_trips INT DEFAULT 0,
    earnings DECIMAL(12,2) DEFAULT 0, is_online BOOLEAN DEFAULT FALSE,
    latitude DECIMAL(10,7), longitude DECIMAL(10,7),
    verified BOOLEAN DEFAULT FALSE, status VARCHAR(20) DEFAULT 'active',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE vehicles (
    id BIGSERIAL PRIMARY KEY, driver_id BIGINT NOT NULL REFERENCES drivers(id),
    plate_number VARCHAR(20) NOT NULL, brand VARCHAR(50), model VARCHAR(50),
    year INT, color VARCHAR(20),
    type VARCHAR(20) DEFAULT 'sedan' CHECK (type IN ('sedan','suv','van','luxury','electric')),
    capacity INT DEFAULT 4, verified BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE rides (
    id BIGSERIAL PRIMARY KEY, passenger_id BIGINT NOT NULL REFERENCES users(id),
    driver_id BIGINT REFERENCES drivers(id), vehicle_id BIGINT REFERENCES vehicles(id),
    pickup_lat DECIMAL(10,7) NOT NULL, pickup_lng DECIMAL(10,7) NOT NULL, pickup_addr VARCHAR(200),
    dest_lat DECIMAL(10,7) NOT NULL, dest_lng DECIMAL(10,7) NOT NULL, dest_addr VARCHAR(200),
    distance DECIMAL(8,2), duration_minutes INT,
    fare DECIMAL(10,2), surge_factor DECIMAL(3,2) DEFAULT 1.00,
    status VARCHAR(20) DEFAULT 'requested' CHECK (status IN ('requested','accepted','arrived','in_progress','completed','cancelled')),
    vehicle_type VARCHAR(20) DEFAULT 'sedan', payment_method VARCHAR(20) DEFAULT 'wallet',
    requested_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, accepted_at TIMESTAMP,
    started_at TIMESTAMP, completed_at TIMESTAMP, created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_rides_passenger ON rides(passenger_id);
CREATE INDEX idx_rides_driver ON rides(driver_id);
CREATE INDEX idx_rides_status ON rides(status);

CREATE TABLE wallets (
    id BIGSERIAL PRIMARY KEY, user_id BIGINT UNIQUE NOT NULL REFERENCES users(id),
    balance DECIMAL(12,2) DEFAULT 0, currency VARCHAR(3) DEFAULT 'CNY',
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE transactions (
    id BIGSERIAL PRIMARY KEY, wallet_id BIGINT NOT NULL REFERENCES wallets(id),
    type VARCHAR(20) NOT NULL CHECK (type IN ('topup','payment','withdrawal','refund','earning')),
    amount DECIMAL(12,2) NOT NULL, description VARCHAR(200),
    ride_id BIGINT REFERENCES rides(id), created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_transactions_wallet ON transactions(wallet_id);

CREATE TABLE ratings (
    id BIGSERIAL PRIMARY KEY, ride_id BIGINT NOT NULL REFERENCES rides(id),
    rater_id BIGINT NOT NULL REFERENCES users(id), ratee_id BIGINT NOT NULL REFERENCES users(id),
    score INT NOT NULL CHECK (score BETWEEN 1 AND 5), comment TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE surge_zones (
    id BIGSERIAL PRIMARY KEY, name VARCHAR(100) NOT NULL,
    factor DECIMAL(3,2) DEFAULT 1.00, latitude DECIMAL(10,7), longitude DECIMAL(10,7),
    radius_km DECIMAL(5,2) DEFAULT 1.00, expires_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users (phone, password, name, role) VALUES
('13800000000', '$2a$10$dummyhash', 'Admin', 'admin');
