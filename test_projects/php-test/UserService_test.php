<?php

class UserServiceTest
{
    private $userService;

    public function setUp()
    {
        $database = new Database('localhost', 'testdb', 'root', 'password');
        $database->connect();
        $this->userService = new UserService($database);
    }

    public function testCreateUser()
    {
        $result = $this->userService->createUser('Test User', 'test@example.com');
        assert($result === true);
    }

    public function testGetUserById()
    {
        $user = $this->userService->getUserById(1);
        assert($user !== null);
    }

    public function testDeleteUser()
    {
        $result = $this->userService->deleteUser(1);
        assert($result === true);
    }
}
