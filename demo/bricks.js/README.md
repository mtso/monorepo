# bricks.js demo

[summary]::
Integrate bricks.js masonry layout into React.

This project demos integrating a response image layout engine (bricks.js) into a React/Redux frontend. The Generator class generates random pic objects for mapping the tiles through. The reducer simply updates the pics array by unshifting new pics into the front of the array. 

The bricks.js library is successfully added by setting the `ref` parameter of the tile container to a callback that initializes the bricks.js instance with the passed node.
`.pack()` is then called on the first time and `.update()` when new images are pushed.

In a real app, I think a good place to instantiate the bricks instance is in the Page container that renders the pic tile grid.

In order to animate the tile transition, we set the `position` attribute of the bricks.js instance to false. This means that bricks.js will use `CSS:transform3d` to re-arrange the tiles, which allows us to use set `transition: transform3d 2s` on the tile's class.

It is also possible to listen to the `pack/update/resize` completions that bricks.js emits after each method.
