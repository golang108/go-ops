/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50726
Source Host           : localhost:3306
Source Database       : osp

Target Server Type    : MYSQL
Target Server Version : 50726
File Encoding         : 65001

Date: 2022-05-02 00:13:34
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for agent_info
-- ----------------------------
DROP TABLE IF EXISTS `agent_info`;
CREATE TABLE `agent_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `peerid` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '节点id',
  `name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT 'agent名称',
  `expected_status` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '期望状态',
  `status` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `is_default` int(11) NOT NULL DEFAULT '0',
  `timeout` int(11) DEFAULT NULL COMMENT '启动超时时间',
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  `version` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '版本信息',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Table structure for app
-- ----------------------------
DROP TABLE IF EXISTS `app`;
CREATE TABLE `app` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
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
-- Table structure for check_item
-- ----------------------------
DROP TABLE IF EXISTS `check_item`;
CREATE TABLE `check_item` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `check_item_id` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '检查项名称',
  `desc` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '检查项描述',
  `type` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `content` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '检查项内容',
  `creater` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '创建人',
  `updater` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '更新人',
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Table structure for check_tpl
-- ----------------------------
DROP TABLE IF EXISTS `check_tpl`;
CREATE TABLE `check_tpl` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tid` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '唯一id',
  `name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '名称',
  `description` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '描述',
  `type` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '类型',
  `creater` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '创建人',
  `updater` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '更新人',
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Table structure for check_tpl_detail
-- ----------------------------
DROP TABLE IF EXISTS `check_tpl_detail`;
CREATE TABLE `check_tpl_detail` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tid` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '模板id',
  `cid` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '检查项id',
  `sort` int(11) DEFAULT NULL COMMENT '排序',
  `weight` double DEFAULT NULL COMMENT '权重',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Table structure for cron_task
-- ----------------------------
DROP TABLE IF EXISTS `cron_task`;
CREATE TABLE `cron_task` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cron_uid` varchar(11) COLLATE utf8_unicode_ci DEFAULT NULL,
  `name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '定时任务名称',
  `content` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '定时任务内容',
  `cron_expr` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '定时任务表达式',
  `lastrun_time` datetime DEFAULT NULL COMMENT '最后执行时间',
  `nextrun_time` datetime DEFAULT NULL,
  `status` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '状态 running  stopped',
  `creater` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '创建人',
  `updater` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '更新人',
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  `type` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '任务类型',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Table structure for plugin
-- ----------------------------
DROP TABLE IF EXISTS `plugin`;
CREATE TABLE `plugin` (
  `id` int(11) NOT NULL,
  `uuid` varchar(255) DEFAULT NULL COMMENT '插件uuid',
  `name` varchar(255) DEFAULT NULL COMMENT '插件名',
  `package_name` varchar(255) DEFAULT NULL COMMENT '包名',
  `os` varchar(255) DEFAULT NULL COMMENT '操作系统',
  `arch` varchar(255) DEFAULT NULL COMMENT '架构',
  `md5` varchar(255) DEFAULT NULL COMMENT '包md5名称',
  `creater` varchar(255) DEFAULT NULL COMMENT '创建人',
  `updater` varchar(255) DEFAULT NULL COMMENT '更新人',
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for script
-- ----------------------------
DROP TABLE IF EXISTS `script`;
CREATE TABLE `script` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '命令名称',
  `content` text COLLATE utf8_unicode_ci COMMENT '脚本内容',
  `args` text COLLATE utf8_unicode_ci COMMENT '参数信息',
  `desc` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '描述信息',
  `type` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '脚本类型shell或者powershell',
  `owner_type` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '命令拥有者类型',
  `owner_uid` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '拥有者uid',
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  `script_uid` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=47 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

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
) ENGINE=MyISAM AUTO_INCREMENT=99 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Table structure for task_preset
-- ----------------------------
DROP TABLE IF EXISTS `task_preset`;
CREATE TABLE `task_preset` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '任务名',
  `type` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '任务类型',
  `content` text COLLATE utf8_unicode_ci COMMENT '任务内容',
  `peers` text COLLATE utf8_unicode_ci COMMENT '主机列表',
  `creater` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '创建人',
  `created` datetime DEFAULT NULL COMMENT '创建时间',
  `updated` datetime DEFAULT NULL COMMENT '更新时间',
  `updater` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '更新人',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=37 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
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
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uuid` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '主机uuid',
  `name` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `hostname` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `os_type` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '操作系统类型',
  `os_info` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '操作系统信息',
  `hosttype` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '主机类型',
  `networktype` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '网络类型',
  `private_ip` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `public_ip` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `created` datetime DEFAULT NULL,
  `updated` datetime DEFAULT NULL,
  `creater` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL COMMENT '创建人',
  `address` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  `peer_id` varchar(255) COLLATE utf8_unicode_ci DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
