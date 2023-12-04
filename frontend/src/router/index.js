import { createWebHistory, createRouter } from "vue-router"
import Login from "../views/Login.vue";
import Home from "../views/Home.vue";
import PageNotFound from "../views/PageNotFound.vue";

const router = createRouter({
    history: createWebHistory(),
    routes: [
        {
            path: "/",
            name: "Login",
            component: Login,
        },
        {
            path: "/Home",
            name: "Home",
            component: Home,
        },
        {
            path: '/:catchAll(.*)*',
            name: 'PageNotFound',
            component: PageNotFound,
        }
    
    ]
})

export default router