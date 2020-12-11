using System;
using System.IO;
using System.Linq;
using System.Reflection;

namespace FindTreesApp
{
    class Program
    {
        const char tree = '#';
        static void Main(string[] args)
        {
            Console.WriteLine("Find Trees!");
            var treeRows = LoadTreeMap();
            Part1(treeRows);
            Part2(treeRows);

            Console.ReadLine();
        }

        static void Part1(char[][] treeMap)
        {
            var foundTrees = FindTrees(3, 1, treeMap);
            Console.WriteLine($"Part 1 Trees Count {foundTrees}.");
        }

        static void Part2(char[][] treeMap)
        {
            var total = FindTrees(1, 1, treeMap)
                    * FindTrees(3, 1, treeMap)
                    * FindTrees(5, 1, treeMap)
                    * FindTrees(7, 1, treeMap)
                    * FindTrees(1, 2, treeMap);

            Console.WriteLine($"Part 2 Result: {total}");
        }

        static long FindTrees(int posX, int posY, char[][]rows)
        {
            var numberOfTrees = 0;
            int colLength = rows[0].Length;

            for (int x = 0, y = 0; y < rows.Length; x += posX, y += posY)
            {
                if(rows[y][x % colLength] == tree)
                {
                    numberOfTrees++;
                }
            }

            return numberOfTrees;
        }



        static char[][] LoadTreeMap()
        {
            var path = Path.GetDirectoryName(Assembly.GetExecutingAssembly().Location);
            var lines = File.ReadLines(path + "\\input\\day3input.dat");
            var rows = lines.ToList().Select(s => s.Trim('\n').Trim('\r').ToCharArray()).ToArray();
            return rows;
        }


    }
}
