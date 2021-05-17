# icka

irccloud keep-alive

## usage

```sh
# build and install
$ make
$ make install

# execute with command line arguments
$ icka -email <email> -password <password>

# provide credentials from environment
$ ICKA_EMAIL=<email> ICKA_PASSWORD=<password> icka

# run once every hour, without exiting
$ icka -email <email> -password <password> -forever
```
