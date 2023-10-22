<template>
    <div>
      <h2>CPU Usage</h2>
      <p>Usage: {{ usage.toFixed(2) }}%</p>
    </div>
  </template>

  <script setup lang="ts">
  import { ref, onMounted, onBeforeUnmount } from 'vue';

  const usage = ref(0);

  async function getCpuUsage() {
    try {
      const res = await fetch('/cpu');
      const data = await res.json();

      if (data && data.name === 'cpu' && typeof data.value === 'number') {
        usage.value = (data.value * 100); // Update usage with the value as a percentage
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

  </style>