const routers = [
    {
        path: '/',
        meta: {
            title: ''
        },
        component: (resolve) => require(['./views/login.vue'], resolve)
    }
];
export default routers;