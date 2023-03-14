# PosGraph

Posgraph is a little program that I built for my Optimisation course at the University of Ljubljana. 
Posgraph takes a image path as an argument and give you an interface to create points and name them.

Right click to position the point, and Enter to validate. Then enter the name of the point on the terminal. When you are done creating points, close the GUI window. 
The program will print a python code using matplotlib and the networkx library.

```shell
git clone https://github.com/slashformotion/posgraph
go get
go build
./posgraph france.jpg
```

Please note that I used Ebiten as a graphic engine, thus the binary can't be build statically.
