import React from 'react'
import Bricks from 'bricks.js'
import { connect } from 'react-redux'
import { addPics } from '../'

const App = ({ pics, initBricks, addPics }) => (
  <div>
    <p>picdemo</p>
    <button onClick={addPics}>Add 5</button>
    <div id='bricks' ref={(node) => {
      initBricks(node)
    }}>
      { pics.map((p) => (
        <div key={p.id} style={{width:300, height:450}} className='tile'>
          <img src={p.image_url} />
          <p>{p.id}: {p.caption}</p>
        </div>
      )) }
    </div>
  </div>
)

const mapStateToProps = ({ pics }) => ({
  pics,
})

const mapDispatchToProps = (dispatch, { bricks }) => ({
  addPics: () => dispatch(addPics()),
})

export default connect(
  mapStateToProps,
  mapDispatchToProps
)(App)
