-- 用户表

CREATE TABLE `user`
(
 `user_id`      int unsigned NOT NULL AUTO_INCREMENT ,
 `name`         varchar(100) NOT NULL DEFAULT '' ,
 `email`        varchar(100) NOT NULL DEFAULT '' ,
 `password`     varchar(100) NOT NULL DEFAULT '' ,
 `introduction` varchar(1000) NOT NULL DEFAULT '' ,
 `github`       varchar(100) NOT NULL DEFAULT '' ,
 `person_url`   varchar(100) NOT NULL DEFAULT '' ,
 `created_at`   datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
 `updated_at`   datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
 -- 用户状态 -1 表示已删除； 0 表示未激活； 1 表示正常
 `status`       tinyint NOT NULL DEFAULT 0 ,
 `data`         text NULL ,

PRIMARY KEY (`user_id`)
) COMMENT='用户表' CHARSET=utf8mb4;
w
-- 问答表

CREATE TABLE `card`
(
 `card_id`    int unsigned NOT NULL AUTO_INCREMENT ,
 `question`   varchar(100) NOT NULL ,
 `answer`     text NOT NULL ,
 `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
 `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
 `auther_id`  int unsigned NOT NULL ,
 `status`     tinyint NOT NULL DEFAULT 0 ,
 `data`       text NULL ,

PRIMARY KEY (`card_id`),
CONSTRAINT `FK_card_user_id` FOREIGN KEY (`auther_id`) REFERENCES `user` (`user_id`)
) COMMENT='问题表' CHARSET=utf8mb4;

-- 标签表

CREATE TABLE `tag`
(
 `tag_id`     int unsigned NOT NULL AUTO_INCREMENT ,
 `name`       varchar(100) NOT NULL ,
 `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
 `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
 `status`     tinyint NOT NULL DEFAULT 0 ,
 `data`       text NULL ,

PRIMARY KEY (`tag_id`)
) CHARSET=utf8mb4;

-- 分类表

CREATE TABLE `category`
(
 `category_id` int unsigned NOT NULL AUTO_INCREMENT ,
 `name`        varchar(100) NOT NULL ,
 `created_at`  datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
 `updated_at`  datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
 `status`      tinyint NOT NULL DEFAULT 0 ,
 `user_id`     int unsigned NOT NULL ,
 `data`        text NULL ,

PRIMARY KEY (`category_id`),
CONSTRAINT `FK_category_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`)
) CHARSET=utf8mb4;


-- 分类问答表

CREATE TABLE `catagory_card`
(
 `category_id` int unsigned NOT NULL ,
 `card_id`     int unsigned NOT NULL ,
 `created_at`  datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
 `updated_at`  datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
 `data`        text NULL ,

PRIMARY KEY (`category_id`),
CONSTRAINT `FK_catagory_card_category_id` FOREIGN KEY (`category_id`) REFERENCES `category` (`category_id`),
CONSTRAINT `FK_catagory_card_card_id` FOREIGN KEY (`card_id`) REFERENCES `card` (`card_id`)
) CHARSET=utf8mb4;

-- 问答贡献表

CREATE TABLE `card_contributer`
(
 `id`         int unsigned NOT NULL AUTO_INCREMENT ,
 `card_id`    int unsigned NOT NULL ,
 `user_id`    int unsigned NOT NULL ,
 `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
 `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
 `data`       text NULL ,

PRIMARY KEY (`id`),
CONSTRAINT `card_contributer_card_id` FOREIGN KEY (`card_id`) REFERENCES `card` (`card_id`),
CONSTRAINT `card_contributer_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`)
) CHARSET=utf8mb4;

-- 问答标签表

