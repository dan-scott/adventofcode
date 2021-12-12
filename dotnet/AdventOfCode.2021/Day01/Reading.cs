using System.Collections;

namespace AdventOfCode._2021.Day01;

internal class Reading : IEnumerable<Reading>
{
    private Reading? Previous { get; set; }
    private Reading? Next { get; set; }
    private int Value { get; set; }

    private static Reading Scan(IList<int> readings)
    {
        var first = new Reading {Value = readings.First()};
        var _ = readings.Skip(1).Aggregate(first, (prev, value) => prev.MakeNext(value));
        return first;
    }

    public static IEnumerable<Reading> ScanInput(IList<int> readings) => Scan(readings);

    public static IEnumerable<Reading> ScanWindows(IList<int> readings)
    {
        var first = Scan(readings);
        return first.Next?.Next ?? throw new ArgumentException("Readings list too short", nameof(readings));
    }

    public bool HasIncreased => PreviousValue < Value;

    public bool HasWindowIncreased => PreviousWindowValue < WindowValue;

    private int PreviousValue => Previous?.Value ?? 0;

    private int PreviousWindowValue => Previous?.WindowValue ?? 0;

    private int WindowValue
        => Value
           + (Previous?.Value ?? 0)
           + (Previous?.Previous?.Value ?? 0);

    private Reading MakeNext(int value)
        => Next = new Reading
        {
            Previous = this,
            Value = value,
        };

    private class Enumerator : IEnumerator<Reading>
    {
        private readonly Reading _first;
        private Reading _current;

        public Enumerator(Reading first)
        {
            _first = first;
            _current = first;
        }

        public bool MoveNext()
        {
            if (_current.Next == null)
            {
                return false;
            }

            _current = _current.Next;
            return true;
        }

        public void Reset()
        {
            _current = _first;
        }

        public Reading Current => _current;

        object IEnumerator.Current => Current;

        public void Dispose()
        {
        }
    }

    public IEnumerator<Reading> GetEnumerator()
    {
        return new Enumerator(this);
    }

    IEnumerator IEnumerable.GetEnumerator()
    {
        return GetEnumerator();
    }
}