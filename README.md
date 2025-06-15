# LearnGo
The path of Learning the Go Programming Language

# Introduction to Go (Golang)

Go (also known as Golang) is an open-source, compiled, statically-typed programming language developed by Google. It was first released in November 2009 by Rob Pike, Ken Thompson, and Robert Griesemer.

Go is designed for simplicity, efficiency, and strong concurrency support, making it a great choice for scalable web apps, command-line tools, desktop apps, and even mobile applications.

## Why Choose Go?

- **Simple Syntax:** Easy to read, write, and maintain.
- **Built-in Concurrency:** Goroutines and channels make multithreading easy.
- **Compiled Language:** Produces native binaries for better performance.
- **Fast Compilation:** The Go compiler is designed for speed.
- **Static Linking:** All dependencies can be bundled into a single binary.
- **Powerful Tooling:**
  - `gofmt`: Formats Go source code automatically.
  - `vet`: Analyzes code and detects potential issues.
  - `staticcheck`: Enforces code style and quality.
- **Garbage Collection:** Automatic memory management simplifies development.
- **Simple Specification:** Easy to learn and fully documented.
- **Open Source:** Free to use and contribute to.

## Popular Projects Using Go

- **Kubernetes** (Google)
- **Docker**
- **Dropbox** (performance-critical parts migrated to Go)
- **Infoblox**

## Installing Go

Go can be installed on Mac, Windows, and Linux.

### MacOS
1. Download the installer: [https://go.dev/dl/](https://go.dev/dl/)
2. Double-click the installer and follow the prompts.
3. Go will be installed at `/usr/local/go`, and `/usr/local/go/bin` will be added to your PATH.

### Windows
1. Download the MSI installer: [https://go.dev/dl/](https://go.dev/dl/)
2. Double-click the installer and follow the prompts.
3. Go will be installed at `C:\Go`, and `C:\Go\bin` will be added to your PATH.

### Linux
1. Download the tarball: [https://go.dev/dl/](https://go.dev/dl/)
2. Extract to `/usr/local/`:
   ```bash
   sudo tar -C /usr/local -xzf go<version>.linux-amd64.tar.gz
   
Add /usr/local/go/bin to your PATH.

Verifying Installation
To check if Go is installed correctly, run:

```bash
go version
go version go1.19.2 linux/amd64
```
You are now ready to start writing Go programs! 🚀

# Learn Go - First Program 🚀

---

## 📂 Project Setup

### 1️⃣ Create Project Directory

```bash
mkdir ~/Documents/learngo/
cd ~/Documents/learngo/
```

2️⃣ Initialize Go Module
```bash
go mod init learngo
```
This creates a go.mod file:

```go
module learngo
go 1.21.0
```

✨ Write Your First Go Program
Create a file named main.go:
```main.go
package main

import "fmt"

func main() {
	fmt.Println("Hello World")
}
```

##🚀 Ways to Run the Program
#1️⃣ go run (Quick Testing)
```bash
go run main.go
```
Compiles and runs directly.

No binary file is created.

#2️⃣ go build (Build Local Binary)
```bash
go build
./learngo
```
Creates binary in current directory.

Run the binary manually.

#3️⃣ go install (Global Install)
```bash
export GOBIN=~/go/bin/
go install
~/go/bin/learngo
```

Installs the binary globally.

(Optional) Add to PATH:

```bash
export PATH=$PATH:~/go/bin
```
Then you can simply run:

```bash
learngo
```

#4️⃣ Go Playground (Online)
Use Go Playground to run Go code online without local installation.

##🔍 Internal Behavior
go run → Compiles to a temporary directory and runs.

go build → Compiles and creates binary in current directory.

go install → Compiles and installs binary to GOBIN directory.

To see where go run builds temporary files:

```bash
go run --work main.go
```

#📖 Code Explanation
package main → Entry package for standalone Go programs.

import "fmt" → Imports the fmt package for I/O operations.

func main() → Program execution starts here.

fmt.Println("Hello World") → Prints to console.
