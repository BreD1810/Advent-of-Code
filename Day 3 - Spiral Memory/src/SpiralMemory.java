import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;

public class SpiralMemory
{

    private int[][] memory;

    private void createMemory(int maxValue)
    {
        //Calculate the number of circles needed around the center
        int circleCount = 0;
        int largestValue = 1;
        int previousLargest = 1;
        while(largestValue < maxValue)
        {
            circleCount++;
            largestValue = previousLargest + 4 +(4*((2*circleCount)-1));
            previousLargest = largestValue;
        }

        System.out.println("Number of circles needed: " + circleCount);

        //Calculate the middle of the spiral
        int middle = circleCount;

        //Add the center value
        memory = new int[(circleCount*2)+1][(circleCount*2)+1];
        memory[middle][middle] = 1;

        //Add in the values to the spiral
        circleCount = 0;
        largestValue = 1;
        previousLargest = 1;
        while(largestValue < maxValue)
        {
            circleCount++;

            int sideLength = (2*circleCount) + 1;

            largestValue = previousLargest + 4 +(4*((2*circleCount)-1));
            previousLargest = largestValue;
            int tempLargest = largestValue;

            //Place the corner piece
            memory[middle+circleCount][middle+circleCount] = largestValue;
            //Add values to the bottom side
            for(int i = 1; i < sideLength; i++)
            {
                memory[middle+circleCount][middle+circleCount-i] = largestValue - i;
                tempLargest = largestValue - i;
            }
            tempLargest--;

            //Add values to the left side
            for(int i = 1; i < sideLength; i++)
            {
                memory[middle+circleCount-i][middle+circleCount-sideLength+1] = tempLargest;
                tempLargest--;
            }

            //Add values to the top side
            for(int i = 1; i < sideLength; i++)
            {
                memory[middle-circleCount][middle-circleCount+i] = tempLargest;
                tempLargest--;
            }

            //Add values for the right side
            for(int i = 1; i < sideLength-1; i++)
            {
                memory[middle-circleCount+i][middle+circleCount] = tempLargest;
                tempLargest--;
            }
        }
    }

    private void drawMemory()
    {
        //Draw the memory row by row
        for(int i = 0; i < memory.length; i++)
        {
            for(int j = 0; j < memory[0].length; j++)
            {
                //Print the number 3 digits wide (leading 0s) and space after
                System.out.print(String.format("%06d", memory[i][j]) + " ");
            }
            System.out.println();
        }
    }

    private int calculateLength(int value)
    {
        int iLocation = 0, jLocation = 0;
        //Find the location in the array
        for ( int i = 0; i < memory.length; ++i ) {
            for ( int j = 0; j < memory[0].length; ++j ) {
                if ( memory[i][j] == value) {
                    iLocation = i;
                    jLocation = j;
                }
            }
        }

        int middle = ((memory.length/2) + 1);
        int vertical = middle - iLocation - 1;
        int horizontal = middle - jLocation - 1;
        return Math.abs(vertical) + Math.abs(horizontal);
    }

    private void part2(int value)
    {
        //Set all values to 0
        for(int i = 0; i < memory.length; i++)
        {
            for(int j = 0; j < memory[0].length; j++)
            {
                memory[i][j] = 0;
            }
        }

        int middle = ((memory.length/2) + 1);
        int iLocation = middle;
        int jLocation = middle;
        int counter = 1;
        int currentValue = 1;
        memory[middle][middle] = 1;
        while(currentValue < value)
        {
            //Right side values
            iLocation = middle + counter;
            jLocation = middle;
            memory[iLocation][jLocation] = neighbourTotal(iLocation, jLocation);

            //Top values

            //Left side values

            //Bottom values
            counter ++;
        }
        System.out.println("First value larger than " + value + ": " + currentValue);
    }

    private int neighbourTotal(int iLocation, int jLocation)
    {
        int sum = 0;
        sum += memory[iLocation+1][jLocation];
        sum += memory[iLocation+1][jLocation+1];
        sum += memory[iLocation][jLocation+1];
        sum += memory[iLocation-1][jLocation+1];
        sum += memory[iLocation-1][jLocation];
        sum += memory[iLocation-1][jLocation-1];
        sum += memory[iLocation][jLocation-1];
        return sum;
    }

    public static void main(String[] args) throws IOException
    {
        SpiralMemory myProgram = new SpiralMemory();

        //Get the input from the user.
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
        System.out.println("What is the value that you want to get?");
        int value = Integer.valueOf(br.readLine());

        if(value == 1)
        {
            //Create the memory
            myProgram.memory = new int[1][1];
            myProgram.memory[0][0] = 1;
            //Draw the memory in the console
            myProgram.drawMemory();

            //Find the shortest distance (Manhatten distance)
            System.out.println();
            System.out.println("Shortest Distance: 0");

        }
        else
        {
            //Create the memory
            myProgram.createMemory(value);
            //Draw the memory in the console
            myProgram.drawMemory();

            //Find the shortest distance (Manhatten distance)
            System.out.println();
            System.out.println("Shortest Distance: " + myProgram.calculateLength(value));
        }

        myProgram.part2(value);

    }

}
