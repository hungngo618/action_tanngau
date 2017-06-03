-- MySQL dump 10.13  Distrib 5.7.17, for osx10.12 (x86_64)
--
-- Host: localhost    Database: action_qna
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
  `intent` varchar(500) DEFAULT NULL,
  `trinh_answer` text,
  `bom_answer` text,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `intent_answer`
--

LOCK TABLES `intent_answer` WRITE;
/*!40000 ALTER TABLE `intent_answer` DISABLE KEYS */;
INSERT INTO `intent_answer` VALUES (1,'tuoi_trinh','Em là con gái, xin đừng hỏi tuổi em anh ơi','Trinh nó 28 tuổi anh ơi'),(2,'tuoi_bom','Anh đi hỏi nó ấy, em không nhớ','Em chưa 18 anh ạ hihi');
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
) ENGINE=InnoDB AUTO_INCREMENT=15 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sample_questions`
--

LOCK TABLES `sample_questions` WRITE;
/*!40000 ALTER TABLE `sample_questions` DISABLE KEYS */;
INSERT INTO `sample_questions` VALUES (2,'Trinh năm nay bao nhiêu tuổi',1),(3,'Em bao nhiêu tuổi rồi Trinh',1),(4,'Trinh sinh năm bao nhiêu',1),(5,'Trinh ơi em năm nay bao nhiêu tuổi thế',1),(6,'Trinh bao nhiêu tuổi Bờm ơi',1),(7,'Trinh năm nay bao nhiêu tuổi thế Bờm',1),(8,'Bờm ơi cái Trinh nó bao nhiêu tuổi vậy',1),(9,'Bờm bao nhiêu tuổi thế',2),(10,'Bờm ơi chú bao nhiêu tuổi',2),(11,'Năm nay bao nhiêu tuổi rồi Bờm',2),(12,'Trinh ơi Bờm nó bao nhiêu tuổi em',2),(13,'Trinh có biết Bờm nó bao nhiêu tuổi không',2),(14,'Trinh biết bờm nó bao nhiêu tuổi không',2);
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

-- Dump completed on 2017-06-03 21:42:04
