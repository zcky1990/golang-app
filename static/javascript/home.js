import {createApp} from 'vue';
import App from './views/Home.vue'

document.addEventListener('DOMContentLoaded', () => {
    const initialData = window.__INITIAL_DATA__;

    createApp(App, {
        data: initialData
    }).mount('#app');
});