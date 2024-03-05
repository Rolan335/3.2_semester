<?php
session_start();
if($_SERVER['REQUEST_METHOD'] == 'POST'){
        $capcha = $_SESSION['sessionCapcha'];
        $answer = $_POST['answer'];
    if($capcha == $answer){
        echo <<<HTML
        <h1>Вы угадали капчу, проходите.</h1>
        HTML;
    } else{
        echo <<<HTML
        <h1>Вы не угадали капчу</h1>
        <a href="/">Назад.</a>
        HTML;
    }
    }
