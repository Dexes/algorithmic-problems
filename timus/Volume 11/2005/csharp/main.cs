using System;
using System.Linq;

namespace Timus
{
    public static class Program
    {
        private const int Size = 5;

        private static void Main()
        {
            var distances = ReadData();
            var routes = new[]
            {
                new Route
                {
                    Length = distances[0, 1] + distances[1, 2] + distances[2, 3] + distances[3, 4],
                    Points = new[] {1, 2, 3, 4, 5},
                },
                new Route
                {
                    Length = distances[0, 2] + distances[2, 1] + distances[1, 3] + distances[3, 4],
                    Points = new[] {1, 3, 2, 4, 5},
                },
                new Route
                {
                    Length = distances[0, 2] + distances[2, 3] + distances[3, 1] + distances[1, 4],
                    Points = new[] {1, 3, 4, 2, 5},
                },
                new Route
                {
                    Length = distances[0, 3] + distances[3, 2] + distances[2, 1] + distances[1, 4],
                    Points = new[] {1, 4, 3, 2, 5},
                },
            };

            Array.Sort(routes, (x, y) => x.Length.CompareTo(y.Length));
            var bestRoute = routes.First();

            Console.WriteLine(bestRoute.Length);
            Console.WriteLine(string.Join(" ", bestRoute.Points));
        }

        // If you don't like easy way
        // private static void Main()
        // {
        //     var distances = ReadData();
        //     var length = int.MaxValue;
        //     var route = new int[Size];
        //     var currentRoute = new int[Size];
        //     currentRoute[Size - 1] = Size - 1;
        //     BuildRoute(ref length, route, 0, currentRoute, 1, distances);
        //
        //     Console.WriteLine(length);
        //     Console.WriteLine(string.Join(" ", route.Select(x => x + 1)));
        // }
        //
        // private static void BuildRoute(
        //     ref int bestLength,
        //     int[] bestRoute,
        //     int length,
        //     int[] route,
        //     int position,
        //     int[,] distances)
        // {
        //     if (position == Size - 1)
        //     {
        //         length += distances[route[position - 1], Size - 1];
        //         if (length < bestLength)
        //         {
        //             Array.Copy(route, bestRoute, Size);
        //             bestLength = length;
        //             return;
        //         }
        //     }
        //
        //     for (var i = 1; i < Size - 1; i++)
        //     {
        //         if (position == Size - 2 && i == 2) continue;
        //         if (Array.IndexOf(route, i, 0, position) >= 0) continue;
        //         route[position] = i;
        //
        //         var distanceToPoint = distances[route[position - 1], i];
        //         BuildRoute(ref bestLength, bestRoute, length + distanceToPoint, route, position + 1, distances);
        //     }
        // }

        private static int[,] ReadData()
        {
            var result = new int[Size, Size];
            for (var i = 0; i < Size; i++)
            {
                var row = Console.ReadLine().Split(' ');
                for (var j = 0; j < Size; j++)
                    result[i, j] = int.Parse(row[j]);
            }

            return result;
        }
    }

    internal struct Route
    {
        public int Length;
        public int[] Points;
    }
}