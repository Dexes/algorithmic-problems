using System;

namespace Timus
{
    public static class Program
    {
        private static void Main()
        {
            var (rBalance, bBalance, anyBalance) = ReadLine();
            var (rDebt, bDebt, anyDebt) = ReadLine();

            (anyBalance, anyDebt) = ReduceDebt(rBalance, rDebt, anyBalance, anyDebt);
            (anyBalance, anyDebt) = ReduceDebt(bBalance, bDebt, anyBalance, anyDebt);

            if (anyBalance < 0 || anyBalance < anyDebt)
                Console.WriteLine("There are no miracles in life");
            else
                Console.WriteLine("It is a kind of magic");
        }

        private static (int, int) ReduceDebt(int balance, int debt, int anyBalance, int anyDebt)
        {
            if (balance < debt)
                anyBalance -= debt - balance;
            else
                anyDebt -= balance - debt;

            return (anyBalance, anyDebt);
        }

        private static (int, int, int) ReadLine()
        {
            var data = Console.ReadLine().Split(' ');
            return (int.Parse(data[0]), int.Parse(data[1]), int.Parse(data[2]));
        }
    }
}