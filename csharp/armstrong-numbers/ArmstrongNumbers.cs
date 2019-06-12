using System;
using System.Linq;

public static class ArmstrongNumbers
{
    public static bool IsArmstrongNumber(int number)
    {
        int sum = 0;
        int[] numberArray = Array.ConvertAll(number.ToString().ToArray(), x=>(int)x - 48);
        Array.ForEach(numberArray, i => {sum += (int)Math.Pow(i, numberArray.Count());} );
        if (number == sum) {
            return true;
        }

        return false;
    }
}