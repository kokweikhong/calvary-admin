DROP TABLE IF EXISTS outgoing_products, incoming_products, products, users;

DROP FUNCTION IF EXISTS update_standard_quantity();

DROP TRIGGER IF EXISTS update_standard_quantity_trigger ON products;

DROP TRIGGER IF EXISTS update_standard_quantity_trigger ON incoming_products;
