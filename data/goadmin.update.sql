/*
Navicat MySQL Data Transfer

Source Server         : localhost
Source Server Version : 50711
Source Host           : localhost:3306
Source Database       : blog

Target Server Type    : MYSQL
Target Server Version : 50711
File Encoding         : 65001

Date: 2017-08-29 17:17:10
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS=0;

INSERT INTO `system_access` (`role_id`, `module_id`) VALUES ('2', '37');
INSERT INTO `system_access` (`role_id`, `module_id`) VALUES ('2', '38');
INSERT INTO `system_access` (`role_id`, `module_id`) VALUES ('2', '39');

INSERT INTO `system_access` (`role_id`, `module_id`) VALUES (2,40),(2,41);

INSERT INTO `system_module` VALUES (37,'Report','Active','活跃留存',23,3,1,2,1),(38,'Report','Online','在线时长',23,4,1,2,1),(39,'Report','Start','启动次数',23,5,1,2,1);
INSERT INTO `system_module` VALUES (40,'Report','Progress','点击转化',23,6,1,2,1),(41,'Report','Scene','场景漏斗',23,7,1,2,1);

INSERT INTO `system_access` VALUES (828,1,2),(829,1,3),(830,1,5),(831,1,6),(832,1,8),(833,1,9),(834,1,10),(835,1,11),(836,1,12),(837,1,14),(838,1,15),(839,1,16),(840,1,17),(841,1,19),(842,1,20),(843,1,21),(844,1,22),(845,1,24),(851,1,27),(852,1,28),(853,1,29),(854,1,30),(855,1,31),(856,1,32),(857,1,33),(858,1,34),(859,1,35),(860,1,36),(846,1,37),(847,1,38),(848,1,39),(849,1,40),(850,1,41),(861,1,42),(862,1,43),(863,1,44),(864,1,45),(865,1,46),(866,1,47),(810,2,2),(811,2,3),(812,2,24),(818,2,27),(819,2,28),(820,2,29),(821,2,30),(822,2,31),(823,2,32),(824,2,33),(825,2,34),(826,2,35),(827,2,36),(813,2,37),(814,2,38),(815,2,39),(816,2,40),(817,2,41);

INSERT INTO `system_module` VALUES (1,'Admin','','基础服务',0,1,1,1,1),(2,'Admin','Main','主界面',1,1,1,2,1),(3,'Admin','Logout','登出系统',1,3,0,2,1),(4,'Module','','模块管理',0,3,1,1,1),(5,'Module','Index','功能列表',4,1,1,2,1),(6,'Module','Save','添加/修改模块',4,1,0,2,1),(7,'Role','','角色管理',0,2,1,1,1),(8,'Role','Index','角色列表',7,1,1,2,1),(9,'Role','Access','获取角色权限',7,1,0,2,1),(10,'Role','SetAccess','更新角色权限',7,1,0,2,1),(11,'Role','Save','添加/修改角色',7,1,0,2,1),(12,'Role','ViewRole','查看角色用户',7,1,0,2,1),(13,'System','','系统管理',0,2,1,1,1),(14,'System','Index','字典管理',13,4,1,2,1),(15,'System','Deldictvalue','删除字典值',13,4,0,2,1),(16,'System','Savedictvalue','添加字典值',13,5,0,2,1),(17,'System','Adddict','添加字典',13,2,0,2,1),(18,'Sysuser','','用户管理',0,1,1,1,1),(19,'Sysuser','Index','用户列表',18,1,1,2,1),(20,'Sysuser','Save','添加/修改用户',18,2,0,2,1),(21,'Sysuser','updateStatus','禁用/解禁用户',18,3,0,2,1),(22,'Sysuser','resetPasswd','重置用户密码',18,4,0,2,1),(23,'Stat','','统计管理',0,1,1,1,1),(24,'Stat','Index','打点数据',23,2,1,2,1),(25,'IdleCrafting','','游戏管理',0,1,1,1,1),(27,'IdleCrafting','Index','IdleCrafting',25,2,1,2,1),(28,'IdleCrafting','ViewPets','宠物列表',25,3,0,2,1),(29,'IdleCrafting','ViewEffects','效果列表',25,4,0,2,1),(30,'IdleCrafting','ViewMaterials','材料列表',25,5,0,2,1),(31,'IdleCrafting','ViewSkills','技能列表',25,6,0,2,1),(32,'IdleCrafting','PetEdit','修改玩家宠物',25,7,0,2,1),(33,'IdleCrafting','EffectEdit','修改玩家效果',25,8,0,2,1),(34,'IdleCrafting','MaterialEdit','修改玩家材料',25,9,0,2,1),(35,'IdleCrafting','SkillEdit','修改玩家技能',25,10,0,2,1),(36,'IdleCrafting','BaseEdit','修改玩家数据',25,11,0,2,1),(37,'Report','Active','活跃留存',23,3,1,2,1),(38,'Report','Online','在线时长',23,4,1,2,1),(39,'Report','Start','启动次数',23,5,1,2,1),(40,'Report','Progress','点击转化',23,6,1,2,1),(41,'Report','Scene','场景漏斗',23,7,1,2,1),(42,'Smash','Index','粉碎水果',25,12,1,2,1),(43,'Smash','BaseEdit','修改玩家数据',25,13,0,2,1),(44,'Qise','Index','彩虹七巧板',25,14,1,2,1),(45,'Qise','BaseEdit','修改玩家数据',25,15,0,2,1),(46,'Banzi','Index','板子英雄',25,16,1,2,1),(47,'Banzi','BaseEdit','修改玩家数据',25,17,0,2,1);

SET FOREIGN_KEY_CHECKS=1;

INSERT INTO `system_module` VALUES (42,'Smash','Index','粉碎水果',25,12,1,2,1),(43,'Smash','BaseEdit','修改玩家数据',25,13,0,2,1),(44,'Qise','Index','彩虹七巧板',25,14,1,2,1),(45,'Qise','BaseEdit','修改玩家数据',25,15,0,2,1),(46,'Banzi','Index','板子英雄',25,16,1,2,1),(47,'Banzi','BaseEdit','修改玩家数据',25,17,0,2,1);

ALTER TABLE system_module ADD icon VARCHAR(255) NOT NULL;

UPDATE system_module SET icon = "fa-dashboard";

UPDATE system_module SET icon = "fa-dashboard" WHERE module_id = 1;
UPDATE system_module SET icon = "fa-home" WHERE module_id = 2;
UPDATE system_module SET icon = "fa-sign-out" WHERE module_id = 3;
UPDATE system_module SET icon = "fa-object-group" WHERE module_id = 4;
UPDATE system_module SET icon = "fa-navicon" WHERE module_id = 5;
UPDATE system_module SET icon = "fa-pencil" WHERE module_id = 6;
UPDATE system_module SET icon = "fa-users" WHERE module_id = 7;
UPDATE system_module SET icon = "fa-user" WHERE module_id = 8;
UPDATE system_module SET icon = "fa-user-plus" WHERE module_id = 9;
UPDATE system_module SET icon = "fa-user-secret" WHERE module_id = 10;
UPDATE system_module SET icon = "fa-user-times" WHERE module_id = 11;
UPDATE system_module SET icon = "fa-user" WHERE module_id = 12;
UPDATE system_module SET icon = "fa-tv" WHERE module_id = 13;
UPDATE system_module SET icon = "fa-tasks" WHERE module_id = 14;
UPDATE system_module SET icon = "fa-close" WHERE module_id = 15;
UPDATE system_module SET icon = "fa-plus" WHERE module_id = 16;
UPDATE system_module SET icon = "fa-plus" WHERE module_id = 17;

UPDATE system_module SET icon = "fa-paw" WHERE module_id = 18;
UPDATE system_module SET icon = "fa-child" WHERE module_id = 19;
UPDATE system_module SET icon = "fa-paint-brush" WHERE module_id = 20;
UPDATE system_module SET icon = "fa-remove" WHERE module_id = 21;
UPDATE system_module SET icon = "fa-retweet" WHERE module_id = 22;

UPDATE system_module SET icon = "fa-cubes" WHERE module_id = 23;
UPDATE system_module SET icon = "fa-cube" WHERE module_id = 24;

UPDATE system_module SET icon = "fa-gamepad" WHERE module_id = 25;
UPDATE system_module SET icon = "fa-life-bouy" WHERE module_id = 27;

UPDATE system_module SET icon = "fa-line-chart" WHERE module_id = 37;
UPDATE system_module SET icon = "fa-calendar-check-o" WHERE module_id = 38;
UPDATE system_module SET icon = "fa-leaf" WHERE module_id = 39;
UPDATE system_module SET icon = "fa-share-alt" WHERE module_id = 40;
UPDATE system_module SET icon = "fa-hourglass-end" WHERE module_id = 41;

UPDATE system_module SET icon = "fa-pie-chart" WHERE module_id = 42;
UPDATE system_module SET icon = "fa-star-half-o" WHERE module_id = 44;
UPDATE system_module SET icon = "fa-road" WHERE module_id = 46;


INSERT INTO `system_module` VALUES (48,'Project','','项目管理',0,1,1,1,1,'fa-list'),(49,'Project','Index','项目列表',48,2,1,2,1,'fa-list-alt'),(50,'Project','Save','添加/修改项目',48,3,0,2,1,'fa-edit');

--
-- Table structure for table `system_project`
--

DROP TABLE IF EXISTS `system_project`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_project` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `project_id` int(11) NOT NULL DEFAULT '0',
  `name` varchar(255) NOT NULL DEFAULT '',
  `api` varchar(255) NOT NULL DEFAULT '',
  `status` tinyint(4) NOT NULL DEFAULT '0',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  `path` varchar(255) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `project_id` (`project_id`)
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_project`
--

LOCK TABLES `system_project` WRITE;
/*!40000 ALTER TABLE `system_project` DISABLE KEYS */;
INSERT INTO `system_project` VALUES (2,2,'肥鱼向上飞','',0,'2018-10-27 09:32:12','2018-10-27 09:32:12',''),(3,4,'天天蹦一蹦','',0,'2018-10-27 09:32:28','2018-10-27 09:32:28',''),(4,5,'17喵喵喵','',1,'2018-10-27 09:32:39','2018-10-27 09:32:39','/home/gameApi/17mmm/'),(5,6,'小猪变变变','',1,'2018-10-27 09:32:53','2018-10-27 09:32:53','/home/gameApi/xzbbb/'),(6,7,'恋之红线','',1,'2018-10-27 09:33:04','2018-10-27 09:33:04','/home/gameApi/redline/'),(7,9,'小甜甜2','',1,'2018-10-27 09:33:15','2018-10-27 09:33:15','/home/gameApi/xtt2/'),(8,10,'粉碎水果传奇2','https://smash.4g111.com/v2/',1,'2018-10-27 09:33:50','2018-10-27 09:33:50','/home/gameApi/smash2/'),(9,11,'IdleCrafting','https://idlecrafting.4g111.com/',1,'2018-10-27 09:34:18','2018-10-27 09:34:18','/home/gameApi/idlecrafting/'),(10,16,'板子英雄','https://banzi.4g111.com/',1,'2018-10-27 09:34:49','2018-10-27 09:34:49','/home/gameApi/banzi/'),(11,17,'彩虹七巧板','https://qise.4g111.com/',1,'2018-10-27 09:35:12','2018-10-27 09:35:12','/home/gameApi/qise/'),(12,12,'report','https://report.4g111.com/',0,'2018-10-27 09:36:55','2018-10-27 09:36:55','/home/gameApi/report/'),(13,13,'goadmin','',0,'2018-10-27 09:37:09','2018-10-27 09:37:09','/home/goadmin/'),(14,14,'adminReport','',0,'2018-10-27 09:37:21','2018-10-27 09:37:21','/home/adminReport/'),(15,15,'gendata','',0,'2018-10-27 09:37:36','2018-10-27 09:37:36','/home/gendata/');
/*!40000 ALTER TABLE `system_project` ENABLE KEYS */;
UNLOCK TABLES;


