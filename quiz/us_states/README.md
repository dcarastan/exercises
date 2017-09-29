
## Aporeto Quiz

### Problem1: Bash Shell Script

**Script console output**

```
Mini:problem1 doru$ ./us_states -h
./us_states -- Creates a US states file

Author: Doru Carastan (doru@rocketmail.com)

Usage:  ./us_states [--help|-h]
        ./us_states --create-file=<filename> [--no-prompt] [--verbose]
```
```
Mini:problem1 doru$ ./us_states --help
./us_states -- Creates a US states file

Author: Doru Carastan (doru@rocketmail.com)

Usage:  ./us_states [--help|-h]
        ./us_states --create-file=<filename> [--no-prompt] [--verbose]
```
```
Mini:problem1 doru$ ./us_states --foo
Unknown option '--foo'
./us_states -- Creates a US states file

Author: Doru Carastan (doru@rocketmail.com)

Usage:  ./us_states [--help|-h]
        ./us_states --create-file=<filename> [--no-prompt] [--verbose]
```
```
Mini:problem1 doru$ ./us_states
ERROR: Specify the output file using the '--create-file' option
./us_states -- Creates a US states file

Author: Doru Carastan (doru@rocketmail.com)

Usage:  ./us_states [--help|-h]
        ./us_states --create-file=<filename> [--no-prompt] [--verbose]
```
```
Mini:problem1 doru$ ./us_states --create-file=50states
```
```
Mini:problem1 doru$ head -5 50states; echo '...';tail -5 50states
Alabama
Alaska
Arizona
Arkansas
California
...
Virginia
Washington
West Virginia
Wisconsin
Wyoming
```
```
Mini:problem1 doru$ ./us_states --create-file=50states
File exists. Overwrite (y/n) ?
File exists. Overwrite (y/n) ? n
File exists. Overwrite (y/n) ? y
```
```
Mini:problem1 doru$ head -5 50states; echo '...';tail -5 50states
Alabama
Alaska
Arizona
Arkansas
California
...
Virginia
Washington
West Virginia
Wisconsin
Wyoming
```
```
Mini:problem1 doru$ ./us_states --create-file=50states --verbose
File already exists
File exists. Overwrite (y/n) ? y
File removed
File created
```
```
Mini:problem1 doru$ ./us_states --create-file=50states --verbose --no-prompt
File already exists
File removed
File created
```
```
Mini:problem1 doru$ chmod 444 50states
```
```
Mini:problem1 doru$ ./us_states --create-file=50states --verbose
File already exists
File exists. Overwrite (y/n) ? y
override r--r--r--  doru/staff for 50states? n
ERROR: Failed to remove file: 50states
```
```
Mini:problem1 doru$ ./us_states --create-file=50states --verbose
File already exists
File exists. Overwrite (y/n) ? y
override r--r--r--  doru/staff for 50states? y
File removed
File created
```
```
Mini:problem1 doru$ ls -l 50states
-rw-r--r--  1 doru  staff   472B Sep 27 16:26 50states
```
