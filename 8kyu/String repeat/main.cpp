#include <cstddef>
#include <string>

std::string repeat_str(size_t repeat, const std::string& str)
{
    std::string result;

    for (size_t i = 0; i < repeat; i++) {
        result += str;
    }

    return result;
}