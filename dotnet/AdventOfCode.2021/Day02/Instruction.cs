namespace AdventOfCode._2021.Day02;

internal readonly record struct Instruction(Direction Type, int Magnitude)
{
    internal static Instruction Parse(string input) => new(ParseType(input), ParseMagnitude(input));

    private static int ParseMagnitude(string input) => input[^1] - '0';

    private static Direction ParseType(string input) =>
        input[0] switch
        {
            'f' => Direction.Forward,
            'd' => Direction.Down,
            _ => Direction.Up,
        };
}