INSERT INTO `system_module` VALUES (51,'Project','SwitchList','分享开关',48,4,1,2,1,'fa-share-alt-square'),(52,'Project','SwitchEdit','编辑分享开关',48,5,0,2,1,'fa-share-square'),(53,'Project','Files','文件列表',48,6,0,2,1,'fa-file-code-o'),(54,'Project','FileEdit','编辑文件',48,7,0,2,1,'fa-file-code-o'),(55,'Project','FileDelete','删除文件',48,8,0,2,1,'fa-file-code-o'),(56,'Project','FileUpload','上传文件',48,9,0,2,1,'fa-file-code-o');

INSERT INTO `system_module` VALUES (57,'Project','Global','配置列表',48,10,1,2,1,'fa-globe'),(58,'Project','GlobalEdit','添加/修改配置',48,11,0,2,1,' fa-globe');
INSERT INTO `system_module` VALUES (59,'Report','Gate','关卡分布',23,8,1,2,1,'fa-map-signs'),(60,'Report','Pet','宠物解锁',23,9,1,2,1,'fa-bug'),(61,'Report','Useskill','技能使用',23,10,1,2,1,'fa-wrench');
INSERT INTO `system_module` VALUES (62,'Project','Task','任务列表',48,12,1,2,1,'fa-tasks'),(63,'Project','TaskEdit','添加/修改任务',48,13,0,2,1,'fa-tasks');
INSERT INTO `system_module` VALUES (64,'Report','GateGrow','前进关卡',23,11,1,2,1,'fa-line-chart'),(65,'Report','GateStop','关卡停留',23,12,1,2,1,'fa-hand-stop-o'),(66,'Project','Token','令牌签发',48,14,1,2,1,'fa-user-secret');
INSERT INTO `system_module` VALUES (67,'Project','Achiement','成就列表',48,15,1,2,1,'fa-trophy'),(68,'Project','AchiementEdit','添加/修改成就',48,16,0,2,1,'fa-trophy');

