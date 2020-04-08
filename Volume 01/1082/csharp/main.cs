using System;

namespace Timus
{
    public static class Program
    {
        private static void Main()
        {
            var n = int.Parse(Console.ReadLine());

            for (var i = 1; i < n; i++)
                Console.Write(i + " ");

            Console.WriteLine(n);
        }
    }
}