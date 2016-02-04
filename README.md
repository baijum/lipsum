# Serve junk content over HTTP

To get the package:

    go get github.com/baijum/lipsum

To run the server:

    lipsum -port 8080 -size 2

The size in mega bytes can be specified using the `-size` option.

Now any URL in that port is going to serve 2 MB plain text file
with [Lorem Ipsum](https://en.wikipedia.org/wiki/Lorem_ipsum) content.

Try to access these URLs:

- [http://localhost:8080/1.txt](http://localhost:8080/1.txt)
- [http://localhost:8080/2.txt](http://localhost:8080/2.txt)
- [http://localhost:8080/3.txt](http://localhost:8080/3.txt)
