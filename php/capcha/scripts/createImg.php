<?php
include $_SERVER['DOCUMENT_ROOT'].'/util/generateString.php';
session_start();

$path = "../img/noise.jpg";
$img = imagecreatefromjpeg($path);

$x = 20;
$y = 35;
foreach(str_split($_SESSION['sessionCapcha']) as $char){
    imagettftext($img,rand(18,30),rand(-30,30),$x,$y,imagecolorallocate($img, rand(0,255), rand(0,255), rand(0,255)),"../font/NanumPenScript-Regular.ttf",$char);
    $x+=rand(20,40);
}

imagepng($img);