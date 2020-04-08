using System;

namespace Timus
{
    public static class Program
    {
        private static void Main()
        {
            var maximums = GetSequenceMaximums(100000);

            while (true)
            {
                var index = int.Parse(Console.ReadLine());
                if (index == 0) return;

                Console.WriteLine(maximums[index]);
            }
        }

        private static int[] GetSequenceMaximums(int sequenceSize)
        {
            var sequence = new int[sequenceSize];
            var maximums = new int[sequenceSize];

            maximums[1] = sequence[1] = 1;

            var j = 0;
            for (var i = 2; i < sequenceSize; i += 2)
            {
                j++;
                sequence[i] = sequence[j];
                sequence[i + 1] = sequence[j] + sequence[j + 1];

                maximums[i] = Math.Max(maximums[i - 1], sequence[i]);
                maximums[i + 1] = Math.Max(maximums[i], sequence[i + 1]);
            }

            return maximums;
        }
    }
}