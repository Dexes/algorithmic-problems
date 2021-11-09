using System;
using System.Collections.Generic;

namespace Timus
{
    public static class Program
    {
        private static void Main()
        {
            var data = Eratosfen(165000);
            var size = int.Parse(Console.ReadLine());

            for (var i = 0; i < size; i++)
                Console.WriteLine(data[int.Parse(Console.ReadLine()) - 1]);
        }

        private static int[] Eratosfen(int size)
        {
            var notPrimes = new bool[size + 1];
            var result = new List<int>();
            var sqrtSize = Math.Sqrt(size);

            for (var i = 2; i < size; i++)
            {
                if (notPrimes[i]) continue;

                result.Add(i);

                if (i > sqrtSize) continue;
                for (long j = i * i; j < size; j += i)
                    notPrimes[j] = true;
            }

            return result.ToArray();
        }
    }
}