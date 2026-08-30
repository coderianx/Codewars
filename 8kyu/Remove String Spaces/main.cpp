#include <algorithm>
#include <string>

std::string no_space(const std::string& x)
{
    std::string result = x;
    result.erase(std::remove(result.begin(), result.end(), ' '), result.end());
    return result;
}