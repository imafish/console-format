console-format
=======



Usage
-------

the library acts differently when console windows size reaches certain limits:

- When window width < 35 characters: display only prefix
- When window width >= 35 characters:  
  the console window is divided to 3 columns: prefix, divider and suffix
  - prefix is at least 3/5 of the entire width;
  - divider at least 7 characters: 2 spaces followed by at least 3 divider chararacters then followed by 2 spaces.
  - suffix is at most 1/5 of the console width; minimum is 7 characters.
  - if suffix is longer than 1/5 of the console width, suffix will be shortened to 3/5 console width - 2 characters and a trailing ".." is appended.
  - if prefix + suffix is too long, prefix is shortened and ".." appended to fit into the console.

Progress
-------

### todo:

- [ ] StatusLine
- [ ] StatusLineMode
- [ ] SuffixAlignMode
- [ ] TextOverflowMode
- [ ] Windows