INSERT INTO `system_module` VALUES (69,'Admin','Send','系统消息',1,4,0,2,1,' fa-send'),(70,'Admin','Sync','同步消息',1,5,0,2,1,'fa-retweet');
INSERT INTO `system_module` VALUES (71,'IdleCrafting','BaseDelete','删除玩家数据',25,12,0,2,1,'fa-dashboard');
INSERT INTO `system_module` VALUES (72,'Admin','SetProject','项目选择',1,6,0,2,1,'fa-retweet');

INSERT INTO `system_module` VALUES (73,'Project','ApiDoc','接口文档',48,17,1,2,1,'fa-sitemap'),(74,'Project','ApiDocView','查看接口',48,18,0,2,1,'fa-sitemap'),(75,'Project','ApiDocEdit','修改接口',48,19,0,2,1,'fa-sitemap'),(76,'Project','ApiDocAudit','审核接口',48,20,0,2,1,'fa-sitemap'),(77,'Report','GateData','关卡数据',23,13,1,2,1,'fa-th');

INSERT INTO `system_module` VALUES (78,'IdleCrafting','ViewAchs','成就列表',25,19,0,2,1,'fa-archive'),(79,'IdleCrafting','AchEdit','修改玩家成就',25,20,0,2,1,'fa-archive'),(80,'IdleCrafting','ViewTasks','任务列表',25,21,0,2,1,'fa-archive'),(81,'IdleCrafting','TaskEdit','修改玩家任务',25,22,0,2,1,'fa-archive');

