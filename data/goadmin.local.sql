-- MySQL dump 10.13  Distrib 5.1.56, for Win32 (ia32)
--
-- Host: localhost    Database: goadmin
-- ------------------------------------------------------
-- Server version	5.1.56-community-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS=0;

--
-- Table structure for table `system_access`
--

DROP TABLE IF EXISTS `system_access`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_access` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` int(11) NOT NULL COMMENT '角色ID',
  `module_id` int(11) NOT NULL COMMENT '模块ID',
  PRIMARY KEY (`id`),
  UNIQUE KEY `unique_index` (`role_id`,`module_id`)
) ENGINE=InnoDB AUTO_INCREMENT=867 DEFAULT CHARSET=utf8 COMMENT='角色权限';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_access`
--

LOCK TABLES `system_access` WRITE;
/*!40000 ALTER TABLE `system_access` DISABLE KEYS */;
INSERT INTO `system_access` VALUES (828,1,2),(829,1,3),(830,1,5),(831,1,6),(832,1,8),(833,1,9),(834,1,10),(835,1,11),(836,1,12),(837,1,14),(838,1,15),(839,1,16),(840,1,17),(841,1,19),(842,1,20),(843,1,21),(844,1,22),(845,1,24),(851,1,27),(852,1,28),(853,1,29),(854,1,30),(855,1,31),(856,1,32),(857,1,33),(858,1,34),(859,1,35),(860,1,36),(846,1,37),(847,1,38),(848,1,39),(849,1,40),(850,1,41),(861,1,42),(862,1,43),(863,1,44),(864,1,45),(865,1,46),(866,1,47),(810,2,2),(811,2,3),(812,2,24),(818,2,27),(819,2,28),(820,2,29),(821,2,30),(822,2,31),(823,2,32),(824,2,33),(825,2,34),(826,2,35),(827,2,36),(813,2,37),(814,2,38),(815,2,39),(816,2,40),(817,2,41);
/*!40000 ALTER TABLE `system_access` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_module`
--

DROP TABLE IF EXISTS `system_module`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_module` (
  `module_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '模块ID',
  `module_name` varchar(255) NOT NULL COMMENT '模块名',
  `method_name` varchar(255) NOT NULL DEFAULT '' COMMENT '方法名',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '功能描述',
  `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '上级模块ID',
  `sort` int(11) NOT NULL DEFAULT '0' COMMENT '排序',
  `show` tinyint(4) NOT NULL DEFAULT '0' COMMENT '是否界面显示',
  `type` tinyint(4) NOT NULL DEFAULT '0' COMMENT '类型',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态',
  PRIMARY KEY (`module_id`)
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8 COMMENT='系统功能模块';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_module`
--

