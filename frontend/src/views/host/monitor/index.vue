<template>
    <div>
        <RouterButton :buttons="buttons" />
        <LayoutContent>
            <router-view></router-view>
        </LayoutContent>
    </div>
</template>

<script lang="ts" setup>
import i18n from '@/lang';

const buttons = [
    {
        label: i18n.global.t('monitor.monitor'),
        path: '/hosts/monitor/monitor',
    },
    {
        label: i18n.global.t('commons.button.set'),
        path: '/hosts/monitor/setting',
    },
];
</script>
<style scoped>

:deep(.el-radio-button ){
    --el-radio-button-checked-text-color:var(--el-color-primary)
}
</style>
