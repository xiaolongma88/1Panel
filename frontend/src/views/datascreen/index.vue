<template>
    <div class="body-container" @click="exitFullScreen">
        <div style="margin: 10px">
            <el-row>
                <el-col :span="24">
                    <div class="header-container">
                        <img style="width: 3%;margin-right: 10px" src="./assets/logo.png" alt="" />
                        <h1 style="width: 80%">边缘计算终端设备数据大屏展示中心</h1>
                        <div class="now-date">
                            <h1 style="margin: 0;color: #00FFAF">12:00:00</h1>
                            <h4 style="margin: 0">2024-09-01</h4>
                        </div>
                    </div>
                </el-col>
            </el-row>
            <el-row style="margin-top: 10px">
                <el-col :span="7">
                    <div style="height: 100%;display: flex;flex-direction: column">
                        <div style="flex: 33%;">
                            <div class="title-bg">
                                <img src="./assets/title_ico.png" style="width: 5%;height: 5%;text-align: center"
                                     alt="">
                                <h3 style="margin: 0 0 0 8px">基本信息</h3>
                            </div>
                            <div style="height: 90%;min-height: 15rem;margin-top: 5%">
                                <div style="height: 50%;display: flex">
                                    <div class="info bg1">
                                        <h4 style="width: 20%;margin-left: 15%;margin-top: 15%;margin-right: 10px">相机数量</h4>
                                        <div style="width: 60%;text-align: center;padding-top: 15%">
                                            <h3 style="margin:0">{{ form.cameras.length }}</h3>
                                            <div
                                                style="font-size: 12px;font-weight: normal;color: rgba(255, 255, 255, 0.6);">
                                                单位：个
                                            </div>
                                        </div>
                                    </div>
                                    <div class="info bg2">
                                        <h4 style="width: 20%;margin-left: 15%;margin-top: 15%;margin-right: 10px">计划任务</h4>
                                        <div style="width: 60%;text-align: center;padding-top: 15%">
                                            <h3 style="margin:0">{{ baseInfo.cronjobNumber }}</h3>
                                            <div
                                                style="font-size: 12px;font-weight: normal;color: rgba(255, 255, 255, 0.6);">
                                                单位：个
                                            </div>
                                        </div>
                                    </div>
                                </div>
                                <div style="height: 50%;display: flex">
                                    <div class="info bg3">
                                        <h4 style="width: 20%;margin-left: 15%;margin-top: 15%;margin-right: 10px">累计检测</h4>
                                        <div style="width: 60%;text-align: center;padding-top: 15%">
                                            <h3 style="margin:0">206</h3>
                                            <div
                                                style="font-size: 12px;font-weight: normal;color: rgba(255, 255, 255, 0.6);">
                                                单位：次
                                            </div>
                                        </div>
                                    </div>
                                    <div class="info bg4">
                                        <h4 style="width: 20%;margin-left: 15%;margin-top: 15%;margin-right: 10px">累计警告</h4>
                                        <div style="width: 60%;text-align: center;padding-top: 15%">
                                            <h3 style="margin:0">16</h3>
                                            <div
                                                style="font-size: 12px;font-weight: normal;color: rgba(255, 255, 255, 0.6);">
                                                单位：次
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div style="flex: 33%;margin-top: 10px">
                            <div class="title-bg">
                                <img src="./assets/title_ico.png" style="width: 5%;height: 5%;text-align: center"
                                     alt="">
                                <h3 style="margin: 0 0 0 8px">硬件信息</h3>
                            </div>
                            <div style="height: 90%;margin-top: 1%">
                                <div style="position: relative">
                                    <div style="height: 100%;min-height: 10rem;margin-top:0" ref="leftTwoRef"></div>
                                </div>
                            </div>
                        </div>
                        <div style="flex: 33%;margin-top: 15px;">
                            <div class="title-bg">
                                <img src="./assets/title_ico.png" style="width: 5%;height: 5%;text-align: center"
                                     alt="">
                                <h3 style="margin: 0 0 0 8px">流量</h3>
                            </div>
                            <div style="height: 90%;margin-top: 1%;">
                                <div style="height: 100%;min-height: 15rem;" ref="leftThreeRef"></div>
                            </div>
                        </div>
                    </div>
                </el-col>
                <el-col :span="10">
                    <div style="height: 100%;display: flex;flex-direction: column">
                        <div style="flex: 5%;color:#ffffff;display: flex;padding: 0 10px 0 10px;">
                            <div style="flex: 20%;display: flex;align-items: center;">
                                <img src="./assets/s2.png" alt="">
                                <h5 style="margin: 0 5px 0 2px">系统</h5>
                                <span style="font-size: 15px;">{{ baseInfo.platformFamily }}</span>
                            </div>
                            <div style="flex: 45%;display: flex;align-items: center">
                                <img src="./assets/s1.png" alt="">
                                <h5 style="margin: 0 10px 0 5px">启动时间</h5>
                                <span style="font-size: 15px;">{{ currentInfo.timeSinceUptime }}</span>
                            </div>
                            <div style="flex: 45%;display: flex;align-items: center;">
                                <img src="./assets/s3.png" alt="">
                                <h5 style="margin: 0 10px 0 5px">运行时长</h5>
                                <span style="font-size: 15px;"> {{ loadUpTime(currentInfo.uptime) }}</span>
                            </div>
                        </div>
                        <div
                            style="flex: 85%;display: flex;flex-wrap: wrap; row-gap: 20px;column-gap: 15px;padding: 10px">
                            <el-tabs v-model="activeName" @tab-click="handleClick"
                                     style="color: #ffffff;width: 100%;height: 50%">
                                <el-tab-pane v-for="(v,k) in form.cameras" :label="v.camID" :name="v.camID"
                                             style="height: 100%;">
                                    <div style="height: 99%;width: 99%">
                                        <video :src="videoUrl" autoplay
                                               style="width:100%;height:100%;border:1px solid #00FFAF;"></video>
                                    </div>
                                </el-tab-pane>
                            </el-tabs>
                            <div style="width: 100%;height: 20rem;overflow-y: auto;" class="video-normal">
                                <img v-for="url in urls" :src="url" style="width: 20%;margin: 10px" />
                            </div>
                        </div>
                        <div style="flex: 10%;color: #ffffff;display: flex;flex-wrap: wrap">
                            <div style="flex:1;display: flex;flex-direction:column;align-items: center"><img
                                src="./assets/device1.png" alt="" style="width: 30%"><h6 style="margin: 0">HK-001</h6>
                            </div>
                            <div style="flex:1;display: flex;flex-direction:column;align-items: center"><img
                                src="./assets/device1.png" alt="" style="width: 30%"><h6 style="margin: 0">HK-002</h6>
                            </div>
                            <div style="flex:1;display: flex;flex-direction:column;align-items: center"><img
                                src="./assets/device2.png" alt="" style="width: 30%"><h6 style="margin: 0">HK-003</h6>
                            </div>
                            <div style="flex:1;display: flex;flex-direction:column;align-items: center"><img
                                src="./assets/device1.png" alt="" style="width: 30%"><h6 style="margin: 0">HK-004</h6>
                            </div>
                            <div style="flex:1;display: flex;flex-direction:column;align-items: center"><img
                                src="./assets/device2.png" alt="" style="width: 30%"><h6 style="margin: 0">HK-005</h6>
                            </div>
                            <div style="flex:1;display: flex;flex-direction:column;align-items: center"><img
                                src="./assets/device1.png" alt="" style="width: 30%"><h6 style="margin: 0">HK-006</h6>
                            </div>
                        </div>
                    </div>
                </el-col>
                <el-col :span="7">
                    <div style="height: 100%;display: flex;flex-direction: column">
                        <div style="flex: 33%;">
                            <div class="title-bg">
                                <img src="./assets/title_ico.png" style="width: 5%;height: 5%;text-align: center"
                                     alt="">
                                <h3 style="margin: 0 0 0 8px">磁盘IO</h3>
                            </div>
                            <div style="height: 90%;margin-top: 1%;">
                                <div style="height: 100%;min-height: 15rem;" ref="rightOneRef"></div>
                            </div>
                        </div>
                        <div style="flex: 33%;">
                            <div class="title-bg">
                                <img src="./assets/title_ico.png" style="width: 5%;height: 5%;text-align: center"
                                     alt="">
                                <h3 style="margin: 0 0 0 8px">相机统计数据</h3>
                            </div>
                            <div style="height: 90%;margin-top: 1%;">
                                <div style="height: 100%;min-height: 15rem;" ref="rightTwoRef"></div>
                            </div>
                        </div>
                        <div style="flex: 33%;">
                            <div class="title-bg">
                                <img src="./assets/title_ico.png" style="width: 5%;height: 5%;text-align: center"
                                     alt="">
                                <h3 style="margin: 0 0 0 8px">相机统计数据</h3>
                            </div>
                            <div style="height: 90%;margin-top: 1%;">
                                <div style="height: 100%;min-height: 15rem;" ref="rightThreeRef"></div>
                            </div>
                        </div>
                    </div>
                </el-col>
            </el-row>
        </div>
    </div>
