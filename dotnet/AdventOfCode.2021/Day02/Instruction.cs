namespace AdventOfCode._2021.Day02;

internal struct Instruction
{
    internal static Instruction Parse(string input) => new(ParseDirection(input), ParseMagnitude(input));

    private static int ParseMagnitude(string input) => input[^1] - '0';

    private static Direction ParseDirection(string input) =>
        input[0] switch
        {
            'f' => Direction.Forward,
            'd' => Direction.Down,
            _ => Direction.Up,
        };

    private Instruction(Direction operation, int magnitude)
    {
        Operation = operation;
        Magnitude = magnitude;
    }

    public Direction Operation { get; }
    public int Magnitude { get; }

    internal enum Direction
    {
        Forward,
        Down,
        Up
    }
}
