using System;
using System.Globalization;

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

            Array.Sort(data);

            double result = data[size - 1];
            for (var i = size - 2; i >= 0; i--)
                result = 2 * Math.Sqrt(result * data[i]);

            Console.WriteLine(string.Format(NumberFormatInfo.InvariantInfo, "{0:F2}", result));
        }
    }
}