</template>
<script setup lang="ts">
import { onMounted, onUnmounted, ref, reactive, nextTick, onActivated, watch, computed } from "vue";
import { GlobalStore } from "@/store";
import * as echarts from "echarts";
import i18n from "@/lang";
import { loadBaseInfo, loadCurrentInfo } from "@/api/modules/dashboard";
import { getIOOptions, getNetworkOptions } from "@/api/modules/host";
import { getSettingInfo, loadUpgradeInfo } from "@/api/modules/setting";
import { dateFormatForSecond } from "@/utils/util";
import { Dashboard } from "@/api/interface/dashboard";
import { getCameraConfigs, updateCameraConfig, getImages } from "@/api/modules/camera";
import "./jsmpeg.min.js";

const globalStore = GlobalStore();

const leftOneRef = ref();
const leftTwoRef = ref();
const leftThreeRef = ref();
const rightOneRef = ref();
const rightTwoRef = ref();
const rightThreeRef = ref();

let isStatusInit = ref<boolean>(true);
const chartOption = ref("network");
const chartsOption = ref({ ioChart1: null, networkChart: null });
const timeIODatas = ref<Array<string>>([]);
const timeNetDatas = ref<Array<string>>([]);
const ioReadBytes = ref<Array<number>>([]);
const ioWriteBytes = ref<Array<number>>([]);
const netBytesSents = ref<Array<number>>([]);
const netBytesRecvs = ref<Array<number>>([]);

