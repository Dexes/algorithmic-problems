using System;
using System.Globalization;

namespace Timus
{
    public static class Program
    {
        private static void Main()
        {
            var splitChars = new[] {' ', '\t', '\n', '\r'};
            var input = Console.In.ReadToEnd().Split(splitChars, StringSplitOptions.RemoveEmptyEntries);
            for (var i = input.Length - 1; i >= 0; i--)
            {
                var root = Math.Sqrt(double.Parse(input[i], NumberFormatInfo.InvariantInfo));
                Console.WriteLine(string.Format(NumberFormatInfo.InvariantInfo, "{0:F4}", root));
            }
        }
    }
}