LOCK TABLES `system_module` WRITE;
/*!40000 ALTER TABLE `system_module` DISABLE KEYS */;
INSERT INTO `system_module` VALUES (1,'Admin','','基础服务',0,1,1,1,1),(2,'Admin','Main','主界面',1,1,1,2,1),(3,'Admin','Logout','登出系统',1,3,0,2,1),(4,'Module','','模块管理',0,3,1,1,1),(5,'Module','Index','功能列表',4,1,1,2,1),(6,'Module','Save','添加/修改模块',4,1,0,2,1),(7,'Role','','角色管理',0,2,1,1,1),(8,'Role','Index','角色列表',7,1,1,2,1),(9,'Role','Access','获取角色权限',7,1,0,2,1),(10,'Role','SetAccess','更新角色权限',7,1,0,2,1),(11,'Role','Save','添加/修改角色',7,1,0,2,1),(12,'Role','ViewRole','查看角色用户',7,1,0,2,1),(13,'System','','系统管理',0,2,1,1,1),(14,'System','Index','字典管理',13,4,1,2,1),(15,'System','Deldictvalue','删除字典值',13,4,0,2,1),(16,'System','Savedictvalue','添加字典值',13,5,0,2,1),(17,'System','Adddict','添加字典',13,2,0,2,1),(18,'Sysuser','','用户管理',0,1,1,1,1),(19,'Sysuser','Index','用户列表',18,1,1,2,1),(20,'Sysuser','Save','添加/修改用户',18,2,0,2,1),(21,'Sysuser','updateStatus','禁用/解禁用户',18,3,0,2,1),(22,'Sysuser','resetPasswd','重置用户密码',18,4,0,2,1),(23,'Stat','','统计管理',0,1,1,1,1),(24,'Stat','Index','打点数据',23,2,1,2,1),(25,'IdleCrafting','','游戏管理',0,1,1,1,1),(27,'IdleCrafting','Index','IdleCrafting',25,2,1,2,1),(28,'IdleCrafting','ViewPets','宠物列表',25,3,0,2,1),(29,'IdleCrafting','ViewEffects','效果列表',25,4,0,2,1),(30,'IdleCrafting','ViewMaterials','材料列表',25,5,0,2,1),(31,'IdleCrafting','ViewSkills','技能列表',25,6,0,2,1),(32,'IdleCrafting','PetEdit','修改玩家宠物',25,7,0,2,1),(33,'IdleCrafting','EffectEdit','修改玩家效果',25,8,0,2,1),(34,'IdleCrafting','MaterialEdit','修改玩家材料',25,9,0,2,1),(35,'IdleCrafting','SkillEdit','修改玩家技能',25,10,0,2,1),(36,'IdleCrafting','BaseEdit','修改玩家数据',25,11,0,2,1),(37,'Report','Active','活跃留存',23,3,1,2,1),(38,'Report','Online','在线时长',23,4,1,2,1),(39,'Report','Start','启动次数',23,5,1,2,1),(40,'Report','Progress','点击转化',23,6,1,2,1),(41,'Report','Scene','场景漏斗',23,7,1,2,1),(42,'Smash','Index','粉碎水果',25,12,1,2,1),(43,'Smash','BaseEdit','修改玩家数据',25,13,0,2,1),(44,'Qise','Index','七色板',25,14,1,2,1),(45,'Qise','BaseEdit','修改玩家数据',25,15,0,2,1),(46,'Banzi','Index','板子英雄',25,16,1,2,1),(47,'Banzi','BaseEdit','修改玩家数据',25,17,0,2,1);
/*!40000 ALTER TABLE `system_module` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_role`
--

DROP TABLE IF EXISTS `system_role`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_role` (
  `role_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `role_name` varchar(255) NOT NULL COMMENT '角色名',
  `description` varchar(255) NOT NULL COMMENT '描述',
  `status` tinyint(4) NOT NULL COMMENT '状态',
  PRIMARY KEY (`role_id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='系统角色';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_role`
--

LOCK TABLES `system_role` WRITE;
/*!40000 ALTER TABLE `system_role` DISABLE KEYS */;
INSERT INTO `system_role` VALUES (1,'超级管理员','超级管理员',1),(2,'IdleCrafting','IdleCrafting',1);
/*!40000 ALTER TABLE `system_role` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_setting`
--

DROP TABLE IF EXISTS `system_setting`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_setting` (
  `k` varchar(32) NOT NULL COMMENT '键',
  `v` text NOT NULL COMMENT '值',
  `d` varchar(100) NOT NULL DEFAULT '' COMMENT '描述',
  PRIMARY KEY (`k`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='系统设置';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_setting`
--

LOCK TABLES `system_setting` WRITE;
/*!40000 ALTER TABLE `system_setting` DISABLE KEYS */;
INSERT INTO `system_setting` VALUES ('IdleCrafting','{\"initGold\":{\"k\":\"initGold\",\"d\":\"50000\",\"at\":1539392557,\"opt\":\"admin\"}}','IdleCrafting');
/*!40000 ALTER TABLE `system_setting` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `system_user`
--

DROP TABLE IF EXISTS `system_user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `system_user` (
  `user_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
  `username` varchar(255) NOT NULL DEFAULT '' COMMENT '用户名',
  `password` varchar(32) NOT NULL DEFAULT '' COMMENT '密码',
  `email` varchar(255) NOT NULL DEFAULT '' COMMENT '邮箱',
  `realname` varchar(255) NOT NULL DEFAULT '' COMMENT '姓名',
  `role_id` tinyint(4) NOT NULL DEFAULT '0' COMMENT '角色',
  `status` tinyint(4) NOT NULL DEFAULT '0' COMMENT '状态',
  `created` datetime NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '创建时间',
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COMMENT='系统用户';
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `system_user`
--

LOCK TABLES `system_user` WRITE;
/*!40000 ALTER TABLE `system_user` DISABLE KEYS */;
INSERT INTO `system_user` VALUES (1,'admin','21232f297a57a5a743894a0e4a801fc3','admin@misaiyer.com','Admin',1,1,'2017-08-28 00:00:00'),(2,'idlecrafting','9739ceb2f5cb07ae37acfe67bdfbc92f','admin@misaiyer.com','IdleCrafting',2,1,'2018-10-13 01:03:55');
/*!40000 ALTER TABLE `system_user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-10-26 15:27:42
SET FOREIGN_KEY_CHECKS=1;
