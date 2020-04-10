
CREATE TABLE IF NOT EXISTS `google_doodles`(
    `id` int unsigned NOT NULL AUTO_INCREMEN,
    `share_text` varchar(255) NOT NULL,
    `name` varchar(64) NOT NULL,
    `title` varchar(64) NOT NULL,
    `width` varchar(12) NOT NULL ,
    `height` varchar(12) NOT NULL ,
    `date` TIMESTAMP NOT NULL ,
    `url_high` varchar(128) NOT NULL,
    `url` varchar(128) NOT NULL ,
    `alternate_url` varchar(1024) NOT NULL,
    PRIMARY KEY (`id`),
    KEY `idx_title` (`title`),
    KEY `idx_date` (`date`)
)   ENGINE=InnoDB CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci
