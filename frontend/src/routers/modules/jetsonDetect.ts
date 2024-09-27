import { Layout } from '@/routers/constant';
const jetsonDetectRouter= {
    sort: 8,
    path: '/jetsonDetect',
    component: Layout,
    meta: {
        title: 'menu.jetsonDetect',
        icon: 'p-log',
    },
    children: [
        {
            path: '/jetsonDetect/camera',
            name: 'CameraConfig',
            component: () => import('@/views/jetsonDetect/camera/index.vue'),
            meta: {
                title: 'menu.cameraConfig',
                requiresAuth: false,
            },
        },
        {
            path: 'run',
            name: 'RunLog',
            component: () => import('@/views/log/run/index.vue'),
            meta: {
                title: 'menu.appLog',
                requiresAuth: false,
            },
        },
        {
            path: '/jetsonDetect/cronjobs',
            name: 'Cronjob',
            component: () => import('@/views/cronjob/index.vue'),
            meta: {
                title: 'menu.cronjob',
                requiresAuth: false,
            },
        },
        {
            path: 'device',
            name: 'Device',
            component: () => import('@/views/toolbox/device/index.vue'),
            meta: {
                title: 'menu.quickSetting',
                requiresAuth: false,
            },
        },
    ]
}
export default jetsonDetectRouter;
