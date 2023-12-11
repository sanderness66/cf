# cf(1) - count files in given directories and subdirectories

Kozmix Go, 05-JUL-2022

```
cf [-a] [-s] [-c] [-S] [dir ...]
```


<a name="description"></a>

# Description

**cf**
counts the number of files in the given directories (or the current
directory if no argument is given), and all its subdirectories. Bit like
**du**(1)
but for file counts, rather than file sizes.
**cf**
only counts regular files, not directories, symbolic links, device
files or named pipes.

**cf**
has these options:


* **-a**  
  do not ignore entries starting with .
  
* **-s**  
  print only a total file count for each argument
  
* **-c**  
  print a grand total
  
* **-S**  
  don't add subdirectory counts to directory count
  

<a name="bugs"></a>

# Bugs


**cf -S -c**
doesn't work as expected.


<a name="see-also"></a>

# See Also

**du**(1),
**ls**(1)


<a name="author"></a>

# Author

svm
