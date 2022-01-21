CREATE TABLE `accounts` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `account` varchar(50) NOT NULL COMMENT '账户名/邮箱',
  `mobile` varchar(50) DEFAULT NULL COMMENT '手机号',
  `password` varchar(255) NOT NULL COMMENT '密码',
  `gender` varchar(50) DEFAULT NULL COMMENT '性别',
  `birthdate` date DEFAULT NULL COMMENT '生日',
  `status` int(11) DEFAULT NULL COMMENT '用户状态',
  `last_login_ip` varchar(15) DEFAULT NULL COMMENT '上次登录的IP',
  `last_logged_in` date DEFAULT NULL COMMENT '上次登录时间',
  `real_name` varchar(50) DEFAULT NULL COMMENT '真实姓名',
  `nick_name` varchar(50) DEFAULT NULL COMMENT '昵称',
  `paused_at` date DEFAULT NULL,
  `closed_at` date DEFAULT NULL,
  `created_at` date DEFAULT current_timestamp() COMMENT '创建时间',
  `updated_at` date DEFAULT current_timestamp() COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `Account` (`account`(20)) USING BTREE COMMENT '每个邮箱只能注册一次'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;