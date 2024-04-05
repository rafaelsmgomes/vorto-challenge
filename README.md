# Vorto Challenge

## Build instructions

Nothing fancy, just run the following command from the root directory.

```sh
go build -o vrp ./src/*.go
```

## Running instructions

Following the evaluateReadME.txt, just build the project and run the following from the root directory:

```sh
python3 evaluatShared.py --cmd "./vrp" --problemDir training
```

## Notes

This solution can definitely be improved upon.

Right now it is a rather naive solution in which each driver goes to the Nearest Neighbour/Deliver, makes it and judges from there wether to go to the another delivery or to go back to origin (respective the 12 hours constraint).

### Algorithm improvements

- **Nearest Neighbour**: The current solution uses a nearest neighbour algorithm to determine the next delivery to make. This is a rather naive approach and can be improved upon. A better approach would be to use a more advanced algorithm like the Branch and Bound algorithm.

- **Driver assignment**: The current solution assigns drivers to deliveries in a rather naive way. It assigns the driver to the closest delivery. A better approach would be to assign drivers to deliveries in a way that minimizes the total time taken to deliver all the packages.

### Code improvements

- **Multithreading**: I need to think more about this, but there may be a way to assign drivers simultaneously, although right now they would still have to share the same state for knowing which deliveries are already made.

- **Using the Queue in a meaningful way**: There is really no point in the use of the queue right now. I though this would be usefull to implement the Branch and Bound algorithm, but never got there.
