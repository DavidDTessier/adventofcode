using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Reflection;

namespace src
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Loading input file!");
            var numbers = LoadInput();

            Console.WriteLine("Part 1: ----");
            DoPart1(numbers);
            Console.WriteLine("End Part 1: ----");
            Console.WriteLine("Part 2: ----");
            DoPart2(numbers);
            Console.WriteLine("End Part 2: ----");

            Console.Read();

        }

        static void DoPart2(string[] numbers)
        {
            numbers.ToList().ForEach(s =>
            {
                numbers.ToList().Where(s1 => s != s1).ToList().ForEach(s2 =>
                {

                    numbers.ToList().Where(s3 => s2 != s3).ToList().ForEach( s4 =>
                    {
                        int one = int.Parse(s);
                        int two = int.Parse(s2);
                        int three = int.Parse(s4);

                        if (CheckSum2020(one, two, three))
                        {
                            Console.WriteLine($"{one} + {two} + {three} == 2020");
                            Console.WriteLine($"{one} x {two} x {three} == {one * two * three}");
                        }
                    });

                });

            });
        }

        static void DoPart1(string[] numbers)
        {
            numbers.ToList().ForEach(s =>
            {
                numbers.ToList().Where(s1 => s != s1).ToList().ForEach(n =>
                {
                    int one = int.Parse(s);
                    int two = int.Parse(n);

                    if (CheckSum2020(one, two))
                    {
                        Console.WriteLine($"{one} + {two} == 2020");
                        Console.WriteLine($"{one} x {two} == {one * two}");
                    }
                });

            });
        }

        static string[] LoadInput()
        {
            var path = Path.GetDirectoryName(Assembly.GetExecutingAssembly().Location);
            string[] lines = File.ReadAllLines(path + "\\input\\input.txt");
            return lines;
        }

        static bool CheckSum2020(int one, int two, int three = 0)
        {
            return one + two + three == 2020;
        }
    }
}
