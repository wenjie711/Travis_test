-- MySQL Script generated by MySQL Workbench
-- 06/01/15 15:28:44
-- Model: New Model    Version: 1.0
-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='TRADITIONAL,ALLOW_INVALID_DATES';

-- -----------------------------------------------------
-- Schema mydb
-- -----------------------------------------------------
-- -----------------------------------------------------
-- Schema qiniu-share
-- -----------------------------------------------------
DROP SCHEMA IF EXISTS `qiniu-share` ;

-- -----------------------------------------------------
-- Schema qiniu-share
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `qiniu-share` DEFAULT CHARACTER SET utf8 ;
USE `qiniu-share` ;

-- -----------------------------------------------------
-- Table `qiniu-share`.`article`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `qiniu-share`.`article` ;

CREATE TABLE IF NOT EXISTS `qiniu-share`.`article` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '文章id',
  `cid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '栏目id category表相对应id',
  `aid` int(10) unsigned NOT NULL COMMENT '管理员ID 对应发表该文章的管理员',
  `title` varchar(80) NOT NULL DEFAULT '' COMMENT '标题',
  `thumb` varchar(100) NOT NULL DEFAULT '' COMMENT '图片地址',
  `content` text NOT NULL COMMENT '内容',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否发布 0:不发布 1:发布',
  `hits` int(10) NOT NULL DEFAULT '0' COMMENT '点击数',
  `iscomment` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否允许评论',
  `createtime` datetime NOT NULL COMMENT '发布时间',
  `updatetime` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `cid` (`cid`,`id`),
  KEY `istop` (`id`),
  KEY `aid` (`aid`,`id`)
)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8
COMMENT = '内容表';


-- -----------------------------------------------------
-- Table `qiniu-share`.`logs`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `qiniu-share`.`logs` ;

CREATE TABLE IF NOT EXISTS `qiniu-share`.`logs` (
  `id` INT(10) UNSIGNED NOT NULL COMMENT '日志id',
  `utype` TINYINT(1) NOT NULL COMMENT '操作者所属：1.管理员 2.用户',
  `uid` INT(10) NOT NULL COMMENT '操作人id',
  `module` VARCHAR(50) CHARACTER SET 'utf8' COLLATE 'utf8_unicode_ci' NOT NULL DEFAULT '' COMMENT '操作的模块',
  `url` VARCHAR(100) CHARACTER SET 'utf8' COLLATE 'utf8_unicode_ci' NOT NULL DEFAULT '' COMMENT '操作对应的url',
  `action` VARCHAR(100) CHARACTER SET 'utf8' COLLATE 'utf8_unicode_ci' NOT NULL DEFAULT '' COMMENT '操作对应的action',
  `ip` VARCHAR(15) CHARACTER SET 'utf8' COLLATE 'utf8_unicode_ci' NOT NULL DEFAULT '0.0.0.0' COMMENT '操作者所在IP',
  `desc` TEXT CHARACTER SET 'utf8' COLLATE 'utf8_unicode_ci' NOT NULL COMMENT '操作说明',
  `createtime` DATETIME NOT NULL COMMENT '操作时间',
  PRIMARY KEY (`id`),
  INDEX `uid` (`uid` ASC, `module` ASC))
ENGINE = InnoDB
AUTO_INCREMENT = 2
DEFAULT CHARACTER SET = utf8
COLLATE = utf8_unicode_ci
COMMENT = '后台操作日志表';


-- -----------------------------------------------------
-- Table `qiniu-share`.`attachment`
-- -----------------------------------------------------
DROP TABLE IF EXISTS `qiniu-share`.`attachment` ;

CREATE TABLE IF NOT EXISTS `qiniu-share`.`attachment` (
  `id` INT(10) UNSIGNED NOT NULL COMMENT '附件id',
  `aid` INT(10) UNSIGNED NOT NULL COMMENT '关联的文章id',
  `name` VARCHAR(45) NOT NULL COMMENT '显示的文件名',
  `url` VARCHAR(45) NOT NULL COMMENT '对应的存储url',
  `download` INT(10) NOT NULL DEFAULT 0 COMMENT '下载次数',
  `createtime` DATETIME NOT NULL COMMENT '创建时间',
  `updatetime` DATETIME NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`),
  INDEX `article_id_idx` (`aid` ASC),
  CONSTRAINT `article_id`
    FOREIGN KEY (`aid`)
    REFERENCES `qiniu-share`.`article` (`id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB;


SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
