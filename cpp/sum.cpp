#include <iostream>

int main()
{
    int sum = 0;
    int n, vi;

    std::cin >> n;

    for (int i = 0; i < n; i++)
    {
        std::cin >> vi;
        sum += vi;
    }
    std::cout << sum << std::endl;
    return 0;
}