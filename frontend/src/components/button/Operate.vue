<template>
    <div class="x-button" :style="style">
        <slot></slot>
    </div>
</template>

<script setup>
import { reactive } from 'vue';
const props = defineProps({
    color: String,
    width: String,
    height: String
})
const style = reactive({
    backgroundColor: props.color,
    width: props.width,
    height: props.height,
})

</script>


<style lang="less" scoped>
.x-button {
    line-height: 100%;
    border-radius: 50%;
}
</style>