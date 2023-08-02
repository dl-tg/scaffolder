# Scaffolder

Scaffolder is a powerful command-line interface tool written in Golang designed to automate the tedious task of creating barebones for your projects. It allows you to define the necessary directory structure in a reusable config YAML file, making it easy for both humans and parsers to work with.
### TODO
- [ ] Specify default variables in YAML like `$var: 1`
## Features

  - Automate project scaffolding with a YAML config file
  - Reusable templates for various programming languages available in the Scaffold configs repository
  - Supports creating folders and files with custom content using simple YAML syntax
  - Easily initialize a Git repository with the --git flag
  - Define and use variables in the YAML file for flexible project generation
  - Remembers custom config directories using the --remember flag for added convenience

### Quick Example

Suppose we have the following YAML config file named test.yaml:

```yaml

hello1: # Folder name
  hello.txt: | # File name
    Hello World # File content
  hello1.txt: # File name (empty)
  hello2.txt: # File name (empty)
hello2: # Folder name
  hello22.txt: # File name (empty)
```
We can scaffold a project named "hello" using this config:

```bash

$ scaffold --name hello --yaml test
```
The resulting directory structure will be:

```
hello
├── hello1
│   ├── hello.txt
│   ├── hello1.txt
│   └── hello2.txt
└── hello2
    └── hello22.txt
```
The content of hello.txt will be:
```
Hello World
```

## Installation
### Linux / MacOS

1. Change your directory to the home directory:

```bash

$ cd ~
```
2. Create a folder for Scaffolder and navigate there:

```bash

$ mkdir scaffolder
$ cd scaffolder
```
3. Download the latest release:

```bash

# Linux:
$ wget -q -O scaffold https://github.com/cemister/scaffolder/releases/download/stable-v1_1_8/scaffolder_linux
# MacOS:
$ curl -s -o scaffold https://github.com/cemister/scaffolder/releases/download/stable-v1_1_8/scaffolder_macos
```
4. Make the file executable:

```bash

$ chmod +x scaffold
```
5. Add the executable to your PATH:

```bash

$ echo 'export PATH="$HOME/scaffolder:$PATH"' >> ~/.bashrc
$ source ~/.bashrc
```
### Windows

1. Open Command Prompt (cmd) or PowerShell and change the directory to the user's home directory:

```bash

# cmd:
cd %userprofile%

# PowerShell:
cd $env:userprofile
```
2. Create a folder named "scaffolder" and change the current directory to it:

```bash

mkdir scaffolder
cd scaffolder
```
3. Download the latest Scaffolder Windows release (replace stable-v1_1_8 with the tag of the latest version if not updated):

```bash

curl -s -o scaffold.exe https://github.com/cemister/scaffolder/releases/download/stable-v1_1_8/scaffolder_win.exe
```
Add the executable to the PATH in Windows:

  1. Press Win + S to open the Windows search bar
  2. Search for "Environment Variables" and select "Edit the system environment variables"
  3. In the System Properties window, click the "Environment Variables" button
  4. In the Environment Variables window, find the "Path" variable under "User variables" and click "Edit"
  5. Click "New," then enter the full path to the "scaffolder" folder (e.g., C:\Users\YourUsername\scaffolder) and click "OK" to add it to the PATH

Note: Make sure to replace YourUsername with your actual Windows username.
## Building from Source

To build Scaffolder from source, ensure you have Golang installed:

1. Clone the repository (Git should be installed):
```bash
$ git clone https://github.com/cemister/scaffolder.git
```

2. Navigate to the project directory:

```bash
$ cd scaffolder
```
3. Build the project:

```bash
$ go build
```

## Usage

Use the following command to scaffold a project:

```bash
$ scaffold --name <project_name> --yaml <config_name> --configdir? <path_to_custom_config_folder_if_exists> --git? <true/false> --remember? <true/false> --variables? <k:v>
```
Note: remember flag specifies whether to remember custom path specified in configdir, avoiding the need to specify it each time
### YAML Config Syntax

To create a file inside the parent (project's) directory, use a "." collection:

```yaml

.:
  main.txt:
```
To create a subdirectory (folder inside a folder), create a new collection with its name as the name of the parent folder, followed by the name of the needed folder separated by a slash (/):

```yaml

hello1:
  hello.txt:
  hello1.txt:
  hello2.txt:
hello2:
  hello22.txt:
hello2/hello3:
  hello33.txt:
```
To create an empty folder, create an empty collection without values:

```yaml

helloempty:
```
### Variables
Variables were added in 1.1.7. Here's how to use them:

Assume we have the following YAML config named cpp.yaml:

```yaml

src:
  main.cpp: |
    #include <iostream>

    int main() {
        std::cout << "Hello, {name}!" << std::endl;
        return 0;
    }
```
In this example, we've defined a variable {name} within the main.cpp file, which will be replaced with the actual value when we scaffold the project.

Now, let's scaffold a C++ project named "hello-cpp" using this config and provide a value for the {name} variable:

```bash

$ scaffold --name hello-cpp --yaml cpp --variables name:John
```
The resulting directory structure will be:

```
hello-cpp
└── src
    └── main.cpp
```
The content of main.cpp will be:

```cpp

#include <iostream>

int main() {
    std::cout << "Hello, John!" << std::endl;
    return 0;
}
```

As you can see, the {name} variable was replaced with "John" in the final project files. This allows you to customize the generated code or any other content based on the values you provide during scaffolding.

As of 1.1.8, support for variables in folder or file names was added. Just wrap it in double quotes like `"{var}": ...` if using in folder/filenames.

## Contributing
If you want to contribute but are unsure how, refer to the official GitHub guide on Contributing to projects.

## License

This project is licensed under the MIT license. See the [LICENSE](LICENSE) file for details.
