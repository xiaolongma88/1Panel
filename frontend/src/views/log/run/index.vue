<template>
    <div>
        <LayoutContent v-loading="loading" :title="$t('logs.run')">
            <template #search>
                <el-select class="float-left p-w-200" v-model="logConfig.name" @change="search()">
                    <template #prefix>{{ $t('commons.button.log') }}</template>
                    <el-option v-for="(item, index) in fileList" :key="index" :label="item" :value="item" />
                </el-select>
                <div class="watchCheckbox">
                    <el-checkbox border @change="changeTail" v-model="isWatch">
                        {{ $t('commons.button.watch') }}
                    </el-checkbox>
                </div>
            </template>
            <template #main>
                <LogFile
                    ref="logRef"
                    :config="logConfig"
                    :default-button="false"
                    v-if="showLog"
                    v-model:loading="loading"
                    v-model:hasContent="hasContent"
                    :style="'height: calc(100vh - 370px);min-height: 200px'"
                />
            </template>
        </LayoutContent>
    </div>
</template>

<script setup lang="ts">
import { nextTick, onMounted, reactive, ref } from 'vue';
import { useRouter } from 'vue-router';
import { getRunFiles } from '@/api/modules/log';
import LogFile from '@/components/log-file/index.vue';

const router = useRouter();
const loading = ref();
const isWatch = ref();
const fileList = ref();
const logRef = ref();

const hasContent = ref(false);
const logConfig = reactive({
    type: 'detect',
    name: '',
});
const showLog = ref(false);

const changeTail = () => {
    logRef.value.changeTail(true);
};

const loadFiles = async () => {
    const res = await getRunFiles();
    fileList.value = res.data || [];
    if (fileList.value) {
        logConfig.name = fileList.value[0];
        search();
    }
};

const search = () => {
    showLog.value = false;
    nextTick(() => {
        showLog.value = true;
    });
};

const onChangeRoute = async (addr: string) => {
    router.push({ name: addr });
};

onMounted(() => {
    loadFiles();
});
</script>

<style scoped lang="scss">
.watchCheckbox {
    margin-top: 2px;
    margin-bottom: 10px;
    float: left;
    margin-left: 20px;
}
</style>
