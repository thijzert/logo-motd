logo-motd is a small utility that renders a black and white image file to text using the [Unicode Block Elements](https://en.wikipedia.org/wiki/Block_Elements).
Its intended use is to generate a company logo greeter in `/etc/motd`.

Installation
------------
Just run:

    go get -u github.com/thijzert/logo-motd/...

Usage
-----
    logo-motd -o /etc/motd /path/to/image.png

Roadmap
-------
Planned features include adding the ability to scale images, and to automatically compensate for the pixel ratio in terminal windows. (On my machine it's 9:17; YMMV.)
Bonus points if I can detect how wide the terminal window is.

License
-------
This program and its source code are available under the terms of the BSD 3-clause license.
Find out what that means here: https://www.tldrlegal.com/l/bsd3
