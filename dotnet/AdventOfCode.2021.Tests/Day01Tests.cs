using Xunit;

namespace AdventOfCode._2021.Tests;

public class Day01Tests
{
    [Fact]
    public void TestPart1()
    {
        var day = new Day01.Day01();
        day.Open();
        Assert.Equal("1696", day.Part1());
    }
    
    [Fact]
    public void TestPart2()
    {
        var day = new Day01.Day01();
        day.Open();
        Assert.Equal("1737", day.Part2());
    }
}