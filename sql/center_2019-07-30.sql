# ************************************************************
# Sequel Pro SQL dump
# Version 4541
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.7.21-log)
# Database: center
# Generation Time: 2019-07-30 10:11:55 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table imagelist
# ------------------------------------------------------------

DROP TABLE IF EXISTS `imagelist`;

CREATE TABLE `imagelist` (
  `ID` bigint(20) NOT NULL AUTO_INCREMENT,
  `URL` varchar(1024) DEFAULT NULL,
  `Title` varchar(20) DEFAULT NULL,
  `Sort` bigint(100) DEFAULT NULL,
  `Module` tinyint(2) DEFAULT NULL,
  `IsShow` tinyint(1) DEFAULT NULL,
  `Content` text,
  `CreateTime` datetime DEFAULT CURRENT_TIMESTAMP,
  `UpdateTime` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `imagelist` WRITE;
/*!40000 ALTER TABLE `imagelist` DISABLE KEYS */;

INSERT INTO `imagelist` (`ID`, `URL`, `Title`, `Sort`, `Module`, `IsShow`, `Content`, `CreateTime`, `UpdateTime`)
VALUES
	(9,'http://192.168.2.62:8881/static/upload/1564468182.jpg','test',0,1,0,'<p>xixixi</p>','2019-07-30 14:29:45','2019-07-30 17:05:26'),
	(10,'http://192.168.2.62:8881/static/upload/1564468583.jpg','testtt',4,1,1,'<p>ssss</p>','2019-07-30 14:36:25','2019-07-30 18:10:54'),
	(11,'http://192.168.2.62:8881/static/upload/1564469832.png','ere',1,1,1,'<p>eeeeee</p>','2019-07-30 14:57:15','2019-07-30 17:51:34'),
	(12,'http://192.168.2.62:8881/static/upload/1564476621.jpg','23',3,1,1,'<p>4343</p>','2019-07-30 15:35:36','2019-07-30 18:10:52');

/*!40000 ALTER TABLE `imagelist` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
