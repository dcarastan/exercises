## Aporeto Quiz

### Problem2: Python/Go uniquify script

```
uniquify [--help|-h]
uniquify --file=<filename> --output=<output-filename> [—verbose]
```
big_file.txt is input file example

uniquified_file.txt is output generated

Content is uniquified without being sorted.

**Console output**

```
Mini:problem2 doru$ ls
README.md            big_file.txt         out.txt              uniquified_file.txt  uniquify*
Mini:problem2 doru$ rm out.txt
Mini:problem2 doru$ ./uniquify --file big_file.txt --output out.txt -verbose
2016-09-27 23:55:55,241 INFO     Reading from big_file.txt
2016-09-27 23:55:55,241 INFO     Output set to out.txt
Mini:problem2 doru$ diff out.txt uniquified_file.txt
Mini:problem2 doru$
```

```
Mini:problem2 doru$ pylint uniquify | tail -3
No config file found, using default configuration
-----------------
Your code has been rated at 10.00/10 (previous run: 10.00/10, +0.00)
```
