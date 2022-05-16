create table if not exists report_daily_guest
(
    id          bigint unsigned auto_increment primary key,
    date        date            not null comment '報表日期',
    guest_count int             not null comment '訪客數',
    created_at  datetime        not null comment '創建時間'
);

