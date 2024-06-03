<?php
include $_SERVER['DOCUMENT_ROOT'].'/api/scripts/session.php';
?>

<!DOCTYPE html>
<html lang="ru">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Capcha</title>
</head>
<body>
    <img src='scripts/createImg.php' alt="capcha. Включите отображение картинок.">
    <form action="resultPage.php" method="get">
    <input type="text" name="answer" placeholder="Введите капчу">
    <input type="submit">
    </form>
    <?php 
        echo "Подсказка: " . $_SESSION['sessionCapcha'];
    ?>
</body>
</html>