# Lab5
This is the code for lab5, it demonstrates the dining philosophers problem and solution.
In order to prevent cyclical dependencies one "Philosopher" accesses a different "Fork" first. 
For example:

func getForks(index int, forks map[int]chan bool) {
	if index == 0 {
		forks[(index+1)%5] <- true
		forks[index] <- true
	} else {
		forks[index] <- true
		forks[(index+1)%5] <- true
	}
}

Philosopher with the index of 0 will access a different fork to all the others, preventing a deadlock.
