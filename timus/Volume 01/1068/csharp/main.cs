using System;

namespace Timus
{
    public static class Program
    {
        private static void Main()
        {
            var n = int.Parse(Console.ReadLine());
            Console.WriteLine(n > 0 ? PositiveSum(n) : NegativeSum(n));
        }

        private static int PositiveSum(int n)
        {
            var sum = 0;
            for (var i = 1; i <= n; i++)
                sum += i;

            return sum;
        }

        private static int NegativeSum(int n)
        {
            var sum = 0;
            for (var i = 1; i >= n; i--)
                sum += i;

            return sum;
        }
    }
}