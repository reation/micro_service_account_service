CREATE TABLE `user_account_log` (
                                    `id` int(11) NOT NULL AUTO_INCREMENT,
                                    `user_id` int(11) NOT NULL DEFAULT '0' COMMENT '用户ID',
                                    `payment_id` int(11) NOT NULL DEFAULT '0' COMMENT '支付ID',
                                    `order_num` char(22) NOT NULL DEFAULT '' COMMENT '订单号（非订单类收退款此字符为空）',
                                    `op_before_balance` decimal(10,0) NOT NULL DEFAULT '0' COMMENT '操作之前账户余额',
                                    `op_prices` decimal(10,0) NOT NULL DEFAULT '0' COMMENT '操作金额',
                                    `op_after_balance` decimal(10,0) NOT NULL DEFAULT '0' COMMENT '操作之后账户余额',
                                    `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '操作类型 1：支付订单 2：退货 3：补偿发放 4：购买会员',
                                    `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                    PRIMARY KEY (`id`),
                                    KEY `user` (`user_id`) USING BTREE,
                                    KEY `payment` (`payment_id`) USING BTREE,
                                    KEY `order` (`order_num`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户账户操作日志表';

