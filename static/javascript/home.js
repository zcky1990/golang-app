import {createApp} from 'vue';
import App from './views/Home.vue'
import counterStore from './vuex/counter-vuex-example'

document.addEventListener('DOMContentLoaded', () => {
    const initialData = window.__INITIAL_DATA__;
    
    // const app = createApp(App);

    // createApp(App, {
    //     data: initialData
    // }).mount('#app');

    // Create the Vue app
    const app = createApp(App);
    // Provide the Vuex store to all components
    app.provide('store', counterStore);
    // Mount the app to the DOM
    app.mount('#app');

});