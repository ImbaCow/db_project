CREATE TABLE `user` (
  `id` int unsigned NOT NULL AUTO_INCREMENT PRIMARY KEY,
  `login` varchar(255) COLLATE 'utf8_general_ci' NOT NULL,
  `password_hash` varchar(255) COLLATE 'utf8_general_ci' NOT NULL
) ENGINE='InnoDB' COLLATE 'utf8_general_ci';