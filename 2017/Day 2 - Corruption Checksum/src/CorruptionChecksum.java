import java.io.BufferedReader;
import java.io.FileReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class CorruptionChecksum
{

    int[][] spreadsheet;

    private int[][] getInput() throws IOException
    {
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
        BufferedReader fileReader = new BufferedReader(new FileReader("src\\Input.txt"));

        //Get the column count
        System.out.println("How many columns does your spreadsheet have?");
        String columns = reader.readLine();

        //Get the row count
        System.out.println("How many rows does your spreadsheet have?");
        String rows = reader.readLine();

        //Create the 2d array to represent the spreadsheet
        int[][] input = new int[Integer.valueOf(rows)][Integer.valueOf(columns)];

        //Fill in the spreadsheet
        for(int i = 0; i < Integer.valueOf(rows); i++)
        {
            //Split the line on any whitespace character
            String[] row = fileReader.readLine().split("\\s");
            for (int j = 0; j < Integer.valueOf(columns); j++)
            {
                if(j < row.length)
                {
                    input[i][j] = Integer.valueOf(row[j]);
                }
                else
                {
                    input[i][j] = -1;
                }
            }
        }

        reader.close();
        return input;
    }

    private float calculateChecksum(int[][] spreadsheet)
    {
        float total = 0;

        for(int i = 0; i < spreadsheet.length; i++)
        {
            //Set the first value of the row as the smallest and largest value
            int smallest = spreadsheet[i][0];
            int largest = spreadsheet[i][0];
            //Compare the rest of the values to this
            for(int j = 1; j < spreadsheet[i].length; j++)
            {
                if(spreadsheet[i][j] == -1)
                {
                    //This part is ignored as it is empty
                }
                else if(spreadsheet[i][j] < smallest)
                {
                    smallest = spreadsheet[i][j];
                }
                else if(spreadsheet[i][j] > largest)
                {
                    largest = spreadsheet[i][j];
                }
            }
            //Calculate the different for the row, and add it to the total.
            total += (largest-smallest);
        }

        return total;
    }

    private float calculateChecksumDivisible(int[][] spreadsheet)
    {
        float total = 0;

        for(int i = 0; i < spreadsheet.length; i++)
        {
           //Loop through the spreadsheet
            for(int j = 0; j < spreadsheet[i].length; j++)
            {
                if(spreadsheet[i][j] == -1)
                {
                    //This part is ignored as it is empty
                }
                else
                {
                    for(int k = 0; k < spreadsheet[i].length; k++)
                    {
                        if(spreadsheet[i][k] == -1)
                        {
                            //This part is ignored as it is empty
                        }
                        else
                        {
                            if((spreadsheet[i][j] % spreadsheet[i][k] == 0) && j != k)
                            {
                                //If they're divisible, add to the total and break out of the loop.
                                total += (spreadsheet[i][j] / spreadsheet[i][k]);
                                break;
                            }
                        }
                    }
                }

            }
       }

        return total;
    }

    public static void main(String[] args) throws IOException
    {
        CorruptionChecksum myProgram = new CorruptionChecksum();
        myProgram.spreadsheet = myProgram.getInput();
        System.out.println("Part 1: " + myProgram.calculateChecksum(myProgram.spreadsheet));
        System.out.println("Part 2: " + myProgram.calculateChecksumDivisible(myProgram.spreadsheet));
    }

}