INSERT INTO `system_module` VALUES (82,'Report','Realtime','实时数据',23,14,1,2,1,'fa-bar-chart');
INSERT INTO `system_module` VALUES (83,'Report','GateLoss','流失数据',23,15,1,2,1,'fa-sort-amount-desc');
INSERT INTO `system_module` VALUES (84,'Project','Source','来源列表',48,21,1,2,1,'fa-bullseye'),(85,'Project','SourceEdit','添加/修改来源',48,22,0,2,1,'fa-bullseye');
INSERT INTO `system_module` VALUES (86,'Report','Realplay','实时在玩',23,16,1,2,1,'fa-motorcycle');
INSERT INTO `system_module` VALUES (87,'Report','Realtrend','实时趋势',23,17,1,2,1,'fa-line-chart');
INSERT INTO `system_module` VALUES (88,'Report','Share','分享数据',23,18,1,2,1,'fa-share');
INSERT INTO `system_module` VALUES (89,'Report','Realtoday','今日注册',23,19,1,2,1,'fa-eye-slash');
INSERT INTO `system_module` VALUES (90,'Admin','Control','控制台',1,7,1,2,1,'fa-television');
INSERT INTO `system_module` VALUES (91,'Report','GateGold','关卡金币',23,20,1,2,1,'fa-life-ring'),(92,'Report','Draw','画廊分享',23,21,1,2,1,'fa-image');
INSERT INTO `system_module` VALUES (93,'Gunner','Index','神奇枪手',25,23,1,2,1,'fa-plane'),(94,'Gunner','BaseEdit','修改玩家数据',25,24,0,2,1,'fa-plane');
INSERT INTO `system_module` VALUES (95,'Report','Wall','流量墙',23,22,1,2,1,'fa-chrome');
INSERT INTO `system_module` VALUES (96,'Project','Gameuser','玩家配置',48,23,1,2,1,'fa-edit'),(97,'Project','GameuserEdit','添加/修改玩家配置',48,24,0,2,1,'fa-edit'),(98,'Project','GameuserDelete','删除玩家配置',48,25,0,2,1,'fa-edit');
INSERT INTO `system_module` VALUES (99,'Stat','GateRetimes','复活挑战',23,23,1,2,1,'fa-times-circle-o'),(100,'Stat','GateScore','关卡得分',23,24,1,2,1,'fa-star-half-o');
INSERT INTO `system_module` VALUES (101,'Project','GameJson','文件配置',48,26,1,2,1,'fa-file-code-o'),(102,'Project','GameJsonEdit','添加/修改文件配置',48,27,0,2,1,'fa-file-code-o'),(103,'Project','GameJsonDelete','删除文件配置',48,28,0,2,1,'fa-file-code-o');
INSERT INTO `system_module` VALUES (104,'ApiDoc','','接口文档',0,1,1,1,1,'fa-folder-open-o'),(105,'ApiDoc','Show','接口文档展示',104,2,0,2,1,'fa-folder-open-o');
INSERT INTO `system_module` VALUES (106,'Report','Fast','极速模式',23,25,1,2,1,'fa-yc');
INSERT INTO `system_module` VALUES (107,'Project','GameJsonRestore','恢复文件配置',48,29,0,2,1,'fa-file-code-o');
INSERT INTO `system_module` VALUES (108,'Triangle','Index','三角板',25,25,1,2,1,'fa-exclamation-triangle'),(109,'Triangle','BaseEdit','修改三角板玩家数据',25,26,0,2,1,'fa-exclamation-triangle');
INSERT INTO `system_module` VALUES (110,'Server','Index','服务列表',48,30,1,2,1,'fa-server'),(111,'Server','BaseEdit','添加/修改服务',48,31,0,2,1,'fa-server');
INSERT INTO `system_module` VALUES (112,'Server','BaseAct','操作服务',48,32,0,2,1,'fa-server');
INSERT INTO `system_module` VALUES (113,'Report','','图表管理',0,1,1,1,1,'fa-pie-chart'),(114,'Report','','数据管理',0,1,1,1,1,'fa-table');
INSERT INTO `system_module` VALUES (115,'Report','Sign','签到分享',114,26,1,2,1,'fa-map-signs');

