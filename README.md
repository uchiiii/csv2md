# csv2md
![](https://github.com/uchiiii/csv2md/workflows/CI/badge.svg)
Convert csv to markdown.

## Install

```
$ go get -u -t github.com/uchiiii/csv2md
```

## Usage
```
$ csv2md help
NAME:
   csv2md - A new cli application

USAGE:
   csv2md [global options] command [command options] [arguments...]

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --delimiter value, -d value  CSV delimiter, expected values: ',', ';'. Default is , (default: ",")
   --padding value, -p value    The number of spaces to add between table cells and column dividers. Default is 2 spaces. (default: 2)
   --help, -h                   show help (default: false)
```

### Example
```
# here testdata in this reository is used.
$ csv2md ./testdata/test1.csv
```

input: ([csv file](https://github.com/uchiiii/csv2md/blob/master/testdata/test1.csv))
```
First Name,Last Name,Location,Allegiance
Mance,Rayder,North of the Wall,Wildlings
Margaery,Tyrell,The Reach,House Tyrell
Danerys,Targaryen,Meereen,House Targaryen
Tyrion,Lannister,King's Landing,House Lannister
```

output: ([markdown](https://github.com/uchiiii/csv2md/blob/master/testdata/test1_expected.md))

First Name  |  Last Name  |  Location           |  Allegiance     
------------|-------------|---------------------|-----------------
Mance       |  Rayder     |  North of the Wall  |  Wildlings      
Margaery    |  Tyrell     |  The Reach          |  House Tyrell   
Danerys     |  Targaryen  |  Meereen            |  House Targaryen
Tyrion      |  Lannister  |  King's Landing     |  House Lannister

raw output:

```
First Name  |  Last Name  |  Location           |  Allegiance     
------------|-------------|---------------------|-----------------
Mance       |  Rayder     |  North of the Wall  |  Wildlings      
Margaery    |  Tyrell     |  The Reach          |  House Tyrell   
Danerys     |  Targaryen  |  Meereen            |  House Targaryen
Tyrion      |  Lannister  |  King's Landing     |  House Lannister
```


## Features
1. **Newline within cell in csv is converted to `<br/>` which is newline sign in markdown.**

## Contribution
Bug reports, fixes, or features? Feel You are wellcome opening issue or pull request any time.

# License
Copyright (c) 2020 Ryosuke Horiuchi. Licensed under [the MIT License](http://opensource.org/licenses/MIT).


