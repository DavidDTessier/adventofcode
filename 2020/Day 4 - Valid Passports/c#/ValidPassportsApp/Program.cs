using System;
using System.Collections.Generic;
using System.IO;
using System.Linq;
using System.Reflection;
using System.Text;
using System.Text.RegularExpressions;

namespace ValidPassportsApp
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Valid Passport App!");
            var pds = LoadPassportData();
            Console.WriteLine("Part One");
            var totalValidPassPorts = 0;
            var totalValidStrictPassports = 0;
            pds.ToList().ForEach(x => {
                if (x.IsValidPassport()) totalValidPassPorts++;
                if (x.IsValidPassportStrict()) totalValidStrictPassports++;
            });

            Console.WriteLine($"Total Valid Passports: {totalValidPassPorts}");
            Console.WriteLine($"Total Valid Stict Passports: {totalValidStrictPassports}");
            Console.ReadLine();

        }



        static IList<PassportData> LoadPassportData()
        {
            var inputFilePath = Path.GetDirectoryName(Assembly.GetCallingAssembly().Location) + "\\input\\day4input.dat";

            var lines = File.ReadAllLines(inputFilePath);

            var sb = new StringBuilder();
            var passports = new List<PassportData>();

            lines.ToList().ForEach(l =>
            {
                var text = l.Trim('\n').Trim('\r');
                if (string.IsNullOrEmpty(text))
                {
                    passports.Add(PassportData.BuildPassportData(sb.ToString()));
                    sb = new StringBuilder();
                }
                else
                {
                    sb.Append(" ").Append(text);
                }

               
            });

            if(!string.IsNullOrEmpty(sb.ToString().TrimStart().TrimEnd()))
            {
                passports.Add(PassportData.BuildPassportData(sb.ToString()));
            }
            return passports;
        }
    }



    class PassportData
    {
        public string BirthYear { get; set; }
        public string IssueYear { get; set; }
        public string ExpirationYear { get; set; }
        public string HairColor { get; set; }
        public string Height { get; set; }
        public string EyeColor { get; set; }
        public string PassportId { get; set; }
        public string CountryId { get; set; }

        public bool IsValidPassport()
        {
            return HasAllRequiredFields() || Has7RequiredFields();
        }

        public bool IsValidPassportStrict()
        {
            var validBirthYear = !string.IsNullOrEmpty(BirthYear) && (int.Parse(BirthYear) >= 1920 && int.Parse(BirthYear) < 2003);
            var validIssueYear = !string.IsNullOrEmpty(IssueYear) && (int.Parse(IssueYear) >= 2010 && int.Parse(IssueYear) < 2021);
            var validExpYear = !string.IsNullOrEmpty(ExpirationYear) && (int.Parse(ExpirationYear) >= 2020 && int.Parse(ExpirationYear) < 2031);
            var validHieght = !string.IsNullOrEmpty(Height) && (new Regex(@"^(1([5-8][0-9]|9[0-3])cm|(59|6[0-9]|7[0-6])in)$").IsMatch(Height));
            var validHairColor = !string.IsNullOrEmpty(HairColor) && (new Regex(@"^#[0-9a-f]{6}$").IsMatch(HairColor));
            var validEyeColor = !string.IsNullOrEmpty(EyeColor) && (new string[] { "amb", "blu", "brn", "gry", "grn", "hzl", "oth" }).Contains(EyeColor);
            var validPassPortId = !string.IsNullOrEmpty(PassportId) && (new Regex(@"^[0-9]{9}$").IsMatch(PassportId));

            return validBirthYear && validExpYear && validIssueYear && validHieght && validHairColor && validEyeColor && validPassPortId;
        }

        public bool Has7RequiredFields()
        {
            return !string.IsNullOrEmpty(BirthYear)
                && !string.IsNullOrEmpty(IssueYear)
                && !string.IsNullOrEmpty(ExpirationYear)
                && !string.IsNullOrEmpty(HairColor)
                && !string.IsNullOrEmpty(EyeColor)
                && !string.IsNullOrEmpty(PassportId)
                && !string.IsNullOrEmpty(Height);
        }

        public bool HasAllRequiredFields()
        {
            return !string.IsNullOrEmpty(BirthYear)
                && !string.IsNullOrEmpty(IssueYear)
                && !string.IsNullOrEmpty(ExpirationYear)
                && !string.IsNullOrEmpty(HairColor)
                && !string.IsNullOrEmpty(EyeColor)
                && !string.IsNullOrEmpty(PassportId)
                && !string.IsNullOrEmpty(Height)
                && !string.IsNullOrEmpty(CountryId);
        }

        public static PassportData BuildPassportData(string passportDataString)
        {
            var str = passportDataString.TrimStart().TrimEnd().Split(" ");
            return new PassportData
            {
                BirthYear = str.ToList().FirstOrDefault(s => s!=null && s.Contains("byr"))?.Split(":")[1],
                IssueYear = str.ToList().FirstOrDefault(s => s != null && s.Contains("iyr"))?.Split(":")[1],
                ExpirationYear = str.ToList().FirstOrDefault(s => s != null && s.Contains("eyr"))?.Split(":")[1],
                EyeColor = str.ToList().FirstOrDefault(s => s != null && s.Contains("ecl"))?.Split(":")[1],
                HairColor = str.ToList().FirstOrDefault(s => s != null && s.Contains("hcl"))?.Split(":")[1],
                Height = str.ToList().FirstOrDefault(s => s != null && s.Contains("hgt"))?.Split(":")[1],
                PassportId = str.ToList().FirstOrDefault(s => s != null && s.Contains("pid"))?.Split(":")[1],
                CountryId = str.ToList().FirstOrDefault(s => s != null && s.Contains("cid"))?.Split(":")[1],
            };
        }


    }

}
