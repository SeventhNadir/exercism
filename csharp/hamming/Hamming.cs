using System;

public static class Hamming
{
    public static int Distance(string firstStrand, string secondStrand)
    {
        int lenFirst = firstStrand.Length;
        int lenSecond = secondStrand.Length;

    if (lenFirst != lenSecond){
        throw new ArgumentException("Cannot compare two unequal strings");
    }
        int HammingScore = 0;
        for (int i = 0; i < lenFirst; i++)
        {
            if (firstStrand[i] != secondStrand[i]){
                HammingScore += 1;
            }
        }

        return HammingScore;
    }
}