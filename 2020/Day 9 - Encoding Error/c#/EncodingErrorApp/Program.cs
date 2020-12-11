using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Reflection;

namespace EncodingErrorApp
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Hello World!");
            var invalidNumber = FindInvalidNumber();

            Console.WriteLine($"First Invalid Number is: {invalidNumber}");

            var sum = GetWeakness(invalidNumber);
            Console.WriteLine($"Weekness sum is: {sum}");

        }

        static long FindInvalidNumber()
        {
            var data = LoadData();
            var preambleSize = 25;
            var invalidNumber = -1L;
            for (int idx=preambleSize; idx <= data.Count; idx++)
            {
                var startIdx = idx - preambleSize;
                var preambleSet = data.ToArray()[new Range(startIdx,idx)];
                var currentNumber = data[idx];
                if (!PreambleSums(preambleSet.ToList()).Contains(currentNumber))
                {
                    invalidNumber = currentNumber;
                    break;
                }
            }

            return invalidNumber;
        }

        static long GetWeakness(long targetNumber)
        {
            var input = LoadData().ToList();
            var result = 0l;
            for( var idx = 0; idx < input.Count; idx++)
            {
                if (input[idx] == targetNumber)
                    continue;

                var sum = 0l;
                var nextIdx = idx;
                while (sum < targetNumber)
                {
                    sum += input[nextIdx];
                    nextIdx++;
                }

                if(sum == targetNumber)
                {
                    var set = input.ToArray()[new Range(idx, nextIdx-1)];
                    result = set.Min() + set.Max();
                    break;
                }
            }

            return result;
        }


        static IList<long> PreambleSums(IList<long> preambleSet)
        {
            var validSum = new List<long>();
            preambleSet.ToList().ForEach(l =>
            {
                preambleSet.ToList().Where(l1=> l1!=l).ToList().ForEach(l2 =>
                {
                    validSum.Add(l + l2);
                });
            });

            return validSum;
        }

        static IList<long> LoadData()
        {
            var inputFilePath = Path.GetDirectoryName(Assembly.GetCallingAssembly().Location)
                + Path.DirectorySeparatorChar + "input" + Path.DirectorySeparatorChar + "input.dat";

            var instructions = File.ReadAllLines(inputFilePath).ToList().Select(l => long.Parse(l)).ToList();
            return instructions;
        }

    }
}
