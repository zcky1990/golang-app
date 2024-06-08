import { createStore } from 'vuex'

// Create a new store instance.
export default createStore({
  state () {
    return {
      count: 0,
      data : ''
    }
  },
  mutations: {
    increment (state) {
      state.count++
    },
    addMessage (state, message){
      state.data = `${state.data} ${message}`
    }
  }
})