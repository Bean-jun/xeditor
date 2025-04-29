<template>
  <div ref="container" class="monaco-editor"></div>
</template>

<script setup>
import * as monaco from "monaco-editor";
import { onMounted, ref } from "vue";
const props = defineProps({
  value: {
    type: String,
    default: "",
  },
  language: {
    type: String,
    default: "python",
  },
  theme: {
    type: String,
    default: "vs-dark",
  },
  fontSize: {
    type: Number,
    default: 20,
  },
});

const container = ref(null);
let editor = null;

onMounted(() => {
  editor = monaco.editor.create(container.value, {
    value: props.value,
    language: props.language,
    theme: props.theme,
    automaticLayout: true,
    fontSize: `${props.fontSize}px`,
  });

  editor.onDidChangeModelContent(() => {
    const event = new CustomEvent("update:value", {
      detail: editor.getValue(),
    });
    container.value.dispatchEvent(event);
  });
});

</script>

<style scoped>
.monaco-editor {
  text-align: left;
  min-width: 100%;
  min-height: 100%;
}
</style>
