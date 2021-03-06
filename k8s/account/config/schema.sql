CREATE TABLE `users` (
  `id` BIGINT UNSIGNED NOT NULL,
  `name` VARCHAR(32) CHARSET ascii COLLATE ascii_bin NOT NULL,
  `password_hash` VARCHAR(254) CHARSET latin1 COLLATE latin1_bin NOT NULL,

  `created_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
  `updated_at` DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),

  PRIMARY KEY (`id`),
  UNIQUE KEY (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
