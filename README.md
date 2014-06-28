# Seesaw
You never know when an idea will be in the sky? No?

## Installation

        # To pull the project from github easily
        go get github.com/souleiman/seesaw

### Go Lang
 - Mac: brew install go
 - Linux: Install [go](http://golang.org/doc/install) from binary or distro repository 
 - Windows: Download and run the latest version of [go](http://golang.org/doc/install)

### Dependencies
 - There are go dependencies that will be needed for set up.

        # With GOPATH and GOROOT properly setup
        go get github.com/coopernurse/gorp #gorp for easy to use db orm
        go get github.com/lib/pq #postgresql driver for golang

## Usage
 - Setting up database: (This assumes you have postgresql installed and you know how to use the cli)
 
        cd $GOPATH/src/github.com/souleiman/seesaw/sql/
        psql -U <your_username> postgres -f db.sql
        psql -U <your_username> imagine -f tables.sql
        
        # To add mock data
        psql -U <your_username> imagine -f mock.sql
        
        # In the near future, we may want to scrap this and fully utilize gorp
    
 - How to run:
 
        # To build and compile your project
        go install github.com/souleiman/seesaw/webserver/server
        
        #and to run, 
        server [-U psqlUser] [-D psqlDbname] [-R api root] [-r root dir] [-p port]
        
        #Or, you can
        go run $GOPATH/src/github.com/souleiman/seesaw/webserver/server/server.go [-U psqlUser] [-D psqlDbname] [-R api root] [-r root dir] [-p port]
 
## Notes
 - Make sure you have GOPATH and GOROOT setup. (Add $GOPATH/bin to your $PATH for your convenience)

