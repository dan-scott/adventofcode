namespace AdventOfCode.Base;

public interface IDay : IDisposable
{ 
    int Number { get; }

    void Open();

    string Part1();

    string Part2();
}