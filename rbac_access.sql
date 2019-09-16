SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for rbac_access
-- ----------------------------
DROP TABLE IF EXISTS `rbac_access`;
CREATE TABLE `rbac_access` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` smallint(6) unsigned NOT NULL,
  `api_id` smallint(6) unsigned NOT NULL,
  `level` tinyint(1) NOT NULL,
  `module` varchar(50) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `groupId` (`role_id`),
  KEY `nodeId` (`api_id`)
) ENGINE=MyISAM AUTO_INCREMENT=141 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of rbac_access
-- ----------------------------
INSERT INTO `rbac_access` VALUES ('133', '1', '7', '2', '模块');
INSERT INTO `rbac_access` VALUES ('132', '1', '10', '3', '操作');
INSERT INTO `rbac_access` VALUES ('131', '1', '4', '2', '模块');
INSERT INTO `rbac_access` VALUES ('130', '1', '12', '3', '操作');
INSERT INTO `rbac_access` VALUES ('129', '1', '9', '3', '操作');
INSERT INTO `rbac_access` VALUES ('128', '1', '3', '2', '模块');
INSERT INTO `rbac_access` VALUES ('127', '1', '13', '3', '操作');
INSERT INTO `rbac_access` VALUES ('126', '1', '6', '3', '操作');
INSERT INTO `rbac_access` VALUES ('125', '1', '5', '3', '操作');
INSERT INTO `rbac_access` VALUES ('124', '1', '8', '3', '操作');
INSERT INTO `rbac_access` VALUES ('123', '1', '14', '3', '操作');
INSERT INTO `rbac_access` VALUES ('64', '15', '9', '3', '操作');
INSERT INTO `rbac_access` VALUES ('63', '15', '3', '2', '模块');
INSERT INTO `rbac_access` VALUES ('62', '15', '8', '3', '操作');
INSERT INTO `rbac_access` VALUES ('140', '2', '14', '3', '操作');
INSERT INTO `rbac_access` VALUES ('139', '2', '13', '3', '操作');
INSERT INTO `rbac_access` VALUES ('61', '15', '6', '3', '操作');
INSERT INTO `rbac_access` VALUES ('60', '15', '2', '2', '模块');
INSERT INTO `rbac_access` VALUES ('59', '15', '1', '1', '项目');
INSERT INTO `rbac_access` VALUES ('65', '15', '12', '3', '操作');
INSERT INTO `rbac_access` VALUES ('66', '15', '7', '2', '模块');
INSERT INTO `rbac_access` VALUES ('122', '1', '2', '2', '模块');
INSERT INTO `rbac_access` VALUES ('121', '1', '1', '1', '项目');
INSERT INTO `rbac_access` VALUES ('138', '2', '8', '3', '操作');
INSERT INTO `rbac_access` VALUES ('137', '2', '6', '3', '操作');
INSERT INTO `rbac_access` VALUES ('136', '2', '5', '3', '操作');
INSERT INTO `rbac_access` VALUES ('135', '2', '2', '2', '模块');
INSERT INTO `rbac_access` VALUES ('134', '2', '1', '1', '项目');
INSERT INTO `rbac_access` VALUES ('113', '3', '7', '2', '模块');
INSERT INTO `rbac_access` VALUES ('112', '3', '10', '3', '操作');
INSERT INTO `rbac_access` VALUES ('111', '3', '4', '2', '模块');
INSERT INTO `rbac_access` VALUES ('110', '3', '12', '3', '操作');
INSERT INTO `rbac_access` VALUES ('109', '3', '9', '3', '操作');
INSERT INTO `rbac_access` VALUES ('108', '3', '3', '2', '模块');
INSERT INTO `rbac_access` VALUES ('107', '3', '1', '1', '项目');
INSERT INTO `rbac_access` VALUES ('114', '16', '1', '1', '项目');
INSERT INTO `rbac_access` VALUES ('115', '16', '3', '2', '模块');
INSERT INTO `rbac_access` VALUES ('116', '16', '9', '3', '操作');
INSERT INTO `rbac_access` VALUES ('117', '16', '12', '3', '操作');
INSERT INTO `rbac_access` VALUES ('118', '16', '4', '2', '模块');
INSERT INTO `rbac_access` VALUES ('119', '16', '10', '3', '操作');
INSERT INTO `rbac_access` VALUES ('120', '16', '7', '2', '模块');
