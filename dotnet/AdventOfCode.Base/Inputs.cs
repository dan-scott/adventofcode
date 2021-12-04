using System.Reflection;

namespace AdventOfCode.Base;

public static class Inputs
{
    private static readonly string RootDir = GetRootDir();

    private static string GetRootDir()
    {
        var rootFolder = System.Environment.GetEnvironmentVariable("ADVENT_OF_CODE_ROOT")
                         ?? Path.Join(Assembly.GetExecutingAssembly().Location, "../../");

        return Path.Join(rootFolder, "inputs");
    }

    private static string InputFileName(int year, int day) => Path.Join(RootDir, year.ToString(), $"{day}.txt");

    public static string Day(int year, int day) => File.ReadAllText(InputFileName(year, day));

    public static IEnumerable<string> Lines(int year, int day) => File.ReadLines(InputFileName(year, day));

    public static IEnumerable<int> LinesAsInt(int year, int day) => Lines(year, day).Select(int.Parse);

}