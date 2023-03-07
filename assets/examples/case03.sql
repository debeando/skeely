CREATE TABLE `properties_email_templates` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `property_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `email_template_type` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `filename` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT 'unique filename without extension',
  `created_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `property_email_template_type_uq` (`property_id`,`email_template_type`),
  KEY `property_idx` (`property_id`),
  KEY `created_at_idx` (`created_at`),
  KEY `email_template_typex_idx` (`email_template_type`)
) ENGINE=InnoDB ADEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci ROW_FORMAT=DYNAMIC