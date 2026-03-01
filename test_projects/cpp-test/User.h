#include <iostream>
#include <string>

class User {
private:
    int id;
    std::string name;
    std::string email;

public:
    User(int id, const std::string& name, const std::string& email)
        : id(id), name(name), email(email) {}

    int getId() const {
        return id;
    }

    std::string getName() const {
        return name;
    }

    std::string getEmail() const {
        return email;
    }

    void setName(const std::string& name) {
        this->name = name;
    }

    void setEmail(const std::string& email) {
        this->email = email;
    }

    void display() const {
        std::cout << "ID: " << id << ", Name: " << name << ", Email: " << email << std::endl;
    }
};
