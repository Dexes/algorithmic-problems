using System;

namespace Timus
{
    public static class Program
    {
        private const double DiagonalLength = 141.421356237;
        private const double StraightLength = 100;

        private static void Main()
        {
            var (n, m, data) = ReadData();
            var result = new double[n, m];

            for (var i = 1; i < n; i++)
                result[i, 0] = i * StraightLength;

            for (var i = 1; i < m; i++)
                result[0, i] = i * StraightLength;

            for (var i = 1; i < n; i++)
            for (var j = 1; j < m; j++)
            {
                result[i, j] = Math.Min(result[i - 1, j] + StraightLength, result[i, j - 1] + StraightLength);
                if (data[i - 1, j - 1])
                    result[i, j] = Math.Min(result[i, j], result[i - 1, j - 1] + DiagonalLength);
            }

            Console.WriteLine(Math.Round(result[n - 1, m - 1]));
        }

        private static (int, int, bool[,]) ReadData()
        {
            var line = Console.ReadLine().Split(' ');
            var (n, m) = (int.Parse(line[0]), int.Parse(line[1]));
            var count = int.Parse(Console.ReadLine());

            var data = new bool[n, m];
            for (var i = 0; i < count; i++)
            {
                line = Console.ReadLine().Split(' ');
                data[int.Parse(line[0]) - 1, int.Parse(line[1]) - 1] = true;
            }

            return (n + 1, m + 1, data);
        }
    }
}