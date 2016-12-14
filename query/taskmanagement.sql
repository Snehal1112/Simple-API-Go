-- MySQL dump 10.13  Distrib 5.7.12, for linux-glibc2.5 (x86_64)
--
-- Host: localhost    Database: OkPortofolio
-- ------------------------------------------------------
-- Server version	5.7.16-0ubuntu0.16.04.1

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

--
-- Table structure for table `OkAkun`
--

DROP TABLE IF EXISTS `OkAkun`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `OkAkun` (
  `id_akun` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `password` varchar(200) COLLATE utf8_unicode_ci NOT NULL,
  `token` text COLLATE utf8_unicode_ci NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `role` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`id_akun`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `OkAkun`
--

LOCK TABLES `OkAkun` WRITE;
/*!40000 ALTER TABLE `OkAkun` DISABLE KEYS */;
INSERT INTO `OkAkun` VALUES (1,'abangs','iman','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InVzZXJuYW1lIjoiYWJhbmdzIiwicGhvbmVfbnVtYmVyIjoiIiwiZW1haWwiOiIiLCJuYW1lIjoiSW1hbiBUdW1vcmFuZyIsInJvbGUiOiJhZG1pbiIsImFrdW5faWQiOjEsInByb2ZpbGVfaWQiOjF9fQ.zH414c9sBtFy7izTaDXboJX_iHFoJ83oywQkkNwPYHM','2016-12-13 12:19:30','2016-12-13 12:19:30','admin'),(2,'adeks','emy','eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJkYXRhIjp7InVzZXJuYW1lIjoiYWRla3MiLCJwaG9uZV9udW1iZXIiOiIiLCJlbWFpbCI6IiIsIm5hbWUiOiJFbWluYXJ0eSAgU2lhbnR1cmkiLCJyb2xlIjoiZGV2ZWxvcGVyIiwiYWt1bl9pZCI6MiwicHJvZmlsZV9pZCI6Mn19.2jaO_dCz_hgApXrkjn1uktltVBHw4SqYIVKlFY_VKW8','2016-12-13 12:19:30','2016-12-14 10:05:13','developer'),(3,'adeks2','hosea','sampletoken','2016-12-13 12:19:30','2016-12-13 12:19:30','developer'),(4,'adeks3','gomgom','sampletoken','2016-12-13 12:19:30','2016-12-13 12:19:30','tester');
/*!40000 ALTER TABLE `OkAkun` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `OkProfile`
--

DROP TABLE IF EXISTS `OkProfile`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `OkProfile` (
  `id_profile` int(11) NOT NULL AUTO_INCREMENT,
  `akun_id` int(11) NOT NULL,
  `first_name` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `last_name` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `photo` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `current_status` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id_profile`),
  KEY `fk_akun_profile_idx` (`akun_id`),
  CONSTRAINT `fk_akun_profile` FOREIGN KEY (`akun_id`) REFERENCES `OkAkun` (`id_akun`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `OkProfile`
--

LOCK TABLES `OkProfile` WRITE;
/*!40000 ALTER TABLE `OkProfile` DISABLE KEYS */;
INSERT INTO `OkProfile` VALUES (1,1,'Iman','Tumorang','photo.png','work',NULL,NULL),(2,2,'Eminarty ','Sianturi','photo.png','work',NULL,NULL),(3,3,'Hosea','Juntaks','photo.png','work',NULL,NULL),(4,4,'Gomgom','Purba','photo.png','work',NULL,NULL);
/*!40000 ALTER TABLE `OkProfile` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `OkProject`
--

DROP TABLE IF EXISTS `OkProject`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `OkProject` (
  `id_project` int(11) NOT NULL AUTO_INCREMENT,
  `nama_project` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `desc` text COLLATE utf8_unicode_ci,
  `git_url` text COLLATE utf8_unicode_ci NOT NULL,
  PRIMARY KEY (`id_project`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `OkProject`
--

LOCK TABLES `OkProject` WRITE;
/*!40000 ALTER TABLE `OkProject` DISABLE KEYS */;
INSERT INTO `OkProject` VALUES (1,'Resep Makanan','This is the android version of Resep Makanan','https://sample.github.com'),(2,'Resep API','This is the backend of Resep Makanan','http://sample-backend.github.com');
/*!40000 ALTER TABLE `OkProject` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `OkProjectTask`
--

DROP TABLE IF EXISTS `OkProjectTask`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `OkProjectTask` (
  `id_p_task` int(11) NOT NULL AUTO_INCREMENT,
  `project_id` int(11) NOT NULL,
  `task_id` int(11) NOT NULL,
  `assignedto_id` int(11) NOT NULL,
  `assigner_id` int(11) NOT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `status` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `due_date` datetime DEFAULT NULL,
  PRIMARY KEY (`id_p_task`),
  KEY `fk_project_idx` (`project_id`),
  KEY `fk_task_idx` (`task_id`),
  KEY `fk_assigner_idx` (`assigner_id`),
  KEY `fk_assignedto_idx` (`assignedto_id`),
  CONSTRAINT `fk_assignedto` FOREIGN KEY (`assignedto_id`) REFERENCES `OkProfile` (`id_profile`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_assigner` FOREIGN KEY (`assigner_id`) REFERENCES `OkProfile` (`id_profile`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_project` FOREIGN KEY (`project_id`) REFERENCES `OkProject` (`id_project`) ON DELETE NO ACTION ON UPDATE NO ACTION,
  CONSTRAINT `fk_task` FOREIGN KEY (`task_id`) REFERENCES `OkTask` (`id_task`) ON DELETE NO ACTION ON UPDATE NO ACTION
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `OkProjectTask`
--

LOCK TABLES `OkProjectTask` WRITE;
/*!40000 ALTER TABLE `OkProjectTask` DISABLE KEYS */;
INSERT INTO `OkProjectTask` VALUES (1,1,1,2,1,NULL,NULL,'working',NULL),(2,2,2,1,1,NULL,NULL,'working',NULL);
/*!40000 ALTER TABLE `OkProjectTask` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `OkTask`
--

DROP TABLE IF EXISTS `OkTask`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `OkTask` (
  `id_task` int(11) NOT NULL AUTO_INCREMENT,
  `title_task` varchar(200) COLLATE utf8_unicode_ci NOT NULL,
  `urgency` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `difficulty` varchar(45) COLLATE utf8_unicode_ci NOT NULL,
  `desc_task` longtext COLLATE utf8_unicode_ci,
  PRIMARY KEY (`id_task`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `OkTask`
--

LOCK TABLES `OkTask` WRITE;
/*!40000 ALTER TABLE `OkTask` DISABLE KEYS */;
INSERT INTO `OkTask` VALUES (1,'Membuat Layout Login','5','4','Create Login Layout'),(2,'Membuat API Login','5','5','Create Login API');
/*!40000 ALTER TABLE `OkTask` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Dumping events for database 'OkPortofolio'
--

--
-- Dumping routines for database 'OkPortofolio'
--
/*!50003 DROP PROCEDURE IF EXISTS `get_all_task_by_project` */;
/*!50003 SET @saved_cs_client      = @@character_set_client */ ;
/*!50003 SET @saved_cs_results     = @@character_set_results */ ;
/*!50003 SET @saved_col_connection = @@collation_connection */ ;
/*!50003 SET character_set_client  = utf8 */ ;
/*!50003 SET character_set_results = utf8 */ ;
/*!50003 SET collation_connection  = utf8_general_ci */ ;
/*!50003 SET @saved_sql_mode       = @@sql_mode */ ;
/*!50003 SET sql_mode              = 'ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_AUTO_CREATE_USER,NO_ENGINE_SUBSTITUTION' */ ;
DELIMITER ;;
CREATE DEFINER=`root`@`localhost` PROCEDURE `get_all_task_by_project`(IN ProjectId INT)
BEGIN
	Select 
	c.title_task , c.difficulty , c.urgency , c.desc_task,
	d.first_name as 'Pemberi' , e.first_name as 'Penerima'
	from OkProjectTask a join OkProject b  on a.project_id = b.id_project
	join OkTask c on a.task_id = c.id_task 
	join OkProfile d  on a.assigner_id = d.id_profile
	join OkProfile e  on a.assignedto_id= e.id_profile 
    WHERE b.id_project = ProjectId
    ;

END ;;
DELIMITER ;
/*!50003 SET sql_mode              = @saved_sql_mode */ ;
/*!50003 SET character_set_client  = @saved_cs_client */ ;
/*!50003 SET character_set_results = @saved_cs_results */ ;
/*!50003 SET collation_connection  = @saved_col_connection */ ;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2016-12-14 18:54:43
