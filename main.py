import sys


def main():
    program, command, arg1, arg2 = sys.argv
  
    if not command:
        print("Usage: task-cli <command> [argument 1] [argument 2]")
        sys.exit(1)

    if (command):
        print(f"You want to do: \"{command}\"")

    if (arg1):
        print(f"Given \"{arg1}\"")

    if (arg2):
        print(f"And given \"{arg2}\"")
        
if __name__ == "__main__":
    main()
