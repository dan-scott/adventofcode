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

    public string Part1() => Run(Sub.New());

    public string Part2() => Run(AimedSub.New());
    
    private string Run(ISub sub) => Lines
        .Select(Instruction.Parse)
        .Aggregate(sub, (s, instruction) => s.RunInstruction(instruction))
        .Position.ProductOfParts().ToString();

}