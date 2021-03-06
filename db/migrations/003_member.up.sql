create table if not exists member
(
    id            bigint unsigned auto_increment primary key,
    type          tinyint      not null comment '用戶類型 1一般用戶 2訪客',
    name          varchar(20)  not null comment '會員名稱',
    device_id     varchar(255) not null comment '設備號',
    real_name     varchar(255) not null default '' comment '真實姓名',
    mobile        varchar(30)  not null default '' comment '手機號',
    email         varchar(255) not null default '' comment '電子信箱',
    online_status tinyint      not null default 2 comment '會員狀態 1在線 2離線',
    created_at    datetime     not null comment '創建時間',
    updated_at    datetime     not null comment '更新時間',
    constraint idx_name
        unique (name)
);

create index idx_device
    on member (device_id);