const urls = ref([]);
const timer = ref(null);
let DataTimer: NodeJS.Timer | null = null;
let isActive = ref(true);
const searchInfo = reactive({
    ioOption: "all",
    netOption: "all"
});
const form = reactive({
    kafka: {},
    cameras: [],
    labels: []
});
const state = reactive({
    myCharts: []
});
const baseInfo = ref<Dashboard.BaseInfo>({
    websiteNumber: 0,
    databaseNumber: 0,
    cronjobNumber: 0,
    appInstalledNumber: 0,

    hostname: "",
    os: "",
    platform: "",
    platformFamily: "",
    platformVersion: "",
    kernelArch: "",
    kernelVersion: "",
    virtualizationSystem: "",

    cpuCores: 0,
    cpuLogicalCores: 0,
    cpuModelName: "",
    currentInfo: null
});
const currentInfo = ref<Dashboard.CurrentInfo>({
    uptime: 0,
    timeSinceUptime: "",
    procs: 0,

    load1: 0,
    load5: 0,
    load15: 0,
    loadUsagePercent: 0,

    cpuPercent: [] as Array<number>,
    cpuUsedPercent: 0,
    cpuUsed: 0,
    cpuTotal: 0,

    memoryTotal: 0,
    memoryAvailable: 0,
    memoryUsed: 0,
    memoryUsedPercent: 0,
    swapMemoryTotal: 0,
    swapMemoryAvailable: 0,
    swapMemoryUsed: 0,
    swapMemoryUsedPercent: 0,

    ioReadBytes: 0,
    ioWriteBytes: 0,
    ioCount: 0,
    ioReadTime: 0,
    ioWriteTime: 0,

    diskData: [],
    gpuData: [],

    netBytesSent: 0,
    netBytesRecv: 0,

    shotTime: new Date()
});
const currentChartInfo = reactive({
    ioReadBytes: 0,
    ioWriteBytes: 0,
    ioCount: 0,
    ioTime: 0,

    netBytesSent: 0,
    netBytesRecv: 0
});

