# Linux

## Useful links

* [List of Unix Commands](https://en.wikipedia.org/wiki/List_of_Unix_commands)

## File types

| File type       | Cmd to crate       | `ls -l` |
| --------------- | ------------------ | ------- |
| Regular File    | touch              | `-`     |
| Directory       | mkdir              | `d`     |
| Block Files     | fdisk              | `b`     |
| Character Files | mknod              | `c`     |
| Pipe Files      | mkfifo             | `p`     |
| Symlink Files   | ln                 | `l`     |
| Socket Files    | `socket()` syscall | `s`     |

## chmod

| #   | Sum                | rwx | Permission              |
| --- | ------------------ | --- | ----------------------- |
| 7   | 4(r) + 2(w) + 1(x) | rwx | read, write and execute |
| 6   | 4(r) + 2(w)        | rw- | read and write          |
| 5   | 4(r)        + 1(x) | r-x | read and execute        |
| 4   | 4(r)               | r-- | read only               |
| 3   | 2(w) + 1(x)        | -wx | write and execute       |
| 2   | 2(w)               | -w- | write only              |
| 1   | 1(x)               | --x | execute only            |
| 0   | 0                  | --- | none                    |

## Jobs

Suspend the process while in the foreground

```
Ctrl Z
```

Start a job directly into background

```
name [args] &
```

Kill a job by number (replace `1` by the number of the job when listed using `jobs`)

```
kill %1
```

* `fg` - sends a job to the foreground
* `bg` - sends a job to the background
* `jobs` - list jobs
