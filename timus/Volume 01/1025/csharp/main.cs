using System;

namespace Timus
{
    public static class Program
    {
        private static void Main()
        {
            var (size, data) = ReadData();
            Console.WriteLine(CalculateResult(size, data));
        }

        private static int CalculateResult(int size, int[] data)
        {
            Array.Sort(data);

            var end = size / 2 + 1;
            var result = 0;
            for (var i = 0; i < end; i++)
                result += data[i] / 2 + 1;

            return result;
        }

        private static (int, int[]) ReadData()
        {
            var size = int.Parse(Console.ReadLine());
            var result = new int[size];

            for (var i = 0; i < size; i++)
                result[i] = ReadInt();

            return (size, result);
        }

        private static int ReadInt()
        {
            var result = 0;
            var current = Console.Read();

            while (current >= '0' && current <= '9')
            {
                result = result * 10 + (current - '0');
                current = Console.Read();
            }

            return result;
        }
    }
}