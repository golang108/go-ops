/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50726
Source Host           : localhost:3306
Source Database       : osp

Target Server Type    : MYSQL
Target Server Version : 50726
File Encoding         : 65001

Date: 2022-04-10 13:22:50
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for app
-- ----------------------------
DROP TABLE IF EXISTS `app`;
CREATE TABLE `app` (
  `id` int(11) NOT NULL,
  `appid` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `api_key` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `sec_key` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `owner` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '应用名',
  `status` tinyint(4) DEFAULT NULL COMMENT '1启用 0 禁用',
  `owner_uid` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '拥有者uid',
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Table structure for script
-- ----------------------------
DROP TABLE IF EXISTS `script`;
CREATE TABLE `script` (
  `id` int(11) NOT NULL,
  `name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '命令名称',
  `content` text COLLATE utf8_unicode_ci COMMENT '脚本内容',
  `args` text COLLATE utf8_unicode_ci COMMENT '参数信息',
  `desc` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '描述信息',
  `type` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '脚本类型shell或者powershell',
  `owner_type` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '命令拥有者类型',
  `owner_uid` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '拥有者uid',
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Table structure for task
-- ----------------------------
DROP TABLE IF EXISTS `task`;
CREATE TABLE `task` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `task_id` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '任务id',
  `type` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '任务类型',
  `content` text COLLATE utf8_unicode_ci COMMENT '任务内容',
  `name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '任务名称',
  `reqid` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '请求id',
  `parent_id` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `status` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  `creater` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '创建人',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=9 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL,
  `uid` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `username` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `passwd` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `email` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  `phone` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Table structure for vm
-- ----------------------------
DROP TABLE IF EXISTS `vm`;
CREATE TABLE `vm` (
  `id` int(11) NOT NULL,
  `uuid` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '主机uuid',
  `name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `hostname` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `az` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '可用区',
  `os_type` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '操作系统类型',
  `os_info` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '操作系统信息',
  `hosttype` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '主机类型',
  `networktype` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '网络类型',
  `private_ip` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `public_ip` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  `creater` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '创建人',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
