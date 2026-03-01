<?php

class UserService
{
    private $database;

    public function __construct(Database $database)
    {
        $this->database = $database;
    }

    public function createUser($name, $email)
    {
        $sql = "INSERT INTO users (name, email) VALUES (:name, :email)";
        $params = [':name' => $name, ':email' => $email];
        return $this->database->execute($sql, $params);
    }

    public function getUserById($id)
    {
        $sql = "SELECT * FROM users WHERE id = :id";
        $params = [':id' => $id];
        $result = $this->database->query($sql, $params);
        return $result ? $result[0] : null;
    }

    public function getUserByEmail($email)
    {
        $sql = "SELECT * FROM users WHERE email = :email";
        $params = [':email' => $email];
        $result = $this->database->query($sql, $params);
        return $result ? $result[0] : null;
    }

    public function getAllUsers()
    {
        $sql = "SELECT * FROM users";
        return $this->database->query($sql);
    }

    public function updateUser($id, $name, $email)
    {
        $sql = "UPDATE users SET name = :name, email = :email WHERE id = :id";
        $params = [':id' => $id, ':name' => $name, ':email' => $email];
        return $this->database->execute($sql, $params);
    }

    public function deleteUser($id)
    {
        $sql = "DELETE FROM users WHERE id = :id";
        $params = [':id' => $id];
        return $this->database->execute($sql, $params);
    }
}
