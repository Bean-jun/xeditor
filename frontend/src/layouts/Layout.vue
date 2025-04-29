<template>
    <header>
        <Header />
    </header>

    <body>
        <router-view></router-view>
    </body>

    <footer>
    </footer>
</template>

<script setup>
import Header from '../components/header/Header.vue';
</script>

<style lang="less" scoped>
body {
    height: 100%;
}
</style>