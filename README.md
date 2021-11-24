rm .git/index
rm -r *

### Go

Follow instructions in [../README.md](../README.md) to install protoc. Then
install the Go protoc plugin (protoc-gen-go):

    $ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

The "go install" command will install protoc-gen-go into the GOBIN
directory.  You can set the $GOBIN environment variable before
running "go install" to change the install location.  Make sure the
install directory is in your shell $PATH.

Build the Go samples with "make go".  This creates the following
executable files in the current directory:

    add_person_go      list_people_go

To run the example:

    ./add_person_go addressbook.data

to add a person to the protocol buffer encoded file addressbook.data.  The file
is created if it does not exist.  To view the data, run:

    ./list_people_go addressbook.data

Observe that the C++, Python, Java, and Dart examples in this directory run in a
similar way and can view/modify files created by the Go example and vice
versa.