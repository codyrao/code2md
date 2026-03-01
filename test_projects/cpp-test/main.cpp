#include "User.h"
#include "Vector.h"
#include <iostream>

int main() {
    Vector<User> users;

    User user1(1, "John Doe", "john@example.com");
    User user2(2, "Jane Smith", "jane@example.com");
    User user3(3, "Bob Johnson", "bob@example.com");

    users.push_back(user1);
    users.push_back(user2);
    users.push_back(user3);

    std::cout << "All users:" << std::endl;
    for (const auto& user : users) {
        user.display();
    }

    std::cout << "\nTotal users: " << users.size() << std::endl;

    users.pop_back();
    std::cout << "\nAfter removing last user:" << std::endl;
    for (const auto& user : users) {
        user.display();
    }

    return 0;
}
