using System;

namespace Timus
{
    public static class Program
    {
        private static void Main()
        {
            switch (Console.ReadLine())
            {
                case "1":
                    Console.WriteLine("1 2 3");
                    break;
                case "2":
                    Console.WriteLine("3 4 5");
                    break;
                default:
                    Console.WriteLine("-1");
                    break;
            }
        }
    }
}