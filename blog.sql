-- phpMyAdmin SQL Dump
-- version 4.4.12
-- http://www.phpmyadmin.net
--
-- Anamakine: 127.0.0.1
-- Üretim Zamanı: 28 Haz 2017, 15:59:37
-- Sunucu sürümü: 5.6.25
-- PHP Sürümü: 5.6.11

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Veritabanı: `blog`
--

-- --------------------------------------------------------

--
-- Tablo için tablo yapısı `kategori`
--

CREATE TABLE IF NOT EXISTS `kategori` (
  `Id` int(11) NOT NULL,
  `Kategori` varchar(150) CHARACTER SET utf8 COLLATE utf8_turkish_ci NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

--
-- Tablo döküm verisi `kategori`
--

INSERT INTO `kategori` (`Id`, `Kategori`) VALUES
(1, 'Golang'),
(2, 'Java'),
(3, 'Deneme');

-- --------------------------------------------------------

--
-- Tablo için tablo yapısı `post`
--

CREATE TABLE IF NOT EXISTS `post` (
  `Id` int(11) NOT NULL,
  `Baslik` varchar(250) CHARACTER SET utf8 COLLATE utf8_turkish_ci NOT NULL,
  `Ozet` text CHARACTER SET utf8 COLLATE utf8_turkish_ci NOT NULL,
  `Detay` text CHARACTER SET utf8 COLLATE utf8_turkish_ci NOT NULL,
  `Kategori` int(11) NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=latin1;

--
-- Tablo döküm verisi `post`
--

INSERT INTO `post` (`Id`, `Baslik`, `Ozet`, `Detay`, `Kategori`) VALUES
(1, 'Deneme', 'Test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test', 'Test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test test', 1),
(2, 'Deneme2', 'test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 ', 'test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 ', 2),
(3, 'Deneme2', 'test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 ', 'test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 test2 ', 2),
(6, 'Deneme Güncel', 'dj ldjlaksjdlasjd asdjl', 's dhdj d sdkjs\n<img class="circle" src="/static/img/gopher.png">\n gggg', 1);

-- --------------------------------------------------------

--
-- Tablo için tablo yapısı `user`
--

CREATE TABLE IF NOT EXISTS `user` (
  `Id` int(11) NOT NULL,
  `Name` varchar(50) CHARACTER SET utf8 COLLATE utf8_turkish_ci NOT NULL,
  `Mail` varchar(50) CHARACTER SET utf8 COLLATE utf8_turkish_ci NOT NULL,
  `Pass` varchar(50) CHARACTER SET utf8 COLLATE utf8_turkish_ci NOT NULL
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=latin1;

--
-- Tablo döküm verisi `user`
--

INSERT INTO `user` (`Id`, `Name`, `Mail`, `Pass`) VALUES
(2, 'Zeyit BAŞAR', 'zeyit17@gmail.com', '202cb962ac59075b964b07152d234b70');

--
-- Dökümü yapılmış tablolar için indeksler
--

--
-- Tablo için indeksler `kategori`
--
ALTER TABLE `kategori`
  ADD PRIMARY KEY (`Id`);

--
-- Tablo için indeksler `post`
--
ALTER TABLE `post`
  ADD PRIMARY KEY (`Id`);

--
-- Tablo için indeksler `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`Id`);

--
-- Dökümü yapılmış tablolar için AUTO_INCREMENT değeri
--

--
-- Tablo için AUTO_INCREMENT değeri `kategori`
--
ALTER TABLE `kategori`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=4;
--
-- Tablo için AUTO_INCREMENT değeri `post`
--
ALTER TABLE `post`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=8;
--
-- Tablo için AUTO_INCREMENT değeri `user`
--
ALTER TABLE `user`
  MODIFY `Id` int(11) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=4;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
