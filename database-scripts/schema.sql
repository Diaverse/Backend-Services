-- MySQL Script generated by MySQL Workbench
-- Sun Mar 22 15:35:20 2020
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering
​
SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';
​
-- -----------------------------------------------------
-- Schema diaverse
-- -----------------------------------------------------
​
-- -----------------------------------------------------
-- Schema diaverse
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `diaverse` DEFAULT CHARACTER SET utf8 ;
USE `diaverse` ;
​
-- -----------------------------------------------------
-- Table `diaverse`.`users`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `diaverse`.`users` (
  `username` VARCHAR(255) NOT NULL,
  `password` VARCHAR(45) NULL,
  `email` VARCHAR(45) NULL,
  `first_name` VARCHAR(45) NULL,
  `last_name` VARCHAR(45) NULL,
  `user_bio` VARCHAR(255) NULL,
  PRIMARY KEY (`username`))
ENGINE = InnoDB;
​
​
-- -----------------------------------------------------
-- Table `diaverse`.`hardwareToken`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `diaverse`.`hardwareToken` (
  `username` VARCHAR(255) NOT NULL,
  `hardwareToken` VARCHAR(255) CHARACTER SET 'utf8' COLLATE 'utf8_general_mysql500_ci' NULL,
  PRIMARY KEY (`username`),
  UNIQUE INDEX `hardwareToken_UNIQUE` (`hardwareToken` ASC),
  CONSTRAINT `username`
    FOREIGN KEY (`username`)
    REFERENCES `diaverse`.`users` (`username`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
​
​
-- -----------------------------------------------------
-- Table `diaverse`.`scripts`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `diaverse`.`scripts` (
  `script_id` INT NOT NULL,
  `hardwareToken` VARCHAR(255) CHARACTER SET 'utf8' COLLATE 'utf8_general_mysql500_ci' NOT NULL,
  `script` VARCHAR(255) NULL,
  PRIMARY KEY (`script_id`),
  UNIQUE INDEX `script_id_UNIQUE` (`script_id` ASC),
  INDEX `hardwareToken_idx` (`hardwareToken` ASC),
  CONSTRAINT `hardwareToken`
    FOREIGN KEY (`hardwareToken`)
    REFERENCES `diaverse`.`hardwareToken` (`hardwareToken`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;
​
​
SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
