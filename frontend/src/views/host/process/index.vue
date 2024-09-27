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
        label: i18n.global.t('menu.process'),
        path: '/hosts/process/process',
    },
    {
        label: i18n.global.t('menu.network'),
        path: '/hosts/process/network',
    },
];
</script>
<style lang="scss" scoped>
:deep(.el-radio-button ){
    --el-radio-button-checked-text-color:var(--el-color-primary)
}
</style>
