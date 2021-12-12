using System.Numerics;
using AdventOfCode.Base;

namespace AdventOfCode._2021.Day02;

public class Day02 : IDay
{
    private List<string>? _lines;
    private List<string> Lines => _lines ?? new List<string>();
    
    public void Dispose()
    {
        _lines = null;
    }

    public int Number => 2;

    public void Open()
    {
        _lines = Inputs.Lines(2021, 2).ToList();
    }

    public string Part1()
    {
        throw new NotImplementedException();
    }

    public string Part2()
    {
        throw new NotImplementedException();
    }
}

internal class Sub
{
    private Vec2 _pos = Vec2.Z;
}