// SPA entry point (Rails: app/javascript/application.js).
//
// Vite is configured with this file as the only input (vite.config.js), and the
// Go shell links the hashed output resolved from .vite/manifest.json.
import '../../static/css/tailwind.css';

import { createApp } from 'vue';

import App from './App.vue';
import router from './router';
import store from '@legacy/vuex/counter-vuex-example';

// Props injected by the Go shell (app/views/layouts/application.html).
// They contain the route, its params, the locale and the environment - never a
// token: authentication travels in an httpOnly cookie.
function readProps() {
    const element = document.getElementById('app-props');
    if (!element) {
        return {};
    }
    try {
        return JSON.parse(element.textContent || '{}');
    } catch (error) {
        console.warn('Unable to parse app props', error);
        return {};
    }
}

const props = readProps();

const app = createApp(App);

store.commit('setState', props);
app.provide('store', store);
app.provide('props', props);

app.use(router);
app.mount('#app');
