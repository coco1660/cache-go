CREATE TABLE tables (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(64) NOT NULL UNIQUE
);

CREATE TABLE cache_items (
     id BIGINT AUTO_INCREMENT PRIMARY KEY,

     table_id BIGINT NOT NULL,

     `key` VARCHAR(128) NOT NULL,
     `value` TEXT NOT NULL,

     expire_at DATETIME NULL,

     created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
     updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

     access_count BIGINT DEFAULT 0,

-- 唯一约束（最关键）
     UNIQUE KEY uk_table_key (table_id, `key`),

-- 普通索引
     KEY idx_table_id (table_id),
     KEY idx_expire (expire_at)
);