INSERT INTO `system_module` VALUES (116,'Project','','项目管理',0,12,1,1,1,'fa-list'),(117,'Project','ApiReq','接口测试',116,33,1,2,1,'fa-wrench');
INSERT INTO `system_module` VALUES (117,'Project','ApiReq','接口测试',116,33,1,2,1,'fa-wrench');
INSERT INTO `system_module` VALUES (118,'Log','','日志管理',0,13,1,1,1,'fa-sticky-note-o'),(119,'Log','Login','登录日志',118,1,1,2,1,'fa-file-text-o'),(120,'Log','Register','注册日志',118,2,1,2,1,'fa-certificate'),(121,'Log','SendClick','分享图点击日志',118,3,1,2,1,'fa-cube'),(122,'Log','Progress','点击转化',118,4,1,2,1,'fa-share-alt'),(123,'Log','SceneAct','场景漏斗',118,5,1,2,1,'fa-hourglass-end'),(124,'Log','Playtime','在线时长',118,6,1,2,1,'fa-calendar-check-o'),(125,'Log','Tip','关卡提示',118,7,1,2,1,'fa-th'),(126,'Log','Wall','流量墙',118,8,1,2,1,'fa-chrome'),(127,'Log','GateGrow','前进关卡',118,9,1,2,1,'fa-line-chart'),(128,'Log','Gold','金币日志',118,10,1,2,1,'fa-life-ring'),(129,'Log','Gate','关卡日志',118,11,1,2,1,'fa-map-signs');
INSERT INTO `system_module` VALUES (130,'Report','Astrology','占星运势',114,27,1,2,1,'fa-shield');
INSERT INTO `system_module` VALUES (131,'Electro','Index','电音三角板',25,27,1,2,1,'fa-music'),(132,'Electro','BaseEdit','修改电音三角板玩家数据',25,28,0,2,1,'fa-music');
INSERT INTO `system_module` VALUES (133,'Stat','GateLoss','新手关卡流失',114,28,1,2,1,'fa-list-alt'),(134,'Stat','PetLevel','宠物主角升级',114,29,1,2,1,'fa-arrow-up');
INSERT INTO `system_module` VALUES (135,'Stat','GateWeapon','合成武器等级',114,30,1,2,1,'fa-fighter-jet');
INSERT INTO `system_module` VALUES (136,'Stat','GateGoldDis','金币产出消耗',114,31,1,2,1,'fa-sitemap'),(137,'Stat','GateGoldDep','金币存量',114,32,1,2,1,'fa-sliders');
INSERT INTO `system_module` VALUES (138,'FuDai','FuDai','福券配置',48,30,1,2,1,'fa-gift'),(139,'FuDai','FuDaiEdit','添加/修改福券配置',48,31,0,2,1,'fa-gift'),(140,'FuDai','FuGate','通关福袋配置',48,32,1,2,1,'fa-glass'),(141,'FuDai','FuGateEdit','添加/修改通关福袋配置',48,33,0,2,1,'fa-glass'),(142,'FuDai','FuDay','每天福袋配置',48,34,1,2,1,'fa-inbox'),(143,'FuDai','FuDayEdit','添加/修改每天福袋配置',48,35,0,2,1,'fa-inbox');
INSERT INTO `system_module` VALUES (144,'Electro','FuDai','获取福袋数据',25,29,0,2,1,'fa-music'),(145,'Electro','FuDaiEdit','修改福袋数据',25,30,0,2,1,'fa-music');
INSERT INTO `system_module` VALUES (146,'Electro','BaseDelete','删除电音玩家',25,31,0,2,1,'fa-music');
INSERT INTO `system_module` VALUES (147,'Project','GameJsonSync','同步文件配置',48,36,0,2,1,'fa-file-code-o');
INSERT INTO `system_module` VALUES (148,'Stat','WeaponLev','武器等级停留',114,33,1,2,1,'fa-map-pin');
INSERT INTO `system_module` VALUES (149,'Card2048','Index','card2048',25,32,1,2,1,'fa-magnet'),(150,'Card2048','BaseEdit','修改card2048玩家数据',25,33,0,2,1,'fa-magnet');
INSERT INTO `system_module` VALUES (151,'Electro','RankDelete','删除电音积分排行',25,34,0,2,1,'fa-music');
INSERT INTO `system_module` VALUES (152,'Card2048','BaseDelete','删除card2048玩家',25,35,0,2,1,'fa-music');
INSERT INTO `system_module` VALUES (153,'Stat','PropList','局数/分数/道具统计',114,34,1,2,1,'fa-anchor');
INSERT INTO `system_module` VALUES (154,'Pintu','Index','拼图大师',25,36,1,2,1,'fa-cubes'),(155,'Pintu','BaseEdit','修改拼图大师玩家数据',25,37,0,2,1,'fa-cubes'),(156,'Pintu','FuDai','获取拼图福袋数据',25,38,0,2,1,'fa-cubes'),(157,'Pintu','FuDaiEdit','修改拼图大师福袋数据',25,39,0,2,1,'fa-cubes'),(158,'Pintu','BaseDelete','删除拼图大师玩家',25,40,0,2,1,'fa-cubes'),(159,'Pintu','RankDelete','删除拼图积分排行',25,41,0,2,1,'fa-cubes');

