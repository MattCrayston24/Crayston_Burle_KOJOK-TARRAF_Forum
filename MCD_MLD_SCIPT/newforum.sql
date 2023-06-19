-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Hôte : 127.0.0.1:3306
-- Généré le : lun. 19 juin 2023 à 23:07
-- Version du serveur : 8.0.31
-- Version de PHP : 8.0.26

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de données : `newforum`
--

-- --------------------------------------------------------

--
-- Structure de la table `categorie`
--

DROP TABLE IF EXISTS `categorie`;
CREATE TABLE IF NOT EXISTS `categorie` (
  `ID_CATEGORIE` int NOT NULL AUTO_INCREMENT,
  `TITRE` varchar(255) COLLATE latin1_bin NOT NULL,
  PRIMARY KEY (`ID_CATEGORIE`),
  UNIQUE KEY `ID_CATEGORIE` (`ID_CATEGORIE`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

--
-- Déchargement des données de la table `categorie`
--

INSERT INTO `categorie` (`ID_CATEGORIE`, `TITRE`) VALUES
(1, 'Programmes'),
(2, 'Alimentation'),
(3, 'Produits');

-- --------------------------------------------------------

--
-- Structure de la table `definir`
--

DROP TABLE IF EXISTS `definir`;
CREATE TABLE IF NOT EXISTS `definir` (
  `ID_TOPIC` int NOT NULL,
  `ID_CATEGORIE` int NOT NULL,
  PRIMARY KEY (`ID_TOPIC`,`ID_CATEGORIE`),
  UNIQUE KEY `ID_TOPIC` (`ID_TOPIC`),
  KEY `ID_CATEGORIE` (`ID_CATEGORIE`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

--
-- Déchargement des données de la table `definir`
--

INSERT INTO `definir` (`ID_TOPIC`, `ID_CATEGORIE`) VALUES
(1, 1),
(2, 1),
(3, 2),
(4, 2),
(5, 3),
(6, 3),
(7, 3),
(8, 1),
(9, 2);

-- --------------------------------------------------------

--
-- Structure de la table `message`
--

DROP TABLE IF EXISTS `message`;
CREATE TABLE IF NOT EXISTS `message` (
  `ID_MESSAGE` int NOT NULL AUTO_INCREMENT,
  `CONTENU` varchar(255) COLLATE latin1_bin NOT NULL,
  `ID_MESSAGE_1` int DEFAULT NULL,
  `ID_TOPIC` int NOT NULL,
  `ID_UTILISATEUR` int NOT NULL,
  PRIMARY KEY (`ID_MESSAGE`),
  UNIQUE KEY `ID_MESSAGE` (`ID_MESSAGE`),
  KEY `ID_MESSAGE_1` (`ID_MESSAGE_1`),
  KEY `ID_TOPIC` (`ID_TOPIC`),
  KEY `ID_UTILISATEUR` (`ID_UTILISATEUR`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

--
-- Déchargement des données de la table `message`
--

INSERT INTO `message` (`ID_MESSAGE`, `CONTENU`, `ID_MESSAGE_1`, `ID_TOPIC`, `ID_UTILISATEUR`) VALUES
(1, 'Message about Programmes 1', NULL, 1, 1),
(2, 'Message about Programmes 2', 1, 1, 2),
(3, 'Message about Alimentation 1', NULL, 3, 3),
(4, 'Message about Alimentation 2', 3, 3, 4),
(5, 'Message about Produits 1', NULL, 5, 5),
(6, 'Message about Produits 2', 5, 5, 1);

-- --------------------------------------------------------

--
-- Structure de la table `role`
--

DROP TABLE IF EXISTS `role`;
CREATE TABLE IF NOT EXISTS `role` (
  `ID_ROLE` int NOT NULL AUTO_INCREMENT,
  `GRADE` varchar(255) COLLATE latin1_bin NOT NULL,
  PRIMARY KEY (`ID_ROLE`),
  UNIQUE KEY `ID_ROLE` (`ID_ROLE`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

--
-- Déchargement des données de la table `role`
--

INSERT INTO `role` (`ID_ROLE`, `GRADE`) VALUES
(1, 'Administrateur'),
(2, 'Moderateur'),
(3, 'Utilisateur');

-- --------------------------------------------------------

--
-- Structure de la table `topic`
--

DROP TABLE IF EXISTS `topic`;
CREATE TABLE IF NOT EXISTS `topic` (
  `ID_TOPIC` int NOT NULL AUTO_INCREMENT,
  `TITRE` varchar(255) COLLATE latin1_bin NOT NULL,
  `ID_UTILISATEUR` int NOT NULL,
  PRIMARY KEY (`ID_TOPIC`),
  UNIQUE KEY `ID_TOPIC` (`ID_TOPIC`),
  UNIQUE KEY `TITRE` (`TITRE`),
  KEY `ID_UTILISATEUR` (`ID_UTILISATEUR`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

--
-- Déchargement des données de la table `topic`
--

INSERT INTO `topic` (`ID_TOPIC`, `TITRE`, `ID_UTILISATEUR`) VALUES
(1, 'Question about Programmes 1', 1),
(2, 'Question about Programmes 2', 2),
(3, 'Question about Alimentation 1', 3),
(4, 'Question about Alimentation 2', 4),
(5, 'Question about Produits 1', 5),
(6, 'Question about Produits 2', 1),
(7, 'Quel est le meilleur produit pour la prise de masse?', 1),
(8, 'Quels exercices pour les pecs?', 1),
(9, 'Quel regime pour la perte de gras?', 1);

-- --------------------------------------------------------

--
-- Structure de la table `utilisateur`
--

DROP TABLE IF EXISTS `utilisateur`;
CREATE TABLE IF NOT EXISTS `utilisateur` (
  `ID_UTILISATEUR` int NOT NULL AUTO_INCREMENT,
  `NOM_UTILISATEUR` varchar(255) COLLATE latin1_bin NOT NULL,
  `ADRESSE_MAIL` varchar(255) COLLATE latin1_bin NOT NULL,
  `MOT_DE_PASSE` varchar(255) COLLATE latin1_bin NOT NULL,
  `ID_ROLE` int NOT NULL,
  `SESSION_TOKEN` varchar(255) COLLATE latin1_bin DEFAULT NULL,
  PRIMARY KEY (`ID_UTILISATEUR`),
  UNIQUE KEY `ID_UTILISATEUR` (`ID_UTILISATEUR`),
  UNIQUE KEY `NOM_UTILISATEUR` (`NOM_UTILISATEUR`),
  UNIQUE KEY `ADRESSE_MAIL` (`ADRESSE_MAIL`),
  UNIQUE KEY `NOM_UTILISATEUR_2` (`NOM_UTILISATEUR`),
  UNIQUE KEY `ADRESSE_MAIL_2` (`ADRESSE_MAIL`),
  UNIQUE KEY `SESSION_TOKEN` (`SESSION_TOKEN`),
  KEY `ID_ROLE` (`ID_ROLE`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=latin1 COLLATE=latin1_bin;

--
-- Déchargement des données de la table `utilisateur`
--

INSERT INTO `utilisateur` (`ID_UTILISATEUR`, `NOM_UTILISATEUR`, `ADRESSE_MAIL`, `MOT_DE_PASSE`, `ID_ROLE`, `SESSION_TOKEN`) VALUES
(1, 'Admin', 'admin@mail.com', 'password1', 1, NULL),
(2, 'Moderator', 'moderator@mail.com', 'password2', 2, NULL),
(3, 'User1', 'user1@mail.com', 'password3', 3, '9ZAyh0BjIT92OVZBMMuZp2GmhpH6I0kZOLh0CZ0i2rA='),
(4, 'User2', 'user2@mail.com', 'password4', 3, NULL),
(5, 'User3', 'user3@mail.com', 'password5', 3, NULL),
(7, 'KhaledT7', 'khaled23mdjkmk@gmail.com', '$2a$10$XcEnmcYg2ZVTf3gTU3OTkuFaeG7HM3Ks706ureHCO2FZ0udTl6MeS', 3, 'pXlTVqJ1HLYErIcUZEXQmC2NRXz1PwcaKD6JVCECwIk='),
(8, 'jesuis', 'jesuis@gmail.com', '$2a$10$vLOSIGbXSJQ9pzGbz1xSlORm4zfncoHRr5tlNZrZcX61Uz1JgAkwy', 3, 'UUfT/bg9j4pBMCraAT10tbDFuLjPHVnujR65y81wJSs=');

--
-- Contraintes pour les tables déchargées
--

--
-- Contraintes pour la table `definir`
--
ALTER TABLE `definir`
  ADD CONSTRAINT `definir_ibfk_1` FOREIGN KEY (`ID_TOPIC`) REFERENCES `topic` (`ID_TOPIC`),
  ADD CONSTRAINT `definir_ibfk_2` FOREIGN KEY (`ID_CATEGORIE`) REFERENCES `categorie` (`ID_CATEGORIE`);

--
-- Contraintes pour la table `message`
--
ALTER TABLE `message`
  ADD CONSTRAINT `message_ibfk_1` FOREIGN KEY (`ID_MESSAGE_1`) REFERENCES `message` (`ID_MESSAGE`),
  ADD CONSTRAINT `message_ibfk_2` FOREIGN KEY (`ID_TOPIC`) REFERENCES `topic` (`ID_TOPIC`),
  ADD CONSTRAINT `message_ibfk_3` FOREIGN KEY (`ID_UTILISATEUR`) REFERENCES `utilisateur` (`ID_UTILISATEUR`);

--
-- Contraintes pour la table `topic`
--
ALTER TABLE `topic`
  ADD CONSTRAINT `topic_ibfk_1` FOREIGN KEY (`ID_UTILISATEUR`) REFERENCES `utilisateur` (`ID_UTILISATEUR`);

--
-- Contraintes pour la table `utilisateur`
--
ALTER TABLE `utilisateur`
  ADD CONSTRAINT `utilisateur_ibfk_1` FOREIGN KEY (`ID_ROLE`) REFERENCES `role` (`ID_ROLE`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
