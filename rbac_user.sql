SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for rbac_user
-- ----------------------------
DROP TABLE IF EXISTS `rbac_user`;
CREATE TABLE `rbac_user` (
  `id` int(8) NOT NULL AUTO_INCREMENT,
  `employee_number` int(8) NOT NULL COMMENT '工号',
  `username` varchar(16) NOT NULL COMMENT '用户名',
  `real_name` varchar(16) NOT NULL DEFAULT '0' COMMENT '真实姓名',
  `display_name` varchar(16) NOT NULL DEFAULT '0' COMMENT '显示姓名',
  `first_name` varchar(16) NOT NULL DEFAULT '0' COMMENT '名',
  `second_name` varchar(16) NOT NULL DEFAULT '0' COMMENT '姓',
  `password` char(64) NOT NULL DEFAULT '' COMMENT '密码',
  `phone` varchar(16) NOT NULL DEFAULT '0' COMMENT '手机号码',
  `email` varchar(32) NOT NULL DEFAULT 'xxxx@xx.cn' COMMENT '邮箱',
  `company` varchar(64) NOT NULL DEFAULT 'xxx有限公司' COMMENT '公司',
  `department` varchar(64) NOT NULL,
  `position` varchar(64) NOT NULL DEFAULT '' COMMENT '岗位',
  `role_id` int(11) NOT NULL,
  `onboard_time` varchar(32) DEFAULT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NULL DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`,`employee_number`),
  UNIQUE KEY `username` (`username`),
  UNIQUE KEY `employee_number` (`employee_number`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=63 DEFAULT CHARSET=utf8mb4 COMMENT='用户信息表';

-- ----------------------------
-- Records of rbac_user
-- ----------------------------
INSERT INTO `rbac_user` VALUES ('1', '100001', 'sheldon.lu', '陆sheldon', '陆sheldon', 'sheldon', '陆', '', '15100000000', 'xxx@xx.cn', 'xxxx有限公司', 'xxx', 'xxx', '0', '2010-01-01', '2019-09-09 00:00:00', '2019-09-09 15:44:28');