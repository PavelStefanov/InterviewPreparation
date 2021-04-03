using System;
using System.Collections.Generic;
using System.Linq;

namespace PalindromePermutation
{
    class Program
    {
        static void Main(string[] args)
        {
            var input = "carerac";

            var result = FindPalindromePermutation(input);

            Console.WriteLine(result);
        }

        static bool FindPalindromePermutation(string str)
        {
            return MakePermutation(str.ToList(), "");
        }

        static bool MakePermutation(List<char> chars, string perm)
        {
            if (chars.Count == 0)
                return IsPalindrome(perm);

            for (var i = 0; i < chars.Count; i++)
            {
                var p = perm + chars[i];
                var c = chars.ToList();
                c.RemoveAt(i);

                var res = MakePermutation(c, p);
                if (res)
                    return res;
            }

            return false;
        }

        static bool IsPalindrome(string str)
        {
            var mid = str.Length / 2;
            for (var i = 0; i <= mid - 1; i++)
            {
                if (str[i] != str[str.Length - 1 - i])
                    return false;
            }

            return true;
        }
    }
}
