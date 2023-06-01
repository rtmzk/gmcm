CREATE TABLE IF NOT EXISTS `hosts` (
    `id` integer,
    `createdAt` datetime,
    `updateAt` datetime,
    `ip` text NOT NULL,
    `type` text DEFAULT "storage",
    `role` text NOT NULL,
    `sshUser` text NOT NULL,
    `sshPort` text,
    `loginType` text DEFAULT "nopass",
    PRIMARY KEY (`id`));

CREATE TABLE  IF NOT EXISTS `users` (
    `id` integer,
    `createdAt` datetime,
    `updateAt` datetime,
    `uuid` text,
    `name` text,
    `password` text,
    `nick_name` text DEFAULT "系统用户",
    `header_img` text,
    `authority_id` text DEFAULT "888",
    PRIMARY KEY (`id`));

INSERT INTO "users" VALUES (1,'2022-02-14 11:59:31.205527005+08:00','2022-02-14 11:59:31.205527005+08:00','8353e620-8d4a-11ec-b0da-0050569ecf76','admin','edoc2','系统用户','','888');
