CREATE TABLE `test_table` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT "comment id",
  `title` varchar(100) NOT NULL COMMENT "comment title",
  `author` varchar(40) NOT NULL COMMENT "comment author",
  `created_at` datetime NOT NULL COMMENT "comment created_at",
  `updated_at` datetime NOT NULL COMMENT "comment updated_at",
  `author_id` bigint NOT NULL COMMENT "comment author_id",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;

CREATE TABLE `test_table2` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT "comment id2",
    `title` varchar(100) NOT NULL COMMENT "comment title2",
    `author` varchar(40) NOT NULL COMMENT "comment author2",
    `created_at` datetime NOT NULL COMMENT "comment created_at2",
    `updated_at` datetime NOT NULL COMMENT "comment updated_at2",
    `author_id` bigint NOT NULL COMMENT "comment author_id2",
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3;