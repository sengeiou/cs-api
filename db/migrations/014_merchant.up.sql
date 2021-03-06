create table if not exists merchant
(
    id         bigint unsigned auto_increment primary key,
    name       varchar(255)      not null comment '商戶名稱',
    code       varchar(255)      not null comment '商戶號',
    `key`      varchar(255)      not null comment '商戶金鑰',
    status     tinyint default 1 not null comment '狀態 1開啟 2關閉',
    created_by bigint unsigned   not null comment '創建管理員',
    created_at datetime          not null comment '創建時間',
    updated_by bigint unsigned   not null comment '更新管理員',
    updated_at datetime          not null comment '更新時間',
    constraint idx_code
        unique (code),
    constraint idx_key
        unique (`key`)
) comment '商家資料表' COLLATE utf8mb4_general_ci;