using System;

namespace Timus
{
    public static class Program
    {
        private static void Main()
        {
            var (size, data, sum) = ReadData();
            long result = 0;
            Array.Sort(data);

            for (var i = 0; i < size; i++)
            {
                result += sum * data[i];
                sum -= data[i];
                result += sum * data[i];
            }

            Console.WriteLine(result);
        }

        private static (long, long[], long) ReadData()
        {
            var size = long.Parse(Console.ReadLine());
            var data = Console.ReadLine().Split(' ');
            var result = new long[size];
            long sum = 0;

            for (var i = 0; i < size; i++)
            {
                sum += result[i] = long.Parse(data[i]);
            }

            return (size, result, sum);
        }
    }
}