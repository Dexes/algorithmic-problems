using System;

namespace Timus
{
    public static class Program
    {
        private static void Main()
        {
            var size = int.Parse(Console.ReadLine());
            var shelf = ReadBox();

            for (var i = 1; i < size; i++)
            {
                shelf = InsertBox(shelf, ReadBox());
                if (shelf != null) continue;

                Console.WriteLine("NO");
                return;
            }

            Console.WriteLine("YES");
        }

        private static Box ReadBox()
        {
            var data = Console.ReadLine().Split(' ');
            var count = int.Parse(data[0]);

            var previous = int.Parse(data[1]);
            for (var i = 2; i <= count; i++)
            {
                var current = int.Parse(data[i]);
                if (current < previous)
                {
                    Console.WriteLine("NO");
                    Environment.Exit(0);
                }

                previous = current;
            }

            return new Box
            {
                FirstSize = int.Parse(data[1]),
                LastSize = int.Parse(data[count]),
            };
        }

        private static Box InsertBox(Box shelf, Box box)
        {
            if (box.LastSize <= shelf.FirstSize)
            {
                box.NextNode = shelf;
                return box;
            }

            var current = shelf;
            while (true)
            {
                if (current.NextNode == null)
                {
                    if (current.LastSize <= box.FirstSize)
                    {
                        current.NextNode = box;
                        return shelf;
                    }

                    return null;
                }

                if (current.LastSize <= box.FirstSize && box.LastSize <= current.NextNode.FirstSize)
                {
                    box.NextNode = current.NextNode;
                    current.NextNode = box;
                    return shelf;
                }

                current = current.NextNode;
            }
        }
    }

    internal class Box
    {
        public int FirstSize;
        public int LastSize;
        public Box NextNode;
    }
}