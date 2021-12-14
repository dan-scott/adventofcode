using AdventOfCode.Base;

namespace AdventOfCode._2021.Day02;

internal class AimedSub : ISub
{
    public static ISub New() => new AimedSub();
    public Vec2 Position { get; private set; } = Vec2.Z;
    private Vec2 _aim = Vec2.Z;

    public ISub RunInstruction(Instruction instruction)
    {
        var (direction, mag) = instruction;
        switch (direction)
        {
            case Direction.Forward:
                Position += _aim * mag + Vec2.Of(mag, 0);
                break;
            case Direction.Down:
                _aim += Vec2.Of(0, mag);
                break;
            case Direction.Up:
                _aim += Vec2.Of(0, -mag);
                break;
        }

        return this;
    }
}