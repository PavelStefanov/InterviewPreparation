using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

namespace PermutationInString
{
    class Program
    {
        static void Main(string[] args)
        {
            var s1 = "ab";
            var s2 = "eidbaooo";

            var result = CheckInclusion(s1, s2);

            Console.WriteLine(result);
        }

        static bool CheckInclusion(string s1, string s2)
        {
            if (s1.Length > s2.Length)
                return false;

            var s1Arr = new int[26];
            foreach (var c in s1)
                s1Arr[c - 'a']++;

            var s2Arr = new int[26];
            for (var i = 0; i < s2.Length; i++)
            {
                if (i < s1.Length - 1)
                {
                    s2Arr[s2[i] - 'a']++;
                    continue;
                }
                else if (i == s1.Length - 1)
                {
                    s2Arr[s2[i] - 'a']++;
                }
                else
                {
                    s2Arr[s2[i - s1.Length ] - 'a']--;
                    s2Arr[s2[i] - 'a']++;
                }


                if (CompareArray(s1Arr, s2Arr))
                    return true;
            }

            return false;
        }

        static bool CompareArray(int[] arr1, int[] arr2)
        {
            for (var i = 0; i < arr1.Length; i++)
                if (arr1[i] != arr2[i])
                    return false;

            return true;
        }
    }
}
