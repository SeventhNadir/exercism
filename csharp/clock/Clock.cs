using System;

public struct Clock
{
    private int _minutes;
    public Clock(int hours, int minutes)
    {
        _minutes = Mod((hours * 60) + minutes, 24*60);
    }

    public Clock Add(int minutesToAdd)
    {
        return new Clock(0, _minutes + minutesToAdd);
    }

    public Clock Subtract(int minutesToSubtract)
    {
        return new Clock(0, _minutes - minutesToSubtract);
    }

    public override string ToString()
    {
        return $"{_minutes/60:00}:{_minutes%60:00}";
    }

    private static int Mod(double x, double y)
    {
        return (int)((x % y + y) % y);
    }

}