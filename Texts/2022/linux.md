# Linux

## File types

| File type | Cmd to crate | `ls -l` |
| --------- | ------------ | ------------ |
| Regular File | touch | `-` |
| Directory | mkdir | `d` |
| Block Files | fdisk | `b` |
| Character Files | mknod | `c` |
| Pipe Files | mkfifo | `p` |
| Symlink Files | ln | `l` |
| Socket Files | `socket()` syscall | `s` |

## chmod

| # | Sum | rwx | Permission |
| - | --- | --- | ---------- |
| 7 | 4(r) + 2(w) + 1(x)|rwx|read, write and execute |
| 6 | 4(r) + 2(w)|rw-|read and write |
| 5 | 4(r)        + 1(x)|r-x|read and execute |
| 4 | 4(r)|r--|read only |
| 3 |        2(w) + 1(x)|-wx|write and execute |
| 2 |        2(w)|-w-|write only |
| 1 |               1(x)|--x|execute only |
| 0 | 0|---|none |
