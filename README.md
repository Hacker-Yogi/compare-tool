# compare-tool

This is a command-line tool that performs various operations on text files, such as sorting lines alphabetically, printing only newly added lines, and printing only unique lines. The tool accepts multiple file paths as input and processes each file individually.

Installation:-
To use this tool, you need to have Go installed on your system. If Go is not installed, you can download and install it from the official Go website: https://golang.org/

Once you have Go installed, you can follow these steps to install and run the tool:

Clone the repository or download the source code.

Navigate to the project directory.

Open a terminal and run the following command to build the tool:

"go build compare-tool.go"


Usage:-
The tool supports the following command-line arguments:

-s or --sort: Sort lines alphabetically.
-na or --newly-added: Print only the newly added lines.
-r or --unique: Print only the unique lines.
-h or --help: Show available arguments.
If you run the tool without any arguments, it will display the result of processing the files according to the default behavior, which is to print all lines in the order they appear in the files.

Example:
./compare-tool -s -na file1.txt file2.txt

This command will sort the lines alphabetically, filter out any duplicates, and print only the newly added lines from the files file1.txt and file2.txt.


