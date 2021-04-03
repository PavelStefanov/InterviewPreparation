using System;

namespace Permutation
{
    class Program
    {
        static void Main(string[] args)
        {
            var Permutation = GetAllPermutation("World");

            foreach (var p in Permutation)
                Console.WriteLine(p);
        }

        static string[] GetAllPermutation(string str)
        {
            var result = new List<string>();

            
        }
    }
}
