SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for rbac_role
-- ----------------------------
DROP TABLE IF EXISTS `rbac_role`;
CREATE TABLE `rbac_role` (
  `id` smallint(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `pid` smallint(6) DEFAULT NULL,
  `status` tinyint(1) unsigned DEFAULT NULL,
  `remark` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `pid` (`pid`),
  KEY `status` (`status`)
) ENGINE=MyISAM AUTO_INCREMENT=17 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of rbac_role
-- ----------------------------
INSERT INTO `rbac_role` VALUES ('1', '超级管理员', '0', '1', '');
INSERT INTO `rbac_role` VALUES ('2', '普通会员', '1', '1', '');
INSERT INTO `rbac_role` VALUES ('16', '普通用户', '1', '1', 'test');
