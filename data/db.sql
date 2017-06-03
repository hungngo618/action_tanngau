-- MySQL dump 10.13  Distrib 5.7.17, for osx10.12 (x86_64)
--
-- Host: localhost    Database: action_tanngau
-- ------------------------------------------------------
-- Server version	5.7.17

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
-- Table structure for table `intent_answer`
--

DROP TABLE IF EXISTS `intent_answer`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `intent_answer` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `instance` tinyint(1) DEFAULT '0',
  `intent` varchar(500) DEFAULT NULL,
  `answer` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `intent_answer`
--

LOCK TABLES `intent_answer` WRITE;
/*!40000 ALTER TABLE `intent_answer` DISABLE KEYS */;
INSERT INTO `intent_answer` VALUES (1,0,'trinh_quequan','Em quê Bến Tre'),(2,1,'bom_quequan','Quê tớ ở Hà Nội'),(3,0,'trinh_tuoi','Sao con trai ai cũng thích hỏi tuổi em hoài à?'),(4,1,'bom_tuoi','Em chưa 18 anh ơi'),(5,2,'trinh_bom_thichnhau','Không tiền chó nó yêu');
/*!40000 ALTER TABLE `intent_answer` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sample_questions`
--

DROP TABLE IF EXISTS `sample_questions`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `sample_questions` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `question` varchar(500) NOT NULL,
  `intent_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=51 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sample_questions`
--

LOCK TABLES `sample_questions` WRITE;
/*!40000 ALTER TABLE `sample_questions` DISABLE KEYS */;
INSERT INTO `sample_questions` VALUES (1,'Trinh ơi em quê ở đâu',1),(2,'Trinh ơi nhà em ở đâu',1),(3,'Quê em ở đâu Trinh',1),(4,'Trinh ơi nhà em ở đâu vậy',1),(9,'Trinh ơi em bao nhiêu tuổi',3),(10,'Trinh ơi em bao nhiêu tuổi rồi',3),(11,'Trinh ơi tuổi em bao nhiêu',3),(12,'Bờm ơi em bao nhiêu tuổi',4),(13,'Bờm ơi chú bao nhiêu tuổi',4),(14,'Bờm bao nhiêu tuổi',4),(15,'Bờm ơi em bao nhiêu tuổi rồi',4),(39,'Bờm ơi, em quê ở đâu',0),(40,'Bờm ơi, chú nhà ở đâu',0),(41,'Bờm ơi quên em ở đâu thế',0),(42,'Quê ở đâu vậy Bờm',0),(43,'Bờm ơi, em quê ở đâu',2),(44,'Bờm ơi, chú nhà ở đâu',2),(45,'Bờm ơi quên em ở đâu thế',2),(46,'Quê ở đâu vậy Bờm',2),(47,'Trinh Bờm yêu nhau không',5),(48,'Trinh ơi yêu Bờm không',5),(49,'Bờm ơi yêu Trinh không',5),(50,'Trinh ơi 2 đứa yêu nhau không',5);
/*!40000 ALTER TABLE `sample_questions` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-06-03 13:48:32
