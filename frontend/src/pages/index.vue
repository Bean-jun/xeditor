<template>
  <div class="code-editor">
    <MonacoEditor v-model:value="code" language="python" theme="vs-dark" />
  </div>
</template>

<script setup>
import { ref, reactive, onMounted } from "vue";
import MonacoEditor from "../components/editor/MonacoEditor.vue";

const code = ref("import sys\n\nprint(sys)");
</script>

<style lang="less" scoped>
.code-editor {
  height: 100%;
}
</style>
