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

export const appendPics = () => {
  const newpics = gen.getFive()
  return {
    type: 'APPEND_PICS',
    pics: newpics,
  }
}

const sizes = [
  { columns: 2, gutter: 10 },
  { mq: '940px', columns: 3, gutter: 10 },
  { mq: '1250px', columns: 4, gutter: 10},
  { mq: '1560px', columns: 5, gutter: 10},
  { mq: '1870px', columns: 6, gutter: 10},  
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
  bricks.resize(true)

  if (!isPacked) {
    bricks.pack()
  } else {
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
