using System;

namespace Timus
{
    public static class Program
    {
        private static void Main()
        {
            var (grille, symbols) = ReadData();

            for (var i = 0; i < 4; i++)
            {
                PrintSymbols(symbols, grille);
                grille = RotateGrille(grille);
            }
        }

        private static void PrintSymbols(char[,] symbols, bool[,] grille)
        {
            for (var i = 0; i < 4; i++)
            for (var j = 0; j < 4; j++)
                if (grille[i, j])
                    Console.Write(symbols[i, j]);
        }

        private static bool[,] RotateGrille(bool[,] grille)
        {
            var result = new bool[4, 4];

            for (var i = 0; i < 4; i++)
            for (var j = 0; j < 4; j++)
                result[i, j] = grille[3 - j, i];

            return result;
        }

        private static (bool[,], char[,]) ReadData()
        {
            var grille = new bool[4, 4];
            var symbols = new char[4, 4];

            for (var i = 0; i < 4; i++)
            {
                var line = Console.ReadLine();
                for (var j = 0; j < 4; j++)
                    grille[i, j] = line[j] == 'X';
            }

            for (var i = 0; i < 4; i++)
            {
                var line = Console.ReadLine();
                for (var j = 0; j < 4; j++)
                    symbols[i, j] = line[j];
            }

            return (grille, symbols);
        }
    }
}