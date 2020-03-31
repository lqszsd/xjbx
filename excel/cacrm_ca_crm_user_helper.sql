/*
 Navicat Premium Data Transfer

 Source Server         : 爱启测试
 Source Server Type    : MySQL
 Source Server Version : 50725
 Source Host           : 139.196.174.234:3306
 Source Schema         : Aiqitest

 Target Server Type    : MySQL
 Target Server Version : 50725
 File Encoding         : 65001

 Date: 02/07/2019 13:24:00
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for cacrm_ca_crm_user_helper
-- ----------------------------
DROP TABLE IF EXISTS `cacrm_ca_crm_user_helper`;
CREATE TABLE `cacrm_ca_crm_user_helper`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `code` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '额外码',
  `job_code` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '工号',
  `phone` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '手机',
  `password` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '密码',
  `headimg` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '头像',
  `name` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '名称',
  `type` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '类型',
  `is_modified_password` int(11) NOT NULL DEFAULT 0 COMMENT '是否修改过密码',
  `back_code` varchar(128) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
  `one_level` int(11) NOT NULL DEFAULT 0 COMMENT '一级',
  `two_level` int(11) NOT NULL DEFAULT 0 COMMENT '二级',
  `three_level` int(11) NOT NULL DEFAULT 0 COMMENT '三级',
  `four_level` int(11) NOT NULL DEFAULT 0 COMMENT '四级',
  `bottom_level` int(11) NOT NULL DEFAULT 0 COMMENT '最底级',
  `is_exist` int(11) NOT NULL DEFAULT 0 COMMENT '是否存在',
  `is_transfer` int(11) NOT NULL DEFAULT 0 COMMENT '是否转移',
  `source` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '导入' COMMENT '来源',
  `state` varchar(4) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '已审核' COMMENT '审核状态',
  `created_at` datetime(0) NOT NULL,
  `updated_at` timestamp(0) NOT NULL DEFAULT CURRENT_TIMESTAMP(0) ON UPDATE CURRENT_TIMESTAMP(0),
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3026 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
