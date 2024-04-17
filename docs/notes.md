# some notes #

## open tcp ports
netstat -anp tcp | awk 'NR<3 || /LISTEN/'

or 

lsof -PiTCP -sTCP:LISTEN

### extract PID of myapp
lsof -PiTCP -sTCP:LISTEN | grep myapp | awk '{ print $2 }'

- use backticks to evaluate and used by kill
kill `lsof -PiTCP -sTCP:LISTEN | grep myapp | awk '{ print $2 }'`