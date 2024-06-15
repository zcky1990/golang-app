import { createStore } from 'vuex'

// Create a new store instance.
const snackBarStore = createStore({
  state () {
    return {
      count: 0,
      data : {}
    }
  },
  mutations: {
    increment (state) {
      state.count++
    },
    setState(state, payload) {
      state.data = payload
    }
  }
})

export default snackBarStore;