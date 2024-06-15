import {createApp} from 'vue';
import App from './views/Home.vue'
import snackBarStore from './vuex/counter-vuex-example'

document.addEventListener('DOMContentLoaded', () => {
    // Create the Vue app
    const app = createApp(App);
    // Provide the Vuex store to all components
    snackBarStore.commit('setState', window.__INITIAL_DATA__)
    
    app.provide('store', snackBarStore);
    // Mount the app to the DOM
    app.mount('#app');

});