CREATE TABLE IF NOT EXISTS trips (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    start_location VARCHAR(255) NOT NULL,
    end_location VARCHAR(255) NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP,
    driver_info VARCHAR(255),
    license_plate VARCHAR(255),
    status VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS buses (
    id SERIAL PRIMARY KEY,
    trip_id INT REFERENCES trips(id),
    eta TIMESTAMP,
    current_location VARCHAR(255),
    status VARCHAR(50) NOT NULL
);
