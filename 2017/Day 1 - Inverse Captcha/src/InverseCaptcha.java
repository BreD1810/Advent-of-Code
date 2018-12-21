import java.io.*;

public class InverseCaptcha
{
    private int[] input;

    private void getinput() throws IOException
    {
        //Get the user input from the commandline.
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
        System.out.println("Please enter your string of captcha numbers.");
        String userInput = reader.readLine();
        reader.close();

        //Convert all of the characters into integers.
        char[] tempArray = userInput.toCharArray();
        input = new int[tempArray.length];
        for(int i = 0; i < tempArray.length; i++)
        {
            input[i] = tempArray[i] - '0';
        }
    }

    private int calculateSumNeighbour()
    {
        int sum = 0;

        for(int i = 0; i < input.length; i++)
        {
            if(i == input.length - 1)
            {
                if(input[i] == input[0])
                {
                    sum += input[i];
                }
            }
            else
            {
                if(input[i] == input[i+1])
                {
                    sum += input[i];
                }
            }
        }

        return sum;
    }

    private int calculateSumHalfway()
    {
        int sum = 0;
        int halfway = input.length/2;

        for(int i = 0; i < input.length; i++)
        {
            if(i + halfway > input.length-1)
            {
                if(input[i] == input[i-halfway])
                {
                    sum += input[i];
                }
            }
            else
            {
                if(input[i] == input[i+halfway])
                {
                    sum += input[i];
                }
            }
        }

        return sum;
    }

    public static void main(String[] args) throws IOException
    {
        InverseCaptcha myProgram = new InverseCaptcha();
        myProgram.getinput();
        System.out.println("Part 1: " + myProgram.calculateSumNeighbour()); //Correct answer = 1158 for the challenge input
        System.out.println("Part 2: " + myProgram.calculateSumHalfway()); //Correct answer = 1132 for the challenge input
    }

}
