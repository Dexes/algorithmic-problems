using System;

namespace Timus
{
    public static class Program
    {
        private static void Main()
        {
            var (size, current) = ReadRow();
            var (alienSum, humanSum) = (current, (size + 1) * 2);

            for (var i = 0; i < size; i++)
            {
                var (alien, human) = ReadRow();
                alienSum += alien;
                humanSum += human;
            }

            var result = alienSum - humanSum;
            if (result > 0)
                Console.WriteLine(result);
            else
                Console.WriteLine("Big Bang!");
        }

        private static (int, int) ReadRow()
        {
            var data = Console.ReadLine().Split(' ');
            return (int.Parse(data[0]), int.Parse(data[1]));
        }
    }
}