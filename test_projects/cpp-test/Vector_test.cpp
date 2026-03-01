#include "Vector.h"
#include <cassert>

void testVectorPushBack() {
    Vector<int> vec;
    vec.push_back(1);
    vec.push_back(2);
    vec.push_back(3);
    assert(vec.size() == 3);
}

void testVectorPopBack() {
    Vector<int> vec;
    vec.push_back(1);
    vec.push_back(2);
    vec.pop_back();
    assert(vec.size() == 1);
}

void testVectorEmpty() {
    Vector<int> vec;
    assert(vec.empty() == true);
    vec.push_back(1);
    assert(vec.empty() == false);
}

int main() {
    testVectorPushBack();
    testVectorPopBack();
    testVectorEmpty();
    return 0;
}
