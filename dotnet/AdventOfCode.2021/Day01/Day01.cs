using AdventOfCode.Base;

namespace AdventOfCode._2021.Day01;

public class Day01 : IDay
{
    private List<int>? _readings;

    private List<int> Readings => _readings ?? new List<int>();
    
    public void Dispose()
    {
        _readings = null;
    }

    public int Number => 1;
    public void Open()
    {
        _readings = Inputs.LinesAsInt(2021, 1).ToList();
    }

    public string Part1() => Reading.ScanInput(Readings).Count(r => r.HasIncreased).ToString();

    public string Part2() => Reading.ScanWindows(Readings).Count(r => r.HasWindowIncreased).ToString();

}