using System;
using System.Collections.Generic;

public class CircularBuffer<T>
{
    private int writeIndex;
    private int readIndex;
    private int bufferSize;
    private List<T> circularBuffer;

    public CircularBuffer(int capacity)
    {
        circularBuffer = new List<T> ( new T[capacity] );
        bufferSize = capacity;
        writeIndex = 0;
        readIndex = 0;
    }

    private int normIndex(int pointer)
    {
        return pointer % bufferSize;
    }
    public T Read()
    {
        if (circularBuffer[normIndex(readIndex)].Equals(0))
        {
            throw new InvalidOperationException("Current index is empty");
        }
        T value = circularBuffer[normIndex(readIndex)];
        circularBuffer[normIndex(readIndex)] = default(T);
        readIndex++;
        return value;
    }

    public void Write(T value)
    {
        if (!circularBuffer[normIndex(writeIndex)].Equals(0))
        {
            throw new InvalidOperationException("Buffer is full");
        }
        circularBuffer[normIndex(writeIndex)] = value;
        writeIndex++;
    }

    public void Overwrite(T value)
    {
        if (!circularBuffer[normIndex(writeIndex)].Equals(0))
        {
            readIndex++;
        }
        circularBuffer[normIndex(writeIndex)] = value;
        writeIndex++;

    }

    public void Clear()
    {
        circularBuffer[normIndex(readIndex)] = default(T);
    }
}