const initConfig = async () => {
    const res = await getCameraConfigs();
    form.kafka = res.data.KafkaConfig;
    form.cameras = res.data.CamereConfig.cameraList;
    form.labels = res.data.LabelConfig.labelList;

};
const initLeftTwo = () => {
    const myChart = echarts.init(leftTwoRef.value);
    let percentText = String(formatNumber(currentInfo.value.cpuUsedPercent)).split(".");
    let memory = String(formatNumber(currentInfo.value.memoryUsedPercent)).split(".");
    let load = String(formatNumber(currentInfo.value.loadUsagePercent)).split(".");
    const option = {
        title: [
            {
                text: `{a|${memory [0]}.}{b|${memory [1] || 0} %}`,
                textStyle: {
                    rich: {
                        a: {
                            fontSize: "22"
                        },
                        b: {
                            fontSize: "14",
                            padding: [5, 0, 0, 0]
                        }
                    },

                    color: "#ffffff",
                    lineHeight: 25,
                    // fontSize: 20,
                    fontWeight: 500
                },
                left: "15%",
                top: "40%",
                subtext: "内存",
                subtextStyle: {
                    color: "#fff",
                    fontSize: 13
                },
                textAlign: "center"
            },
            {
                text: `{a|${percentText[0]}.}{b|${percentText[1] || 0} %}`,
                textStyle: {
                    rich: {
                        a: {
                            fontSize: "22"
                        },
                        b: {
                            fontSize: "14",
                            padding: [5, 0, 0, 0]
                        }
                    },

                    color: "#ffffff",
                    lineHeight: 25,
                    // fontSize: 20,
                    fontWeight: 500
                },
                left: "49%",
                top: "40%",
                subtext: "CPU",
                subtextStyle: {
                    color: "#fff",
                    fontSize: 13
                },
                textAlign: "center"
            },
            {
                text: `{a|${load [0]}.}{b|${load[1] || 0} %}`,
                textStyle: {
                    rich: {
                        a: {
                            fontSize: "22"
                        },
                        b: {
                            fontSize: "14",
                            padding: [5, 0, 0, 0]
                        }
                    },

                    color: "#ffffff",
                    lineHeight: 25,
                    // fontSize: 20,
                    fontWeight: 500
                },
                right: "4%",
                top: "40%",
                subtext: "负载",
                subtextStyle: {
                    color: "#fff",
                    fontSize: 13
                },
                textAlign: "center"
            }
        ],
        polar: [
            {
                id: "polar1",
                radius: ["61%", "65%"],
                center: ["15%", "58%"]
            },
            {
                id: "polar2",
                radius: ["61%", "65%"],
                center: ["50%", "58%"]
            },
            {
                id: "polar3",
                radius: ["61%", "65%"],
                center: ["82%", "58%"]
            }
        ],
        angleAxis: [
            {
                polarId: "polar1",
                max: 100,
                show: false
            },
            {
                polarId: "polar2",
                max: 100,
                show: false
            },
            {
                polarId: "polar3",
                max: 100,
                show: false
            }
        ],
        radiusAxis: [
            {
                polarId: "polar1",
                type: "category",
                show: true,
                axisLabel: {
                    show: false
                },
                axisLine: {
                    show: false
                },
                axisTick: {
                    show: false
                }
            },
            {
                polarId: "polar2",
                type: "category",
                show: true,
                axisLabel: {
                    show: false
                },
                axisLine: {
                    show: false
                },
                axisTick: {
                    show: false
                }
            },
            {
                polarId: "polar3",
                type: "category",
                show: true,
                axisLabel: {
                    show: false
                },
                axisLine: {
                    show: false
                },
                axisTick: {
                    show: false
                }
            }
        ],
        series: [
            {
                type: "bar",
                roundCap: true,
                barWidth: 20,
                showBackground: true,
                coordinateSystem: "polar",
                radius: ["40%", "58%"],
                polarIndex: 0,
                backgroundStyle: {
                    color: "rgba(0, 255, 175, 0.3)"
                },
                color: [
                    new echarts.graphic.LinearGradient(0, 1, 0, 0, [
                        {
                            offset: 1,
                            color: "#00FFAF"
                        }
                    ])
                ],
                label: {
                    show: false
                },
                data: [formatNumber(currentInfo.value.memoryUsedPercent)]
            },
            {
                type: "bar",
                roundCap: true,
                barWidth: 20,
                showBackground: true,
                coordinateSystem: "polar",
                radius: ["40%", "58%"],
                polarIndex: 1,
                backgroundStyle: {
                    color: "rgba(0, 255, 175, 0.3)"
                },
                color: [
                    new echarts.graphic.LinearGradient(0, 1, 0, 0, [
                        {
                            offset: 1,
                            color: "#00FFAF"
                        }
                    ])
                ],
                label: {
                    show: false
                },
                data: [formatNumber(currentInfo.value.cpuUsedPercent)]
            },
            {
                type: "bar",
                roundCap: true,
                barWidth: 20,
                showBackground: true,
                coordinateSystem: "polar",
                radius: ["40%", "58%"],
                polarIndex: 2,
                backgroundStyle: {
                    color: "rgba(0, 255, 175, 0.3)"
                },
                color: [
                    new echarts.graphic.LinearGradient(0, 1, 0, 0, [
                        {
                            offset: 1,
                            color: "#00FFAF"
                        }
                    ])
                ],
                label: {
                    show: false
                },
                data: [formatNumber(currentInfo.value.loadUsagePercent)]
            },
            {
                type: "pie",
                radius: ["49%", "50%"],
                center: ["15%", "58%"],
                label: {
                    show: false
                },
                color: "#00FFAF",
                data: [
                    {
                        value: 0,
                        itemStyle: {
                            shadowColor: "rgba(0, 255, 175, 0.4)",
                            shadowBlur: 5
                        }
                    }
                ]
            },
            {
                type: "pie",
                radius: ["49%", "50%"],
                center: ["50%", "58%"],
                label: {
                    show: false
                },
                color: "#00FFAF",
                data: [
                    {
                        value: 0,
                        itemStyle: {
                            shadowColor: "rgba(0, 255, 175, 0.4)",
                            shadowBlur: 5
                        }
                    }
                ]
            },
            {
                type: "pie",
                radius: ["49%", "50%"],
                center: ["82%", "58%"],
                label: {
                    show: false
                },
                color: "#00FFAF",
                data: [
                    {
                        value: 0,
                        itemStyle: {
                            shadowColor: "rgba(0, 255, 175, 0.4)",
                            shadowBlur: 5
                        }
                    }
                ]
            }
        ]
    };
    myChart.setOption(option);
    state.myCharts.push(myChart);
};
const initLeftThree = () => {
    const myChart = echarts.init(leftThreeRef.value);
    const option = {
        grid: {
            left: "10%",
            right: "10%",
            top: "10%",
            bottom: "10%"
        },
        xAxis: {
            type: "category",
            axisLabel: {
                fontSize: 12,
                color: "#fff"
            },
            axisLine: {
                lineStyle: {
                    type: "dashed"
                }
            },
            data: timeNetDatas.value
        },
        yAxis: {
            type: "value",
            axisLabel: {
                fontSize: 12,
                color: "#fff"
            },
            axisLine: {
                lineStyle: {
                    type: "dashed"
                }
            }
        },
        legend: {
            data: [i18n.global.t('monitor.up'), i18n.global.t('monitor.down')],
            textStyle: {
                color: "#fff"
            }
        },
        series: [
            {
                name:i18n.global.t('monitor.up'),
                data: netBytesSents.value,
                type: "line",
                itemStyle: {
                    color: "#fff",
                    borderColor: "#1fef99",
                    borderWidth: 4,
                    radius: 8
                },
                lineStyle: {
                    width: 3,
                    color: "#00FFAF"
                },

            },
            {
                name:i18n.global.t('monitor.down'),
                data: netBytesRecvs.value,
                type: "line",
                itemStyle: {
                    color: "#fff",
                    borderColor: "#00ff00",
                    borderWidth: 4,
                    radius: 8
                },
                lineStyle: {
                    width: 3,
                    color: "#05caf6"
                },

            }
        ]
    };
    myChart.setOption(option);
    state.myCharts.push(myChart);
};
const initRightOne = () => {
    const myChart = echarts.init(rightOneRef.value);
    const option = {
        grid: {
            left: "10%",
            right: "10%",
            top: "10%",
            bottom: "10%"
        },
        xAxis: {
            type: "category",
            axisLabel: {
                fontSize: 12,
                color: "#fff"
            },
            axisLine: {
                lineStyle: {
                    type: "dashed"
                }
            },
            data: timeIODatas.value
        },
        yAxis: {
            type: "value",
            axisLabel: {
                fontSize: 12,
                color: "#fff"
            },
            splitLine: {
                show: false
            }
        },
        legend: {
            data: [i18n.global.t('monitor.read'), i18n.global.t('monitor.write')],
            textStyle: {
                color: "#fff"
            }
        },
        series: [
            {
                name: i18n.global.t('monitor.read'),
                data: ioReadBytes.value,
                type: "line",
                itemStyle: {
                    color: "#fff",
                    borderColor: "#00ff00",
                    borderWidth: 4,
                    radius: 8
                },
                lineStyle: {
                    width: 3,
                    color: "#00FFAF"
                },
                areaStyle: {
                    color: {
                        x: 0,
                        y: 0,
                        x2: 0,
                        y2: 1,
                        colorStops: [
                            {
                                offset: 0,
                                color: "#00FFAF"
                            },
                            {
                                offset: 1,
                                color: "rgba(255, 255, 255, 0)"
                            }
                        ]
                    }
                }
            },
            {
                name: i18n.global.t('monitor.write'),
                data: ioWriteBytes.value,
                type: "line",
                itemStyle: {
                    color: "#fff",
                    borderColor: "#00ff00",
                    borderWidth: 4,
                    radius: 8
                },
                lineStyle: {
                    width: 3,
                    color: "#05caf6"
                },
                areaStyle: {
                    color: {
                        x: 0,
                        y: 0,
                        x2: 0,
                        y2: 1,
                        colorStops: [
                            {
                                offset: 0,
                                color: "#05caf6"
                            },
                            {
                                offset: 1,
                                color: "rgba(255, 255, 255, 0)"
                            }
                        ]
                    }
                }
            }
        ]
    };
    myChart.setOption(option);
    state.myCharts.push(myChart);
};
const initRightTwo = () => {
    const myChart = echarts.init(rightTwoRef.value);
    const option = {
        grid: {
            left: "13%",
            right: "10%",
            top: "10%",
            bottom: "10%"
        },
        xAxis: {
            show: false
        },
        yAxis: {
            data: ["HK-001", "HK-002", "HK-003", "HK-004", "HK-005", "HK-006"],
            axisLabel: {
                fontSize: 12,
                color: "#fff"
            }
        },
        series: [
            {
                name: "数量",
                type: "bar",
                barWidth: 15,
                label: {
                    show: true,
                    position: "right",
                    valueAnimation: true,
                    formatter: function(params) {
                        return params.value;
                    },
                    textStyle: {
                        color: function(params) {
                            if (params.value >= 900) {
                                return "#00FFAF";
                            } else {
                                return "#fff";
                            }
                        }
                    }
                },
                itemStyle: {
                    color: function(params) {
                        if (params.value >= 900) {
                            return "#00FFAF";
                        } else {
                            return "rgba(255, 255, 255, 0.6)";
                        }
                    }
                },
                data: [900, 480, 18, 80, 8, 1880]
            }
        ]
    };
    myChart.setOption(option);
    state.myCharts.push(myChart);
};
const initRightThree = () => {
    const myChart = echarts.init(rightThreeRef.value);
    const option = {
        grid: {
            left: "13%",
            right: "10%",
            top: "10%",
            bottom: "10%",
            lineStyle: {
                type: "dashed"  // 将网格线设置为虚线
            }
        },
        xAxis: {
            type: "category",
            boundaryGap: false,
            axisLabel: {
                fontSize: 12,
                color: "#fff"
            },
            data: ["05-04", "05-05", "05-06", "05-07", "05-09", "05-10", "05-11"]
        },
        yAxis: {
            type: "value",
            axisLabel: {
                fontSize: 12,
                color: "#fff"
            },
            splitLine: {
                show: true
            }
        },
        series: [
            {
                type: "line",
                smooth: true,
                data: [20, 10, 30, 60, 50, 80, 30],
                itemStyle: {
                    color: "#ffffff"
                },
                lineStyle: {
                    color: "rgb(13,229,162)"
                },
                areaStyle: {
                    color: "rgb(9,148,106)"
                }
            }
        ]
    };
    myChart.setOption(option);
    state.myCharts.push(myChart);
};
const onLoadBaseInfo = async (isInit: boolean, range: string) => {
    if (range === "all" || range === "io") {
        ioReadBytes.value = [];
        ioWriteBytes.value = [];
        timeIODatas.value = [];
    } else if (range === "all" || range === "network") {
        netBytesSents.value = [];
        netBytesRecvs.value = [];
        timeNetDatas.value = [];
    }
    const res = await loadBaseInfo(searchInfo.ioOption, searchInfo.netOption);
    baseInfo.value = res.data;
    currentInfo.value = baseInfo.value.currentInfo;
    await onLoadCurrentInfo();
    isStatusInit.value = false;
    if (isInit) {
        DataTimer = setInterval(async () => {
            if (isActive.value && !globalStore.isOnRestart) {
                await onLoadCurrentInfo();
            }
        }, 3000);
    }
};
const onLoadCurrentInfo = async () => {
    const res = await loadCurrentInfo(searchInfo.ioOption, searchInfo.netOption);
    currentInfo.value.timeSinceUptime = res.data.timeSinceUptime;

    let timeInterval = Number(res.data.uptime - currentInfo.value.uptime) || 3;
    currentChartInfo.netBytesSent =
        res.data.netBytesSent - currentInfo.value.netBytesSent > 0
            ? Number(((res.data.netBytesSent - currentInfo.value.netBytesSent) / 1024 / timeInterval).toFixed(2))
            : 0;
    netBytesSents.value.push(currentChartInfo.netBytesSent);
    if (netBytesSents.value.length > 20) {
        netBytesSents.value.splice(0, 1);
    }

    currentChartInfo.netBytesRecv =
        res.data.netBytesRecv - currentInfo.value.netBytesRecv > 0
            ? Number(((res.data.netBytesRecv - currentInfo.value.netBytesRecv) / 1024 / timeInterval).toFixed(2))
            : 0;
    netBytesRecvs.value.push(currentChartInfo.netBytesRecv);
    if (netBytesRecvs.value.length > 20) {
        netBytesRecvs.value.splice(0, 1);
    }

    currentChartInfo.ioReadBytes =
        res.data.ioReadBytes - currentInfo.value.ioReadBytes > 0
            ? Number(((res.data.ioReadBytes - currentInfo.value.ioReadBytes) / 1024 / 1024 / timeInterval).toFixed(2))
            : 0;
    ioReadBytes.value.push(currentChartInfo.ioReadBytes);
    if (ioReadBytes.value.length > 20) {
        ioReadBytes.value.splice(0, 1);
    }

    currentChartInfo.ioWriteBytes =
        res.data.ioWriteBytes - currentInfo.value.ioWriteBytes > 0
            ? Number(((res.data.ioWriteBytes - currentInfo.value.ioWriteBytes) / 1024 / 1024 / timeInterval).toFixed(2))
            : 0;
    ioWriteBytes.value.push(currentChartInfo.ioWriteBytes);
    if (ioWriteBytes.value.length > 20) {
        ioWriteBytes.value.splice(0, 1);
    }
    currentChartInfo.ioCount = Math.round(Number((res.data.ioCount - currentInfo.value.ioCount) / timeInterval));
    let ioReadTime = res.data.ioReadTime - currentInfo.value.ioReadTime;
    let ioWriteTime = res.data.ioWriteTime - currentInfo.value.ioWriteTime;
    let ioChoose = ioReadTime > ioWriteTime ? ioReadTime : ioWriteTime;
    currentChartInfo.ioTime = Math.round(Number(ioChoose / timeInterval));

    timeIODatas.value.push(dateFormatForSecond(res.data.shotTime));
    if (timeIODatas.value.length > 20) {
        timeIODatas.value.splice(0, 1);
    }
    timeNetDatas.value.push(dateFormatForSecond(res.data.shotTime));
    if (timeNetDatas.value.length > 20) {
        timeNetDatas.value.splice(0, 1);
    }
    loadData();
    currentInfo.value = res.data;
    initLeftTwo();
    initLeftThree()
    initRightOne()
    results()
};
const loadData = async () => {
    if (chartOption.value === "io") {
        chartsOption.value["ioChart"] = {
            xData: timeIODatas.value,
            yData: [
                {
                    name: i18n.global.t("monitor.read"),
                    data: ioReadBytes.value
                },
                {
                    name: i18n.global.t("monitor.write"),
                    data: ioWriteBytes.value
                }
            ],
            formatStr: "MB"
        };
    } else {
        chartsOption.value["networkChart"] = {
            xData: timeNetDatas.value,
            yData: [
                {
                    name: i18n.global.t("monitor.up"),
                    data: netBytesSents.value
                },
                {
                    name: i18n.global.t("monitor.down"),
                    data: netBytesRecvs.value
                }
            ],
            formatStr: "KB/s"
        };
    }
};

