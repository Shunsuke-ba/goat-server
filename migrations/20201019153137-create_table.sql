-- +migrate Up

-- ------------------------------------------------
-- Table`user`
-- ------------------------------------------------
CREATE TABLE IF NOT EXISTS `user` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `email` VARCHAR(254) NOT NULL DEFAULT 'NONE',
    `password` VARCHAR(100) NOT NULL DEFAULT 'NONE',
    `social_id` VARCHAR(50) NOT NULL DEFAULT 'NONE',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    `deleted_at` DATETIME NOT NULL DEFAULT 'NONE',
    PRIMARY KEY (`id`))
ENGINE = InnoDB;

-- ------------------------------------------------
-- Table`login`
-- ------------------------------------------------
CREATE TABLE IF NOT EXISTS `login` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `last_login` DATETIME NOT NULL,
    `continuation` INT NOT NULL DEFAULT 0,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`))
ENGINE = InnoDB;

-- ------------------------------------------------
-- Table`profile`
-- ------------------------------------------------
CREATE TABLE IF NOT EXISTS `profile` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `user_id` INT NOT NULL,
    `name` VARCHAR(255) NOT NULL,
    `avatar` VARCHAR(255) NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`))
ENGINE = InnoDB;

-- ------------------------------------------------
-- Table`point`
-- ------------------------------------------------
CREATE TABLE IF NOT EXISTS `point` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `match_id` INT NOT NULL,
    `user_id` INT NOT NULL,
    `bet_id` INT NOT NULL,
    `origin_id` INT NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`))
ENGINE = InnoDB;

-- originはlog関数作成で破棄

-- ------------------------------------------------
-- Table`origin`
-- ------------------------------------------------
CREATE TABLE IF NOT EXISTS `origin` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(50) NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`))
ENGINE = InnoDB;

-- scoreとmatchはAPIのResponse内容で変更確定

-- ------------------------------------------------
-- Table`score`
-- ------------------------------------------------
CREATE TABLE IF NOT EXISTS `score` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `sport_id` INT NOT NULL,
    `match_id` INT NOT NULL,
    `home_score` INT NOT NULL,
    `away_score` INT NOT NULL,
    `period` VARCHAR(10) NOT NULL,
    `ended` BOOLEAN NOT NULL DEFAULT FALSE,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`))
ENGINE = InnoDB;

-- ------------------------------------------------
-- Table`match`
-- ------------------------------------------------
CREATE TABLE IF NOT EXISTS `match` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `league` VARCHAR(10) NOT NULL,
    `tournament` VARCHAR(30) NOT NULL,
    `home_team` VARCHAR(30) NOT NULL,
    `away_team` VARCHAR(30) NOT NULL,
    `venue` VARCHAR(30) NOT NULL,
    `country` VARCHAR(30) NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`))
ENGINE = InnoDB;

-- ------------------------------------------------
-- Table`sport`
-- ------------------------------------------------
CREATE TABLE IF NOT EXISTS `sport` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(20) NOT NULL,
    PRIMARY KEY (`id`))
ENGINE = InnoDB;

-- ------------------------------------------------
-- Table`bet`
-- ------------------------------------------------
CREATE TABLE IF NOT EXISTS `bet` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `sport_id` INT NOT NULL,
    `match_id` INT NOT NULL,
    `latch` INT NOT NULL,
    `odds` DECIMAL NOT NULL,
    `pay_done` BOOLEAN NOT NULL DEFAULT FALSE,
    -- `content_id` INT NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `deleted_at` DATETIME NULL DEFAULT NULL,
    PRIMARY KEY (`id`))
ENGINE = InnoDB;

-- 保留

-- ------------------------------------------------
-- Table`bet_contents`
-- ------------------------------------------------
-- CREATE TABLE IF NOT EXISTS `bet_contents` (
--     `id` INT NOT NULL AUTO_INCREMENT,
--     `name` VARCHAR(50) NOT NULL,
--     `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
--     PRIMARY KEY (`id`))
-- ENGINE = InnoDB;
