<?php

require_once 'User.php';
require_once 'Database.php';
require_once 'UserService.php';

$database = new Database('localhost', 'testdb', 'root', 'password');
$database->connect();

$userService = new UserService($database);

$userService->createUser('John Doe', 'john@example.com');
$userService->createUser('Jane Smith', 'jane@example.com');

$users = $userService->getAllUsers();
foreach ($users as $user) {
    echo "ID: {$user['id']}, Name: {$user['name']}, Email: {$user['email']}\n";
}
