<template>
    <div class="nav">
        <div class="common-operation">
            <div class="cell" v-for="item in commonOperation" :key="item.key" @click="item.callback()">
                <!-- {{ item.title }} -->
                <Operate :color="item.color" width="15px" height="15px" ></Operate>
            </div>
        </div>
    </div>
</template>

<script setup>
import Operate from "../button/Operate.vue"
import { WindowMinimise, WindowToggleMaximise, Quit } from "../../../wailsjs/runtime"

const commonOperation = [
    {
        "key": 1,
        "title": "最小化",
        "color": "green",
        "callback": WindowMinimise,
    },
    {
        "key": 2,
        "title": "最大化",
        "color": "yellow",
        "callback": WindowToggleMaximise,
    },
    {
        "key": 3,
        "title": "关闭",
        "color": "red",
        "callback": Quit,
    },
]

</script>

<style lang="less" scoped>
.nav {
    color: white;
    background-color: #000000;
    height: 30px;
    --wails-draggable: drag;

    .common-operation {
        display: flex;
        flex-direction: row;
        justify-content: flex-end;
        align-items: center;
        height: 100%;

        .cell {
            margin: 0 5px;
            cursor: pointer;
        }

        // .cell:hover {
        //     background-color: red;
        // }
    }
}
</style>