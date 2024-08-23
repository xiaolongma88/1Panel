import { Layout } from '@/routers/constant';

const dataScreenRouter = {
    sort: 1,
    path: '/data_screen',
    component: Layout,
    redirect: '/data_screen/index',
    meta: {
        icon: 'p-database',
        title: 'menu.data',
    },
    children: [
        {
            path: 'index',
            name: 'BigData',
            component: () => import('@/views/datascreen/index.vue'),
            hidden: true,
            meta: {
                requiresAuth: false,
            },
        },
    ],
};

export default dataScreenRouter;
