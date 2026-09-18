#include <algorithm>

int get_sum(int a, int b)
{
    //Good Luck!
    int start = std::min(a, b);
    int end = std::max(a, b);

    int sum = 0;

    for (int i = start; i <= end; i++) {
        sum += i;
    }

    return sum;
}