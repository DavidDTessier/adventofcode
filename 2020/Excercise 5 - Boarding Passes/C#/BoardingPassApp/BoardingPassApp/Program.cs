using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Reflection;
using System.Text;
using System.Text.RegularExpressions;

namespace BoardingPassApp
{
    class Program
    {
        

        static void Main(string[] args)
        {
            Console.WriteLine("Part 1!");
            var bps = LoadBoardingPasses();
            var highestSeat = bps.ToList().Max(bp => bp.GetSeatId());
            Console.WriteLine($"Highest Seat ID : {highestSeat}");

            Console.WriteLine("Part 2:");
            var seatIds = bps.OrderBy(b => b.GetSeatId()).Select(b => b.GetSeatId()).ToList();
            var minSeat = seatIds.First() - 1;
            var maxSeat = seatIds.Last();
            var sumSeats = seatIds.Sum();
            var possibleSeat = ((maxSeat * (maxSeat + 1) / 2) - (minSeat * (minSeat + 1) / 2) - sumSeats);
            Console.WriteLine($"My Seat is: {possibleSeat}");
        }

        static IList<BoardingPass> LoadBoardingPasses()
        {
            var inputFilePath = Path.GetDirectoryName(Assembly.GetCallingAssembly().Location) + "/input/input.dat";

            var lines = File.ReadAllLines(inputFilePath);

            var sb = new StringBuilder();
            var bps = new List<BoardingPass>();

            lines.ToList().ForEach(l =>
            {
                var text = l.Trim('\n').Trim('\r').TrimEnd();

                bps.Add(new BoardingPass
                {
                    PassNumber = text,
                    Row = DecodeSeatRow(text.Substring(0,7)),
                    Column = DecodeSeatRow(text.Substring(7,3))
                });


            });

            return bps;
        }

        static int DecodeSeatRow(string row)
        {
            return Convert.ToInt32(Regex.Replace(Regex.Replace(row, "[BR]", "1"), "[FL]", "0"), 2);
        }
    }

    class BoardingPass
    {
        public string PassNumber { get; set; }
        public int Row { get; set; }
        public int Column { get; set; }

        public int GetSeatId()
        {
            return Row * 8 + Column;
        }
    }
}
