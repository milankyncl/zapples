ALTER TABLE `features` ADD COLUMN `enabled_since` DATETIME AFTER `enabled` DEFAULT NULL;
ALTER TABLE `features` ADD COLUMN `enabled_until` DATETIME AFTER `enabled` DEFAULT NULL;