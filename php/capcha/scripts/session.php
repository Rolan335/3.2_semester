<?php
include $_SERVER['DOCUMENT_ROOT'].'/util/generateString.php';
session_start();
$_SESSION['sessionCapcha'] = $capcha;