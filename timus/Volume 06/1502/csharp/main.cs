using System;

namespace Timus
{
    public static class Program
    {
        private static void Main()
        {
            long n = int.Parse(Console.ReadLine());
            Console.WriteLine((n + 2) * (n + 1) * n / 2);
        }
    }
}