CREATE TABLE `card_tag`
(
 `id`         int unsigned NOT NULL AUTO_INCREMENT ,
 `card_id`    int unsigned NOT NULL ,
 `tag_id`     int unsigned NOT NULL ,
 `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
 `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
 `data`       text NULL ,

PRIMARY KEY (`id`),
CONSTRAINT `card_tag_card_id` FOREIGN KEY (`card_id`) REFERENCES `card` (`card_id`),
CONSTRAINT `card_tag_tag_id` FOREIGN KEY (`tag_id`) REFERENCES `tag` (`tag_id`)
) CHARSET=utf8mb4;

-- 评论表

CREATE TABLE `comment`
(
 `id`                int unsigned NOT NULL AUTO_INCREMENT ,
 `card_id`           int unsigned NOT NULL ,
 `target_comment_id` int unsigned NULL ,
 `user_id`           int unsigned NOT NULL ,
 `content`           text NOT NULL ,
 `like_count`        int NOT NULL DEFAULT 0 COMMENT '评论被点赞的数量' ,
 `status`            tinyint NOT NULL DEFAULT 0 ,
 `created_at`        datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
 `updated_at`        datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
 `data`              text NULL ,

PRIMARY KEY (`id`),
CONSTRAINT `comment_id` FOREIGN KEY (`target_comment_id`) REFERENCES `comment` (`id`),
CONSTRAINT `comment_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`),
CONSTRAINT `comment_card_id` FOREIGN KEY (`card_id`) REFERENCES `card` (`card_id`)
) CHARSET=utf8mb4;

-- 用户问答表

CREATE TABLE `user_card`
(
 `id`         int unsigned NOT NULL AUTO_INCREMENT ,
 `card_id`    int unsigned NOT NULL ,
 `user_id`    int unsigned NOT NULL ,
 `factor`     decimal(1) NOT NULL DEFAULT 0.0 ,
 `schedule`   int NOT NULL DEFAULT 0 ,
 `status`     tinyint NOT NULL DEFAULT 0 ,
 `like`       tinyint NOT NULL DEFAULT 0 ,
 `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
 `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
 `data`       text NULL ,

PRIMARY KEY (`id`),
CONSTRAINT `user_card_card_id` FOREIGN KEY (`card_id`) REFERENCES `card` (`card_id`),
CONSTRAINT `user_card_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`)
) CHARSET=utf8mb4;

-- 用户分类表

CREATE TABLE `user_catagory`
(
 `id`          int unsigned NOT NULL AUTO_INCREMENT ,
 `user_id`     int unsigned NOT NULL ,
 `category_id` int unsigned NOT NULL ,
 `created_at`  datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
 `updated_at`  datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
 `data`        text NULL ,

PRIMARY KEY (`id`),
CONSTRAINT `user_catagory_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`),
CONSTRAINT `user_catagory_category_id` FOREIGN KEY (`category_id`) REFERENCES `category` (`category_id`)
) CHARSET=utf8mb4;

-- 用户粉丝表

CREATE TABLE `user_like`
(
 `id`           int unsigned NOT NULL AUTO_INCREMENT ,
 `user_id`      int unsigned NOT NULL ,
 `like_user_id` int unsigned NOT NULL ,
 `created_at`   datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
 `updated_at`   datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
 `data`         text NULL ,

PRIMARY KEY (`id`),
CONSTRAINT `user_like_user_id_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`),
CONSTRAINT `user_like_user_id_2` FOREIGN KEY (`like_user_id`) REFERENCES `user` (`user_id`)
) CHARSET=utf8mb4;

-- 用户标签表

CREATE TABLE `user_tag`
(
 `id`         int unsigned NOT NULL AUTO_INCREMENT ,
 `user_id`    int unsigned NOT NULL ,
 `tag_id`     int unsigned NOT NULL ,
 `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
 `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ,
 `data`       text NULL ,

PRIMARY KEY (`id`),
CONSTRAINT `user_tag_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`user_id`),
CONSTRAINT `user_tag_tag_id` FOREIGN KEY (`tag_id`) REFERENCES `tag` (`tag_id`)
) CHARSET=utf8mb4;
