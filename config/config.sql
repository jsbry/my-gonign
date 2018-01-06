

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


CREATE TABLE `tbls` (
  `tbl_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'テーブルID',
  `project_id` bigint(20) UNSIGNED NOT NULL COMMENT 'プロジェクトID',
  `db_id` bigint(20) UNSIGNED NOT NULL COMMENT 'データベースID',
  `tbl_name` varchar(255) NOT NULL DEFAULT '' COMMENT 'データベース名',
  `tbl_engine` varchar(255) NOT NULL DEFAULT '' COMMENT 'ストレージエンジン',
  `tbl_charset` varchar(255) NOT NULL DEFAULT '' COMMENT '文字コード',
  `tbl_collation` varchar(255) NOT NULL DEFAULT '' COMMENT '照合順序',
  `tbl_comment` text NOT NULL DEFAULT '' COMMENT 'テーブルコメント',
  `is_deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT '削除フラグ',
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `cols` (
  `col_id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY COMMENT 'カラムID',
  `project_id` bigint(20) UNSIGNED NOT NULL COMMENT 'プロジェクトID',
  `db_id` bigint(20) UNSIGNED NOT NULL COMMENT 'データベースID',
  `tbl_id` bigint(20) UNSIGNED NOT NULL COMMENT 'テーブルID',
  `col_physical_name` varchar(255) NOT NULL DEFAULT '' COMMENT '物理名',
  `col_logical_name` varchar(255) NOT NULL DEFAULT '' COMMENT '論理名',
  `col_length` varchar(255) NOT NULL DEFAULT '' COMMENT '桁数',
  `col_default` varchar(255) NOT NULL DEFAULT '' COMMENT 'デフォルト値',
  `col_collation` varchar(255) NOT NULL DEFAULT '' COMMENT '照合順序',
  `col_attribute` varchar(255) NOT NULL DEFAULT '' COMMENT '属性',
  `col_not_null` tinyint(1) NOT NULL DEFAULT 0 COMMENT 'NOT NULL',
  `col_comment` TEXT NOT NULL DEFAULT '' COMMENT 'カラムコメント',
  `is_deleted` tinyint(4) NOT NULL DEFAULT '0' COMMENT '削除フラグ',
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '登録日時',
  `updated` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日時'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
