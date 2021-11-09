CREATE TABLE IF NOT EXISTS `questions` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `number` varchar(100) DEFAULT NULL,
  `question` varchar(100) DEFAULT NULL,
  `answer` varchar(100) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;