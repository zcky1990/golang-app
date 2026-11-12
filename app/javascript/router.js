// Client routes. This table must mirror config/routes.go: every path the Go
// router answers with the shell needs an entry here, otherwise a page refresh
// renders nothing.
import { createRouter, createWebHistory } from 'vue-router';

const routes = [
    {
        path: '/',
        name: 'home',
        component: () => import('./views/HomeView.vue'),
    },
    {
        path: '/wedding/:id',
        name: 'wedding',
        component: () => import('./views/WeddingView.vue'),
        props: true,
    },
    {
        path: '/:pathMatch(.*)*',
        name: 'not-found',
        component: () => import('./views/NotFoundView.vue'),
    },
];

export default createRouter({
    history: createWebHistory(),
    routes,
    scrollBehavior() {
        return { top: 0 };
    },
});
