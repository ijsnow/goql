# GoQL

A port from the [graphql/graphql-js](https://github.com/graphql/graphql-js) library to Go.

## Reasoning

All current Go implementations of GraphQL servers are incomplete and either unmaintained or unstable.

## Goal

Create an implementation of a GraphQL server in Go that is as complete as the official JS library itself and use best practices for maintaining an open sourced library.

## Course of action

First, create as close of a mirror as we can to the JS lib. To do this, 
we'll nearly mirror the directory structure and essentially copy the files and
convert the code to the closest possible go implementation, even if the code doesn't
end up being the cleanest and there is a lot of repitition.

Then, "go-ify" the code. Reorganize, clean the code up and hopefully remove a lot of code that is
unneccessary due to Go's awesome type system.

Finally, make cool stuff.