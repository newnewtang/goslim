# goslim
### An implementation of FitNesse SliM in golang
### Author:  newnewtang 

# note
goslim comes from cslim. 
                  cslim is writted by `Robert Martin` and I am his fans.  

## Fixtures

Fixtures in goslim are sets of functions and a data structure that binds them
together.   

Note that the struct Division is used to hold data for all the functions in
this fixture, rather like an object <grin>.

Each function returns a `string`,  so if you want to return an integer you'll 
have to convert it to a `string`.  

You can cause a slim exception in any function by returning a message with the
`slimprotocol.EXCEPTION(reason)` function.  This will turn the cell being executed 
yellow. 

## Registration

Notice the `RegisterFixture` business in the demo/main.go file, like this:

        goslim.RegisterFixture("Division", division.Division_Create)

This creates a registration function that will be called by the main program
of the Slim Server.  It is your responsibility to register your fixture.

You donn't need to register your functions, but if you wang to use alias name
to call function, you can register like this:

        goslim.RegisterMethod("TestSlim", "echo", "OneArg")
        
The middle paramet is the alias method name.


## Running the example fitnesse test

You need to download FitNesse from fitnesse.org.  Example pages are in
src/demo/pages.  Fire up fitnesse and create a new page like
CslimExample.  Add this to your CslimExample page:

        !contents -R2 -g -p -f -h
        !define TEST_SYSTEM {slim}
        !define TEST_RUNNER {<path>/cslim/Cslim_cslim}
        !define COMMAND_PATTERN {%m}

You should be able to see the CounterTest, DivisionTest, etc.


## Communications

goslim communicates with FitNesse over a stream socket.  A tcp server
 is provided.  

