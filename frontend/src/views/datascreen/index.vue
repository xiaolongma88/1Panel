<template>
    <div class="body-container" @click="exitFullScreen">
        <div style="margin: 10px">
            <div class="header-container">
                <img style="width: 3%;margin-right: 10px" src="./assets/logo.png" alt=""/>
                <h1>边缘计算终端设备数据大屏展示中心</h1>
            </div>
            <div style="flex: 95%">主体</div>
        </div>
    </div>
</template>
<script setup lang="ts">
import {onMounted} from 'vue'
import { GlobalStore } from "@/store";

const globalStore = GlobalStore();

const exitFullScreen = (event) => {
    const rect = event.target.getBoundingClientRect();
    const clientX = event.clientX;
    const clientY = event.clientY;
    // 假设右上角 20x20 像素区域为目标区域
    if (!globalStore.isFullScreen && clientX >= rect.width&& clientY <= 30) {
        // 执行点击右上角伪元素时的逻辑
        globalStore.isFullScreen = true
    }
    if (globalStore.isFullScreen && clientX >= rect.right - 30 && clientY <= 30) {
        globalStore.isFullScreen = false
    }
}
onMounted(() => {

})
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
    background-image: url('./assets/Reset.svg');  /* 替换为您的转换 icon 图片路径 */
    background-size: cover;
    width: 20px;
    height: 20px;
    position: absolute;
    top: 10px;
    right: 10px;
}
.header-container {
    flex: 4%;
    width: 100%;
    display: flex;
    align-items: center;
    background-color: rgba(255, 0, 0, 0);
    background-image: url('./assets/bg-heade.png');
    background-size: 100% 100%;
    background-repeat: no-repeat;
    color: white
}
</style>
