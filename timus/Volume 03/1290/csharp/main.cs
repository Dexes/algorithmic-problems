using System;

namespace Timus
{
    public static class Program
    {
        private static void Main()
        {
            var size = int.Parse(Console.ReadLine());
            var data = new int[size];

            for (var i = 0; i < size; i++)
                data[i] = int.Parse(Console.ReadLine());

            Array.Sort(data, (first, second) => second.CompareTo(first));

            for (var i = 0; i < size; i++)
                Console.WriteLine(data[i]);
        }
    }
}