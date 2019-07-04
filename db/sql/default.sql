--
-- 表的结构 `micro_admin_menu`
--

CREATE TABLE IF NOT EXISTS `micro_admin_menu` (
  `id` int(10) UNSIGNED NOT NULL ,
  `parent_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '父菜单id',
  `type` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '菜单类型:0-只作为菜单,1-有界面可访问菜单,2-无界面可访问菜单,3-有界面非显示菜单',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '状态:1-显示,0-不显示',
  `list_order` float NOT NULL DEFAULT '10000' COMMENT '排序',
  `url` varchar(50) NOT NULL DEFAULT '' COMMENT '路径',
  `func` varchar(50) NOT NULL DEFAULT '' COMMENT '控制器函数',
  `method` varchar(100) NOT NULL DEFAULT '' COMMENT '规则方法(大写)GET、POST、PUT、 PUT | GET',
  `param` varchar(50) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '额外参数',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单名称',
  `icon` varchar(20) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '菜单图标',
  `remark` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  KEY `status` (`status`),
  KEY `parent_id` (`parent_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='后台菜单表';

--
-- 转存表中的数据 `micro_admin_menu`
--
INSERT INTO `micro_admin_menu` (`id`, `parent_id`, `type`, `status`, `list_order`, `url`, `func`, `method`, `param`, `name`, `icon`, `remark`) VALUES
(1,     0,   0, 1, 0,     'admin/setting', 'ROOT',                'GET',    '', '设置',                '&#xe6ae', '系统设置入口'),
(100,   1,   1, 1, 100,   'admin/setting', 'Mail',                'GET',    '', '邮箱配置',             '&#xe69f', '邮箱配置'),
(10000, 100, 2, 1, 10000, 'admin/setting', 'MailConf',            'POST',   '', '邮箱配置提交保存',      '', '邮箱配置提交保存'),
(10001, 100, 2, 1, 10000, 'admin/setting', 'MailTest',            'POST',   '', '邮件发送测试',          '', '邮件发送测试'),
(101,   1,   1, 1, 101,   'admin/setting', 'Link',                'GET',    '', '友情链接',             '&#xe6f7', '友情链接管理'),
(10100, 101, 3, 1, 10000, 'admin/setting', 'LinkEdit',            'GET',    '', '编辑友情链接',         '', '编辑友情链接'),
(10101, 101, 2, 1, 10000, 'admin/setting', 'LinkEditPost',        'POST',   '', '编辑友情链接提交保存',  '', '编辑友情链接提交保存'),
(10102, 101, 2, 1, 10000, 'admin/setting', 'LinkDelete',          'POST',   '', '删除友情链接',         '', '删除友情链接'),
(10103, 101, 2, 1, 10000, 'admin/setting', 'LinkStatusChange',    'POST',   '', '友情链接显示隐藏',      '', '友情链接显示隐藏'),
(10104, 101, 2, 1, 10000, 'admin/setting', 'LinkList',            'GET',    '', '友情链接列表',          '', '友情链接列表'),
(102,   1,   1, 1, 0,     'admin/setting', 'Site',                'GET',    '', '网站配置',             '&#xe6fc', '网站配置'),
(10200, 102, 2, 1, 10000, 'admin/setting', 'SitePost',            'POST',   '', '网站信息设置提交',      '', '网站信息设置提交'),
(104,   1,   1, 1, 10000, 'admin/setting', 'Upload',              'GET',    '', '上传设置',             '&#xe71d', '上传设置'),
(10400, 104, 2, 1, 10000, 'admin/setting', 'UploadPost',          'POST',   '', '上传设置提交',          '', '上传设置提交'),

(2,     0,   0, 1, 10,    'admin/user',    'Root',                'GET',    '', '管理员管理',             '&#xe6b8', '用户管理'),
(200,   2,   1, 1, 10000, 'admin/user',    'RoleIndex',           'GET',    '', '角色管理',             '&#xe6f5', '角色管理'),
(20000, 200, 3, 1, 10000, 'admin/user',    'RoleEdit',            'GET',    '', '编辑角色',              '', '编辑角色'),
(20001, 200, 2, 1, 10000, 'admin/user',    'RoleEditPost',        'POST',   '', '编辑角色提交',           '', '编辑角色提交'),
(20002, 200, 2, 1, 10000, 'admin/user',    'RoleDelete',          'POST',   '', '删除角色',               '', '删除角色'),
(20003, 200, 2, 1, 10000, 'admin/user',    'RoleStatusChange',    'POST',   '', '角色状态设置',            '', '角色状态设置'),
(20004, 200, 2, 1, 10000, 'admin/user',    'RoleList',            'GET',    '', '角色列表',               '', '角色列表'),
(201,   2,   1, 1, 10000, 'admin/user',    'UserIndex',           'GET',    '', '管理员',               '&#xe726', '管理员'),
(20100, 201, 3, 1, 10000, 'admin/user',    'UserEdit',            'GET',    '', '编辑用户',               '', '编辑用户'),
(20101, 201, 2, 1, 10000, 'admin/user',    'UserEditPost',        'POST',   '', '编辑用户提交',           '', '编辑用户提交'),
(20102, 201, 2, 1, 10000, 'admin/user',    'UserDelete',          'POST',   '', '删除用户',               '', '删除用户'),
(20103, 201, 2, 1, 10000, 'admin/user',    'UserStatusChange',    'POST',   '', '用户状态设置',            '', '用户状态设置'),
(20104, 201, 2, 1, 10000, 'admin/user',    'UserList',            'GET',    '', '用户列表',               '', '用户列表'),

(3,     0,   0, 1, 30,    'admin/member',  'Root',                'GET',    '', '会员管理',            '&#xe70b', '会员管理'),
(300,   3,   1, 1, 10000, 'admin/member',  'Index',               'GET',    '', '会员',               '&#xe6ba', '会员'),
(30000, 300, 2, 1, 10000, 'admin/member',  'StatusChange',        'POST',   '', '会员状态设置',          '', '会员状态设置'),
(30001, 300, 2, 1, 10000, 'admin/member',  'List',                'GET',    '', '会员列表',               '', '会员列表'),
(30002, 300, 1, 1, 10000, 'admin/member',  'Edit',                'GET',    '', '编辑会员',               '', '编辑会员'),
(30003, 300, 2, 1, 10000, 'admin/member',  'EditPost',            'POST',   '', '编辑会员提交',           '', '编辑会员提交'),
(30004, 300, 2, 1, 10000, 'admin/member',  'Delete',              'POST',   '', '删除会员',               '', '删除会员'),

(4,     0,   0, 1, 30,    'admin/portal',  'Root',                'GET',    '', '门户管理',              '&#xe6b4', '门户管理'),
(400,   4,   1, 1, 10000, 'admin/portal',  'ArticleIndex',        'GET',    '', '文章管理',              '&#xe705', '文章列表'),
(40000, 400, 1, 1, 10000, 'admin/portal',  'ArticleEdit',         'GET',    '', '编辑文章',              '', '编辑文章'),
(40001, 400, 2, 1, 10000, 'admin/portal',  'ArticleEditPost',     'POST',   '', '编辑文章提交',           '', '编辑文章提交'),
(40002, 400, 2, 1, 10000, 'admin/portal',  'ArticleDelete',       'POST',   '', '文章删除',              '', '文章删除'),
(40003, 400, 2, 1, 10000, 'admin/portal',  'ArticleStatusChange', 'POST',   '', '文章状态变更',           '', '文章状态变更'),
(40004, 400, 2, 1, 10000, 'admin/portal',  'ArticleList',         'GET',    '', '文章列表',               '', '文章列表'),
(401,   4,   1, 1, 10000, 'admin/portal',  'CateIndex',           'GET',    '', '分类管理',              '&#xe699', '分类列表'),
(40100, 401, 1, 1, 10000, 'admin/portal',  'CateEdit',            'GET',    '', '编辑分类',              '', '编辑分类'),
(40101, 401, 2, 1, 10000, 'admin/portal',  'CateEditPost',        'POST',   '', '编辑分类提交',        '', '编辑分类提交'),
(40102, 401, 2, 1, 10000, 'admin/portal',  'CateStatusChange',    'POST',   '', '更新分类状态',            '', '更新分类状态'),
(40103, 401, 2, 1, 10000, 'admin/portal',  'CateDelete',          'POST',   '', '删除分类',               '', '删除分类'),
(40104, 401, 2, 1, 10000, 'admin/portal',  'CateRestore',         'GET',    '', '恢复分类',               '', '恢复分类'),
(402,   4,   1, 1, 10000, 'admin/portal',  'TagIndex',            'GET',    '', '标签管理',               '&#xe6f2', '标签管理'),
(40200, 402, 1, 1, 10000, 'admin/portal',  'TagAdd',              'GET',    '', '添加标签',               '', '添加文签'),
(40201, 402, 2, 1, 10000, 'admin/portal',  'TagEditPost',         'POST',   '', '添加标签提交',           '', '添加文签提交'),
(40202, 402, 2, 1, 10000, 'admin/portal',  'TagStatusChange',     'POST',   '', '更新标签状态',           '', '更新标签状态'),
(40203, 402, 2, 1, 10000, 'admin/portal',  'TagDelete',           'POST',   '', '删除标签',               '', '删除标签'),
(40204, 402, 2, 1, 10000, 'admin/portal',  'TagList',             'GET',    '', '标签列表',               '', '标签列表');

-- --------------------------------------------------------

--
-- 表的结构 `micro_asset`
--

CREATE TABLE IF NOT EXISTS `micro_asset` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户id',
  `file_size` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '文件大小,单位B',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '上传时间',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态:1-可用,0-不可用',
  `download_times` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '下载次数',
  `file_key` varchar(64) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '文件惟一码',
  `filename` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '文件名',
  `file_path` varchar(100) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '文件路径,相对于upload目录,可以为url',
  `file_md5` varchar(32) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '文件md5值',
  `file_sha1` varchar(40) CHARACTER SET utf8 NOT NULL DEFAULT '',
  `suffix` varchar(10) NOT NULL DEFAULT '' COMMENT '文件后缀名,不包括点',
  `more` text COMMENT '其它详细信息,JSON格式',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资源表';


--
-- 表的结构 `micro_user_like_201904`
--

CREATE TABLE IF NOT EXISTS `micro_casbin_rule` (
  `p_type` varchar(100) NOT NULL DEFAULT '' COMMENT '规则类型',
  `v0` varchar(100) COMMENT '规则0',
  `v1` varchar(100) COMMENT '规则1',
  `v2` varchar(100) COMMENT '规则2',
  `v3` varchar(100) COMMENT '规则3',
  `v4` varchar(100) COMMENT '规则4',
  `v5` varchar(100) COMMENT '规则4',
  UNIQUE KEY `rule_key` (`p_type`, `v0`, `v1`, `v2`, `v3`, `v4`, `v5`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COMMENT='授权表';

-- --------------------------------------------------------

--
-- 表的结构 `micro_comment`
--

CREATE TABLE IF NOT EXISTS `micro_comment` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `parent_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '被回复的评论id',
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '发表评论的用户id',
  `to_user_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '被评论的用户id',
  `object_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '评论内容 id',
  `like_count` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '点赞数',
  `dislike_count` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '吐槽数',
  `floor` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '楼层数',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '评论时间',
  `delete_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态:1-已审核,0-未审核',
  `type` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '评论类型:1-实名评论',
  `tab_name` varchar(64) NOT NULL DEFAULT '' COMMENT '评论内容所在表,不带表前缀',
  `full_name` varchar(50) NOT NULL DEFAULT '' COMMENT '评论者昵称',
  `url` text COMMENT '原文地址',
  `content` text CHARACTER SET utf8mb4 COMMENT '评论内容',
  `more` text CHARACTER SET utf8mb4 COMMENT '扩展属性',
  PRIMARY KEY (`id`),
  KEY `table_id_status` (`tab_name`,`object_id`,`status`),
  KEY `object_id` (`object_id`) USING BTREE,
  KEY `status` (`status`) USING BTREE,
  KEY `parent_id` (`parent_id`) USING BTREE,
  KEY `create_time` (`create_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='评论表';

-- --------------------------------------------------------

--
-- 表的结构 `micro_link`
--

CREATE TABLE IF NOT EXISTS `micro_link` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态:1-显示,2-不显示',
  `rating` int(11) NOT NULL DEFAULT '0' COMMENT '友情链接评级',
  `list_order` float NOT NULL DEFAULT '10000' COMMENT '排序',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新时间',
  `description` varchar(255) NOT NULL DEFAULT '' COMMENT '友情链接描述',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '友情链接地址',
  `name` varchar(30) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '友情链接名称',
  `image` varchar(100) NOT NULL DEFAULT '' COMMENT '友情链接图标',
  `target` varchar(10) NOT NULL DEFAULT '' COMMENT '友情链接打开方式',
  `rel` varchar(50) NOT NULL DEFAULT '' COMMENT '链接与网站的关系',
  PRIMARY KEY (`id`),
  KEY `status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='友情链接表';

--
-- 转存表中的数据 `micro_link`
--

INSERT INTO `micro_link` (`id`, `status`, `rating`, `list_order`, `create_time`, `update_time`,`description`, `url`, `name`, `image`, `target`, `rel`) VALUES
(1, 1, 1, 8, 1329633709, 1329633709,'百度', 'http://www.baidu.com', 'baidu', '', '_blank', '');

-- --------------------------------------------------------

--
-- 表的结构 `micro_category`
--

CREATE TABLE IF NOT EXISTS `micro_category` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `parent_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '分类父id',
  `level` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '级别:1-一级,2-二级,3-三级',
  `article_count` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '分类文章数',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态:1-发布,2-不发布',
  `delete_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间',
  `list_order` float NOT NULL DEFAULT '10000' COMMENT '排序',
  `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '分类名称',
  `description` varchar(255) NOT NULL DEFAULT '' COMMENT '分类描述',
  `seo_title` varchar(100) NOT NULL DEFAULT '',
  `seo_keywords` varchar(255) NOT NULL DEFAULT '',
  `seo_description` varchar(255) NOT NULL DEFAULT '',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新时间',
  `one_tpl` varchar(50) NOT NULL DEFAULT '' COMMENT '分类文章页模板',
  `more` text COMMENT '扩展属性',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='分类表';

-- --------------------------------------------------------

--
-- 表的结构 `micro_category_article`
--

CREATE TABLE IF NOT EXISTS `micro_category_article` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `article_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '文章id',
  `category_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '分类id',
  `list_order` float NOT NULL DEFAULT '10000' COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `term_taxonomy_id` (`category_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='分类对应表';

-- --------------------------------------------------------

--
-- 表的结构 `micro_article`
--

CREATE TABLE IF NOT EXISTS `micro_article` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `article_type` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '类型:1-文章,2-页面',
  `article_format` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '内容格式:1-html,2-md',
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '发表者用户id',
  `article_status` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态:1-已发布,0-未发布',
  `comment_status` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '评论状态:1-允许,0-不允许',
  `is_top` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否置顶:1-置顶,2-不置顶',
  `recommended` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否推荐:1-推荐,2-不推荐',
  `article_hits` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '查看数',
  `article_favorites` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '收藏数' ,
  `article_like` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '点赞数',
  `comment_count` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '评论数',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新时间',
  `published_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '发布时间',
  `delete_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间',
  `article_title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'article',
  `article_keywords` varchar(150) NOT NULL DEFAULT '' COMMENT 'seo keywords',
  `article_excerpt` varchar(500) NOT NULL DEFAULT '' COMMENT 'article',
  `thumbnail` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '缩略图',
  `article_content` text COMMENT '文章内容',
  `more` text COMMENT '扩展属性:如缩略图,格式为json',
  PRIMARY KEY (`id`),
  KEY `type_status_date` (`article_type`,`article_status`,`create_time`,`id`),
  KEY `user_id` (`user_id`),
  KEY `create_time` (`create_time`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='文章表' ROW_FORMAT=COMPACT;

-- --------------------------------------------------------

--
-- 表的结构 `micro_tag`
--

CREATE TABLE IF NOT EXISTS `micro_tag` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '标签id',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态:1-发布,2-不发布',
  `recommended` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '是否推荐:1-推荐,2-不推荐',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新时间',
  `name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '标签名称',
  UNIQUE KEY `name` (`name`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='portal应用 文章标签表';

-- --------------------------------------------------------

--
-- 表的结构 `micro_tag_article`
--

CREATE TABLE IF NOT EXISTS `micro_tag_article` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `tag_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '标签 id',
  `article_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '文章 id',
  PRIMARY KEY (`id`),
  KEY `article_id` (`article_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='标签文章对应表';

-- --------------------------------------------------------

--
-- 表的结构 `micro_role`
--

CREATE TABLE IF NOT EXISTS `micro_role` (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `parent_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '父角色ID',
  `role_status` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '状态:1-正常,2-禁用',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '更新时间',
  `list_order` float NOT NULL DEFAULT '0' COMMENT '排序',
  `role_name` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`),
  UNIQUE KEY (`role_name`),
  KEY `parent_id` (`parent_id`),
  KEY `role_status` (`role_status`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

INSERT INTO `micro_role` (`id`, `parent_id`, `role_status`, `create_time`, `update_time`, `list_order`, `role_name`, `remark`) VALUES
(1, 0, 1, 1329633709, 1329633709, 0,  'super', '拥有网站最高管理员权限！'),
(2, 1, 1, 1329633709, 1329633709, 0,  'admin', '管理员权限！'),
(3, 1, 1, 1329633709, 1329633709, 10, 'star', 'star'),
(4, 1, 1, 1329633709, 1329633709, 10, 'ad', 'ad');

-- --------------------------------------------------------

--
-- 表的结构 `micro_role_user`
--

CREATE TABLE IF NOT EXISTS `micro_role_user` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `role_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '角色 id',
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户id',
  PRIMARY KEY (`id`),
  KEY `role_id` (`role_id`),
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='用户角色对应表';
INSERT INTO `micro_role_user` (`id`, `role_id`, `user_id`) VALUES
(1, 1, 1),
(2, 2, 2),
(3, 3, 3),
(4, 4, 4);
-- --------------------------------------------------------

--
-- 表的结构 `micro_user`
--

CREATE TABLE IF NOT EXISTS `micro_user` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `bind_id` varchar(32) NOT NULL DEFAULT ''  COMMENT '绑定第三方账号ID',
  `user_type` tinyint(3) UNSIGNED NOT NULL DEFAULT '2' COMMENT '用户类型:1-admin,2-star,3-ad',
  `sex` tinyint(2) NOT NULL DEFAULT '1' COMMENT '性别:1-保密,2-男,3-女',
  `birthday` int(11) NOT NULL DEFAULT '0' COMMENT '生日',
  `last_login_time` int(11) NOT NULL DEFAULT '0' COMMENT '最后登录时间',
  `score` int(11) NOT NULL DEFAULT '0' COMMENT '用户等级',
  `coin` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '金币',
  `balance` decimal(16,2) NOT NULL DEFAULT '0.00' COMMENT '余额',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '注册时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `user_status` tinyint(3) UNSIGNED NOT NULL DEFAULT '3' COMMENT '用户状态:1-正常,2-冻结,3-禁用',
  `user_login` varchar(32) NOT NULL DEFAULT ''  COMMENT '用户名',
  `user_pass` varchar(64) NOT NULL DEFAULT '' COMMENT '登录密码:micro_password加密',
  `pay_pass` varchar(64) NOT NULL DEFAULT '' COMMENT '支付密码:micro_password加密',
  `freeze_time` int(11) NOT NULL DEFAULT '0' COMMENT '冻结时间',
  `user_nickname` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '用户昵称',
  `user_email` varchar(100) NOT NULL DEFAULT '' COMMENT '用户登录邮箱',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '用户头像',
  `signature` varchar(255) NOT NULL DEFAULT '' COMMENT '个性签名',
  `last_login_ip` varchar(15) NOT NULL DEFAULT '' COMMENT '最后登录ip',
  `user_activation_key` varchar(60) NOT NULL DEFAULT '' COMMENT '激活码',
  `mobile` VARCHAR(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '' COMMENT '中国手机不带国家代码.国际手机号格式为：国家代码-手机号',
  `more` text COMMENT '扩展属性',
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_login` (`user_login`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

INSERT INTO `micro_user` (`id`, `user_type`,`score`, `create_time`, `update_time`, `user_status`,`user_login`, `user_pass`, `user_nickname`) VALUES
(1, 1, 99, 1329633709, 1329633709, 1,'super','87aa030ee466817a62dee4d8dc396891', '超级管理员'),
(2, 1, 99, 1329633709, 1329633709, 1,'admin','87aa030ee466817a62dee4d8dc396891', '管理员'),
(3, 2, 1, 1329633709, 1329633709, 1,'wh','87aa030ee466817a62dee4d8dc396891', 'star'),
(4, 3, 1, 1329633709, 1329633709, 1,'ad','87aa030ee466817a62dee4d8dc396891', 'ad');

-- --------------------------------------------------------
--
-- 表的结构 `micro_bind_user`
--

CREATE TABLE IF NOT EXISTS `micro_bind_user` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` varchar(20) NOT NULL COMMENT '绑定账号',
  `start` int(11) NOT NULL DEFAULT '2' COMMENT '粉丝数',
  `region_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '地域',
  `class_id` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '分类',
  `bind_status` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '绑定状态:1-已绑定,2-未绑定,3-审核中',
  `pass_time` int(11) NOT NULL DEFAULT '0' COMMENT '自动审核通过日期',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='绑定账号表';

-- --------------------------------------------------------

--
-- 表的结构 `micro_category`
--

CREATE TABLE IF NOT EXISTS `micro_category` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '分类id',
  `parent_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '分类父id',
  `level` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态:1一级,2-二级,3-三级',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '状态:1-发布,2-不发布,3-软删除',
  `delete_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '删除时间',
  `list_order` float NOT NULL DEFAULT '10000' COMMENT '排序',
  `name` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '分类名称',
  `description` varchar(200) NOT NULL DEFAULT '' COMMENT '分类描述',
  `seo_title` varchar(100) NOT NULL DEFAULT '',
  `seo_keywords` varchar(200) NOT NULL DEFAULT '',
  `seo_description` varchar(200) NOT NULL DEFAULT '',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '注册时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `one_tpl` varchar(50) NOT NULL DEFAULT '' COMMENT '分类模板',
  UNIQUE KEY `name` (`name`),
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='分类表';


-- --------------------------------------------------------
--
-- 表的结构 `micro_task_list_201904`
--

CREATE TABLE IF NOT EXISTS `micro_task_list_201904` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT '1' COMMENT '发布者id',
  `class_id` bigint(20) UNSIGNED NOT NULL DEFAULT '1'  COMMENT '分类id',
  `count` int(15) NOT NULL DEFAULT '1' COMMENT '任务数量',
  `consume_count` int(15) NOT NULL DEFAULT '0' COMMENT '已接单任务数量',
  `check_count` int(15) NOT NULL DEFAULT '0' COMMENT '提交审核任务数量',
  `finish_count` int(15) NOT NULL DEFAULT '0' COMMENT '完成任务数量',
  `balance` decimal(16,2) NOT NULL DEFAULT '0.00' COMMENT '任务单价',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '发布日期',
  `end_time` int(11) NOT NULL DEFAULT '0' COMMENT '结束日期',
  `list_order` float NOT NULL DEFAULT '10000' COMMENT '排序',
  `region_rule` bigint(10) UNSIGNED NOT NULL DEFAULT '0'  COMMENT '地域id限制',
  `class_rule` tinyint(2) UNSIGNED NOT NULL DEFAULT '0' COMMENT '类别限制:0-不限制,1-同类别限制',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '发布状态:0-未发布,1-审核中,2-已发布',
  `task_title` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '任务标题',
  `task_describe` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '任务简述',
  `thumbnail` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '缩略图',
  `task_content` text COMMENT '任务内容',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='任务表';

-- --------------------------------------------------------
--
-- 表的结构 `micro_task_action_201904`
--

CREATE TABLE IF NOT EXISTS `micro_task_action_201904` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `task_id` bigint(20) UNSIGNED NOT NULL DEFAULT '1' COMMENT '任务id',
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT '1' COMMENT '用户id',
  `region_id` int(10) UNSIGNED NOT NULL DEFAULT '1' COMMENT '用户地域',
  `class_id` int(10) UNSIGNED NOT NULL DEFAULT '1' COMMENT '用户分类',
  `balance` decimal(16,2) NOT NULL DEFAULT '0.00' COMMENT '任务单价',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '接单日期',
  `end_time` int(11) NOT NULL DEFAULT '0' COMMENT '结束日期',
  `status` tinyint(3) UNSIGNED NOT NULL DEFAULT '0' COMMENT '状态:0-未完成,1-自动审核中,2-已完成,3-复审,4-无效',
  `check_count` tinyint(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '复审次数',
  `qr_url` varchar(100) NOT NULL DEFAULT 'noset' COMMENT '二维码路径',
  `comment_level` tinyint(3) UNSIGNED NOT NULL DEFAULT '4' COMMENT '评论等级',
  `comment_text` text COMMENT '评论内容',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8 COMMENT='任务接单表';

-- --------------------------------------------------------
--
-- 表的结构 `micro_task_check_log_201904`
--

CREATE TABLE IF NOT EXISTS `micro_task_check_log_201904` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT '1' COMMENT '用户id',
  `task_id` bigint(20) UNSIGNED NOT NULL DEFAULT '1' COMMENT '任务id',
  `reason_describe` varchar(200) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '原因描述',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='任务复审表';


-- --------------------------------------------------------

--
-- 表的结构 `micro_user_action_log_201904`
--

CREATE TABLE IF NOT EXISTS `micro_user_action_log_201904` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户id',
  `count` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '访问次数',
  `last_visit_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '最后访问时间',
  `parameter` varchar(100) NOT NULL DEFAULT '' COMMENT '参数',
  `action` varchar(50) NOT NULL DEFAULT '' COMMENT '操作名称:url',
  `ip` varchar(15) NOT NULL DEFAULT '' COMMENT '用户ip',
  PRIMARY KEY (`id`),
  KEY `user_parameter_action` (`user_id`,`parameter`,`action`),
  KEY `user_parameter_action_ip` (`user_id`,`parameter`,`action`,`ip`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='访问记录表';

-- --------------------------------------------------------

--
-- 表的结构 `micro_user_balance_log_201904`
--

CREATE TABLE IF NOT EXISTS `micro_user_balance_log_201904` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户 id',
  `create_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '创建时间',
  `change` decimal(16,2) NOT NULL DEFAULT '0.00' COMMENT '更改余额',
  `balance` decimal(16,2) NOT NULL DEFAULT '0.00' COMMENT '更改后余额',
  `fee` decimal(16,2) NOT NULL DEFAULT '0.00' COMMENT '手续费',
  `description` varchar(255) NOT NULL DEFAULT '' COMMENT '描述',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户余额变更日志表';

-- --------------------------------------------------------

--
-- 表的结构 `micro_user_favorite`
--

CREATE TABLE IF NOT EXISTS `micro_user_favorite` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户 id',
  `title` varchar(100) NOT NULL DEFAULT '' COMMENT '收藏内容的标题',
  `thumbnail` VARCHAR(100) NOT NULL DEFAULT '' COMMENT '缩略图',
  `url` VARCHAR(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '收藏内容的地址，JSON格式',
  `description` TEXT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '收藏内容的描述',
  `create_time` int(10) UNSIGNED DEFAULT '0' COMMENT '收藏时间',
  PRIMARY KEY (`id`),
  KEY `uid` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户收藏表';

-- --------------------------------------------------------

--
-- 表的结构 `micro_user_login_attempt`
--

CREATE TABLE IF NOT EXISTS `micro_user_login_attempt` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) UNSIGNED NOT NULL DEFAULT '0' COMMENT '用户 id',
  `login_attempts` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '尝试次数',
  `attempt_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '尝试登录时间',
  `locked_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '锁定时间',
  `ip` varchar(15) NOT NULL DEFAULT '' COMMENT '用户 ip',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户登录尝试表' ROW_FORMAT=COMPACT;

-- --------------------------------------------------------

--
-- 表的结构 `micro_user_score_log_201904`
--

CREATE TABLE IF NOT EXISTS `micro_user_score_log_201904` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) NOT NULL DEFAULT '0' COMMENT '用户 id',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `action` varchar(50) NOT NULL DEFAULT '' COMMENT '用户操作名称',
  `score` int(11) NOT NULL DEFAULT '0' COMMENT '更改积分,可以为负',
  `coin` int(11) NOT NULL DEFAULT '0' COMMENT '更改金币,可以为负',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户操作奖励日志表';


-- --------------------------------------------------------

--
-- 表的结构 `micro_verification_code_201904`
--

CREATE TABLE IF NOT EXISTS `micro_verification_code_201904` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '表id',
  `count` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '当天已经发送成功的次数',
  `send_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '最后发送成功时间',
  `expire_time` int(10) UNSIGNED NOT NULL DEFAULT '0' COMMENT '验证码过期时间',
  `code` varchar(8) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '最后发送成功的验证码',
  `account` varchar(100) CHARACTER SET utf8 NOT NULL DEFAULT '' COMMENT '手机号或者邮箱',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='手机邮箱数字验证码表';


-- --------------------------------------------------------

--
-- 表的结构 `micro_user_like_201904`
--
CREATE TABLE IF NOT EXISTS `micro_user_like_201904` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '用户 id',
  `object_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '内容原来的主键id',
  `create_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建时间',
  `url` varchar(255) NOT NULL DEFAULT '' COMMENT '内容的地址',
  `title` varchar(100) NOT NULL DEFAULT '' COMMENT '内容的标题',
  `thumbnail` varchar(100) NOT NULL DEFAULT '' COMMENT '缩略图',
  `description` text COMMENT '内容的描述',
  PRIMARY KEY (`id`),
  KEY `uid` (`user_id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户点赞表';


-- --------------------------------------------------------

--
-- 表的结构 `micro_option`
--
CREATE TABLE IF NOT EXISTS `micro_option` (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `autoload` tinyint(3) UNSIGNED NOT NULL DEFAULT '1' COMMENT '是否自动加载:1-自动加载,0-不自动加载',
  `option_name` varchar(64) NOT NULL DEFAULT '' COMMENT '配置名',
  `option_value` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci COMMENT '配置值',
  PRIMARY KEY (`id`),
  UNIQUE KEY `option_name` (`option_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='全站配置表' ROW_FORMAT=COMPACT;