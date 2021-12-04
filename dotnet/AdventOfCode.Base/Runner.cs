using System.Diagnostics;

namespace AdventOfCode.Base;

public static class Runner
{
    public static void Run(IEnumerable<IDay> days)
    {
        var overall = new Stopwatch();
        var daySw = new Stopwatch();

        foreach (var day in days)
        {
            daySw.Restart();
            Console.Write($"Loading day {day.Number}...");
            day.Open();
            Console.WriteLine($" done ({daySw.Elapsed})");
            
            daySw.Restart();
            Console.Write("\tSolving part 1...");
            var result = day.Part1();
            Console.WriteLine($"done ({daySw.Elapsed}) {result, 20}");
            
            daySw.Restart();
            Console.Write("\tSolving part 2...");
            result = day.Part2();
            Console.WriteLine($"done ({daySw.Elapsed}) {result, 20}\n");
        }
        
        Console.WriteLine($"Solutions complete in {overall.Elapsed}");
    }
}