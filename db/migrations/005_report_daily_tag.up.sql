create table if not exists report_daily_tag
(
    id         bigint unsigned auto_increment primary key,
    date       date            not null comment '報表日期',
    tag_id     bigint unsigned not null comment '標籤ID',
    count      int              not null comment '人數',
    created_at datetime        not null comment '創建時間'
);

