# Simple Adventure Scripting Language Manual

## 1. Basic Variables

Variables are used to store and manipulate data. Ensure there are spaces between names, numbers, and signs. For example, `gold=gold+10` is incorrect. Use `gold = gold + 10` instead.

### Variable Types

The type of a variable is automatically determined based on the value assigned to it.

#### Defining Variables:

- **Numbers** - `SET gold = 0`
- **Strings** - `SET name = "Alex"`

#### Operations on Variables

- You can assign one variable to another - `SET newGold = gold`
- You can perform simple arithmetic operations like addition and subtraction - `SET gold = gold + 10`

#### Built-in Variables

- **PlayerResponse** - This variable stores the player's response when they are asked a question or given a decision.

## 2. Commands

### Asking Questions:

**ASK**: This asks the player a question and stores their answer in PlayerResponse.

- _Example:_ `ASK "What is your name?" > PlayerResponse`

### Showing Messages:

**SHOW**: This shows a message to the player. A message can also include a variable.

- _Example:_ `SHOW "Welcome to the game!"`
- _Example:_ `SHOW "Hello, {$PlayerResponse}!"`

### Going to different cards

**WHEN**: This goes to a different SCENE based on the player's choice.

- _Example:_ `WHEN "Left" > 3`

### Defining Cards

**SCENE**: This starts a new part of the game.

```
SCENE 1
// Game part here
END
```

### Ending the Game

**END**: This ends the game.

- _Example:_ `END`

## Example Game

```
ASK "What is your name?" > PlayerResponse
SHOW "Welcome, {$PlayerResponse}!"

SCENE 1
SHOW "You found a treasure chest!"
gold = 10
SHOW "You have {$gold} gold."
END SCENE

SCENE 2
DECIDE "Do you go left or right?" > PlayerResponse
WHEN "Left" > 3
WHEN "Right" > 4
END SCENE

SCENE 3
SHOW "You found more gold!"
gold = gold + 20
SHOW "You now have {$gold} gold."
END SCENE

SCENE 4
SHOW "You encountered a dragon!"
END SCENE

END
```
