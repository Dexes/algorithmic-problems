using System;

namespace Timus
{
    public static class Program
    {
        private static void Main()
        {
            Console.ReadLine();
            var n = uint.Parse(Console.ReadLine());

            var bytes = new[] {n & 0xFF, n & (0xFF << 8), n & (0xFF << 16), n & (0xFF << 24)};
            Console.WriteLine(bytes[0] << 24 | bytes[1] << 8 | bytes[2] >> 8 | bytes[3] >> 24);
        }
    }
}