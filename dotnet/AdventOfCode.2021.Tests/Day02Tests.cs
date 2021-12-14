using Xunit;

namespace AdventOfCode._2021.Tests;

public class Day02Tests
{
    [Fact]
    public void TestPart1()
    {
        var day = new Day02.Day02();
        day.Open();
        Assert.Equal("1727835", day.Part1());
    }
    
    [Fact]
    public void TestPart2()
    {
        var day = new Day02.Day02();
        day.Open();
        Assert.Equal("1544000595", day.Part2());
    }
}