function loadUpTime(uptime: number) {
    if (uptime <= 0) {
        return "-";
    }
    let days = Math.floor(uptime / 86400);
    let hours = Math.floor((uptime % 86400) / 3600);
    let minutes = Math.floor((uptime % 3600) / 60);
    let seconds = uptime % 60;
    if (days !== 0) {
        return (
            days +
            i18n.global.t("commons.units.day") +
            " " +
            hours +
            i18n.global.t("commons.units.hour") +
            " " +
            minutes +
            i18n.global.t("commons.units.minute") +
            " " +
            seconds +
            i18n.global.t("commons.units.second")
        );
    }
    if (hours !== 0) {
        return (
            hours +
            i18n.global.t("commons.units.hour") +
            " " +
            minutes +
            i18n.global.t("commons.units.minute") +
            " " +
            seconds +
            i18n.global.t("commons.units.second")
        );
    }
    if (minutes !== 0) {
        return minutes + i18n.global.t("commons.units.minute") + " " + seconds + i18n.global.t("commons.units.second");
    }
    return seconds + i18n.global.t("commons.units.second");
}

function formatNumber(val: number) {
    return Number(val.toFixed(2));
}

const results = async () => {
    getImages().then((res) => {
        urls.value = [];
        res.data.map(item => {
            if (item != "") urls.value.push("http://localhost:9999/api/v1/camera/config/res?imageName=" + item);
        });
    });
};
// 批量设置 echarts resize
const initEchartsResizeFun = () => {
    nextTick(() => {
        for (let i = 0; i < state.myCharts.length; i++) {
            state.myCharts[i].resize();
        }
    });
};
// 批量设置 echarts resize
const initEchartsResize = () => {
    window.addEventListener("resize", initEchartsResizeFun);
};
// 由于页面缓存原因，keep-alive
onActivated(() => {
    initEchartsResizeFun();
});
//全屏切换
const exitFullScreen = (event) => {
    const rect = event.target.getBoundingClientRect();
    const clientX = event.clientX;
    const clientY = event.clientY;
    // 假设右上角 20x20 像素区域为目标区域
    if (!globalStore.isFullScreen && clientX >= rect.width && clientY <= 30) {
        // 执行点击右上角伪元素时的逻辑
        globalStore.isFullScreen = true;
        initEchartsResizeFun();
    }
    if (globalStore.isFullScreen && clientX >= rect.right - 30 && clientY <= 30) {
        globalStore.isFullScreen = false;
        initEchartsResizeFun();
    }
};

