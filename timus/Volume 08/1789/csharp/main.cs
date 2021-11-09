using System;

namespace Timus
{
    public static class Program
    {
        private static void Main()
        {
            var n = int.Parse(Console.ReadLine());
            var m = 2 * n - 1;
            var steps = new int[m];
            var counter = 0;

            for (var i = 1; i <= n; i++)
                steps[counter++] = i;

            for (var i = n; i >= 2; i--)
                steps[counter++] = i;

            Console.WriteLine(m);
            Console.WriteLine(string.Join(" ", steps));
        }
    }
}