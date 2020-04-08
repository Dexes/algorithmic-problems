using System;

namespace Timus
{
    internal enum Direction
    {
        Right,
        Left,
    }

    public static class Program
    {
        private static void Main()
        {
            var (count, data) = ReadData();
            var changes = CalculateChanges(count, data);

            Console.WriteLine(Math.Min(
                Math.Min(changes[0], changes[1]),
                Math.Min(changes[2], changes[3])
            ));
        }

        private static int[] CalculateChanges(int count, Direction[] data)
        {
            // cases:
            // 0 - < < < < > > > >
            // 1 - > > > > < < < <
            // 2 - < > < > < > < >
            // 3 - > < > < > < > <

            var half = count / 2;
            var changes = new int[4];

            // first and second cases
            for (var i = 0; i < half; i++)
            {
                var caseIndex = data[i] == Direction.Right ? 0 : 1;
                changes[caseIndex]++;
            }

            for (var i = half; i < count; i++)
            {
                var caseIndex = data[i] == Direction.Left ? 0 : 1;
                changes[caseIndex]++;
            }

            // third and fourth cases
            for (var i = 0; i < count; i += 2)
            {
                var caseIndex = data[i] == Direction.Right ? 2 : 3;
                changes[caseIndex]++;
            }
            
            for (var i = 1; i < count; i += 2)
            {
                var caseIndex = data[i] == Direction.Left ? 2 : 3;
                changes[caseIndex]++;
            }

            return changes;
        }

        private static (int, Direction[]) ReadData()
        {
            var count = int.Parse(Console.ReadLine());
            var data = Console.ReadLine();
            var result = new Direction[count];

            for (var i = 0; i < count; i++)
                result[i] = data[i * 5] == '<' ? Direction.Left : Direction.Right;

            return (count, result);
        }
    }
}