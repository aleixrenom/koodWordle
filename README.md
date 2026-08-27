# kood / Wordle instructions

## The Situation

In this project you are tasked with creating a command-line game in Go that mimics the popular [Wordle](https://www.nytimes.com/games/wordle/index.html). The user will guess a hidden word, and your program will provide feedback on their guesses.

## Functional Requirements

### Game Rules

The player has to **guess a 5-letter** word and can make up to **six guesses** to try to identify the secret word. After each guess, the game **provides feedback** indicating which letters are correct:

- **Green** for letters in the correct position.
- **Yellow** for correct letters in the wrong position.
- **White** for incorrect letters.

If the player fails to guess the word after **six attempts**, the game should reveal the **secret word**.

### User Interaction

The game must be launched from the command line with a **command-line argument** that signifies the **index** of the word in the provided word list. If the command-line argument is **missing** or **invalid**, the program should handle it **gracefully without crashing**. So the execution command should look like this:

go run . 10


After each guess, the game **provides feedback** and shows the **remaining number of attempts** and a **sorted list of remaining letters** (A-Z) that have not been guessed incorrectly, with letters in the feedback **capitalized**. You have to use **ANSI escape codes** in the format of `"\u001B[32m"` (this example is ANSI green) to color the output.

After each game, the player is presented with an option to **view their stats**.

### Data Management

The word list must be read from a file named `wordle-words.txt` at **project root**. A list with words can be found on the internet, for example [here](https://gist.github.com/dracos/dd0668f281e685bad51479e5acaadb93). **Do not push** the word list file to the repository - it must be included in the `.gitignore` file.

The stats are written to a CSV file named `stats.csv` and include a row of **comma-separated values** for each game in the following order:

- `username`
- `secret word`
- `number of attempts`
- `"win" or "loss"`

The stats file must **persist over games**. **Do not push** the stats file to the repository - it must be included in the `.gitignore` file.

If the word list file is missing, the program should handle it **without crashing**.

### Technical Implementation

`main.go` needs to be located at the **root** of your `koodWordle` repository. This is the entry point of the application, handling command-line arguments and starting the game. Your **Go module name** should match your repository name: `koodWordle`.

You must use `bufio.Scanner` to read from standard input (`stdin`). Handle **EOF** (Ctrl+D) **gracefully** by breaking out of input loops when `scanner.Scan()` returns false.

**NB! All the specific requirements and messages are not provided in the description. You have to discover them from the provided video and automated test feedback.**

### Key Output Formats

Note that the standard output (`stdout`) of your code needs to match the expected `stdout`. To help you out a bit, here is a list with some of the required output formats. You'll need to figure out any missing formats and where to use them by watching the video and reading the test results.

- **Start screen**:

Enter your username:


- **After login**:

Welcome to Wordle! Guess the 5-letter word.
Enter your guess:


- **Valid guess screen**:

Feedback: 
Remaining letters: 
Attempts remaining:  
Enter your guess:


- **After game**:

Do you want to see your stats? (yes/no):


- **After stats question - if `yes` is entered**:

Stats for :
Games played: 
Games won: 
Average attempts per game: 
Press Enter to exit...


### Possible Package structure

An example package structure for the project. This is not compulsory, but it is good to think through and have a clear structure for your projects.

koodWordle
├── main.go           // Entry point, processes arguments, starts game
├── game/             // Game logic and mechanics
│   └── game.go       // Core game functionality and feedback generation
├── io/               // Input/output operations
│   └── io.go         // File handling for words and statistics
└── model/            // Data structures
└── user.go       // User entity and statistics tracking  


## Video Example

Video: kood Wordle demo

## Bonus Functionality

If you were inspired by this task, you may implement the following additional features after submitting the initial version:

- **Instant replay**: Allow players to restart the next game without relaunching the app.
- **Random word selection**: Instead of reading the word number from the command line, randomly choose it each time.
- **Multi-Length Word Support**: Allow the player to choose the word length at the start of the game (e.g., 5-letter, 6-letter, or 7-letter words).

**NB! Do not submit these functionalities for automated tests.**

## Useful Links

- [Command-line arguments in Go](https://gobyexample.com/command-line-arguments)
- [Reading user input](https://gosamples.dev/read-user-input/)
- [File operations](https://pkg.go.dev/os)
- [CSV](https://en.wikipedia.org/wiki/Comma-separated_values)
- [ANSI escape codes](https://en.wikipedia.org/wiki/ANSI_escape_code)
Actually, give it to me as a downloadable markdown file insteadd.

go run . 10


After each guess, the game **provides feedback** and shows the **remaining number of attempts** and a **sorted list of remaining letters** (A-Z) that have not been guessed incorrectly, with letters in the feedback **capitalized**. You have to use **ANSI escape codes** in the format of `"\\u001B[32m"` (this example is ANSI green) to color the output.

After each game, the player is presented with an option to **view their stats**.

### Data Management

The word list must be read from a file named `wordle-words.txt` at **project root**. A list with words can be found on the internet, for example [here](https://gist.github.com/dracos/dd0668f281e685bad51479e5acaadb93). **Do not push** the word list file to the repository - it must be included in the `.gitignore` file.

The stats are written to a CSV file named `stats.csv` and include a row of **comma-separated values** for each game in the following order:

- `username`
- `secret word`
- `number of attempts`
- `"win" or "loss"`

The stats file must **persist over games**. **Do not push** the stats file to the repository - it must be included in the `.gitignore` file.

If the word list file is missing, the program should handle it **without crashing**.

### Technical Implementation

`main.go` needs to be located at the **root** of your `koodWordle` repository. This is the entry point of the application, handling command-line arguments and starting the game. Your **Go module name** should match your repository name: `koodWordle`.

You must use `bufio.Scanner` to read from standard input (`stdin`). Handle **EOF** (Ctrl+D) **gracefully** by breaking out of input loops when `scanner.Scan()` returns false.

**NB! All the specific requirements and messages are not provided in the description. You have to discover them from the provided video and automated test feedback.**

### Key Output Formats

Note that the standard output (`stdout`) of your code needs to match the expected `stdout`. To help you out a bit, here is a list with some of the required output formats. You'll need to figure out any missing formats and where to use them by watching the video and reading the test results.

- **Start screen**:

Enter your username:


- **After login**:

Welcome to Wordle! Guess the 5-letter word.
Enter your guess:


- **Valid guess screen**:

Feedback: 
Remaining letters: 
Attempts remaining:  
Enter your guess:


- **After game**:

Do you want to see your stats? (yes/no):


- **After stats question - if `yes` is entered**:

Stats for :
Games played: 
Games won: 
Average attempts per game: 
Press Enter to exit...


### Possible Package structure

An example package structure for the project. This is not compulsory, but it is good to think through and have a clear structure for your projects.

koodWordle
├── main.go           // Entry point, processes arguments, starts game
├── game/             // Game logic and mechanics
│   └── game.go       // Core game functionality and feedback generation
├── io/               // Input/output operations
│   └── io.go         // File handling for words and statistics
└── model/            // Data structures
└── user.go       // User entity and statistics tracking


## Video Example

Video: kood Wordle demo

## Bonus Functionality

If you were inspired by this task, you may implement the following additional features after submitting the initial version:

- **Instant replay**: Allow players to restart the next game without relaunching the app.
- **Random word selection**: Instead of reading the word number from the command line, randomly choose it each time.
- **Multi-Length Word Support**: Allow the player to choose the word length at the start of the game (e.g., 5-letter, 6-letter, or 7-letter words).

**NB! Do not submit these functionalities for automated tests.**

## Useful Links

- [Command-line arguments in Go](https://gobyexample.com/command-line-arguments)
- [Reading user input](https://gosamples.dev/read-user-input/)
- [File operations](https://pkg.go.dev/os)
- [CSV](https://en.wikipedia.org/wiki/Comma-separated_values)
- [ANSI escape codes](https://en.wikipedia.org/wiki/ANSI_escape_code)
"""

filename = "kood_wordle_instructions.md"
with open(filename, "w", encoding="utf-8") as f:
    f.write(markdown_content)

print(f"File created: {filename}")

Your Markdown file is ready

MD icon
kood_wordle_instructions
 MD 
Open



Gemini is AI and can make mistakes, including about people. Your privacy and GeminiOpens in a new window

Analysing
# kood / Wordle

## The Situation

In this project you are tasked with creating a command-line game in Go that mimics the popular [Wordle](https://www.nytimes.com/games/wordle/index.html). The user will guess a hidden word, and your program will provide feedback on their guesses.

## Functional Requirements

### Game Rules

The player has to **guess a 5-letter** word and can make up to **six guesses** to try to identify the secret word. After each guess, the game **provides feedback** indicating which letters are correct:

- **Green** for letters in the correct position.
- **Yellow** for correct letters in the wrong position.
- **White** for incorrect letters.

If the player fails to guess the word after **six attempts**, the game should reveal the **secret word**.

### User Interaction

The game must be launched from the command line with a **command-line argument** that signifies the **index** of the word in the provided word list. If the command-line argument is **missing** or **invalid**, the program should handle it **gracefully without crashing**. So the execution command should look like this:

```
go run . 10
```

After each guess, the game **provides feedback** and shows the **remaining number of attempts** and a **sorted list of remaining letters** (A-Z) that have not been guessed incorrectly, with letters in the feedback **capitalized**. You have to use **ANSI escape codes** in the format of `"\u001B[32m"` (this example is ANSI green) to color the output.

After each game, the player is presented with an option to **view their stats**.

### Data Management

The word list must be read from a file named `wordle-words.txt` at **project root**. A list with words can be found on the internet, for example [here](https://gist.github.com/dracos/dd0668f281e685bad51479e5acaadb93). **Do not push** the word list file to the repository - it must be included in the `.gitignore` file.

The stats are written to a CSV file named `stats.csv` and include a row of **comma-separated values** for each game in the following order:

- `username`
- `secret word`
- `number of attempts`
- `"win" or "loss"`

The stats file must **persist over games**. **Do not push** the stats file to the repository - it must be included in the `.gitignore` file.

If the word list file is missing, the program should handle it **without crashing**.

### Technical Implementation

`main.go` needs to be located at the **root** of your `koodWordle` repository. This is the entry point of the application, handling command-line arguments and starting the game. Your **Go module name** should match your repository name: `koodWordle`.

You must use `bufio.Scanner` to read from standard input (`stdin`). Handle **EOF** (Ctrl+D) **gracefully** by breaking out of input loops when `scanner.Scan()` returns false.

**NB! All the specific requirements and messages are not provided in the description. You have to discover them from the provided video and automated test feedback.**

### Key Output Formats

Note that the standard output (`stdout`) of your code needs to match the expected `stdout`. To help you out a bit, here is a list with some of the required output formats. You'll need to figure out any missing formats and where to use them by watching the video and reading the test results.

- **Start screen**:

```
Enter your username:
```

- **After login**:

```
Welcome to Wordle! Guess the 5-letter word.
Enter your guess:
```

- **Valid guess screen**:

```
Feedback: <guess word in uppercase>
Remaining letters: <included and untested uppercase letters separated with whitespace> 
Attempts remaining:  <number>
Enter your guess:
```

- **After game**:

```
Do you want to see your stats? (yes/no):
```

- **After stats question - if `yes` is entered**:

```
Stats for <username>:
Games played: <number>
Games won: <number>
Average attempts per game: <float number>
Press Enter to exit...
```

### Possible Package structure

An example package structure for the project. This is not compulsory, but it is good to think through and have a clear structure for your projects.

```
koodWordle
├── main.go           // Entry point, processes arguments, starts game
├── game/             // Game logic and mechanics
│   └── game.go       // Core game functionality and feedback generation
├── io/               // Input/output operations
│   └── io.go         // File handling for words and statistics
└── model/            // Data structures
    └── user.go       // User entity and statistics tracking
```

## Video Example

Video: kood Wordle demo

## Bonus Functionality

If you were inspired by this task, you may implement the following additional features after submitting the initial version:

- **Instant replay**: Allow players to restart the next game without relaunching the app.
- **Random word selection**: Instead of reading the word number from the command line, randomly choose it each time.
- **Multi-Length Word Support**: Allow the player to choose the word length at the start of the game (e.g., 5-letter, 6-letter, or 7-letter words).

**NB! Do not submit these functionalities for automated tests.**

## Useful Links

- [Command-line arguments in Go](https://gobyexample.com/command-line-arguments)
- [Reading user input](https://gosamples.dev/read-user-input/)
- [File operations](https://pkg.go.dev/os)
- [CSV](https://en.wikipedia.org/wiki/Comma-separated_values)
- [ANSI escape codes](https://en.wikipedia.org/wiki/ANSI_escape_code)
kood_wordle_instructions.md
Displaying kood_wordle_instructions.md.
