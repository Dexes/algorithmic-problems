using System;

namespace Timus
{
    public static class Program
    {
        private static void Main()
        {
            var data = Console.In.ReadToEnd();
            var longest = LongestPalindrome(data);
            if (longest.Length == 1)
                Console.WriteLine(data[0]);
            else
                Console.WriteLine(longest);
        }

        private static string LongestPalindrome(string text)
        {
            // https://medium.com/hackernoon/manachers-algorithm-explained-longest-palindromic-substring-22cb27a5e96f
            var n = text.Length;
            n = 2 * n + 1;
            var l = new int[n + 1];
            (l[0], l[1]) = (0, 1);

            var c = 1;
            var r = 2;
            int i;
            var maxLpsLength = 0;
            var maxLpsCenterPosition = 0;

            for (i = 2; i < n; i++)
            {
                var iMirror = 2 * c - i;
                l[i] = 0;
                var diff = r - i;

                if (diff > 0)
                    l[i] = Math.Min(l[iMirror], diff);

                while (i + l[i] + 1 < n && i - l[i] > 0 &&
                       ((i + l[i] + 1) % 2 == 0 ||
                        text[(i + l[i] + 1) / 2] ==
                        text[(i - l[i] - 1) / 2]))
                {
                    l[i]++;
                }

                if (l[i] > maxLpsLength)
                {
                    maxLpsLength = l[i];
                    maxLpsCenterPosition = i;
                }

                if (i + l[i] > r)
                {
                    c = i;
                    r = i + l[i];
                }
            }

            var start = (maxLpsCenterPosition - maxLpsLength) / 2;
            var end = start + maxLpsLength - 1;

            return text.Substring(start, end - start + 1);
        }
    }
}