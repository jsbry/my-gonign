


CREATE TABLE `projects` (
  `project_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'プロジェクトID',
  `project_name` text NOT NULL COMMENT 'プロジェクト名',
  `is_deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT '削除フラグ',
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `dbs` (
  `db_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'データベースID',
  `project_id` bigint(20) UNSIGNED NOT NULL COMMENT 'プロジェクトID',
  `db_name` varchar(255) NOT NULL DEFAULT '' COMMENT 'データベース名',
  `db_engine` varchar(255) NOT NULL DEFAULT '' COMMENT 'ストレージエンジン',
  `db_charset` varchar(255) NOT NULL DEFAULT '' COMMENT '文字コード',
  `is_deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT '削除フラグ',
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
