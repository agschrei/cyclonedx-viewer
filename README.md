# CycloneDx Viewer

The main purpose of this project is to teach myself how to build CLI apps in golang.  

Beyond that, it might also be useful if you are looking for a small application that can parse a [CycloneDx BOM](https://cyclonedx.org/use-cases/) and visualize it.  

Right now, this is just a skeleton, but the plan is to support the following features:  
- [ ] serve a dynamic web page that renders the BOM contents
- [ ] write rendered HTML to the local FS for offline consumption
- [ ] provide at-a-glance statistics about parsed BOMs on the CLI

## Why write it in golang?

- it ships with a small http server that does exactly what I need
- I've worked with hugo before and really like golang's html templating  
- I want to learn it and the best way to do so is to build something that might actually be useful to somebody
- It can compile to a single binary that can be distributed standalone


## Which itch does this tool attempt to scratch?  

Enterprise software is **messy**, understanding dependency trees that are *ten layers deep* by looking at a JSON is hard.  

So my goal here is to build a small tool that can present the information contained in a Software BOM in a more human-readable manner.  

I'm aware that the target audience for that kind of thing isn't huge, but it's a good project to dive a bit deeper into the world of golang.  

I am first and foremost trying to solve my own problems here and learn something new along the way, but if, by any means, this little tool is useful to you then you are more than welcome to adopt it.
