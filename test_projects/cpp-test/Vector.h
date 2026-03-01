#include <vector>
#include <algorithm>

template <typename T>
class Vector {
private:
    std::vector<T> data;

public:
    void push_back(const T& value) {
        data.push_back(value);
    }

    void pop_back() {
        if (!data.empty()) {
            data.pop_back();
        }
    }

    T& operator[](int index) {
        return data[index];
    }

    const T& operator[](int index) const {
        return data[index];
    }

    int size() const {
        return data.size();
    }

    bool empty() const {
        return data.empty();
    }

    void clear() {
        data.clear();
    }

    void sort() {
        std::sort(data.begin(), data.end());
    }

    typename std::vector<T>::iterator begin() {
        return data.begin();
    }

    typename std::vector<T>::iterator end() {
        return data.end();
    }

    typename std::vector<T>::const_iterator begin() const {
        return data.begin();
    }

    typename std::vector<T>::const_iterator end() const {
        return data.end();
    }
};
