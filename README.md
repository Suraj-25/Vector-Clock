# Vector-Clock
Implementation of Vector Clock in distributed system using Golang


Vector Clock is an algorithm which generally creates the partial ordering of the events and used to detects the causality violations in a distributed system. These clocks generally expand on Scalar time to facilitate a casually consistent view of the distributed system. They detect whether a contributed event has caused another event in the distributed system. It essentially captures all the casual relationships of that system. This algorithm helps us to mark every process with a vector for each local clock of every process within the system.

Here we are implementing the Lamport Logical Clock Synchronization (Vector Clock) which shows us the timestamp for every process which took place in the system. The complete step for implementation of vector clock is given below: -

1. First of all, we need to create a VectorC.go file.
To create the VectorC.go file we first define the Timestamp type and four processes in it.

We also defined the increment or changes in the timestamp as the message sends and received by the different processes.

We defined the sleep time after which each process will respond with their respective timestamps.

We also defined several go routines which will run simultaneously for sending and receiving the message for the occurrence of the event.

The VectorClock.txt file is there from which the program will read and gives the value of the timestamps.

2. After saving this code in file VectorC.go file, place the VectorClock.txt file in the same folder and then run the code using the command go run VectorC.go

