import http from '@/api';
import { ReqPage, ResPage, UpdateByFile } from '../interface';
import { Camera } from '../interface/camera';

export const getCameraConfigs = () => {
    return http.get('/camera/conf/info');
}
export const updateCameraConfig = (req: Camera.CameraInfo) => {
    return http.post('/camera/conf/update', req);
}
export const getImages = (camId:string) => {
    return http.get('/camera/conf/img',{camId:camId});
}

export const parseRTSP = (req:Camera.RtspInfo) => {
    return http.post('/camera/stream',req)
}
