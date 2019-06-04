#include <iostream>

int main()
{
    int sum = 0;
    int n, vi;

    std::ios_base::sync_with_stdio(false);

    std::scanf("%d", &n);

    for (int i = 0; i < n; i++)
    {
        std::scanf("%d", &vi);
        sum += vi;
    }
    std::printf("%d\n", sum);
    return 0;
}