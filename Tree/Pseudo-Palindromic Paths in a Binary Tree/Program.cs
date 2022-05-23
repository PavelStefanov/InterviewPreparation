using System;
using System.Collections.Generic;
using System.Linq;

namespace Pseudo_Palindromic_Paths_in_a_Binary_Tree
{
    class Program
    {
        static void Main(string[] args)
        {
            Console.WriteLine("Hello World!");
        }
    }


    //  Definition for a binary tree node.
    public class TreeNode
    {
        public int val;
        public TreeNode left;
        public TreeNode right;
        public TreeNode(int val = 0, TreeNode left = null, TreeNode right = null)
        {
            this.val = val;
            this.left = left;
            this.right = right;
        }
    }

    public class Solution
    {
        public int PseudoPalindromicPaths(TreeNode root)
        {
            var c = 0;
            var path = new List<int>();
            var stack = new Stack<TreeNode>();
            stack.Push(root);

            while (stack.Peek() != null)
            {
                var node = stack.Pop();

                path.Add(node.val);
                if (node.left == null && node.right == null)
                {
                    if (IsPseudoPalindromic(path))
                    {
                        c++;
                    }
                    path.RemoveAt(path.Count - 1);
                }
                else
                {
                    if (node.right != null)
                        stack.Push(node.right);
                    if (node.left != null)
                        stack.Push(node.left);
                }
            }


            return c;
            // var pathes = GetPathes(root, new List<int>());
            // return pathes.Where(IsPseudoPalindromic).Count();
        }

        private List<List<int>> GetPathes(TreeNode node, List<int> prev)
        {
            var res = new List<List<int>>();

            var cur = new List<int>();
            cur.AddRange(prev);
            cur.Add(node.val);

            if (node.left == null && node.right == null)
            {
                res.Add(cur);
            }
            else
            {
                if (node.left != null)
                {
                    res.AddRange(GetPathes(node.left, cur));
                }
                if (node.right != null)
                {
                    res.AddRange(GetPathes(node.right, cur));
                }
            }

            return res;
        }

        private bool IsPseudoPalindromic(List<int> nums)
        {
            int isPalindrom = 0;

            for (int i = 1; i < 10; ++i)
            {
                if (nums.Where(n => n == i).Count() % 2 == 1)
                {
                    isPalindrom++;
                    if (isPalindrom > 1)
                    {
                        return false;
                    }
                }
            }
            return true;
        }
    }
}
