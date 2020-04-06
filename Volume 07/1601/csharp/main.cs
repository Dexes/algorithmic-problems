using System;
using System.Collections.Generic;

namespace Timus
{
    public static class Program
    {
        private static void Main()
        {
            var isFirst = true;
            foreach (var ch in Input())
            {
                Console.Write(isFirst ? ch : char.ToLower(ch));
                if (isFirst)
                    isFirst = !char.IsLetter(ch);
                else
                    isFirst = ch == '.' || ch == '?' || ch == '!';
            }
        }

        private static IEnumerable<char> Input()
        {
            while (true)
            {
                var ch = Console.Read();
                if (ch < 0) break;

                yield return (char) ch;
            }
        }
    }
}