using System;

public class SpiralMatrix
{
    enum Direction { Right, Down, Left, Up };

    public static int[,] GetMatrix(int size)
    {
        int colIndex = 0;
        int rowIndex = 0;
        int number = 0;
        int requiredIterations = size * 2 - 1;
        int[,] spiral = new int[size, size];
        Direction direction = Direction.Right;

        for (int i = 0; i < requiredIterations; i++)
        {
            for (int j = 0; j < (requiredIterations - i + 1) / 2; j++)
            {
                if (direction == Direction.Right)
                {
                    if (colIndex < size - 1)
                    {
                        colIndex++;
                    }
                }

                if (direction == Direction.Down)
                {
                    if (rowIndex < size - 1)
                    {
                        rowIndex++;
                    }
                }

                if (direction == Direction.Left)
                {
                        colIndex--;
                }

                if (direction == Direction.Up)
                {
                    rowIndex--;
                }
                
                if (spiral[rowIndex, colIndex] == 0)
                {
                    number++;
                    spiral[rowIndex, colIndex] = number;
                }
            }
            switch (direction)
            {
                case Direction.Right:
                {
                    direction = Direction.Down;
                    continue;
                }
                case Direction.Down:
                {
                    direction = Direction.Left;
                    continue;
                }
                case Direction.Left:
                {
                    direction = Direction.Up;
                    continue;
                }
                case Direction.Up:
                {
                    direction = Direction.Right;
                    continue;
                }
            }

        }

        // Each value in the spiral is off by one.
        // Iterate through each value and increment by one.
        if (size > 1)
        {
            for (int col = 0; col < spiral.GetLength(0); col++)
                for (int row = 0; row < spiral.GetLength(1); row++)
                    spiral[col, row] += 1;

        }
        
        return spiral;
    }
}
