<template>
    <header>
        <Header/>
    </header>

    <body>
        <router-view></router-view>
    </body>

    <footer>
        <h1>footer</h1>
    </footer>
</template>

<script setup>
import Header from '../components/header/Header.vue';
</script>

<style lang="less" scoped>

</style>