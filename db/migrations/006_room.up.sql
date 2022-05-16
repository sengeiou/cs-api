create table if not exists room
(
    id         bigint unsigned auto_increment primary key,
    staff_id   bigint unsigned           not null comment '職員ID',
    member_id  bigint unsigned           not null comment '會員ID',
    tag_id     bigint unsigned default 0 not null comment '標籤ID',
    score      tinyint         default 0 not null comment '評分 1-5',
    status     tinyint         default 1 not null comment '客服房狀態 1等待中 2服務中 3已關閉',
    created_at datetime                  not null comment '創建時間',
    updated_at datetime                  not null comment '更新時間',
    closed_at  datetime                  null comment '關閉時間'
);

