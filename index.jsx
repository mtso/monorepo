import React from 'react'
import { render } from 'react-dom'
import { Provider } from 'react-redux'
import { createStore } from 'redux'
import rootReducer from './reducers'
import App from './components/App'
import './bricks.scss'

import { Generator } from './data'
import Bricks from 'bricks.js'

const gen = new Generator()
const pics = gen.getInitialPics()
const initialState = {
  pics,
}

export const addPics = () => {
  const newpics = gen.getFive()
  return {
    type: 'ADD_PICS',
    pics: newpics,
  }
}

const sizes = [
  { columns: 3, gutter: 10 },
]

let bricks = null
let isPacked = false

const initBricks = (node) => {
  bricks = new Bricks({
    container: '#bricks',
    packed: 'data-packed',
    sizes,
    position: false,
  })

  if (!isPacked) {
    bricks.pack()
  } else {
    console.log('updating')
    bricks.update()
  }
}

const store = createStore(rootReducer, initialState)
const mount = document.createElement('div')
document.body.appendChild(mount)

render(
  <Provider store={store}>
    <App initBricks={initBricks} />
  </Provider>,
  mount
)
