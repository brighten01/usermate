/*
 Navicat Premium Dump SQL

 Source Server         : 本地环境
 Source Server Type    : MySQL
 Source Server Version : 90200 (9.2.0)
 Source Host           : localhost:3306
 Source Schema         : mate

 Target Server Type    : MySQL
 Target Server Version : 90200 (9.2.0)
 File Encoding         : 65001

 Date: 15/04/2025 16:34:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for order_detail
-- ----------------------------
DROP TABLE IF EXISTS `order_detail`;
CREATE TABLE `order_detail` (
  `id` int NOT NULL AUTO_INCREMENT,
  `gender` tinyint DEFAULT NULL COMMENT '1 男 2 女',
  `level` int DEFAULT NULL COMMENT '等级',
  `duration` tinyint DEFAULT NULL COMMENT '时长',
  `service_category_name` varchar(255) COLLATE utf8mb4_general_ci NOT NULL COMMENT '服务名称',
  `service_category_id` int NOT NULL COMMENT '服务id',
  `wechat` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '微信等联系方式',
  `note` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '备注',
  `createdAt` datetime DEFAULT NULL,
  `updatedAt` datetime DEFAULT NULL,
  `order_id` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '主订单id',
  PRIMARY KEY (`id`,`order_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='订单详情\n';

-- ----------------------------
-- Records of order_detail
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for orders
-- ----------------------------
DROP TABLE IF EXISTS `orders`;
CREATE TABLE `orders` (
  `id` int NOT NULL,
  `order_id` varchar(100) COLLATE utf8mb4_general_ci NOT NULL COMMENT '订单id',
  `uid` int NOT NULL COMMENT '用户id',
  `service_category` int DEFAULT NULL COMMENT '订单类型 各种服务类型',
  `start_time` datetime DEFAULT NULL COMMENT '开始时间',
  `end_time` datetime DEFAULT NULL COMMENT '结束时间',
  `status` tinyint DEFAULT NULL COMMENT '订单状态 1 开始 2 进行中 3 已完成 4 退单 5 取消 6 关闭\n',
  `amount` decimal(10,2) DEFAULT NULL COMMENT '订单金额',
  `createdAt` datetime DEFAULT NULL COMMENT '下单时间',
  `payment` int DEFAULT NULL COMMENT '1 余额支付 2 支付宝 3 微信',
  `discount` decimal(10,2) DEFAULT NULL COMMENT '优惠券扣减',
  `updatedAt` datetime DEFAULT NULL COMMENT '更新时间',
  `avatar` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '图片地址',
  `link_url` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '主页地址',
  `is_order_after` tinyint DEFAULT NULL COMMENT '是否续单 1 续 2 不续',
  `user_mate_id` int NOT NULL COMMENT '接单人id',
  PRIMARY KEY (`id`,`order_id`,`uid`,`user_mate_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='订单主表';

-- ----------------------------
-- Records of orders
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for server_category
-- ----------------------------
DROP TABLE IF EXISTS `server_category`;
CREATE TABLE `server_category` (
  `id` int NOT NULL,
  `category_name` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `base_amount` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '基础金额',
  `parent_id` int NOT NULL COMMENT '上级服务',
  `status` tinyint DEFAULT NULL COMMENT '1开启 2 关闭',
  `seven_days_price` decimal(10,2) DEFAULT NULL COMMENT '7天价格',
  `one_day_price` decimal(10,2) DEFAULT NULL COMMENT '1天价格',
  `month_price` decimal(10,2) DEFAULT NULL COMMENT '单月价格',
  `createdAt` timestamp NULL DEFAULT NULL,
  `updatedAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`,`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='服务类型,需要录入信息，最终计算';

-- ----------------------------
-- Records of server_category
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user_info
-- ----------------------------
DROP TABLE IF EXISTS `user_info`;
CREATE TABLE `user_info` (
  `id` int NOT NULL,
  `username` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `vip_level` tinyint(1) DEFAULT NULL COMMENT 'vip等级',
  `progress` tinyint DEFAULT NULL COMMENT '成长值 用于画滚动条',
  `balance` decimal(10,3) DEFAULT NULL COMMENT '可用余额',
  `status` tinyint DEFAULT '0' COMMENT '用户状态0 正常 1 拉黑',
  `avatar` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '头像地址',
  `is_hidden` tinyint DEFAULT NULL COMMENT '是否隐藏 1 隐藏 2 不隐藏',
  `order_num` tinyint DEFAULT NULL COMMENT '订单数',
  `createdAt` timestamp NULL DEFAULT NULL,
  `updatedAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='用户表\n';

-- ----------------------------
-- Records of user_info
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user_mate_feed
-- ----------------------------
DROP TABLE IF EXISTS `user_mate_feed`;
CREATE TABLE `user_mate_feed` (
  `id` int NOT NULL AUTO_INCREMENT,
  `source_url` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '资源地址\n',
  `datetime` datetime DEFAULT NULL COMMENT '发布时间',
  `text` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '描述信息',
  `clicks` tinyint DEFAULT NULL COMMENT '点赞数',
  `nickname` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '昵称',
  `source_type` tinyint DEFAULT NULL COMMENT '1 图片 2 视频',
  `createdAt` timestamp NULL DEFAULT NULL,
  `updatedAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='店员动态\n';

-- ----------------------------
-- Records of user_mate_feed
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user_mate_level
-- ----------------------------
DROP TABLE IF EXISTS `user_mate_level`;
CREATE TABLE `user_mate_level` (
  `id` int NOT NULL AUTO_INCREMENT,
  `level` tinyint DEFAULT NULL COMMENT '等级',
  `level_name` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '等级名称',
  `status` tinyint DEFAULT NULL COMMENT '状态 1 上线 2 下线',
  `radio` tinyint DEFAULT NULL COMMENT '系数',
  `createdAt` timestamp NULL DEFAULT NULL,
  `updatedAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='系数级别表\n';

-- ----------------------------
-- Records of user_mate_level
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for user_mates
-- ----------------------------
DROP TABLE IF EXISTS `user_mates`;
CREATE TABLE `user_mates` (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `group_id` tinyint NOT NULL COMMENT '组id',
  `real_name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '真实姓名',
  `tags` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '标签随着标签表的更更新而更新逗号分隔',
  `birthday` date DEFAULT NULL COMMENT '出生日期',
  `hobby` varchar(10) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '个人爱好',
  `avatar` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '上传头像地址',
  `nickname` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '昵称',
  `images` text COLLATE utf8mb4_general_ci COMMENT '图片地址json,多张图',
  `age` tinyint NOT NULL DEFAULT '0' COMMENT '年龄',
  `province` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地区',
  `sign` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '个性签名',
  `videourl` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '视频地址',
  `favorates` int DEFAULT '0' COMMENT '收藏次数',
  `is_online` tinyint NOT NULL DEFAULT '1' COMMENT '1在线2 离线',
  `is_employee` tinyint NOT NULL DEFAULT '0' COMMENT '1在职 离职',
  `is_approv` tinyint NOT NULL DEFAULT '1' COMMENT '1 通过审核 2 审核未通过',
  `is_deleted` int NOT NULL DEFAULT '0' COMMENT '是否删除 0 正常1 删除',
  `createdAt` timestamp NULL DEFAULT NULL,
  `updatedAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`,`group_id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='个人信息表\n';

-- ----------------------------
-- Records of user_mates
-- ----------------------------
BEGIN;
INSERT INTO `user_mates` (`id`, `username`, `group_id`, `real_name`, `tags`, `birthday`, `hobby`, `avatar`, `nickname`, `images`, `age`, `province`, `sign`, `videourl`, `favorates`, `is_online`, `is_employee`, `is_approv`, `is_deleted`, `createdAt`, `updatedAt`) VALUES (4, '小林', 1, '琳琳', 'aaa,bbb', '2004-01-01', '音乐', NULL, 'test', 'https://aliyun-soudcloud.com/usermate/20250/04/14/1231-sddfsfd.jpg', 22, '山东', 'adfafafd', 'https://aliyun-soudcloud.com/usermate/20250/04/14/1231-sddfsfd.m3u8', 0, 1, 0, 1, 0, '2025-04-14 21:26:12', '2025-04-14 21:26:12');
COMMIT;

-- ----------------------------
-- Table structure for user_online_duration
-- ----------------------------
DROP TABLE IF EXISTS `user_online_duration`;
CREATE TABLE `user_online_duration` (
  `id` int NOT NULL AUTO_INCREMENT,
  `login_time` datetime DEFAULT NULL COMMENT '登入时间',
  `logout_time` datetime DEFAULT NULL COMMENT '登出时间',
  `duration` int DEFAULT NULL COMMENT '时长',
  `device_type` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '设备类型',
  `createdAt` timestamp NULL DEFAULT NULL,
  `updatedAt` timestamp NULL DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='统计在线时长\n';

-- ----------------------------
-- Records of user_online_duration
-- ----------------------------
BEGIN;
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
