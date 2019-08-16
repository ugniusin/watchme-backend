-- +migrate Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `event` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
    `from_date` datetime NOT NULL,
    `to_date` datetime NOT NULL,
    `event_type` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    KEY `event_type` (`event_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- +migrate Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `event`;