const activeName = ref("001");
const videoUrl = ref("/src/views/datascreen/video/test1.mp4");
const handleClick = (tab: TabsPaneContext, event: Event) => {
    // console.log(tab, event)
};

onMounted(() => {
    initConfig();
    onLoadBaseInfo(true, "all");
    initLeftTwo();
    initLeftThree();
    initRightOne();
    initRightTwo();
    initRightThree();
    initEchartsResize();
});
onUnmounted(() => {
    timer.value = null;
});
// 监听 pinia 中的 tagsview 开启全屏变化，重新 resize 图表，防止不出现/大小不变等
watch(
    () => globalStore.isFullScreen,
    () => {
        initEchartsResizeFun();
    }
);
watch(
    () => activeName.value,
    () => {
        switch (activeName.value) {
            case "001":
                videoUrl.value = "http://localhost:9999/api/v1/camera/test/video?videoName=test1.mp4";
                break;
            case "002":
                videoUrl.value = "http://localhost:9999/api/v1/camera/test/video?videoName=test2.mp4";
                break;
            case "003":
                videoUrl.value = "http://localhost:9999/api/v1/camera/test/video?videoName=test3.mp4";
                break;
            case "004":
                videoUrl.value = "http://localhost:9999/api/v1/camera/test/video?videoName=test4.mp4";
                break;
            case "005":
                videoUrl.value = "http://localhost:9999/api/v1/camera/test/video?videoName=test5.mp4";
                break;
            case "006":
                videoUrl.value = "http://localhost:9999/api/v1/camera/test/video?videoName=test6.mp4";
                break;
            case "007":
                videoUrl.value = "http://localhost:9999/api/v1/camera/test/video?videoName=test1.mp4";
                break;
            default:
                videoUrl.value = "http://localhost:9999/api/v1/camera/test/video?videoName=test2.mp4";
        }
    }
);
</script>
<style scoped lang="scss">
.body-container {
    width: 100%;
    height: 100%;
    margin: 0;
    background-image: url('./assets/bg2.png');
    background-size: cover;
    display: flex;
    flex-direction: column;
}

