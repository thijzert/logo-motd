logo-motd is a small utility that renders a black and white image file to text using the [Unicode Block Elements](https://en.wikipedia.org/wiki/Block_Elements).
Its intended use is to generate a company logo greeter in `/etc/motd`.

Installation
------------
Just run:

    go get -u github.com/thijzert/logo-motd/...

Usage
-----
    logo-motd -o /etc/motd /path/to/image.png

License
-------
This program and its source code are available under the terms of the BSD 3-clause license.
Find out what that means here: https://www.tldrlegal.com/l/bsd3
