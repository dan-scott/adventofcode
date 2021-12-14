using AdventOfCode.Base;

namespace AdventOfCode._2021.Day02;

internal class Sub : ISub
{
    public static ISub New() => new Sub();
    public Vec2 Position { get; private set; } = Vec2.Z;

    public ISub RunInstruction(Instruction instruction)
    {
        Position += instruction switch
        {
            (Direction.Forward, var mag) => Vec2.Of(mag, 0),
            (Direction.Down, var mag) => Vec2.Of(0, mag),
            (Direction.Up, var mag) => Vec2.Of(0, -mag),
            _ => Vec2.Z
        };
        return this;
    }
}