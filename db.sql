CREATE DATABASE IF NOT EXISTS shorturl;

CREATE TABLE IF NOT EXISTS `urls` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `full_url` varchar(1000) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `urls_full_url_IDX` (`full_url`)
) ENGINE=InnoDB;