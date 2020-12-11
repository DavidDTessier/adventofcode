using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Reflection;
using System.Text.RegularExpressions;

namespace BaggageApp
{
    class Program
    {
        const string myBag = "shiny gold";
        static void Main(string[] args)
        {
            Console.WriteLine("Part 1: --");
            var bags = LoadBags();
            var count = 0;
            bags.ToList().ForEach(b =>
            {
                if(ContainsBagColors(myBag, b.Key, bags))
                {
                    count++;
                }

            });

            Console.WriteLine($"Count of bag colors that could contain {myBag} : {count}");

            Console.WriteLine("Part 2: ---");
            count = CountBags(myBag, bags);
            Console.WriteLine($"Number of bag that {myBag} can contain: {count}");

        }

        static int CountBags(string bagName, IDictionary<string, Dictionary<string, int>> bags)
        {
            var count = 0;
            foreach(var kv in bags[bagName])
            {
                count += kv.Value + kv.Value * (CountBags(kv.Key, bags));
            }
            return count;
        }

        static bool ContainsBagColors(string desiredBagColor, string bagName, IDictionary<string, Dictionary<string, int>> bags)
        {
            foreach( var kv in bags[bagName])
            {
                if (kv.Key.Equals(desiredBagColor))
                {
                    return true;
                }
                else if(ContainsBagColors(desiredBagColor, kv.Key, bags))
                {
                    return true;
                }
            }

            return false;
        }
        

        static IDictionary<string, Dictionary<string, int>> LoadBags()
        {
            var inputFilePath = Path.GetDirectoryName(Assembly.GetCallingAssembly().Location)
                + Path.DirectorySeparatorChar + "input.dat";

            var lines = File.ReadAllLines(inputFilePath);
            var bags = new Dictionary<string, Dictionary<string, int>>();
            lines.ToList().ForEach(l =>
            {
                var str = l.Split("contain");
                string bag = str[0].Replace(" bags ", "");
                var childBags = new Dictionary<string, int>();
                var cbs = str[1].Split(",");
                cbs.ToList().Where(s=>s.Trim() != "no other bags.").ToList().ForEach(cb =>
                {
                    var cn = int.Parse(Regex.Match(cb.Trim(), @"\d+").Value);
                    var bag = Regex.Replace(cb.Trim(), @"\d+", "").Replace(".", "").Replace("bags","").Replace("bag","").Trim();
                    childBags.Add(bag, cn);
                });

                bags.Add(bag,childBags);
            });

            return bags;
        }
    }
}

