import React from 'react'
import Bricks from 'bricks.js'
import { connect } from 'react-redux'
import { addPics, appendPics } from '../'

const App = ({ pics, initBricks, addPics, appendPics }) => (
  <div>
    <h1>picdemo <button onClick={addPics}>Add 5 Pics</button></h1>
    
    <div id='bricks' ref={(node) => {
      initBricks(node)
    }}>
      { pics.map((p) => (
        <div key={p.id} className='tile'>
          <img src={p.image_url} />
          <p>{p.id}: {p.caption}</p>
        </div>
      )) }
    </div>

    <button onClick={appendPics}>Add 5 Pics</button>
  </div>
)

const mapStateToProps = ({ pics }) => ({
  pics,
})

const mapDispatchToProps = (dispatch, { bricks }) => ({
  addPics: () => dispatch(addPics()),
  appendPics: () => dispatch(appendPics()),
})

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(App)
