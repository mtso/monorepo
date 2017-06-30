import { combineReducers } from 'redux'

const pics = (state = [], action) => {
  switch(action.type) {
    case 'ADD_PICS':
      return [
        ...action.pics,
        ...state,
      ]
      
    default:
      return state
  }
}

export default combineReducers({
  pics,
})
