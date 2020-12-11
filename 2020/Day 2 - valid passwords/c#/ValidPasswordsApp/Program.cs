using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Reflection;

namespace ValidPasswordsApp
{
    class Program
    {
      
        static void Main(string[] args)
        {
            Console.WriteLine("Load Password Policies...");
            var policies = LoadPolicies();
            Console.WriteLine("Start of Part 1: -----");
            DoPart1(policies);
            Console.WriteLine("End Of Part 1 : -----");
            Console.WriteLine("Start Part 2: ------");
            DoPart2(policies);
            Console.WriteLine("End of Part 2.");
        }

        static bool IsValidCharAt(string text, int idx, char val)
        {
            var zeroIdx = idx - 1;
            return zeroIdx >= 0 && zeroIdx <= (text.Length-1) && text.ElementAt(zeroIdx) == val;
        }

        private static void DoPart2(List<Policy> policies)
        {
            int validPaswordIdx = 0;
            policies.ToList().ForEach(p =>
            {
                
                var isMinValid = IsValidCharAt(p.Password, p.Min, p.TargetCharacter);
                var isMaxValid = IsValidCharAt(p.Password, p.Max, p.TargetCharacter);
                if (isMaxValid ^ isMinValid)
                    validPaswordIdx += 1;
            });

            Console.WriteLine($"Total Valid Password Indexes: {validPaswordIdx}");
        }

        static void DoPart1(IList<Policy> policies)
        {
            int validPaswords = 0;
            policies.ToList().ForEach(p =>
            {
                var count = p.Password.Count(f => f == p.TargetCharacter);
                if (count >= p.Min && count <= p.Max)
                    validPaswords += 1;
            });

            Console.WriteLine($"Total Valid Passwords: {validPaswords}");
        }

        static List<Policy> LoadPolicies()
        {
            var path = Path.GetDirectoryName(Assembly.GetExecutingAssembly().Location);
            var lines =  File.ReadLines(path + "\\input\\day2input.txt");
            var policies = new List<Policy>();
            lines.ToList().ForEach(s =>
            {
                var strArray = s.Split(" ");

                policies.Add(new Policy
                {
                    Min = int.Parse(strArray[0].Split("-")[0]),
                    Max = int.Parse(strArray[0].Split("-")[1]),
                    TargetCharacter = char.Parse(strArray[1].Trim(':')),
                    Password = strArray[2]
                });
            });
            return policies;
        }
    }

    class Policy
    {
        public int Min { get; set; }
        public int Max { get; set; }
        public char TargetCharacter { get; set; }
        public string Password { get; set; }

    }
}
