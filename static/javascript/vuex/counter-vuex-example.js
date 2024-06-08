import { createStore } from 'vuex'

// Create a new store instance.
export default createStore({
  state () {
    return {
      snackbar:{
        show: false,
        message: '',
      },
      count: 0,
      data : ''
    }
  },
  mutations: {
    setSnackbarMessage(state, message) {
      state.snackbar.message = message
    },
    showMessage(state) {
      state.snackbar.show = true
    },
    hideMessage(state) {
      state.snackbar.show = false
    },
    increment (state) {
      state.count++
    },
    addMessage (state, message){
      state.data = `${state.data} ${message}`
    }
  }
})