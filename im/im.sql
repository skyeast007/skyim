DROP TABLE IF EXISTS `t_user`;
CREATE TABLE `t_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL COMMENT '用户名',
  `account` char(15) NOT NULL COMMENT '用户账号',
  `mobile` char(11) DEFAULT '' COMMENT '手机号',
  `sign` varchar(100) DEFAULT '' COMMENT '签名',
  `password` char(32) NOT NULL COMMENT '密码',
  `gender` tinyint(4) NOT NULL DEFAULT '0' COMMENT '性别，0未知，１男，２女',
  `email` varchar(64) DEFAULT '' COMMENT '邮箱',
  `avatar` varchar(128) DEFAULT '' COMMENT '用户头像路径，标准长度为１１８个字符',
  `status` tinyint(4) DEFAULT NULL COMMENT '0正常，1冻结',
  `create_time` int(11) NOT NULL DEFAULT 0 COMMENT '建立时的间戳',
  `delete_time` int(11) NOT NULL DEFAULT 0 COMMENT '标记删除，并记录删除时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_account` (`account`),
  UNIQUE KEY `idx_mobile` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户基本信息';
insert into t_user values(
  null,'李华','tttlkkkl','18000000000','我们需要一些切实的行动...','skyim',1,'yehong0000@163.com','',0,unix_timestamp(),unix_timestamp(),now()
);
insert into t_user values(
  null,'东哥','skyim','18000000001','我们需要一些切实的行动...','skyim',1,'yehong0000@163.com','',0,unix_timestamp(),unix_timestamp(),now()
);
insert into t_user values(
  null,'虎子','skyim1','18000000002','我们需要一些切实的行动...','skyim',1,'yehong0000@163.com','',0,unix_timestamp(),unix_timestamp(),now()
);
insert into t_user values(
  null,'测试1','skyim2','18000000003','我们需要一些切实的行动...','skyim',1,'yehong0000@163.com','',0,unix_timestamp(),unix_timestamp(),now()
);
insert into t_user values(
  null,'测试2','skyim3','18000000004','我们需要一些切实的行动...','skyim',1,'yehong0000@163.com','',0,unix_timestamp(),unix_timestamp(),now()
);
insert into t_user values(
  null,'测试3','skyim4','18000000005','我们需要一些切实的行动...','skyim',1,'yehong0000@163.com','',0,unix_timestamp(),unix_timestamp(),now()
);