-- create table `system_chat`
-- --------------------------------------------------
--  Table Structure for `goadmin/models.SystemChat`
-- --------------------------------------------------

-- ALTER DATABASE goadmin DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
-- ALTER TABLE `system_chat` CONVERT TO CHARACTER SET UTF8;

--
-- Table structure for table `system_chat`
--

DROP TABLE IF EXISTS `system_chat`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_chat` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL DEFAULT '0',
  `username` varchar(255) NOT NULL DEFAULT '',
  `message` varchar(255) DEFAULT NULL,
  `content` longtext,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_chat`
--

LOCK TABLES `system_chat` WRITE;
/*!40000 ALTER TABLE `system_chat` DISABLE KEYS */;
INSERT INTO `system_chat` VALUES (1,1,'admin','hello world!','','2018-11-06 09:20:37');
/*!40000 ALTER TABLE `system_chat` ENABLE KEYS */;
UNLOCK TABLES;

create table `system_apidoc`
-- --------------------------------------------------
--  Table Structure for `goadmin/models.ApiDoc`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `system_apidoc` (
    `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `project_id` integer NOT NULL DEFAULT 0 ,
    `type` integer NOT NULL DEFAULT 0 ,
    `status` integer NOT NULL DEFAULT 0 ,
    `sort` integer NOT NULL DEFAULT 0 ,
    `url` varchar(100) NOT NULL DEFAULT '' ,
    `method` integer NOT NULL DEFAULT 1 ,
    `name` varchar(100) NOT NULL DEFAULT '' ,
    `detail` longtext,
    `doc_id` integer NOT NULL DEFAULT 0 ,
    `create_id` integer NOT NULL DEFAULT 0 ,
    `update_id` integer NOT NULL DEFAULT 0 ,
    `audit_id` integer NOT NULL DEFAULT 0 ,
    `audit_time` datetime,
    `created_at` datetime NOT NULL,
    `updated_at` datetime NOT NULL
) ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

ALTER TABLE system_apidoc CHANGE `rank` `sort` integer;

ALTER TABLE `system_role` ADD `project_ids` VARCHAR(255) NULL;

create table `system_server`
-- --------------------------------------------------
--  Table Structure for `goadmin/models.Server`
-- --------------------------------------------------
CREATE TABLE IF NOT EXISTS `system_server` (
    `id` integer AUTO_INCREMENT NOT NULL PRIMARY KEY,
    `name` varchar(30) NOT NULL DEFAULT '' ,
    `host` varchar(20) NOT NULL DEFAULT '' ,
    `port` varchar(20) NOT NULL DEFAULT '' ,
    `path` varchar(255) NOT NULL DEFAULT '' ,
    `status` integer unsigned NOT NULL DEFAULT 0 ,
    `project_id` integer NOT NULL DEFAULT 0 ,
    `jwt_secret` varchar(43) NOT NULL DEFAULT '' ,
    `token` varchar(43),
    `appid` varchar(18),
    `app_secret` varchar(32),
    `group` varchar(20),
    `prefixs` varchar(255),
    `default_db` varchar(255),
    `default_prefix` varchar(20),
    `logger_db` varchar(255),
    `logger_prefix` varchar(20),
    `updated_at` datetime NOT NULL,
    `created_at` datetime NOT NULL,
    UNIQUE (`name`, `host`)
) ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
