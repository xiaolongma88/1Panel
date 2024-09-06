import http from '@/api';
import { ReqPage, ResPage, UpdateByFile } from '../interface';
import { Camera } from '../interface/camera';

export const getCameraConfigs = () => {
    return http.get('/camera/conf/info');
}
export const updateCameraConfig = (req: Camera.CameraInfo) => {
    return http.post('/camera/conf/update', req);
}
export const getImages = () => {
    return http.get('/camera/conf/img');
}
