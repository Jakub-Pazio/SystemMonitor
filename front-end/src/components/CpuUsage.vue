<template>
  <div>
    <h2>CPU Usage</h2>
    <p :class="{ hot: isHot }">Usage: {{ usage.toFixed(2) }}%</p>
    <CpuStack :usage="usage"/>
  </div>
</template>

<script setup lang="ts">
import {onBeforeUnmount, onMounted, ref} from 'vue';
import CpuStack from './CpuStack.vue';

const usage = ref(0);
const isHot = ref(false);

async function getCpuUsage() {
  try {
    const res = await fetch('/cpu');
    const data = await res.json();

    if (data && data.name === 'cpu' && typeof data.value === 'number') {
      usage.value = (data.value * 100); // Update usage with the value as a percentage
      isHot.value = data.value > 0.7;
    } else {
      console.error('Invalid data format received from the API.');
      usage.value = 0; // Return a default value
    }
  } catch (error) {
    console.error('Error fetching CPU usage:', error);
    usage.value = 0; // Return a default value in case of an error
  }
}

onMounted(() => {
  // Call getCpuUsage every second
  const updateInterval = setInterval(getCpuUsage, 1000);

  // Stop the interval when the component is unmounted (cleanup)
  onBeforeUnmount(() => {
    clearInterval(updateInterval);
  });
});
</script>

<style scoped>
.hot {
  color: red;
}
</style>