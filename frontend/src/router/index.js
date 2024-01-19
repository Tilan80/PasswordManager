import { createWebHistory, createRouter } from "vue-router"
import Login from "../views/Login.vue";
import Home from "../views/Home.vue";
import PageNotFound from "../views/PageNotFound.vue";
import AddNew from "../components/AddNew.vue";
import AddOwn from "../components/AddOwn.vue";
import Delete from "../components/Delete.vue";
import Get from "../components/Get.vue";
import GetAll from "../components/GetAll.vue";

export const router = createRouter({
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
        },
        {
            path: "/AddNew",
            name: "AddNew",
            component: AddNew,
        },
        {
            path: "/AddOwn",
            name: "AddOwn",
            component: AddOwn,
        },
        {
            path: "/Delete",
            name: "Delete",
            component: Delete,
        },
        {
            path: "/Get",
            name: "Get",
            component: Get,
        },
        {
            path: "/GetAll",
            name: "GetAll",
            component: GetAll,
        }
    
    ]
})