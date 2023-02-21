DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
  `id`            bigint(20)   NOT NULL AUTO_INCREMENT COMMENT '评论id，自增主键',
  `user_id`       bigint(20)   NOT NULL COMMENT '评论发布用户id',
  `video_id`      bigint(20)   NOT NULL COMMENT '评论视频id',
  `comment_text`  varchar(255) NOT NULL COMMENT '评论内容',
  `create_date`   datetime     NOT NULL COMMENT '评论发布时间',
  `action_type`        tinyint(4)   NOT NULL DEFAULT '1' COMMENT '默认评论发布为1，取消后为2',
  PRIMARY KEY (`id`),
  KEY `videoIdIdx` (`video_id`) USING BTREE COMMENT '评论列表使用视频id作为索引-方便查看视频下的评论列表'
) ENGINE=InnoDB AUTO_INCREMENT=1206 DEFAULT CHARSET=utf8 COMMENT='评论表';