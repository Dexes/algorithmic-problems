using System;
using System.Globalization;

namespace Timus
{
    public static class Program
    {
        private static void Main()
        {
            var (size, radius, points) = ReadData();
            var result = CalculateResult(size, radius, points);
            Console.WriteLine(Math.Round(result, 2));
        }

        private static double CalculateResult(int size, double radius, Point[] points)
        {
            var last = size - 1;
            var result = points[0].DistanceTo(points[last]);

            for (var i = 0; i < last; i++)
                result += points[i].DistanceTo(points[i + 1]);

            return result + 2.0 * Math.PI * radius;
        }

        private static (int, double, Point[]) ReadData()
        {
            var data = Console.ReadLine().Split(' ');
            var size = int.Parse(data[0]);
            var radius = double.Parse(data[1], NumberFormatInfo.InvariantInfo);
            var points = new Point[size];

            for (var i = 0; i < size; i++)
            {
                data = Console.ReadLine().Split(' ');
                points[i] = new Point
                {
                    X = double.Parse(data[0], NumberFormatInfo.InvariantInfo),
                    Y = double.Parse(data[1], NumberFormatInfo.InvariantInfo),
                };
            }

            return (size, radius, points);
        }
    }

    internal class Point
    {
        public double X;
        public double Y;

        public double DistanceTo(Point point)
        {
            var dx = X - point.X;
            var dy = Y - point.Y;

            return Math.Sqrt(dx * dx + dy * dy);
        }
    }
}