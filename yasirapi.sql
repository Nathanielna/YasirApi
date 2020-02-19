/*
 Navicat MySQL Data Transfer

 Source Server         : HUAWEI
 Source Server Type    : MySQL
 Source Server Version : 50640
 Source Host           : 139.9.57.122:3306
 Source Schema         : yasirapi

 Target Server Type    : MySQL
 Target Server Version : 50640
 File Encoding         : 65001

 Date: 19/02/2020 21:33:22
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for baidutop
-- ----------------------------
DROP TABLE IF EXISTS `baidutop`;
CREATE TABLE `baidutop`  (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '百度排行榜',
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '标题',
  `time` datetime(0) NULL DEFAULT NULL COMMENT '时间',
  `sort` int(11) NULL DEFAULT NULL COMMENT '顺序',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 11 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of baidutop
-- ----------------------------
INSERT INTO `baidutop` VALUES (1, '第40届全英音乐奖', '2020-02-19 20:03:35', 1);
INSERT INTO `baidutop` VALUES (2, '李兰迪做蛋糕', '2020-02-19 20:03:35', 2);
INSERT INTO `baidutop` VALUES (3, '教育部第1号预警', '2020-02-19 20:03:35', 3);
INSERT INTO `baidutop` VALUES (4, '我在北京等你定档', '2020-02-19 20:03:35', 4);
INSERT INTO `baidutop` VALUES (5, '2020北京车展延期', '2020-02-19 20:03:35', 5);
INSERT INTO `baidutop` VALUES (6, '小区出入证大赏', '2020-02-19 20:03:35', 6);
INSERT INTO `baidutop` VALUES (7, '浓眉哥罚球绝杀', '2020-02-19 20:03:35', 7);
INSERT INTO `baidutop` VALUES (8, '东京奥运如期举行', '2020-02-19 20:03:35', 8);
INSERT INTO `baidutop` VALUES (9, '李明博获刑17年', '2020-02-19 20:03:35', 9);
INSERT INTO `baidutop` VALUES (10, '梁咏琪西班牙探亲', '2020-02-19 20:03:35', 10);

SET FOREIGN_KEY_CHECKS = 1;
