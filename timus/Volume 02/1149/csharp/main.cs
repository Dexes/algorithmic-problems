using System;

namespace Timus
{
    public static class Program
    {
        private static void Main()
        {
            var n = int.Parse(Console.ReadLine());
            var result = new string('(', n - 1);

            for (var i = 0; i < n; i++)
                result += A(i + 1) + "+" + (n - i) + ")";

            Console.WriteLine(result.TrimEnd(')'));
        }

        private static string A(int n)
        {
            var result = "sin(1";

            for (var i = 2; i <= n; i++)
                result += IsOdd(i) ? "+sin(" + i : "-sin(" + i;

            return result + new string(')', n);
        }

        private static bool IsOdd(int number)
        {
            return (number | 1) == number;
        }
    }
}