CREATE DATABASE IF NOT EXISTS shorturl;

CREATE TABLE IF NOT EXISTS `urls3` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `full_url` varchar(600) NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `urls_full_url_IDX` (`full_url`)
) ENGINE=InnoDB;