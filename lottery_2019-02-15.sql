# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.23)
# Database: lottery
# Generation Time: 2019-02-15 02:45:14 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table lt_blackip
# ------------------------------------------------------------

DROP TABLE IF EXISTS `lt_blackip`;

CREATE TABLE `lt_blackip` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `ip` varchar(50) NOT NULL DEFAULT '' COMMENT 'ip地址',
  `blacktime` int(11) NOT NULL DEFAULT '0' COMMENT '限制到期时间',
  `sys_created` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `sys_updated` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table lt_code
# ------------------------------------------------------------

DROP TABLE IF EXISTS `lt_code`;

CREATE TABLE `lt_code` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `gift_id` int(11) NOT NULL DEFAULT '0' COMMENT '奖品 id，关联 lt_gift 表',
  `code` varchar(255) NOT NULL DEFAULT '' COMMENT '虚拟券编码',
  `sys_created` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `sys_updated` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `sys_status` smallint(11) NOT NULL DEFAULT '0' COMMENT '状态，0-正常，1-删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table lt_gift
# ------------------------------------------------------------

DROP TABLE IF EXISTS `lt_gift`;

CREATE TABLE `lt_gift` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '奖品名称',
  `prize_num` int(11) NOT NULL DEFAULT '0' COMMENT '奖品数量，0-无限量，>0-限量，<0-无奖品',
  `left_num` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '剩余奖品数量',
  `prize_code` varchar(50) NOT NULL DEFAULT '0-0' COMMENT '0-9999表示100%，0-0表示万分之一',
  `prize_time` int(10) NOT NULL DEFAULT '0' COMMENT '发奖周期，天',
  `img` varchar(255) NOT NULL DEFAULT '' COMMENT '奖品图片',
  `displayorder` int(11) NOT NULL DEFAULT '0' COMMENT '位置序号，小的排在前面',
  `gtype` int(11) NOT NULL DEFAULT '0' COMMENT '奖品类型，0-虚拟币，1-虚拟券，2-实物小奖，3-实物大奖',
  `gdata` varchar(255) NOT NULL DEFAULT '' COMMENT '拓展数据，如：虚拟币数量',
  `time_begin` int(11) NOT NULL DEFAULT '0' COMMENT '开始时间',
  `time_end` int(11) NOT NULL DEFAULT '0' COMMENT '结束时间',
  `prize_data` mediumtext NOT NULL COMMENT '发奖计划，[[时间1, 数量1], [时间2, 数量2]]',
  `prize_begin` int(11) NOT NULL DEFAULT '0' COMMENT '发奖周期的开始',
  `prize_end` int(11) NOT NULL DEFAULT '0' COMMENT '发奖周期的结束',
  `sys_status` smallint(11) NOT NULL DEFAULT '0' COMMENT '状态，0-正常，1-删除',
  `sys_created` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `sys_updated` int(11) NOT NULL DEFAULT '0' COMMENT '修改时间',
  `sys_ip` varchar(50) NOT NULL DEFAULT '' COMMENT '操作人 ip',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table lt_result
# ------------------------------------------------------------

DROP TABLE IF EXISTS `lt_result`;

CREATE TABLE `lt_result` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `gift_id` int(11) NOT NULL DEFAULT '0' COMMENT '奖品 id，关联 lt_gift 表',
  `gift_name` varchar(255) NOT NULL DEFAULT '' COMMENT '奖品名称',
  `gift_type` int(11) NOT NULL DEFAULT '0' COMMENT '奖品类型，同 lt_gift.type',
  `uid` int(11) NOT NULL DEFAULT '0' COMMENT '用户 id',
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `prize_code` int(11) NOT NULL DEFAULT '0' COMMENT '抽奖号码（4位的随机数）',
  `gift_data` varchar(255) NOT NULL DEFAULT '' COMMENT '获奖信息',
  `sys_created` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `sys_ip` varchar(50) NOT NULL DEFAULT '' COMMENT '用户抽奖的 ip',
  `sys_status` smallint(11) NOT NULL DEFAULT '0' COMMENT '状态，0-正常，1-删除',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table lt_user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `lt_user`;

CREATE TABLE `lt_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(50) NOT NULL DEFAULT '' COMMENT '用户名',
  `blacktime` int(11) NOT NULL DEFAULT '0' COMMENT '黑名单限制到期时间',
  `realname` varchar(50) NOT NULL DEFAULT '' COMMENT '联系人',
  `mobile` varchar(50) NOT NULL DEFAULT '' COMMENT '手机号',
  `address` varchar(255) NOT NULL DEFAULT '' COMMENT '联系地址',
  `sys_created` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `sys_update` int(11) NOT NULL DEFAULT '0' COMMENT '修改时间',
  `sys_ip` varchar(50) NOT NULL DEFAULT '' COMMENT 'ip 地址',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;



# Dump of table lt_userday
# ------------------------------------------------------------

DROP TABLE IF EXISTS `lt_userday`;

CREATE TABLE `lt_userday` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT '0' COMMENT '用户id',
  `day` int(11) NOT NULL DEFAULT '0' COMMENT '日期',
  `num` int(11) NOT NULL DEFAULT '0' COMMENT '次数',
  `sys_created` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `sys_updated` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;




/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
