<script setup lang="ts">
import CpuUsage from './components/CpuUsage.vue';
import TempCheck from './components/TempCheck.vue';
import MemUsage from './components/MemUsage.vue';
</script>

<template>
  <div class="app-container">
    <h1>PC Monitor alfa</h1>
    <div class="left-panel">
      <CpuUsage />
      <TempCheck />
    </div>
    <div class="right-panel">
      <MemUsage />
    </div>
  </div>
</template>

<style scoped>
.app-container {
  align-items: center;
  text-align: center;
}

.left-panel, .right-panel {
  display: inline-block;
  vertical-align: top;
  margin: 20px;
}
</style>