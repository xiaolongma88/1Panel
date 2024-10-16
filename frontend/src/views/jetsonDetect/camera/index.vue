<script setup lang="ts">
import { reactive, ref } from "vue";
import {getCameraConfigs,updateCameraConfig} from '@/api/modules/camera'
import { onMounted } from "@vue/runtime-core";
import {
    Plus,
    Delete,
} from '@element-plus/icons-vue'
import { ElMessage, ElMessageBox } from "element-plus";

const loading = ref(false);
const form = reactive({
    kafka:{},
    cameras:[],
    labels:[],
});
const initConfig = async () => {
    const res = await getCameraConfigs()
    form.kafka = res.data.KafkaConfig
    form.cameras = res.data.CamereConfig.cameraList
    form.labels = res.data.LabelConfig.labelList
    console.log(form.cameras);
}

const addOne = () =>{
    form.cameras.push({camID:'',rtspPath:'',width:600,height:400})
}
const delOne = (key) => {
    form.cameras.splice(key,1)
}
const refresh = () => {
    location.reload();
}
const saveReload = async () => {
    const data = {
        KafkaConfig:form.kafka,
        CamereConfig:{
            cameraList:form.cameras
        },
        LabelConfig:{
            labelList:form.labels
        }
    }

    ElMessageBox.confirm( '保存后不可恢复，确认继续？',
        '确认提示',
        {
            confirmButtonText: '确定',
            cancelButtonText: '取消',
            type: 'warning'
        }
    )
        .then(async () => {
            loading.value = true;
            await updateCameraConfig(data).then(res => {
                loading.value = false
                if (res.code == 200) ElMessage.success('修改成功');
            })
        })
        .catch(() => {
            //location.reload();
        });

}
onMounted(() => {
    initConfig()
})
</script>

<template>
<div v-loading="loading">
    <LayoutContent :title="$t('toolbox.camera.toolbox')" :divider="true">
        <template #main>
            <el-row style="margin-top: 20px">
                <el-col :span="1"><br /></el-col>
                <el-col :xs="24" :sm="20" :md="20" :lg="10" :xl="10">
                    <h4>Kafka配置</h4>
                    <el-form :model="form.kafka" label-position="left" ref="formRef" label-width="120px">
                        <el-form-item label="状态" prop="enable">
                            <el-switch v-model="form.kafka.enable" />
                        </el-form-item>
                        <el-form-item label="IP" prop="ip">
                            <el-input v-model="form.kafka.ip"></el-input>
                        </el-form-item>
                        <el-form-item label="密码" prop="password">
                            <el-input v-model="form.kafka.password"></el-input>
                        </el-form-item>
                        <el-form-item label="端口" prop="port">
                            <el-input v-model="form.kafka.port"></el-input>
                        </el-form-item>
                        <el-form-item label="用户" prop="user">
                            <el-input v-model="form.kafka.user"></el-input>
                        </el-form-item>
                    </el-form>
                </el-col>
                <el-col :span="1"><br /></el-col>
                <el-col :xs="24" :sm="20" :md="20" :lg="10" :xl="10">
                    <h4>标签配置</h4>
                    <el-form  label-position="left" ref="formRef" label-width="120px">
                        <el-select
                            disabled
                            v-model="form.labels"
                            multiple
                            placeholder="Select"
                            style="width: 100%"
                        >
                            <el-option
                                v-for="(item,key) in form.labels"
                                :key="key"
                                :label="item"
                                :value="item"
                            />
                        </el-select>
                    </el-form>
                </el-col>
            </el-row>
            <el-row style="margin-top: 20px">
                <el-col :span="1"><br /></el-col>
                <el-col :span="24">
                    <h4>相机列表</h4>
                    <el-row>
                        <el-col :xs="24" :sm="20" :md="20" :lg="10" :xl="10" v-for="(item,k) in form.cameras" :key="k" style="margin-right: 20px">
                            <el-col :span="1"><br /></el-col>
                            <el-row >
                                <el-col :span="20">
                                    <el-form-item label="相机ID">
                                        <el-input v-model="form.cameras[k].camID"></el-input>
                                    </el-form-item>
                                    <el-form-item label="宽">
                                        <el-input-number v-model="form.cameras[k].width"></el-input-number>
                                    </el-form-item>
                                    <el-form-item label="高">
                                        <el-input-number v-model="form.cameras[k].height"></el-input-number>
                                    </el-form-item>
                                </el-col>
                                <el-col :span="4">
                                    <el-button v-if="k == form.cameras.length-1" @click="addOne" type="primary" :icon="Plus" circle style="margin-left: 10px"/>
                                    <el-button type="danger" @click="delOne(k)" :icon="Delete" circle style="margin-left: 10px"/>
                                </el-col>
                            </el-row>
                            <el-form-item label="流地址">
                                <el-input v-model="form.cameras[k].rtspPath"></el-input>
                            </el-form-item>

                        </el-col>
                    </el-row>
                </el-col>
            </el-row>
            <el-row style="margin-top: 20px">
                <el-col :span="20"></el-col>
                <el-col :span="4">
                    <el-button type="info" @click="refresh">恢复</el-button>
                    <el-button type="primary" @click="saveReload">保存并重启应用</el-button>
                </el-col>
            </el-row>
        </template>
    </LayoutContent>
</div>
</template>

<style scoped lang="scss">

</style>
