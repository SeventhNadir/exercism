using System;
using System.Collections.Generic;

public static class BracketPush
{
    private const string OpenBrackets = "{[(";
    private const string CloseBrackets = "}])";
    private static readonly string[] ValidPairs = { "{}", "[]", "()" };
    public static bool IsPaired(string input)
    {
        int index = 0;
        Stack<(int, char)> brackets = new Stack<(int, char)>();

        foreach (char c in input)
        {
            if (OpenBrackets.Contains(c))
            {
                brackets.Push((index, c));
            }

            if (CloseBrackets.Contains(c))
            {
                // Closing bracket with no open bracket
                if (brackets.Count == 0) return false;
                var bracketItem = brackets.Peek();
                //  Closing brackets don't match the first item on the stack, incorrect bracket pairing
                if (!Array.Exists(ValidPairs, pairing => pairing == bracketItem.Item2 + c.ToString())) return false;
                 
                var substring = input.Substring(bracketItem.Item1, index - bracketItem.Item1 + 1);
                if (input.Length > substring.Length)
                {   
                    // Recursive call test test the contents of the current bracket set
                    if (BracketPush.IsPaired(substring) == false) return false;  
                }
                brackets.Pop();
            } 

            index++;
        }
        if (brackets.Count != 0) 
        {
            return false;
        }
        return true;
    }
}