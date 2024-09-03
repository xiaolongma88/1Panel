
export namespace Camera {
    export interface CameraInfo {
        KafkaConfig: KafkaConfig;
        CamereConfig:CamereConfig
        LabelConfig: LabelConfig;
    }
    export interface KafkaConfig {
        enable: Boolean;
        ip: String;
        port: String
        user: String;
        password: String;
    }
    export interface CameraItem {
        camID: String;
        rtspPath: String;
    }
    export interface CamereConfig {
        cameraList: Array<CameraItem>;
    }
    export interface LabelConfig {
        labelList: Array<string>;
    }
}