.body-container::after {
    content: "";
    background-image: url('./assets/Reset.svg'); /* 替换为您的转换 icon 图片路径 */
    background-size: cover;
    width: 20px;
    height: 20px;
    position: absolute;
    top: 10px;
    right: 10px;
}

.header-container {
    display: flex;
    align-items: center;
    background-color: rgba(255, 0, 0, 0);
    background-image: url('./assets/bg-heade.png');
    background-size: 100% 100%;
    background-repeat: no-repeat;
    color: white
}

.now-date {
    margin-top: 20px;
    width: 17%;
    height: 100%;
    background-image: url("./assets/date_box.png");
    background-size: 100% 100%;
    text-align: center;
}

.title-bg {
    color: #ffffff;
    background-image: url("./assets/title_bg.png");
    background-size: 100% 100%;
    display: flex;
}

.info {
    flex: 50%;
    margin: 10px 20px 10px 10px;
    color: #ffffff;
    display: flex;
}

.bg1 {
    background-image: url("./assets/infg_bg1.png");
    background-repeat: no-repeat;
}

.bg2 {
    background-image: url("./assets/info_bg2.png");
    background-repeat: no-repeat;
}

.bg3 {
    background-image: url("./assets/info_bg3.png");
    background-repeat: no-repeat;
}

.bg4 {
    background-image: url("./assets/info_bg4.png");
    background-repeat: no-repeat;
}

.sys-info {
    background-image: url("./assets/bg_circle.png");
}

.video-normal {
    background-image: url("./assets/bg-video.png");
    background-size: 100% 100%;
}

:deep(.el-tabs__item) {
    color: #ffffff;
}

:deep(.el-tabs__item:hover) {
    color: #00FFAF;
}

:deep(.el-tabs__item.is-active) {
    color: #00FFAF
}

:deep(.el-tabs__active-bar) {
    background-color: #00FFAF
}

:deep(.el-tabs__nav-next) {
    color: #00FFAF
}

:deep(.el-tabs__nav-prev) {
    color: #00FFAF
}
</style>
