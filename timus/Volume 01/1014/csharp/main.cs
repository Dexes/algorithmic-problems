using System;
using System.Collections.Generic;

namespace Timus
{
    public static class Program
    {
        private static void Main()
        {
            var n = long.Parse(Console.ReadLine());

            if (n == 0)
            {
                Console.WriteLine(10);
                return;
            }

            if (n < 10)
            {
                Console.WriteLine(n);
                return;
            }

            var stack = new Stack<long>();
            for (var i = 9; i >= 2 && n > 1; i--)
            {
                while (n % i == 0)
                {
                    stack.Push(i);
                    n /= i;
                }
            }

            if (n != 1)
            {
                Console.WriteLine(-1);
                return;
            }

            while (stack.Count > 0)
                Console.Write(stack.Pop());

            Console.WriteLine();
        }
    }
}