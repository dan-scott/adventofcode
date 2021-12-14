using AdventOfCode.Base;

namespace AdventOfCode._2021.Day02;

internal interface ISub
{
    Vec2 Position { get; }

    ISub RunInstruction(Instruction instruction);
}