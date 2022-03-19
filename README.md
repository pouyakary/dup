# dup
dup is a tiny and fast command line utility to find the duplicate files within a directory.

## Installing
Have yourself the Golang installed. On UNIX machines you can have the GNU Make and run the `make install`, on the windows, run a `go build dup` and put the `dup.exe` where you wish.

## Using

```
usage: dup [options] <directory>
  -h, help              displays this help message
  -q, quite    stops the software from displaying the results
  -r, remove   removes the duplicates from the directory
  -e, exact    compares the files exactly, without normalization
```

## What is Normalization?
When you run dup with the `-e` flag, it compares the byte information of the files exactly as they are. However, in the normal mode, it compares a more logical content. For example. In the JPEG files, it first removes the EXIF data, so that if you have two duplicate photos with different save dates, the duplication be revealed.

<br><br>

Copyright 2021-present by Pouya Kary (kary@gnu.org)<br>
&mdash; _a dragon's lover_