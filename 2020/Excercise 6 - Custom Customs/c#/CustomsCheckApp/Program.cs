using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Reflection;
using System.Text;

namespace CustomsCheckApp
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Custom Customs App");
       
            var groups = LoadCustomsAnswers();

            Console.WriteLine("Part 1----");
            var totalyes = groups.Sum(g => g.Answers.Distinct().Count());
            Console.WriteLine($"Total Annswered questions: {totalyes}");

            Console.WriteLine("Part 2: -----");
            var uniqueCount = 0;
            groups.ToList().ForEach(g =>
            {
                var countList = g.Answers.GroupBy(c => c).Select(g => new { g.Key, Count = g.Count() })
                .ToDictionary(j => j.Key, h => h.Count).Select(kv => kv.Value).Where(x => x == g.Size).ToList();
                
                uniqueCount += countList.Count();
          
            });

            Console.WriteLine($"Total everyone answered questions: {uniqueCount}");


        }


        static IList<Group> LoadCustomsAnswers()
        {
            var inputFilePath = Path.GetDirectoryName(Assembly.GetCallingAssembly().Location) + Path.DirectorySeparatorChar + "input.dat";

            var lines = File.ReadAllLines(inputFilePath);

            var sb = new StringBuilder();
            var answers = new List<Group>();
            var groupSize = 0;
            lines.ToList().ForEach(l =>
            {
                var text = l.Trim('\n').Trim('\r').TrimEnd();
                if(!string.IsNullOrEmpty(text))
                {
                    sb.Append(text);
                    groupSize++;
                }
                else
                {
                    answers.Add(new Group { Answers = sb.ToString(), Size = groupSize });
                    groupSize = 0;
                    sb = new StringBuilder();
                }


            });

            if (!string.IsNullOrEmpty(sb.ToString())){
                answers.Add(new Group { Answers = sb.ToString(), Size = groupSize });
            }

            return answers;
        }

    }

    class Group
    {
        public string Answers { get; set; }
        public int Size { get; set; }
    }
}
