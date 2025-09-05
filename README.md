# TaskFlow CLI ⚡

A lightning-fast, zero-dependency task management CLI tool that keeps your productivity flowing. Built with Go for speed and simplicity.

## Description

TaskFlow CLI is a minimalist command-line task tracker that stores your tasks locally in JSON format. No accounts, no subscriptions, no cloud dependencies—just pure, efficient task management right from your terminal. Whether you're managing daily todos, tracking project milestones, or organizing your workflow, TaskFlow keeps it simple and fast.

## Why? 🤔

**The Problem:** Most task management tools are bloated with features you don't need, require internet connections, or lock you into proprietary ecosystems. Developers and power users need something that integrates seamlessly with their terminal workflow.

**The Solution:** TaskFlow CLI provides a clean, fast, and reliable way to manage tasks without leaving your terminal. It's built with the philosophy that the best tools are the ones that get out of your way and let you focus on what matters—getting things done.

**Goals:**
- ⚡ **Speed**: Lightning-fast operations with zero startup time
- 🔒 **Privacy**: Your tasks stay on your machine, always
- 🎯 **Simplicity**: Intuitive commands that feel natural
- 🛠️ **Developer-friendly**: Perfect for terminal-based workflows

## Quick Start

### Prerequisites
- Go 1.24.5 or higher

### Installation

1. **Clone the repository:**
   ```bash
   git clone https://github.com/dmitriy-zverev/Task-Tracker-CLI.git
   cd Task-Tracker-CLI
   ```

2. **Build the application:**
   ```bash
   go build -o taskflow main.go
   ```

3. **Add to your PATH (optional but recommended):**
   ```bash
   # For macOS/Linux
   sudo mv taskflow /usr/local/bin/
   
   # Or add to your shell profile
   export PATH=$PATH:$(pwd)
   ```

4. **Start managing tasks:**
   ```bash
   taskflow add "Build something awesome"
   ```

## Usage

### Core Commands

**Add a new task:**
```bash
taskflow add "Buy groceries"
# Output: Task added successfully (ID: 1)
```

**Update an existing task:**
```bash
taskflow update 1 "Buy groceries and cook dinner"
```

**Delete a task:**
```bash
taskflow delete 1
```

### Task Status Management

**Mark task as in progress:**
```bash
taskflow mark-in-progress 1
```

**Mark task as completed:**
```bash
taskflow mark-done 1
```

### Viewing Tasks

**List all tasks:**
```bash
taskflow list
```

**Filter by status:**
```bash
taskflow list done          # Show completed tasks
taskflow list todo          # Show pending tasks
taskflow list in-progress   # Show tasks in progress
```

### Task Structure

Each task contains:
- **id**: Unique identifier
- **description**: Task description
- **status**: `todo`, `in-progress`, or `done`
- **createdAt**: Creation timestamp
- **updatedAt**: Last modification timestamp

### Example Workflow

```bash
# Start your day
taskflow add "Review pull requests"
taskflow add "Update documentation"
taskflow add "Deploy to staging"

# Begin working
taskflow mark-in-progress 1

# Complete tasks
taskflow mark-done 1
taskflow mark-done 2

# Check your progress
taskflow list done
```

## Contributing

We welcome contributions! Here's how you can help make TaskFlow even better:

### Getting Started

1. **Fork the repository**
2. **Create a feature branch:**
   ```bash
   git checkout -b feature/amazing-feature
   ```
3. **Make your changes and test thoroughly**
4. **Commit with clear messages:**
   ```bash
   git commit -m "Add amazing feature that does X"
   ```
5. **Push to your fork and submit a pull request**

### Development Guidelines

- **Code Style**: Follow Go conventions and run `go fmt`
- **Testing**: Add tests for new features
- **Documentation**: Update README for any new commands or features
- **Dependencies**: Keep it dependency-free—that's our superpower!

### Ideas for Contributions

- 🎨 Add color output for better readability
- 📅 Implement due dates and reminders
- 🔍 Add search functionality
- 📊 Create task statistics and reports
- 🔄 Add task templates and recurring tasks
- 📱 Build shell completions for better UX

### Bug Reports

Found a bug? Please open an issue with:
- Your operating system
- Go version
- Steps to reproduce
- Expected vs actual behavior

---

**Built with ❤️ and Go** | **No dependencies, maximum freedom**
