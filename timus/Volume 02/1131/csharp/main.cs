using System;

namespace Timus
{
    public static class Program
    {
        private static void Main()
        {
            var (computers, cables) = ReadData();
            if (computers == 1)
            {
                Console.WriteLine(0);
                return;
            }

            var empty = computers - 1;
            var configured = 1;
            var time = 0;

            while (empty > 0)
            {
                empty -= configured;
                time++;
                configured *= 2;

                if (configured >= cables)
                {
                    time += empty / cables;
                    if (empty % cables > 0) time++;
                    break;
                }
            }

            Console.WriteLine(time);
        }

        private static (int, int) ReadData()
        {
            var data = Console.ReadLine().Split(' ');
            return (int.Parse(data[0]), int.Parse(data[1]));
        }
    }
}