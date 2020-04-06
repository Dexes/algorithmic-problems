using System;
using System.Collections.Generic;
using System.Linq;

namespace Timus
{
    public static class Program
    {
        private static void Main()
        {
            var size = int.Parse(Console.ReadLine());
            var statistic = new Dictionary<string, int>();
            for (var i = 0; i < size; i++)
            {
                var key = Console.ReadLine();
                if (statistic.ContainsKey(key))
                    statistic[key]++;
                else
                    statistic[key] = 1;
            }

            Console.WriteLine(string.Join("\n", statistic.Where(pair => pair.Value > 1).Select(pair => pair.Key)));
        }
    }
}