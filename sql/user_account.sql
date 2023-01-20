CREATE TABLE `user_account` (
                                `id` int(11) NOT NULL AUTO_INCREMENT,
                                `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
                                `pay_password` char(32) NOT NULL DEFAULT '' COMMENT '余额支付密码的MD5',
                                `balance` decimal(10,0) NOT NULL DEFAULT '0' COMMENT '账户余额',
                                `state` tinyint(4) NOT NULL DEFAULT '0' COMMENT '账户状态 1：正常， 21:异常， 31：冻结',
                                `code` int(11) NOT NULL DEFAULT '123456789' COMMENT '修改码',
                                `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                PRIMARY KEY (`id`),
                                KEY `user` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户账户表';

