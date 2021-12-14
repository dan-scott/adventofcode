namespace AdventOfCode.Base;

public readonly record struct Vec2(int X, int Y)
{
    public static readonly Vec2 Z = Of(0, 0);

    public static Vec2 Of(int x, int y) => new(x, y);

    public static Vec2 operator +(Vec2 left, Vec2 right) => left.Add(right);
    public static Vec2 operator *(Vec2 left, int right) => left.Scale(right);

    private Vec2 Add(Vec2 other) => Of(X + other.X, Y + other.Y);

    private Vec2 Scale(int magnitude) => Of(X * magnitude, Y * magnitude);
}