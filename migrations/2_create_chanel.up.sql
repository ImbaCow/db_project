CREATE TABLE `channel` (
  `id` int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `name` varchar(255) COLLATE 'utf8_general_ci' NOT NULL,
  `display_name` varchar(255) COLLATE 'utf8_general_ci' NOT NULL
) ENGINE='InnoDB' COLLATE 'utf8_general_ci';