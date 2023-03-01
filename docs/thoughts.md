### Text vs varchar vs characters
Changed type of name column in database from character(50) to text, because character(n) is a
fixed-length type and it will add blank padding to the end if there are less then n chars. Text field
can store strings of unlimited length. Another aproach is to use varchar(n) instead, so you will have
a limit of n characters, but it still can be less then n.
