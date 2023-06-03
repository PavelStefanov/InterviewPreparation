// See https://aka.ms/new-console-template for more information
var answer = new Solution().DailyTemperatures(new int[] { 34, 80, 80, 34, 34, 80, 80, 80, 80, 34 });
Console.WriteLine("[{0}]", string.Join(", ", answer));


public class Solution
{
    public int[] DailyTemperatures(int[] temperatures)
    {
        var answers = new int[temperatures.Length];

        for (var i = temperatures.Length - 2; i >= 0; i--)
        {
            if (temperatures[i] < temperatures[i + 1])
            {
                answers[i] = 1;
            }
            else if (temperatures[i] == temperatures[i + 1])
            {
                if (answers[i + 1] == 0)
                {
                    answers[i] = 0;
                }
                else
                {
                    answers[i] = answers[i + 1] + 1;
                }
            }
            else
            {
                answers[i] = answers[i + 1];

                var j = i + answers[i + 1] + 1;

                if (temperatures[i] < temperatures[j])
                {
                    answers[i]++;
                    continue;
                }
                if (temperatures[i] > temperatures[j] && answers[j] == 0)
                {
                    answers[i] = 0;
                    continue;
                }

                for (; j < temperatures.Length;)
                {
                    if (temperatures[i] < temperatures[j])
                    {
                        answers[i]++;
                        break;
                    }
                    else if (temperatures[i] >= temperatures[j] && answers[j] == 0)
                    {
                        answers[i] = 0;
                        break;
                    }
                    else
                    {
                        answers[i] += answers[j];
                        j += answers[j];
                    }
                }
            }
        }

        return answers;
    }
}