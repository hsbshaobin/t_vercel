123

-- 邀请码表
drop table if exists ac_invite_code;
create table ac_invite_code(
    ac_invite_code_id bigint unsigned not null auto_increment comment '主键id',
    app_id varchar(32) not null comment '应用ID',
    invite_code char(8) not null comment '邀请码',
    is_used tinyint unsigned not null default 0 comment '是否已使用：0 否；1 是',
    browser_id varchar(32) null comment '浏览器标识',
    used_time DATETIME null comment '使用时间',
    create_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '获取时间',
    update_time DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
    primary key(ac_invite_code_id)
) ENGINE = INNODB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COMMENT '邀请码表';

-- insert into ac_invite_code(app_id,invite_code)
-- values
-- (),
