-- create users table
CREATE TABLE IF NOT EXISTS users (
  id SERIAL PRIMARY KEY,
  username VARCHAR(255) NOT NULL UNIQUE,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  role VARCHAR(255) NOT NULL DEFAULT 'user',
  department VARCHAR(255) NOT NULL DEFAULT '',
  profile_image VARCHAR(255) NOT NULL DEFAULT '',
  is_active BOOLEAN NOT NULL DEFAULT TRUE,
  position VARCHAR(255) NOT NULL DEFAULT '',
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- create inventory products table
CREATE TABLE IF NOT EXISTS products (
  id SERIAL PRIMARY KEY,
  code VARCHAR(255) NOT NULL UNIQUE,
  name VARCHAR(255) NOT NULL,
  brand VARCHAR(255) NOT NULL DEFAULT '',
  supplier VARCHAR(255) NOT NULL DEFAULT '',
  quantity INT NOT NULL DEFAULT 0,
  standard_unit VARCHAR(255) NOT NULL DEFAULT '',
  image VARCHAR(255) NOT NULL DEFAULT '',
  remarks VARCHAR(255) NOT NULL DEFAULT '',
  is_active BOOLEAN NOT NULL DEFAULT TRUE,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- create inventory incoming products table
CREATE TABLE IF NOT EXISTS incoming_products (
  id SERIAL PRIMARY KEY,
  product_id INT REFERENCES products(id) ON DELETE CASCADE,
  status VARCHAR(255) NOT NULL DEFAULT 'pending',
  quantity INT NOT NULL DEFAULT 0,
  length FLOAT NOT NULL DEFAULT 0,
  width FLOAT NOT NULL DEFAULT 0,
  height FLOAT NOT NULL DEFAULT 0,
  unit VARCHAR(255) NOT NULL DEFAULT '',
  -- standard quantity update to product table
  standard_quantity INT NOT NULL DEFAULT 0,
  reference_number VARCHAR(255) NOT NULL DEFAULT '',
  reference_document VARCHAR(255) NOT NULL DEFAULT '',
  cost_price FLOAT NOT NULL DEFAULT 0,
  store_location VARCHAR(255) NOT NULL DEFAULT '',
  store_country VARCHAR(255) NOT NULL DEFAULT '',
  created_by INT REFERENCES users(id) ON DELETE SET NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_by INT REFERENCES users(id) ON DELETE SET NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- create inventory outgoing products table
CREATE TABLE IF NOT EXISTS outgoing_products (
  id SERIAL PRIMARY KEY,
  incoming_product_id INT REFERENCES incoming_products(id) ON DELETE CASCADE,
  product_id INT REFERENCES products(id) ON DELETE CASCADE,
  status VARCHAR(255) NOT NULL DEFAULT 'pending',
  quantity INT NOT NULL DEFAULT 0,
  standard_quantity INT NOT NULL DEFAULT 0,
  reference_number VARCHAR(255) NOT NULL DEFAULT '',
  reference_document VARCHAR(255) NOT NULL DEFAULT '',
  cost_price FLOAT NOT NULL DEFAULT 0,
  created_by INT REFERENCES users(id) ON DELETE SET NULL,
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_by INT REFERENCES users(id) ON DELETE SET NULL,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- create functions to update standard quantity
-- from sum of incoming products and minus sum of outgoing products standard quantity
CREATE OR REPLACE FUNCTION update_standard_quantity()
RETURNS TRIGGER AS $$
BEGIN
  UPDATE products
  SET quantity = (SELECT SUM(standard_quantity) FROM incoming_products WHERE product_id = NEW.product_id) - (SELECT SUM(standard_quantity) FROM outgoing_products WHERE product_id = NEW.product_id)
  WHERE id = NEW.product_id;
  RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER update_standard_quantity_trigger
AFTER INSERT OR UPDATE OR DELETE ON incoming_products
FOR EACH ROW
EXECUTE PROCEDURE update_standard_quantity();

CREATE TRIGGER update_standard_quantity_trigger
AFTER INSERT OR UPDATE OR DELETE ON outgoing_products
FOR EACH ROW
EXECUTE PROCEDURE update_standard_quantity();

