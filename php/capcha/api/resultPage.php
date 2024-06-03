<?php
session_start();
if($_SERVER['REQUEST_METHOD'] == 'GET'){
        $capcha = $_SESSION['sessionCapcha'];
        $answer = $_GET['answer'];
    if($capcha == $answer){
        echo <<<HTML
        <h1>Вы угадали капчу, проходите.</h1>
        HTML;
    } else{
        echo <<<HTML
        <h1>Вы не угадали капчу</h1>
        <a href="index.php">Назад.</a>
        HTML;
    }
    }
