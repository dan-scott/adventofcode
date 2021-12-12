namespace AdventOfCode.Base;

public record struct Vec2<T>(T X, T Y) where T : INumber<T>
{
    public static readonly Vec2<T> Z = Of(T.Zero, T.Zero);
    public static Vec2<T> Of(T x, T y) => new(x, y);

    public Vec2<T> Add(Vec2<T> other) => Of(X + other.X, Y + other.Y);

    public T ManhattanDistanceToOrigin => T.Abs(X) + T.Abs(Y);
}

public record struct Vec2(int X, int Y)
{
    public static readonly Vec2 Z = Of(0, 0);

    public static Vec2 Of(int x, int y) => new(x, y);

    public static Vec2 operator +(Vec2 left, Vec2 right) => left.Add(right);
    
    public Vec2 Add(Vec2 other) => Of(X + other.X, Y + other.Y);

    public int ManhattanDistanceToOrigin => Math.Abs(X) + Math.Abs(Y);
}