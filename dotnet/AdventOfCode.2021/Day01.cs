using AdventOfCode.Base;

namespace AdventOfCode._2021;

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

    public string Part1() => Readings
        .Aggregate(
            (0, int.MaxValue), 
            (agg, curr) => (
                agg.Item2 < curr 
                    ? agg.Item1 + 1 
                    : agg.Item1, 
                curr))
        .Item1
        .ToString();

    public string Part2() => "";
}