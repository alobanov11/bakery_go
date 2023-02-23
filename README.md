## Bakery test task

Write a programm with using design patterns.

Given: 
	- The bakery bakes different types of pastries in t seconds.
	- Bake types are available in the bakes.json file
	- You can bake n different types of pastries at one time.
	- After the time has elapsed, we must wait for the last (their) pastries to finish cooking and show the result.

Example: bakery -t 4 -n 1 -f bakes.json
```
- Cooking Bread
- Cooking Bun
- Cooking Pie
- Done Pie
- Cooking Pie
...
```

Result: Prepared: Bread(1), Buns(2), Pies(4)