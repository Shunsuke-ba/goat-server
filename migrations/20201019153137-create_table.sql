-- +migrate Up

-- ------------------------------------------------
-- Table`users`
-- ------------------------------------------------
CREATE TABLE IF NOT EXISTS `users` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(255) NOT NULL,
    `email` VARCHAR(254) NOT NULL DEFAULT 'NONE',
    `password` VARCHAR(100) NOT NULL DEFAULT 'NONE',
    `social_id` VARCHAR(50) NOT NULL DEFAULT 'NONE',
    `last_login` DATETIME NOT NULL,
    `continue_days` INT NOT NULL DEFAULT 0,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`))
ENGINE = InnoDB;

-- ------------------------------------------------
-- Table`points`
-- ------------------------------------------------
CREATE TABLE IF NOT EXISTS `points` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `match_id` INT NOT NULL,
    `user_id` INT NOT NULL,
    `bet_id` INT NOT NULL,
    `origin_id` INT NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`))
ENGINE = InnoDB;

-- ------------------------------------------------
-- Table`origins`
-- ------------------------------------------------
CREATE TABLE IF NOT EXISTS `origins` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(50) NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`))
ENGINE = InnoDB;

-- ------------------------------------------------
-- Table`scores`
-- ------------------------------------------------
CREATE TABLE IF NOT EXISTS `scores` (
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
-- Table`matches`
-- ------------------------------------------------
CREATE TABLE IF NOT EXISTS `matches` (
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
-- Table`sports`
-- ------------------------------------------------
CREATE TABLE IF NOT EXISTS `sports` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(20) NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`))
ENGINE = InnoDB;

-- ------------------------------------------------
-- Table`bets`
-- ------------------------------------------------
CREATE TABLE IF NOT EXISTS `bets` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `sport_id` INT NOT NULL,
    `match_id` INT NOT NULL,
    `latch` INT NOT NULL,
    `odds` DECIMAL NOT NULL,
    `pay_done` BOOLEAN NOT NULL DEFAULT FALSE,
    `content_id` INT NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`))
ENGINE = InnoDB;

-- ------------------------------------------------
-- Table`bet_contents`
-- ------------------------------------------------
CREATE TABLE IF NOT EXISTS `bet_contents` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(50) NOT NULL,
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`))
ENGINE = InnoDB;
