import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/login.vue'
import Space from '../views/space.vue'

Vue.use(VueRouter)

const routes = [
    {
        path: "/login",
        name: "login",
        component: Login
    },
    {
        path: "/space",
        name: "space",
        component: Space
    }
]

const router = new VueRouter({
    routes
})

router.beforeEach((to, from, next) => {
    if(to.path == '/login')
        next()
    const token = window.localStorage.getItem('token')
    if(to.path == '/space' && !token)
        next('login')
    else
        next()
}) 

export default router
