USE wallet_core;

SET FOREIGN_KEY_CHECKS = 0;

TRUNCATE TABLE gain_done;
TRUNCATE TABLE gain_projection;
TRUNCATE TABLE label;
TRUNCATE TABLE invoice_done;
TRUNCATE TABLE invoice_projection;
TRUNCATE TABLE gain_category;
TRUNCATE TABLE invoice_category;
TRUNCATE TABLE payment_type;

SET FOREIGN_KEY_CHECKS = 1;