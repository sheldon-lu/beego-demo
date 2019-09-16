SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for rbac_api
-- ----------------------------
DROP TABLE IF EXISTS `rbac_api`;
CREATE TABLE `rbac_api` (
  `id` smallint(6) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(20) NOT NULL,
  `title` varchar(50) DEFAULT NULL,
  `status` tinyint(1) DEFAULT '0',
  `remark` varchar(255) DEFAULT NULL,
  `sort` smallint(6) unsigned DEFAULT NULL,
  `pid` smallint(6) unsigned NOT NULL,
  `level` tinyint(1) unsigned NOT NULL,
  `is_show` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `level` (`level`),
  KEY `pid` (`pid`),
  KEY `status` (`status`),
  KEY `name` (`name`)
) ENGINE=MyISAM AUTO_INCREMENT=15 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of rbac_api
-- ----------------------------
INSERT INTO `rbac_api` VALUES ('1', 'RBAC', '后台管理', '1', '', '0', '0', '1', '1');
INSERT INTO `rbac_api` VALUES ('2', 'User', '用户管理', '1', '', '3', '1', '2', '1');
INSERT INTO `rbac_api` VALUES ('3', 'Hack', '主机管理', '1', '', '1', '1', '2', '1');
INSERT INTO `rbac_api` VALUES ('4', 'Week', '主机列表', '1', '', '1', '3', '3', '1');
INSERT INTO `rbac_api` VALUES ('5', 'role/list', '角色管理', '1', '', '3', '2', '3', '1');
INSERT INTO `rbac_api` VALUES ('6', 'apis/list', '权限管理', '1', '', '3', '2', '3', '1');
INSERT INTO `rbac_api` VALUES ('7', 'Auto', '管理1', '1', 'test', '2', '1', '2', '1');
INSERT INTO `rbac_api` VALUES ('8', 'user/list', '用户列表', '1', '', '3', '2', '3', '1');
INSERT INTO `rbac_api` VALUES ('9', 'list', '列表1', '0', '', '2', '7', '3', '1');
INSERT INTO `rbac_api` VALUES ('10', 'list', '列表2', '1', '', '2', '7', '3', '1');
INSERT INTO `rbac_api` VALUES ('12', 'json', '列表3', '1', '', '2', '7', '3', '1');
INSERT INTO `rbac_api` VALUES ('13', 'role/listjson', '角色列表', '1', '', '3', '2', '3', '0');
INSERT INTO `rbac_api` VALUES ('14', 'user/edit', '个人资料修改', '1', null, '3', '2', '3', '1');
