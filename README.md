# japan holiday command for cal

## install

use golang

```
go get -u github.com/satoshun/cal-holiday/command/holiday
```


## usage

combines with cal command

show current month calendar with holiday

```sh
cal | holiday
   November 2016      
Su Mo Tu We Th Fr Sa  
       1  2  3  4  5  
 6  7  8  9 10 11 12  
13 14 15 16 17 18 19  
20 21 22 23 24 25 26  
27 28 29 30   
```

actually, coloring day of 3, 23. 

next, show 2/2016 holidays

```sh
cal 2 2016 | holiday
   February 2016      
Su Mo Tu We Th Fr Sa
    1  2  3  4  5  6  
 7  8  9 10 11 12 13  
14 15 16 17 18 19 20  
21 22 23 24 25 26 27  
28 29     
```

actually, coloring day of 29.


## TODO

- corresponds shunbun, shuubunn day
- more hackable command
