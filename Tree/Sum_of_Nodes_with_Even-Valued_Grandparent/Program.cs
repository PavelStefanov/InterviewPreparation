

Console.WriteLine("Hello, World!");


//   Definition for a binary tree node.
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
    public int SumEvenGrandparent(TreeNode root)
    {
        return SumEvenGrandparent(root, false, false);
    }

    public int SumEvenGrandparent(TreeNode node, bool isGrandParentEven, bool isParentEven)
    {
        var result = 0;
        if (node is null)
        {
            return result;
        }

        var isEven = node.val % 2 == 0;

        if (isGrandParentEven)
        {
            result += node.val;
        }
        result += SumEvenGrandparent(node.left, isParentEven, isEven)
             + SumEvenGrandparent(node.right, isParentEven, isEven);

        return result;
    }
}