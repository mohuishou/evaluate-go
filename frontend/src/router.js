const routers = [
    {
        path: '/',
        meta: {
            title: ''
        },
        component: (resolve) => require(['./views/login.vue'], resolve)
    },
    {
        path: '/lists',
        meta: {
            title: ''
        },
        component: (resolve) => require(['./views/lists.vue'], resolve)
    }
];
export default routers;