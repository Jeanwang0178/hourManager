/*
Navicat MySQL Data Transfer

Source Server         : local
Source Server Version : 50636
Source Host           : 127.0.0.1:3306
Source Database       : sinochem_man_hour

Target Server Type    : MYSQL
Target Server Version : 50636
File Encoding         : 65001

Date: 2018-12-14 13:01:42
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for sys_auth
-- ----------------------------
DROP TABLE IF EXISTS `sys_auth`;
CREATE TABLE `sys_auth` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '自增ID',
  `pid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '上级ID，0为顶级',
  `auth_name` varchar(64) NOT NULL DEFAULT '0' COMMENT '权限名称',
  `auth_url` varchar(255) NOT NULL DEFAULT '0' COMMENT 'URL地址',
  `sort` int(1) unsigned NOT NULL DEFAULT '999' COMMENT '排序，越小越前',
  `icon` varchar(255) NOT NULL,
  `is_show` tinyint(1) unsigned NOT NULL DEFAULT '0' COMMENT '是否显示，0-隐藏，1-显示',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '操作者ID',
  `create_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态，1-正常，0-删除',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=66 DEFAULT CHARSET=utf8mb4 COMMENT='权限因子';

-- ----------------------------
-- Records of sys_auth
-- ----------------------------
INSERT INTO `sys_auth` VALUES ('1', '0', '权限', '/', '1', '', '0', '1', '1', '1', '1', '1505620970', '1505620970');
INSERT INTO `sys_auth` VALUES ('2', '1', '权限管理', '/', '999', 'fa-id-card', '1', '1', '0', '1', '1', '0', '1505622360');
INSERT INTO `sys_auth` VALUES ('3', '2', '用户管理', '/user/list', '1', 'fa-user-o', '1', '1', '1', '1', '1', '1505621186', '1505621186');
INSERT INTO `sys_auth` VALUES ('4', '2', '角色管理', '/role/list', '2', 'fa-user-circle-o', '1', '1', '0', '1', '1', '0', '1505621852');
INSERT INTO `sys_auth` VALUES ('5', '3', '新增', '/user/add', '1', '', '0', '1', '0', '1', '1', '0', '1505621685');
INSERT INTO `sys_auth` VALUES ('6', '3', '修改', '/user/edit', '2', '', '0', '1', '0', '1', '1', '0', '1505621697');
INSERT INTO `sys_auth` VALUES ('7', '3', '删除', '/user/ajaxdel', '3', '', '0', '1', '1', '1', '1', '1505621756', '1505621756');
INSERT INTO `sys_auth` VALUES ('8', '4', '新增', '/role/add', '1', '', '1', '1', '0', '1', '1', '0', '1505698716');
INSERT INTO `sys_auth` VALUES ('9', '4', '修改', '/role/edit', '2', '', '0', '1', '1', '1', '1', '1505621912', '1505621912');
INSERT INTO `sys_auth` VALUES ('10', '4', '删除', '/role/ajaxdel', '3', '', '0', '1', '1', '1', '1', '1505621951', '1505621951');
INSERT INTO `sys_auth` VALUES ('11', '2', '权限因子', '/auth/list', '3', 'fa-list', '1', '1', '1', '1', '1', '1505621986', '1505621986');
INSERT INTO `sys_auth` VALUES ('12', '11', '新增', '/auth/add', '1', '', '0', '1', '1', '1', '1', '1505622009', '1505622009');
INSERT INTO `sys_auth` VALUES ('13', '11', '修改', '/auth/edit', '2', '', '0', '1', '1', '1', '1', '1505622047', '1505622047');
INSERT INTO `sys_auth` VALUES ('14', '11', '删除', '/auth/ajaxdel', '3', '', '0', '1', '1', '1', '1', '1505622111', '1505622111');
INSERT INTO `sys_auth` VALUES ('15', '1', '个人中心', 'profile/edit', '1001', 'fa-user-circle-o', '1', '1', '0', '1', '1', '0', '1506001114');
INSERT INTO `sys_auth` VALUES ('24', '15', '资料修改', '/user/modify', '1', 'fa-edit', '1', '1', '0', '1', '1', '0', '1506057468');
INSERT INTO `sys_auth` VALUES ('52', '1', '项目管理', '/', '100', 'fa-files-o', '1', '1', '0', '1', '1', '0', '1544250948');
INSERT INTO `sys_auth` VALUES ('53', '52', '项目设置', '/project/list', '100', 'fa-files-o', '1', '1', '0', '1', '1', '0', '1544251646');
INSERT INTO `sys_auth` VALUES ('55', '53', '新增', '/project/add', '1', '', '0', '1', '0', '1', '1', '0', '1544250651');
INSERT INTO `sys_auth` VALUES ('56', '53', '修改', '/project/edit', '2', '', '0', '1', '0', '1', '1', '0', '1544250671');
INSERT INTO `sys_auth` VALUES ('57', '53', '删除', '/project/ajaxdel', '3', '', '0', '1', '0', '1', '1', '0', '1544250704');
INSERT INTO `sys_auth` VALUES ('60', '1', '工时管理', '/', '100', 'fa-files-o', '1', '1', '0', '1', '1', '0', '1544250948');
INSERT INTO `sys_auth` VALUES ('61', '60', '工时录入', '/manhour/list', '100', 'fa-files-o', '1', '1', '0', '1', '1', '0', '1544669276');
INSERT INTO `sys_auth` VALUES ('62', '61', '新增', '/manhour/add', '1', '', '0', '1', '0', '1', '1', '0', '1544494387');
INSERT INTO `sys_auth` VALUES ('63', '61', '修改', '/manhour/edit', '2', '', '0', '1', '0', '1', '1', '0', '1544494394');
INSERT INTO `sys_auth` VALUES ('64', '61', '删除', '/manhour/ajaxdel', '3', '', '0', '1', '0', '1', '1', '0', '1544494400');
INSERT INTO `sys_auth` VALUES ('65', '60', '工时管理', '/manhour/listall', '100', 'fa-files-o', '1', '1', '0', '1', '1', '0', '1544669255');

-- ----------------------------
-- Table structure for sys_man_hour
-- ----------------------------
DROP TABLE IF EXISTS `sys_man_hour`;
CREATE TABLE `sys_man_hour` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `project_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '项目ID',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `work_date` int(11) NOT NULL DEFAULT '0' COMMENT '日期',
  `task_target` varchar(1024) NOT NULL DEFAULT '0' COMMENT '当日工作目标',
  `task_progress` varchar(20) NOT NULL DEFAULT '0' COMMENT '任务进展情况',
  `man_hour` decimal(15,5) DEFAULT NULL COMMENT '本日用时',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态，1-正常 0禁用',
  `create_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4 COMMENT='工时统计表';

-- ----------------------------
-- Records of sys_man_hour
-- ----------------------------
INSERT INTO `sys_man_hour` VALUES ('3', '3', '1', '1545955200', '当日工作当日工作当日工作', '务进展情务进展情', '22.00000', '1', '1', '1', '1544578904', '1544669068');
INSERT INTO `sys_man_hour` VALUES ('4', '3', '1', '1544572801', '322', '322', '33.00000', '1', '1', '1', '1544579087', '1544600772');
INSERT INTO `sys_man_hour` VALUES ('5', '3', '1', '1544659200', '42', '22', '4.00000', '1', '1', '1', '1544579186', '1544679850');
INSERT INTO `sys_man_hour` VALUES ('6', '4', '1', '1544400000', '当日工作目当日工作目当日工作目', '任务进展情任务进展情任务进展情任务进展情', '23.00000', '1', '1', '1', '1544604236', '1544604236');
INSERT INTO `sys_man_hour` VALUES ('7', '4', '1', '1543881600', '333', '4444444', '22.00000', '1', '1', '1', '1544604296', '1544604296');
INSERT INTO `sys_man_hour` VALUES ('8', '4', '1', '1545091200', '2', '2', '22.00000', '1', '1', '1', '1544604410', '1544604410');
INSERT INTO `sys_man_hour` VALUES ('9', '4', '1', '1543968000', '22wwww', '任务进展情况任务进展情况', '6.00000', '1', '1', '1', '1544617753', '1544617781');
INSERT INTO `sys_man_hour` VALUES ('10', '3', '3', '1544832000', '打倒资本主义', '小特朗普弹劾中任务进展情况任务进展情况任', '12.00000', '1', '3', '3', '1544757142', '1544757142');
INSERT INTO `sys_man_hour` VALUES ('11', '3', '3', '1544486400', '社会之公主模板目标', '审核中', '8.00000', '1', '3', '3', '1544757187', '1544757187');

-- ----------------------------
-- Table structure for sys_project
-- ----------------------------
DROP TABLE IF EXISTS `sys_project`;
CREATE TABLE `sys_project` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `project_name` varchar(32) NOT NULL DEFAULT '0' COMMENT '项目名称',
  `detail` varchar(255) NOT NULL DEFAULT '0' COMMENT '备注',
  `create_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改这ID',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态1-正常，0-删除',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='项目表';

-- ----------------------------
-- Records of sys_project
-- ----------------------------
INSERT INTO `sys_project` VALUES ('3', 'TMS项目', 'TMS项目--化销', '1', '1', '1', '1544252076', '1544600718');
INSERT INTO `sys_project` VALUES ('4', '壹化网', '测试壹化网项目', '1', '1', '1', '1544601034', '1544601034');

-- ----------------------------
-- Table structure for sys_project_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_project_user`;
CREATE TABLE `sys_project_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `project_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '项目ID',
  `user_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '用户ID',
  `create_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态1-正常，0-删除',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=36 DEFAULT CHARSET=utf8 COMMENT='项目用户表';

-- ----------------------------
-- Records of sys_project_user
-- ----------------------------
INSERT INTO `sys_project_user` VALUES ('31', '3', '1', '1', '1', '1', '1544273682', '1544600952');
INSERT INTO `sys_project_user` VALUES ('32', '3', '3', '1', '1', '1', '1544600980', '1544600980');
INSERT INTO `sys_project_user` VALUES ('33', '3', '4', '1', '1', '1', '1544600980', '1544600980');
INSERT INTO `sys_project_user` VALUES ('34', '4', '3', '1', '1', '1', '1544603455', '1544603455');
INSERT INTO `sys_project_user` VALUES ('35', '4', '1', '1', '1', '1', '1544604214', '1544604214');

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `role_name` varchar(32) NOT NULL DEFAULT '0' COMMENT '角色名称',
  `detail` varchar(255) NOT NULL DEFAULT '0' COMMENT '备注',
  `create_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改这ID',
  `status` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '状态1-正常，0-删除',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '添加时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='角色表';

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES ('1', '普通用户', '普通用户', '0', '1', '1', '1544757036', '1544757036');
INSERT INTO `sys_role` VALUES ('2', '普通管理员', '普通管理员', '0', '4', '1', '1544757461', '1544757461');

-- ----------------------------
-- Table structure for sys_role_auth
-- ----------------------------
DROP TABLE IF EXISTS `sys_role_auth`;
CREATE TABLE `sys_role_auth` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `role_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '角色ID',
  `auth_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '权限ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `role_id` (`role_id`,`auth_id`)
) ENGINE=InnoDB AUTO_INCREMENT=118 DEFAULT CHARSET=utf8mb4 COMMENT='权限和角色关系表';

-- ----------------------------
-- Records of sys_role_auth
-- ----------------------------
INSERT INTO `sys_role_auth` VALUES ('65', '1', '0');
INSERT INTO `sys_role_auth` VALUES ('57', '1', '1');
INSERT INTO `sys_role_auth` VALUES ('63', '1', '15');
INSERT INTO `sys_role_auth` VALUES ('64', '1', '24');
INSERT INTO `sys_role_auth` VALUES ('58', '1', '60');
INSERT INTO `sys_role_auth` VALUES ('59', '1', '61');
INSERT INTO `sys_role_auth` VALUES ('60', '1', '62');
INSERT INTO `sys_role_auth` VALUES ('61', '1', '63');
INSERT INTO `sys_role_auth` VALUES ('62', '1', '64');
INSERT INTO `sys_role_auth` VALUES ('117', '2', '0');
INSERT INTO `sys_role_auth` VALUES ('94', '2', '1');
INSERT INTO `sys_role_auth` VALUES ('106', '2', '2');
INSERT INTO `sys_role_auth` VALUES ('107', '2', '3');
INSERT INTO `sys_role_auth` VALUES ('111', '2', '4');
INSERT INTO `sys_role_auth` VALUES ('108', '2', '5');
INSERT INTO `sys_role_auth` VALUES ('109', '2', '6');
INSERT INTO `sys_role_auth` VALUES ('110', '2', '7');
INSERT INTO `sys_role_auth` VALUES ('112', '2', '8');
INSERT INTO `sys_role_auth` VALUES ('113', '2', '9');
INSERT INTO `sys_role_auth` VALUES ('114', '2', '10');
INSERT INTO `sys_role_auth` VALUES ('115', '2', '15');
INSERT INTO `sys_role_auth` VALUES ('116', '2', '24');
INSERT INTO `sys_role_auth` VALUES ('95', '2', '52');
INSERT INTO `sys_role_auth` VALUES ('96', '2', '53');
INSERT INTO `sys_role_auth` VALUES ('97', '2', '55');
INSERT INTO `sys_role_auth` VALUES ('98', '2', '56');
INSERT INTO `sys_role_auth` VALUES ('99', '2', '57');
INSERT INTO `sys_role_auth` VALUES ('100', '2', '60');
INSERT INTO `sys_role_auth` VALUES ('101', '2', '61');
INSERT INTO `sys_role_auth` VALUES ('102', '2', '62');
INSERT INTO `sys_role_auth` VALUES ('103', '2', '63');
INSERT INTO `sys_role_auth` VALUES ('104', '2', '64');
INSERT INTO `sys_role_auth` VALUES ('105', '2', '65');

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `login_name` varchar(20) NOT NULL DEFAULT '' COMMENT '用户名',
  `real_name` varchar(32) NOT NULL DEFAULT '0' COMMENT '真实姓名',
  `password` char(32) NOT NULL DEFAULT '' COMMENT '密码',
  `role_ids` varchar(255) NOT NULL DEFAULT '0' COMMENT '角色id字符串，如：2,3,4',
  `phone` varchar(20) NOT NULL DEFAULT '0' COMMENT '手机号码',
  `email` varchar(50) NOT NULL DEFAULT '' COMMENT '邮箱',
  `salt` char(10) NOT NULL DEFAULT '' COMMENT '密码盐',
  `last_login` int(11) NOT NULL DEFAULT '0' COMMENT '最后登录时间',
  `last_ip` char(15) NOT NULL DEFAULT '' COMMENT '最后登录IP',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态，1-正常 0禁用',
  `create_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建者ID',
  `update_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改者ID',
  `create_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '修改时间',
  `company_name` varchar(64) NOT NULL DEFAULT '' COMMENT '公司',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_user_name` (`login_name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES ('1', 'system', '超级管理员', '832a38350deb09263a088ce29fe6c613', '0', '13888888889', 'sinochem@163.com', 'WaTR', '1544763448', '127.0.0.1', '1', '0', '1', '0', '1544757248', '超级管理员公司');
INSERT INTO `sys_user` VALUES ('2', 'george518', 'georgeHao', '832a38350deb09263a088ce29fe6c613', '1,2', '13811558899', '12@11.com', 'WaTR', '1506125048', '127.0.0.1', '1', '0', '1', '0', '1544668651', '小赵公司D');
INSERT INTO `sys_user` VALUES ('3', 'test', '普通用户', '832a38350deb09263a088ce29fe6c613', '1', '13811559988', 'hei@123.com', 'WaTR', '1544763302', '127.0.0.1', '1', '1', '3', '1505919245', '1544757084', '小白公司C');
INSERT INTO `sys_user` VALUES ('4', 'admin', '普通管理员', '832a38350deb09263a088ce29fe6c613', '2', '13988009988', '232@124.com', 'WaTR', '1544763452', '127.0.0.1', '1', '1', '1', '1506047337', '1544668564', '小文公司');
