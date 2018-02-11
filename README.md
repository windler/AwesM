# AwesM - Awesome Modeling
`AwesM` is a tool that simplifies the generation of graphviz diagrams using the `awesome modeling language`. This project is still under development. Currently, there is only an option to generate activity diagrams. Code is ugly and not testet. There will be no installation instructions until first major release.

## Examples
`aml-File`:
```
#start
wake up
stand up
..?[not tired]shower
..get clothed
..brush teeth 
..?[tired]drink coffee
....??[still tired]make sports
......|put on socks
......|go to toilet
#end
``` 

Output:  
![testoutput](examples/simple_test.png)