<?php
include $_SERVER['DOCUMENT_ROOT'].'/api/util/generateString.php';
session_start();
$_SESSION['sessionCapcha'] = $capcha;