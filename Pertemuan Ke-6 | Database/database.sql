-- Jalankan di MySQL (phpMyAdmin / CLI) atau lewat Docker:
--   docker exec -i mysql-pemweb mysql -uroot -prootpassword < database.sql
-- Nama database harus sama dengan yang ada di DB_DSN (.env).
CREATE DATABASE IF NOT EXISTS pemograman_web
    CHARACTER SET utf8mb4
    COLLATE utf8mb4_unicode_ci;

USE pemograman_web;

CREATE TABLE IF NOT EXISTS produk (
    id         BIGINT(20)   NOT NULL AUTO_INCREMENT,
    nama       VARCHAR(100) NOT NULL,
    harga      BIGINT(20)   NOT NULL DEFAULT 0,
    stok       INT(11)      NOT NULL DEFAULT 0,
    keterangan TEXT         NULL,
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Data contoh
INSERT INTO produk (nama, harga, stok) VALUES
    ('Iphone Duo', 39000000, 10),
    ('Samsung Fold 8', 30000000, 5);
