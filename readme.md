# Readme #

## Template Setup ##

1. initialise Go module: go mod initi github.com/mjdilworth/<name>
2. initialse git (git init -b main)
3. git add README.md
4. git commit -m "first commit"
5. git remote add origin git@github.com:mjdilworth/<reponame>.git
6. git push -u origin main 
7. git config --global alias.pushfwl "push --force-with-lease"
8. original template is in main branch and i make new branches for each test meeting.
9. git checkout -b <branch name>
10. git push -u origin <branch>

### Some handy go notes to help idomatic principles

- constant declarations are not UPPERCASE
- small varible names
- group variables - var {...}
- functions that panic prefix with "Must"
- Structures - use named initialisation of structure variables
- mutex grouping : in declaration put mutex above what it protects and name it for what it is protecting
- interface - needs to be an "er" for naming, and be composable of other interfaces to minimise require implementation
- function grouping/order : put most important functions at the top of the file, exported ones first
- HTTP handlers : always prefix name with "handle"
- enums : go doesnt have them , but prefix enum constants with type name
- contructor : for 
