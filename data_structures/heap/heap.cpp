#include <iostream>

template<typename T>
struct Interface {
    virtual ~Interface() = default;
    virtual bool less(size_t i, size_t j) const = 0;
    virtual void swap(size_t i, size_t j) = 0;
    virtual int len() const = 0;
    virtual void push(const T& x) = 0;
    virtual T pop() = 0;
};

template<typename T>
void push(Interface<T>& h, T x) {
    h.push(x);
    up(h, h.len() - 1);
}

template<typename T>
T pop(Interface<T>& h) {
    h.swap(0, h.len() - 1);
    T val = h.pop();
    heapify(h, 0);
    return val;
}

template<typename T>
void heapify(Interface<T>& h, size_t i) {
    size_t n = h.len();
    while (true) {
        size_t smallest = i;
        
        size_t lc = 2 * i + 1;
        size_t rc = 2 * i + 2;

        if (lc < n && h.less(lc, smallest)) {
            smallest = lc;
        }
        if (rc < n && h.less(rc, smallest)) {
            smallest = rc;
        }
        if (smallest == i) {
            return;
        }
        h.swap(smallest, i);
        i = smallest;
    }
}

template<typename T>
void up(Interface<T>& h, size_t i) {
    while (i > 0 && h.less(i, (i - 1) / 2)) {
        h.swap(i, (i - 1) / 2);
        i = (i - 1) / 2;
    }
}